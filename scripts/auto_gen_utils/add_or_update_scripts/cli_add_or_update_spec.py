#!/usr/bin/env python

#
# This script manipulates pom.xml to either add new specs or update the versions of existing specs.
#

import click
import xml.etree.ElementTree as ET
import os
from .single_pom_file_add_or_update_spec import DEFAULT_POM_LOCATION
from .add_or_update_spec_utils import AddOrUpdateSpecResult, compute_changed_settings
from .add_or_update_spec_utils import write_xml
from shared.version_utils import is_version_increasing


XMLNS = {"ns":"http://maven.apache.org/POM/4.0.0"}

# This is the template for a child pom.
POM_TEMPLATE = """
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <parent>
    <groupId>com.oracle.bmc.sdk</groupId>
    <artifactId>python-cli</artifactId>
    <version>1.0.0-SNAPSHOT</version>
    <relativePath>../../pom.xml</relativePath>
  </parent>
  
  <groupId>com.oracle.bmc.sdk</groupId>
  <version>1.0.0-SNAPSHOT</version>
  <modelVersion>4.0.0</modelVersion>
  <artifactId>{spec_name}</artifactId>
  <properties>
    <cli-root-dir>${project.basedir}/../../</cli-root-dir>

    <codegen.artifactory.groupId>{group_id}</codegen.artifactory.groupId>
    <codegen.artifactory.artifactId>{artifact_id}</codegen.artifactory.artifactId>
    <codegen.artifactory.version>{version}</codegen.artifactory.version>
    <codegen.spec.name>{spec_file}</codegen.spec.name>
    <codegen.package.name>{spec_name}</codegen.package.name>
    <generated-output-dir>${project.basedir}</generated-output-dir>
    <generate-rootCliGroupOverride>{root_group_override}</generate-rootCliGroupOverride>
  </properties>

  <build>
    <directory>${cli-root-dir}/target</directory>
  </build>
       
</project>
"""


# Get the settings from an existing pom file.
def gather_settings(pom_file):
    settings = {}
    if not os.path.exists(pom_file):
        return settings

    pom = parse_pom(pom_file)

    # Get the spec_name from the top most artifactId tag.
    xpath = "./ns:artifactId"
    artifactId = pom.findall(xpath, XMLNS)
    if len(artifactId) < 1:
        print("spec_name could not be found in existing pom file")
    else:
        settings["spec_name"] = artifactId[0].text

    add_property_setting("codegen.artifactory.groupId", "group_id", pom, settings)
    add_property_setting("codegen.artifactory.artifactId", "artifact_id", pom, settings)
    add_property_setting("codegen.artifactory.version", "version", pom, settings)
    add_property_setting("codegen.spec.name", "relative_spec_path", pom, settings)
    # Not everything will have a root_group_override
    try:
        add_property_setting("generate-rootCliGroupOverride", "root_group_override", pom, settings)
    except:  # noqa: ignore=E722
        pass

    # Check if the current pom is overriding the default build plugins.
    # For example, key_management-pom.xml and streaming-pom.xml do this.
    for child in pom.getroot():
        if "build" in child.tag:
            for build_child in child:
                if "plugins" in build_child.tag:
                    settings["plugins"] = build_child

    return settings


def add_property_setting(xpath_property, setting_name, pom, settings):
    xpath = ".//ns:properties//ns:" + xpath_property
    properties = pom.findall(xpath, XMLNS)
    if len(properties) < 1:
        raise Exception(setting_name + " not found in pom.xml")
    else:
        settings[setting_name] = properties[0].text


def determine_pom_location(artifact_id, spec_name, services_root_dir):
    # spec_name may not be specified if this is an existing spec
    # so we have to look through existing service directories to find a pom.xml with this artifact_id
    for individual_service_directory in os.listdir(services_root_dir):
        individual_service_directory = os.path.join(services_root_dir, individual_service_directory)
        if os.path.isdir(individual_service_directory):
            for f in os.listdir(individual_service_directory):
                file_path = os.path.join(individual_service_directory, f)
                if os.path.isfile(file_path) and file_path.endswith('pom.xml'):
                    pom = parse_pom(file_path)

                    xpath = ".//ns:properties//ns:codegen.artifactory.artifactId"
                    properties = pom.findall(xpath, XMLNS)
                    if len(properties) < 1:
                        raise Exception("{} not found in {}".format(xpath, file_path))

                    artifact_id_from_pom = properties[0].text
                    if artifact_id_from_pom == artifact_id:
                        return file_path

    # there is no existing pom for this artifact id
    # so return the new location where it should be based on spec_name
    if not spec_name:
        raise Exception("Spec name must be specified for new service")

    pom_path = os.path.join(services_root_dir, spec_name, "pom.xml")

    # create the new directory to contain this pom.xml
    pom_dir = os.path.dirname(pom_path)
    if not os.path.exists(pom_dir):
        os.makedirs(pom_dir)

    return pom_path


def parse_pom(pom_file):
    if not os.path.exists(pom_file):
        return None

    pom = ET.parse(pom_file)

    # allow default namespace for output, dont print ns0: prefixes everywhere
    ET.register_namespace('',"http://maven.apache.org/POM/4.0.0")

    return pom


