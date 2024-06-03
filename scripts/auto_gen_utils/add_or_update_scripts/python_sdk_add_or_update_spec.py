#!/usr/bin/env python

#
# This script manipulates pom.xml to either add new specs or update the versions of existing specs.
#

import os
import re
import click
from click.exceptions import MissingParameter

from .add_or_update_spec_utils import convert_camel_to_snake_case
from .add_or_update_spec_utils import parse_pom
from .module_pom_file_add_or_update_spec import module_pom_file_add_or_update_spec
from .module_pom_file_add_or_update_spec import DEFAULT_POM_LOCATION
from .module_pom_file_add_or_update_spec import DEFAULT_GITHUB_WHITELIST_LOCATION
from .module_pom_file_add_or_update_spec import ns
from .module_pom_file_add_or_update_spec import check_args_for_new_service as base_check_args_for_new_service

DEFAULT_MODULE_LOCATION = "poms"

GENERATE_EXECUTION_TEMPLATE = """
<execution>
    <id>python-public-sdk-{artifact_id}</id>
    <phase>compile</phase>
    <goals>
        <goal>generate</goal>
    </goals>
    <configuration>
        <language>${{codegen-language}}</language>
        <specPath>${{preprocessed-temp-dir}}/{artifact_id}/${{{artifact_id}-spec-file}}</specPath>
        <outputDir>src/oci</outputDir>
        <basePackage>OCI</basePackage>
        <specGenerationType>{spec_generation_type}</specGenerationType>
        <additionalProperties>
            <specName>{spec_name}</specName>
            <generateInitFile>true</generateInitFile>
            {regional_non_regional_service_overrides}
        </additionalProperties>
        <featureIdConfigFile>${{feature-id-file}}</featureIdConfigFile>
        <featureIdConfigDir>${{feature-id-dir}}</featureIdConfigDir>
        <serviceSlugNameMappingConfigFile>${{spec-temp-dir}}/serviceSlugNameMapping.yaml</serviceSlugNameMappingConfigFile>
        <isTestGenerationEnabled>true</isTestGenerationEnabled>
    </configuration>
</execution>
"""


CLEAN_ELEMENT_TEMPLATE = """
 <fileset>
    <directory>src/oci/{spec_name}</directory>
    <includes>
        <include>**/*</include>
    </includes>
</fileset>
"""

MODULE_TEMPLATE = '<module>poms/{}/pom.xml</module>'

GITHUB_WHITELIST_TEMPLATE = '\n^src/oci/{}/.*' + r'\.py$'


# Returns was_changed, was_ignored
def update_endpoint(pom, artifact_id, endpoint):
    results = pom.findall(".//ns:plugin[ns:artifactId='bmc-sdk-swagger-maven-plugin']/ns:executions/ns:execution[ns:id='python-public-sdk-{artifact_id}']//ns:additionalProperties/ns:endpoint".format(artifact_id=artifact_id), ns)
    if results:
        # Only allow update if the pom.xml file has the endpoint
        endpoint_node = results[0]
        if endpoint_node.text != endpoint:
            endpoint_node.text = endpoint
            return True, False
        else:
            return False, False

    print('Ignored update to subdomain/endpoint -- service had no setting for this in pom.xml file and is getting that information from spec')
    return False, True


def check_args_for_new_service(locals):
    base_check_args_for_new_service(locals)
    # not checking 'endpoint' anymore; can be specified either in ticket
    # or in spec using 'x-obmcs-endpoint-template'. If neither is specified,
    # this fails in the generator


def gather_settings(pom, artifact_id):
    settings = {}

    xpath = ".//ns:dependencyManagement//ns:dependency[ns:artifactId='{artifact_id}']".format(artifact_id=artifact_id)
    dependency = pom.findall(xpath, ns)[0]
    settings["group_id"] = dependency.find('./ns:groupId', ns).text
    settings["artifact_id"] = artifact_id
    settings["version"] = dependency.find('./ns:version', ns).text

    xpath = ".//ns:properties/ns:{artifact_id}-spec-file".format(artifact_id=artifact_id)
    spec_file_node = pom.findall(xpath, ns)[0]
    settings["relative_spec_path"] = spec_file_node.text

    results = pom.findall(".//ns:plugin[ns:artifactId='bmc-sdk-swagger-maven-plugin']/ns:executions/ns:execution[ns:id='python-public-sdk-{artifact_id}']//ns:additionalProperties/ns:endpoint".format(artifact_id=artifact_id), ns)
    if results:
        endpoint_node = results[0]
        subdomain = endpoint_node.text
        subdomain = re.sub('^.*://', '', subdomain)  # remove protocol and '://'
        subdomain = re.sub(r'\.{domain}.*$', '', subdomain)  # remove '.{domain}' and everything after it
        settings["subdomain"] = subdomain

    settings["spec_name"] = get_spec_name(pom)

    return settings


def get_spec_name(pom):
    xpath_for_spec_name = "./ns:artifactId"
    spec_name = pom.find(xpath_for_spec_name, ns).text

    return spec_name


