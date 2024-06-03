#!/usr/bin/env python

#
# This script manipulates pom.xml to either add new specs or update the versions of existing specs.
#


import click
import re

from .single_pom_file_add_or_update_spec import single_pom_file_add_or_update_spec
from .single_pom_file_add_or_update_spec import DEFAULT_POM_LOCATION
from .single_pom_file_add_or_update_spec import DEFAULT_GITHUB_WHITELIST_LOCATION
from .single_pom_file_add_or_update_spec import ns
from .single_pom_file_add_or_update_spec import check_args_for_new_service as base_check_args_for_new_service


DEFAULT_MAKE_FILE_LOCATION = "../Makefile"


GENERATE_EXECUTION_TEMPLATE = """
<execution>
    <id>go-public-sdk-{artifact_id}</id>
    <phase>compile</phase>
    <goals>
        <goal>generate</goal>
    </goals>
    <configuration>
        <language>oracle-go-sdk</language>
        <specPath>${{preprocessed-temp-dir}}/{artifact_id}/${{{artifact_id}-spec-file}}</specPath>
        <outputDir>${{env.GOPATH}}/src/${{fullyQualifiedProjectName}}</outputDir>
        <basePackage>{spec_name}</basePackage>
        <specGenerationType>${{generationType}}</specGenerationType>
        <additionalProperties>
            <specName>{spec_name}</specName>
            <fqProjectName>${{fullyQualifiedProjectName}}</fqProjectName>
            {regional_non_regional_service_overrides}
        </additionalProperties>
        <featureIdConfigFile>${{feature-id-file}}</featureIdConfigFile>
        <featureIdConfigDir>${{feature-id-dir}}</featureIdConfigDir>
        <serviceSlugNameMappingConfigFile>${{spec-temp-dir}}/serviceSlugNameMapping.yaml</serviceSlugNameMappingConfigFile>
    </configuration>
</execution>
"""

CLEAN_ELEMENT_TEMPLATE = """
 <fileset>
    <directory>lib/oci/{spec_name}</directory>
    <includes>
        <include>**/*</include>
    </includes>
</fileset>
"""

GITHUB_WHITELIST_TEMPLATE = '\n^{}/'


# Returns was_changed, was_ignored
def update_endpoint(pom, artifact_id, endpoint):
    subdomain = endpoint
    subdomain = re.sub('^.*://', '', subdomain)  # remove protocol and '://'
    subdomain = re.sub(r'\.{domain}.*$', '', subdomain)  # remove '.{domain}' and everything after it

    results = pom.findall(".//ns:plugin[ns:artifactId='bmc-sdk-swagger-maven-plugin']/ns:executions/ns:execution[ns:id='go-public-sdk-{artifact_id}']//ns:additionalProperties/ns:serviceHostName".format(artifact_id=artifact_id), ns)
    if results:
        endpoint_node = results[0]

        if endpoint_node.text != subdomain:
            endpoint_node.text = subdomain
            return True, False
        return False, False

    print('Ignored update to subdomain/endpoint -- service had no setting for this in pom.xml file and is getting that information from spec')
    return False, True


def check_args_for_new_service(locals):
    base_check_args_for_new_service(locals)
    # not checking 'endpoint' anymore; can be specified either in ticket
    # or in spec using 'x-obmcs-endpoint-template'. If neither is specified,
    # this fails in the generator


def add_spec_name_to_make_file(spec_name, make_file_location):
    specNameToken = '##SPECNAME##'
    with open(make_file_location) as f:
        newText = f.read().replace(specNameToken, "{} {}".format(spec_name, specNameToken))

    with open(make_file_location, "w") as f:
        f.write(newText)


def goify_specname(name):
    return name.replace('_', '').lower()


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

    results = pom.findall(".//ns:plugin[ns:artifactId='bmc-sdk-swagger-maven-plugin']/ns:executions/ns:execution[ns:id='go-public-sdk-{artifact_id}']//ns:additionalProperties/ns:serviceHostName".format(artifact_id=artifact_id), ns)
    if results:
        settings["subdomain"] = results[0].text

    settings["spec_name"] = pom.findall(".//ns:plugin[ns:artifactId='bmc-sdk-swagger-maven-plugin']/ns:executions/ns:execution[ns:id='go-public-sdk-{artifact_id}']//ns:additionalProperties/ns:specName".format(artifact_id=artifact_id), ns)[0].text

    return settings


def add_or_update_spec(artifact_id=None, group_id=None, spec_name=None, relative_spec_path=None, endpoint=None, subdomain=None, version=None, spec_generation_type=None, regional_sub_service_overrides=None, non_regional_sub_service_overrides=None, pom_location=None, github_whitelist_location=None, makefile_location=None):
    if not spec_generation_type:
        spec_generation_type = 'PUBLIC'

    if endpoint and not subdomain:
        subdomain = endpoint
        subdomain = re.sub('^.*://', '', subdomain)  # remove protocol and '://'
        subdomain = re.sub(r'\.{domain}.*$', '', subdomain)  # remove '.{domain}' and everything after it

    # force format of spec_name by removing underscore and lower case
    if spec_name:
        spec_name = goify_specname(spec_name)

    result = single_pom_file_add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, None, subdomain, version, spec_generation_type,
        regional_sub_service_overrides, non_regional_sub_service_overrides, pom_location, github_whitelist_location, GITHUB_WHITELIST_TEMPLATE,
                       GENERATE_EXECUTION_TEMPLATE, CLEAN_ELEMENT_TEMPLATE, update_endpoint, check_args_for_new_service, gather_settings)

    if result.updated and not result.existing:
        # For new spec only
        add_spec_name_to_make_file(spec_name, makefile_location)

    return result


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
@click.option('--pom-location', type=click.Path(exists=True), default=DEFAULT_POM_LOCATION, help='Location of the pom.xml file to update')
@click.option('--github-whitelist-location', type=click.Path(exists=True), default=DEFAULT_GITHUB_WHITELIST_LOCATION, help='Location of the github.whitelist file to update')
@click.option('--makefile-location', type=click.Path(exists=True), default=DEFAULT_MAKE_FILE_LOCATION, help='Location of the Makefile to update')
def add_or_update_spec_command(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version, spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides, pom_location, github_whitelist_location, makefile_location):
    print(add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version, spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides, pom_location, github_whitelist_location, makefile_location))


if __name__ == '__main__':
    add_or_update_spec_command()
