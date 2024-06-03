#!/usr/bin/env python

#
# This is a library for manipulating a single pom.xml file to either
# add new specs or update the versions of existing specs.
#
# Can be used by Python SDK, CLI, Ruby SDK and Go SDK

import os
import xml.etree.ElementTree as ET
from click.exceptions import MissingParameter, UsageError

from .add_or_update_spec_utils import parse_pom, write_xml, AddOrUpdateSpecResult
from .add_or_update_spec_utils import compute_changed_settings
from .add_or_update_spec_utils import add_spec_module_to_github_whitelist
from shared.version_utils import is_version_increasing

DEFAULT_POM_LOCATION = "pom.xml"
DEFAULT_GITHUB_WHITELIST_LOCATION = "github.whitelist"

SPEC_FILE_PROPERTY_TEMPLATE = """
<{artifact_id}-spec-file>{spec_path_relative_to_jar}</{artifact_id}-spec-file>
"""

UNPACK_EXECUTION_TEMPLATE = """
<execution>
    <id>unpack-{artifact_id}</id>
    <phase>initialize</phase>
    <goals>
        <goal>unpack</goal>
    </goals>
    <configuration>
        <artifactItems>
            <artifactItem>
                <groupId>{group_id}</groupId>
                <artifactId>{artifact_id}</artifactId>
                <type>jar</type>
                <includes>**/*</includes>
                <outputDirectory>${{spec-temp-dir}}/{artifact_id}</outputDirectory>
            </artifactItem>
        </artifactItems>
    </configuration>
</execution>
"""

PREFER_EXECUTION_TEMPLATE = """
<execution>
    <id>spec-conditionals-prefer-{artifact_id}</id>
    <phase>initialize</phase>
    <goals>
        <goal>prefer</goal>
    </goals>
    <configuration>
        <inputFiles>
            <!-- New layout: source/<spec.proto.yaml> -->
            <inputFile>${{spec-temp-dir}}/{artifact_id}/source/${{{artifact_id}-spec-file}}</inputFile>
            <!-- Old layout: ./<spec.proto.yaml> -->
            <inputFile>${{spec-temp-dir}}/{artifact_id}/${{{artifact_id}-spec-file}}</inputFile>
        </inputFiles>
        <outputFile>${{preferred-temp-dir}}/{artifact_id}/${{{artifact_id}-spec-file}}</outputFile>
    </configuration>
</execution>
"""

PREPROCESS_EXECUTION_TEMPLATE = """
<execution>
    <id>spec-conditionals-preprocess-{artifact_id}</id>
    <phase>initialize</phase>
    <goals>
        <goal>preprocess</goal>
    </goals>
    <configuration>
        <inputFile>${{preferred-temp-dir}}/{artifact_id}/${{{artifact_id}-spec-file}}</inputFile>
        <outputFile>${{preprocessed-temp-dir}}/{artifact_id}/${{{artifact_id}-spec-file}}</outputFile>
        <groupFile>${{enabled-groups-file}}</groupFile>
        <groupFileCollectionDir>${{enabled-groups-dir}}</groupFileCollectionDir>
    </configuration>
</execution>
"""


DEPENDENCY_MANAGEMENT_TEMPLATE = """
    <dependency>
        <groupId>{group_id}</groupId>
        <artifactId>{artifact_id}</artifactId>
        <version>{version}</version>
    </dependency>
"""

MODULE_ARTIFACT_ID_TEMPLATE = """{parent_artifact_id}-{spec_name}-codegen"""

ns = {"ns": "http://maven.apache.org/POM/4.0.0"}

# allow default namespace for output, dont print ns0: prefixes everywhere
ET.register_namespace('', "http://maven.apache.org/POM/4.0.0")


def add_artifact_id_to_module(pom, parent_artifact_id, spec_name):
    # Search for an artifactId node in the root of pom.
    xpath = "ns:artifactId"
    artfact_node = pom.findall(xpath, ns)[0]
    artfact_node.text = MODULE_ARTIFACT_ID_TEMPLATE.format(parent_artifact_id=parent_artifact_id,
                                                           spec_name=spec_name)


def get_artifact_id(pom):
    xpath = ".//ns:artifactId"
    artfact_id = pom.findall(xpath, ns)[0].text
    return artfact_id


def generate_and_add_property_element(pom, artifact_id, spec_path_relative_to_jar):
    content = SPEC_FILE_PROPERTY_TEMPLATE.format(
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar)

    property_element = ET.fromstring(content)

    xpath = ".//ns:properties"
    properties = pom.findall(xpath, ns)[0]
    properties.append(property_element)


def add_module_to_parent_pom(pom, module_entry):
    module_element = ET.fromstring(module_entry)
    xpath = ".//ns:modules"
    properties = pom.findall(xpath, ns)[0]
    properties.insert(0, module_element)