def existing_spec(source_path, artifact_id):
    xpath_for_spec_dependency_declaration = ".//ns:dependency[ns:artifactId='{artifact_id}']".format(artifact_id=artifact_id)
    for item in os.listdir(source_path):
        path = os.path.join(source_path, item)
        if os.path.isdir(path):
            pom_location = os.path.join(path, "pom.xml")
            if os.path.exists(pom_location):
                pom = parse_pom(pom_location)
                if (pom.findall(xpath_for_spec_dependency_declaration, ns)):
                    print("Existing service")
                    return pom, pom_location

    return None, None


def add_or_update_spec(artifact_id=None,
                       group_id=None,
                       spec_name=None,
                       relative_spec_path=None,
                       endpoint=None,
                       subdomain=None,
                       version=None,
                       spec_generation_type=None,
                       regional_sub_service_overrides=None,
                       non_regional_sub_service_overrides=None,
                       pom_location=None,
                       github_whitelist_location=None,
                       module_location=None):

    if pom_location and not module_location:
        module_location = os.path.join(os.path.dirname(pom_location), DEFAULT_MODULE_LOCATION)
        print("Location of pom modules: {}".format(module_location))

    # force format of spec_name to (lower) snake case for consistency with standards of python SDK and CLI
    if spec_name:
        spec_name = convert_camel_to_snake_case(spec_name)

    service_pom, service_pom_path = existing_spec(module_location, artifact_id)
    print("service_pom='{}', service_pom_path='{}'".format(service_pom, service_pom_path))

    if not service_pom_path:
        if not spec_name:
            raise MissingParameter('Must specify --spec-name for new spec')
        service_pom_path = os.path.join(module_location, spec_name, "pom.xml")
        os.makedirs(os.path.join(module_location, spec_name))
        print("New module pom path: {}".format(service_pom_path))

    return module_pom_file_add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain,
                                              version, spec_generation_type, regional_sub_service_overrides,
                                              non_regional_sub_service_overrides, pom_location,
                                              github_whitelist_location, GITHUB_WHITELIST_TEMPLATE,
                                              GENERATE_EXECUTION_TEMPLATE, CLEAN_ELEMENT_TEMPLATE, update_endpoint,
                                              check_args_for_new_service, gather_settings, module_pom=service_pom,
                                              module_pom_path=service_pom_path, module_template=MODULE_TEMPLATE)


@click.command()
@click.option('--artifact-id', help='The artifact id for the spec artifact (e.g. coreservices-api-spec')
@click.option('--group-id', help='The group id for the spec artifact (e.g. com.oracle.pic.commons)')
@click.option('--spec-name', help='The name of the spec. This will be used (e.g. core, identity, object_storage). This is also used as the module name.')
@click.option('--relative-spec-path', help='The relative path of the spec within the artifact (e.g. coreservices-api-spec-20160918-external.yaml)')
@click.option('--endpoint', help='The base endpoint for the service (e.g. https://iaas.{domain}/20160918)')
@click.option('--subdomain', help='The subdomain for the service (e.g. if the endpoint is https://iaas.{domain}/20160918), the subdomain is "iaas"')
@click.option('--version', help='The version of the spec artifact (e.g. 0.0.1-SNAPSHOT')
@click.option('--spec-generation-type', help='The generation type: PUBLIC or PREVIEW')
@click.option('--regional-sub-service-overrides', multiple=True, help="""For specs that contain multiple services
(because there are operations with different tags in the spec), which of those services should be considered regional.
Services are considered as regional by default.

This should be the snake_cased name of the tag/service. For example kms_provisioning instead of kmsProvisioning.

This parameter can be provided multiple times""")
@click.option('--non-regional-sub-service-overrides', multiple=True, help="""For specs that contain multiple services
(because there are operations with different tags in the spec), which of those services should be considered non-regional.

This should be the snake_cased name of the tag/service. For example kms_provisioning instead of kmsProvisioning.

This parameter can be provided multiple times""")
@click.option('--pom-location',
              type=click.Path(exists=True),
              default=DEFAULT_POM_LOCATION,
              help='Location of the pom.xml file to update')
@click.option('--github-whitelist-location',
              type=click.Path(exists=True),
              default=DEFAULT_GITHUB_WHITELIST_LOCATION,
              help='Location of the github.whitelist file to update')
@click.option('--module-location',
              type=click.Path(exists=True),
              default=DEFAULT_MODULE_LOCATION,
              help="Parent directory containing the module pom.xml files")
def add_or_update_spec_command(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version,
                               spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides,
                               pom_location, github_whitelist_location, module_location):

    if not artifact_id:
        raise click.exceptions.MissingParameter(param_type='option', param_hint='--artifact-id', message='Artifact id parameter is required')

    if subdomain and endpoint:
        raise click.exceptions.UsageError('Cannot specify both --endpoint and --subdomain')

    print(add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version,
                             spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides,
                             pom_location, github_whitelist_location, module_location))


if __name__ == '__main__':
    add_or_update_spec_command()