def add_or_update_spec(artifact_id=None, group_id=None, spec_name=None, relative_spec_path=None, version=None, spec_generation_type=None, add_sub_groups=None, multiple_services_in_spec=None, pom_location=None, root_group_override=None):
    found = None
    ignored = []
    previous = {}
    changed = []

    if not os.path.exists(pom_location):
        os.mkdir(pom_location)

    if not spec_generation_type:
        spec_generation_type = 'PUBLIC'

    if artifact_id:
        file_name = determine_pom_location(artifact_id, spec_name, pom_location)
        previous = gather_settings(file_name)

        if 'artifact_id' in previous:
            found = previous['artifact_id']

        if group_id is None and 'group_id' in previous:
            group_id = previous['group_id']
        if relative_spec_path is None and 'relative_spec_path' in previous:
            relative_spec_path = previous['relative_spec_path']

        # updating spec_name is not supported in self service so always use previous if it exists
        # if we want to support this in the future, the Python SDK will need to support it as well
        if 'spec_name' in previous:
            if spec_name is not None:
                ignored.append('spec_name')

            spec_name = previous['spec_name']

        previous_version = None
        if 'version' in previous:
            previous_version = previous['version']
        new_version = None
        if previous_version is None or previous_version == version or is_version_increasing(previous_version, version):
            new_version = version
        else:
            new_version = previous_version
            print(spec_name, ': The version was not updated to {}, because it was already at {}.'.format(version, previous_version))
        
        pom_string = POM_TEMPLATE

        root_group_override = ''
        if not multiple_services_in_spec:
            if 'root_group_override' in previous:
                # This preserves the existing root_group_override already found in the pom
                root_group_override = previous['root_group_override']
            else:
                root_group_override = spec_name.replace('_', '-')
        else:
            # Remove the root group override for multiple service specs.
            pom_string = pom_string.replace("<generate-rootCliGroupOverride>{root_group_override}</generate-rootCliGroupOverride>", "")        

        # Note, we use regular string replace instead of format so we don't have to double and triple-escape things in the template
        if artifact_id is not None:
            pom_string = pom_string.replace("{artifact_id}", artifact_id)  # required in 1_process_preview_jira_queue
        if new_version is not None:
            pom_string = pom_string.replace("{version}", new_version)  # required in 1_process_preview_jira_queue
        if spec_name is not None:
            pom_string = pom_string.replace("{spec_name}", spec_name)
        else:
            print("Warning: Could not determine the spec_name from jira or previous pom.")
        if group_id is not None:
            pom_string = pom_string.replace("{group_id}", group_id)
        else:
            print("Warning: Could not determine the group_id from jira or previous pom.")
        if relative_spec_path is not None:
            pom_string = pom_string.replace("{spec_file}", relative_spec_path)
        else:
            print("Warning: Could not determine the spec_file from jira or previous pom.")
        pom_string = pom_string.replace("{root_group_override}", root_group_override)

        root = ET.fromstring(pom_string)
        pom = ET.ElementTree(element=root)
        # If our previous version of the pom was overriding the default build plugins,
        # we need to preserve that.  
        if 'plugins' in previous:
            build = root.find('{http://maven.apache.org/POM/4.0.0}build')
            build.append(previous['plugins'])
        write_xml(file_name, pom)        
        current = gather_settings(file_name)
        changed = compute_changed_settings(previous, current)
    
    return AddOrUpdateSpecResult(
        updated=(not found) or changed != [],  # not found means it's a new spec, or if it is an existing spec, changed needs to be non-empty
        existing=found is not None,
        ignored=ignored,
        previous=previous,
        changed=changed)


@click.command()
@click.option('--artifact-id', help='The artifact id for the spec artifact (e.g. coreservices-api-spec')
@click.option('--group-id', help='The group id for the spec artifact (e.g. com.oracle.pic.commons)')
@click.option('--spec-name', help='The name of the spec. This will be used (e.g. core, identity, object_storage). This is also used as the module name.')
@click.option('--relative-spec-path', help='The relative path of the spec within the artifact (e.g. coreservices-api-spec-20160918-external.yaml)')
@click.option('--version', help='The version of the spec artifact (e.g. 0.0.1-SNAPSHOT')
@click.option('--spec-generation-type', help='The generation type: PUBLIC or PREVIEW')
@click.option('--add-sub-groups', is_flag=True, help='For new specs this will always be true (without it, nothing from the new spec would be added). For existing specs, providing this value causes all commands in the spec to be automatically added to the CLI.')
@click.option('--multiple-services-in-spec', is_flag=True, help='Provide this flag if the spec contains multiple services (e.g. the Core spec has Compute, Block Storage and Virtual Networking). This will disable behaviour such as root group overrides as otherwise multiple services would try and use the same root group and commands would be suppressed')
@click.option('--pom-location', type=click.Path(exists=True), default=DEFAULT_POM_LOCATION, help='Location of the pom.xml file to update')
def add_or_update_spec_command(artifact_id, group_id, spec_name, relative_spec_path, version, spec_generation_type, add_sub_groups, multiple_services_in_spec, pom_location):
    print(add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, version, spec_generation_type, add_sub_groups, multiple_services_in_spec, pom_location))


if __name__ == '__main__':
    pass