def update_relative_spec_path(pom, artifact_id, spec_path_relative_to_jar):
    xpath = ".//ns:properties/ns:{artifact_id}-spec-file".format(artifact_id=artifact_id)
    spec_file_node = pom.findall(xpath, ns)[0]
    if spec_file_node.text != spec_path_relative_to_jar:
        spec_file_node.text = spec_path_relative_to_jar
        return True
    return False


def generate_and_add_unpack_element(pom, group_id, artifact_id, spec_path_relative_to_jar):
    content = UNPACK_EXECUTION_TEMPLATE.format(
        group_id=group_id,
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar)

    unpack_element = ET.fromstring(content)

    # find dex-get-spec-artifact-plugin where unpacking happens
    unpack_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='dex-get-spec-artifact-plugin']/ns:executions", ns)[0]
    unpack_plugin_executions.append(unpack_element)


def generate_and_add_prefer_element(pom, group_id, artifact_id, spec_path_relative_to_jar):
    content = PREFER_EXECUTION_TEMPLATE.format(
        group_id=group_id,
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar)

    unpack_element = ET.fromstring(content)

    # find spec-conditionals-preprocessor-plugin where unpacking happens
    unpack_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='spec-conditionals-preprocessor-plugin']/ns:executions", ns)[0]
    unpack_plugin_executions.append(unpack_element)


def generate_and_add_preprocess_element(pom, group_id, artifact_id, spec_path_relative_to_jar):
    content = PREPROCESS_EXECUTION_TEMPLATE.format(
        group_id=group_id,
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar)

    unpack_element = ET.fromstring(content)

    # find spec-conditionals-preprocessor-plugin where unpacking happens
    unpack_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='spec-conditionals-preprocessor-plugin']/ns:executions", ns)[0]
    unpack_plugin_executions.append(unpack_element)


def generate_and_add_generate_section(pom, spec_name, artifact_id, spec_path_relative_to_jar, spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides, generate_execution_template):
    regional_non_regional_service_overrides_content = ''
    if regional_sub_service_overrides or non_regional_sub_service_overrides:
        if regional_sub_service_overrides:
            for override in regional_sub_service_overrides:
                regional_non_regional_service_overrides_content += '<isRegionalClient.{service_name}>true</isRegionalClient.{service_name}>\n'.format(service_name=override)

        if non_regional_sub_service_overrides:
            for override in non_regional_sub_service_overrides:
                regional_non_regional_service_overrides_content += '<isRegionalClient.{service_name}>false</isRegionalClient.{service_name}>\n'.format(service_name=override)

    content = generate_execution_template.format(
        artifact_id=artifact_id,
        spec_name=spec_name,
        spec_path_relative_to_jar=spec_path_relative_to_jar,
        spec_generation_type=spec_generation_type,
        regional_non_regional_service_overrides=regional_non_regional_service_overrides_content)

    generate_element = ET.fromstring(content)

    # find bmc-sdk-swagger-maven-plugin where generation happens
    generate_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='bmc-sdk-swagger-maven-plugin']/ns:executions", ns)[0]
    generate_plugin_executions.append(generate_element)


def generate_and_add_clean_section(pom, spec_name, clean_element_template):
    if not clean_element_template:
        return

    content = clean_element_template.format(
        spec_name=spec_name)

    clean_element = ET.fromstring(content)

    # find filesetes where clean directory goes
    filesets = pom.findall(".//ns:plugin[ns:artifactId='maven-clean-plugin']//ns:filesets", ns)[0]
    filesets.append(clean_element)


def generate_and_add_dependency_management_section(pom, group_id, artifact_id, version):
    content = DEPENDENCY_MANAGEMENT_TEMPLATE.format(
        group_id=group_id,
        artifact_id=artifact_id,
        version=version)

    dep_mgt_element = ET.fromstring(content)

    # find dependencies where version is specified
    dependencies = pom.findall(".//ns:dependencyManagement/ns:dependencies", ns)[0]
    dependencies.append(dep_mgt_element)


def update_version_of_existing_spec(pom, artifact_id, version):
    xpath = ".//ns:dependencyManagement//ns:dependency[ns:artifactId='{artifact_id}']".format(artifact_id=artifact_id)
    dependency = pom.findall(xpath, ns)[0]
    old_version = dependency.find('./ns:version', ns).text

    if not is_version_increasing(old_version, version):
        return old_version

    dependency.find('./ns:version', ns).text = version
    return None  # the old version was lower


def check_args_for_new_service(locals):
    if not locals['version']:
        raise MissingParameter('Must specify --version for new spec')

    if not locals['group_id']:
        raise MissingParameter('Must specify --group-id for new spec')

    if not locals['spec_name']:
        raise MissingParameter('Must specify --spec-name for new spec')

    if not locals['relative_spec_path']:
        raise MissingParameter('Must specify --relative-spec-path for new spec')


def module_pom_file_add_or_update_spec(artifact_id=None, group_id=None, spec_name=None,
                                       relative_spec_path=None, endpoint=None, subdomain=None, version=None,
                                       spec_generation_type=None, regional_sub_service_overrides=None,
                                       non_regional_sub_service_overrides=None, pom_location=None,
                                       github_whitelist_location=None, github_whitelist_template=None, generate_execution_template=None,
                                       clean_element_template=None, update_endpoint_function=None,
                                       check_args_for_new_service_function=check_args_for_new_service, gather_settings=None,
                                       module_pom=None, module_pom_path=None, module_template=None):

    found = False

    if not generate_execution_template:
        raise ValueError('Must supply generate_execution_template')

    if not artifact_id:
        raise MissingParameter(param_type='option', param_hint='--artifact-id', message='Artifact id parameter is required')

    if subdomain and endpoint:
        raise UsageError('Cannot specify both --endpoint and --subdomain')

    # Parent pom
    pom = parse_pom(pom_location)

    updated_spec = False

    ignored = []
    previous = {}
    changed = []

    # determine if this artifact is already in the spec
    # xpath_for_spec_dependency_declaration = ".//ns:dependency[ns:artifactId='{artifact_id}']".format(artifact_id=artifact_id)
    if module_pom:
        print('Artifact {} already exists in module pom. Updating specified fields...'.format(artifact_id))

        found = True

        previous = gather_settings(module_pom, artifact_id)

        if version:
            # Todo: Determine if we need to find the version again in update_version_of_existing_spec.  It should
            #       already be in `previous["version"]`
            newer_version = update_version_of_existing_spec(module_pom, artifact_id, version)
            if newer_version:
                print('The version was not updated to {}, because it was already at {}.'.format(version, newer_version))
            else:
                updated_spec |= True

        if relative_spec_path:
            updated_spec |= update_relative_spec_path(module_pom, artifact_id, relative_spec_path)

        if update_endpoint_function:
            was_ignored = False
            if endpoint:
                was_changed, was_ignored = update_endpoint_function(module_pom, artifact_id, endpoint)
                updated_spec |= was_changed
            elif subdomain:
                was_changed, was_ignored = update_endpoint_function(module_pom, artifact_id, 'https://{}.{{domain}}'.format(subdomain))
                updated_spec |= was_changed

            if was_ignored:
                ignored.append('subdomain')

        if spec_name:
            ignored.append('spec_name')
        if group_id:
            ignored.append('group_id')

        current = gather_settings(module_pom, artifact_id)
        changed = compute_changed_settings(previous, current)
    else:
        print(os.getcwd())
        module_pom = parse_pom("add_or_update_scripts/templates/pom-template.xml")

        if endpoint:
            print('Ignoring endpoint setting for new services; new services have to get it from the spec')
            ignored.append('endpoint')

        if subdomain:
            print('Ignoring subdomain setting for new services; new services have to get it from the spec')
            ignored.append('subdomain')

        if subdomain and not endpoint:
            endpoint = 'https://{}.{{domain}}'.format(subdomain)

        check_args_for_new_service_function(locals())

        if not spec_generation_type:
            spec_generation_type = 'PUBLIC'

        print('Artifact {} does not exist. Adding it...'.format(spec_name))
        parent_artifact_id = get_artifact_id(pom)
        print("Parent_artifact_id: {}".format(parent_artifact_id))
        print("spec_name: {}".format(spec_name))
        add_artifact_id_to_module(module_pom, parent_artifact_id, spec_name)
        generate_and_add_property_element(module_pom, artifact_id, relative_spec_path)
        generate_and_add_unpack_element(module_pom, group_id, artifact_id, relative_spec_path)
        generate_and_add_prefer_element(module_pom, group_id, artifact_id, relative_spec_path)
        generate_and_add_preprocess_element(module_pom, group_id, artifact_id, relative_spec_path)
        generate_and_add_generate_section(module_pom, spec_name, artifact_id, relative_spec_path, spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides, generate_execution_template)
        generate_and_add_clean_section(module_pom, spec_name, clean_element_template)
        generate_and_add_dependency_management_section(module_pom, group_id, artifact_id, version)
        add_spec_module_to_github_whitelist(spec_name, github_whitelist_location, github_whitelist_template)
        add_module_to_parent_pom(pom, module_template.format(spec_name))

        updated_spec = True

    if updated_spec:
        write_xml(pom_location, pom)
        write_xml(module_pom_path, module_pom)

        print('====== Success! ======')
        print("""
Next Steps
----------
1. Run mvn clean install to update generated code
""")
    else:
        print('===== Spec was not updated =====')

    return AddOrUpdateSpecResult(
        updated=updated_spec,
        existing=found,
        ignored=ignored,
        previous=previous,
        changed=changed
    )
