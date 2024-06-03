#!/usr/bin/env python

#
# This script manipulates the pom.xml tree to either add new specs or update the versions of existing specs.
#


from .spec_updater_base import SpecUpdaterBase
import xml.etree.ElementTree as ET
import re
import click
import os
import string
import uuid
from click.exceptions import UsageError, MissingParameter
from glob import glob

from shared.version_utils import is_version_increasing
from .add_or_update_spec_utils import AddOrUpdateSpecResult, compute_changed_settings, indent
from .add_or_update_spec_utils import add_spec_module_to_github_whitelist, write_xml
from .add_or_update_spec_utils import CommentedTreeBuilder
from .add_or_update_spec_utils import parse_pom


DEFAULT_PARENT_POM_LOCATION = "pom.xml"
DEFAULT_GITHUB_WHITELIST_LOCATION = "github.whitelist"
ns = {"ns":"http://maven.apache.org/POM/4.0.0"}

# The path for the modules is under "oci-dotnet-sdk/Codegen"
DOTNETSDK_CODEGEN_LOCATION = "Codegen"


# Template for include each module in pom.xml
MODULE_TEMPLATE = "<module>{name}</module>"
# Pom.xml template specific to DotNet SDK
DOTNETSDK_POM_FILE_TEMPLATE = """<?xml version='1.0' encoding='UTF-8'?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <parent>
    <groupId>com.oracle.oci.sdk</groupId>
    <artifactId>oci-dotnet-sdk-codegen</artifactId>
    <version>{sdk_version}</version>
    <relativePath>..</relativePath>
  </parent>
  <artifactId>oci-dotnet-sdk-{service_name}-codegen</artifactId>
  <name>Oracle Cloud Infrastructure SDK - {service_friendly_name} Service Codegen</name>
  <description>This project contains the SDK used for Oracle Cloud Infrastructure {service_friendly_name}</description>
  <properties>
    <codegen.artifactory.groupId>{group_id}</codegen.artifactory.groupId>
    <codegen.artifactory.artifactId>{artifact_id}</codegen.artifactory.artifactId>
    <codegen.artifactory.version>{artifact_version}</codegen.artifactory.version>
    <codegen.spec.name>{spec_path_relative_to_jar}</codegen.spec.name>
    <codegen.package.name>{module_name}</codegen.package.name>
    <codegen.generate.waiters>{generate_waiters}</codegen.generate.waiters>
    <codegen.generate.paginators>{generate_paginators}</codegen.generate.paginators>
    <package.name.lower>{service_name}</package.name.lower>
  </properties>
  <build>
    <plugins>
        <plugin>
            <groupId>org.commonjava.maven.plugins</groupId>
            <artifactId>directory-maven-plugin</artifactId>
        </plugin>
        <plugin>
            <groupId>org.codehaus.mojo</groupId>
            <artifactId>build-helper-maven-plugin</artifactId>
        </plugin>
        <plugin>
            <groupId>com.oracle.oci.sdk.utilities</groupId>
            <artifactId>dex-get-spec-artifact-plugin</artifactId>
        </plugin>
        <plugin>
            <groupId>com.oracle.oci.sdk.utilities</groupId>
            <artifactId>spec-conditionals-preprocessor-plugin</artifactId>
        </plugin>
        <plugin>
            <groupId>com.oracle.bmc.sdk</groupId>
            <artifactId>bmc-sdk-swagger-maven-plugin</artifactId>
        </plugin>
        <plugin>
            <groupId>com.mycila</groupId>
            <artifactId>license-maven-plugin</artifactId>
        </plugin>
        <plugin>
            <artifactId>maven-antrun-plugin</artifactId>
        </plugin>
    </plugins>
    </build>
</project>
"""

MODULE_PROJECT_FILE_TEMPLATE = """
<Project Sdk="Microsoft.NET.Sdk">
  <Import Project="..\\General.targets" />
  <PropertyGroup>
    <ProjectGuid>{{{guid}}}</ProjectGuid>
    <OutputType>Library</OutputType>
    <TargetFramework>netstandard2.0</TargetFramework>
    <LangVersion>8.0</LangVersion>
    <PackageId>OCI.DotNetSDK.{module_name}</PackageId>
    <RootNamespace>Oci.{module_name}Service</RootNamespace>
    <Authors>Oracle Cloud Infrastructure</Authors>
    <Company>Oracle</Company>
    <PackageTags>Oracle;OCI;Oracle Cloud;OracleCloud;oci-sdk;OracleCloudInfrastructure;{module_name}</PackageTags>
    <Description>Oracle Cloud Infrastructure Cloud {service_friendly_name} Service</Description>
  </PropertyGroup>
  <ItemGroup>
    <ProjectReference Include="..\\Common\\OCI.DotNetSDK.Common.csproj"/>
  </ItemGroup>
</Project>
"""


# The dotnet_sdk-specific pom.xml template parameters for XML parsing
DOTNETSDK_SPEC_PARAMS_XML_PATH_DICT = {
    'group_id': ".//ns:properties//ns:codegen.artifactory.groupId",
    'artifact_id': ".//ns:properties//ns:codegen.artifactory.artifactId",
    'version':  ".//ns:properties//ns:codegen.artifactory.version",
    'relative_spec_path': ".//ns:properties//ns:codegen.artifact.spec.path",
    'service_name': ".//ns:properties//ns:codegen.service.name",
    'subdomain': ".//ns:properties//ns:codegen.service.group.endpoint"
}


MODULE_README_FORMAT = """
# OCI .NET client for {service_friendly_name} Service

This module enables you to write code to manage resources for {service_friendly_name} Service.

## Requirements

To use this module, you must have the following:

- An Oracle Cloud Infrastructure account.
- A user created in that account, in a group with a policy that grants the desired permissions. This can be a user for yourself, or another person/system that needs to call the API. For an example of how to set up a new user, group, compartment, and policy, see [Adding Users](https://docs.cloud.oracle.com/en-us/iaas/Content/GSG/Tasks/addingusers.htm). For a list of typical policies you may want to use, see [Common Policies](https://docs.cloud.oracle.com/en-us/iaas/Content/Identity/Concepts/commonpolicies.htm).
- A key pair used for signing API requests, with the public key uploaded to Oracle. Only the user calling the API should be in possession of the private key. For more information, see [Configuring Credentials](https://docs.cloud.oracle.com/en-us/iaas/Content/API/SDKDocs/javasdkconfig.htm)

## Installing

Use the following command to install this module:

```
dotnet add package OCI.DotNetSDK.{module_name}
```
"""


PROJECT_ENTRY_TEMPLATE = """Project("{{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}}") = "OCI.DotNetSDK.{module_name}", "{module_name}\\OCI.DotNetSDK.{module_name}.csproj", "{{{guid}}}"
EndProject
"""


PROJECT_PLATEFORM_TEMPLATE = """        {{{guid}}}.Debug|Any CPU.ActiveCfg = Debug|Any CPU
        {{{guid}}}.Debug|Any CPU.Build.0 = Debug|Any CPU
        {{{guid}}}.Debug|x64.ActiveCfg = Debug|Any CPU
        {{{guid}}}.Debug|x64.Build.0 = Debug|Any CPU
        {{{guid}}}.Debug|x86.ActiveCfg = Debug|Any CPU
        {{{guid}}}.Debug|x86.Build.0 = Debug|Any CPU
        {{{guid}}}.Release|Any CPU.ActiveCfg = Release|Any CPU
        {{{guid}}}.Release|Any CPU.Build.0 = Release|Any CPU
        {{{guid}}}.Release|x64.ActiveCfg = Release|Any CPU
        {{{guid}}}.Release|x64.Build.0 = Release|Any CPU
        {{{guid}}}.Release|x86.ActiveCfg = Release|Any CPU
        {{{guid}}}.Release|x86.Build.0 = Release|Any CPU
"""


class DotNetSDKSpecUpdater(SpecUpdaterBase):
    # Override the spec name that is defined in the service pom.xml files as the testing service
    # references folders without snake-case (based on the .NET sdk formatting).
    def format_module_name_for_template(self, module_name):
        return module_name.replace("-", "").replace("_", "").replace(" ","")


def add_or_update_spec(artifact_id=None, group_id=None, spec_name=None, relative_spec_path=None, endpoint=None, subdomain=None, version=None, spec_generation_type=None, generate_waiters=True, generate_paginators=True, pom_location=None, github_whitelist_location=None):
    sdk_dir = os.path.dirname(pom_location)

    found = find_existing_module(sdk_dir, artifact_id)

    ignored = []
    previous = {}
    changed = []
    if found:
        print('Artifact {} already exists in pom.xml. Updating specified fields...'.format(artifact_id))

        previous = gather_settings(sdk_dir, found)

        if version:
            newer_version = update_version_of_existing_spec(sdk_dir, found, version)
            if newer_version:
                print('The version was not updated to {}, because it was already at {}.'.format(version, newer_version))

        if relative_spec_path:
            update_relative_spec_path_of_existing_spec(sdk_dir, found, relative_spec_path)

        was_ignored = False
        if endpoint:
            subdomain = endpoint
            subdomain = re.sub('^.*://', '', subdomain)  # remove protocol and '://'
            subdomain = re.sub(r'\.{domain}.*$', '', subdomain)  # remove '.{domain}' and everything after it
            was_changed, was_ignored = update_subdomain(sdk_dir, found, subdomain)
        elif subdomain:
            was_changed, was_ignored = update_subdomain(sdk_dir, found, subdomain)

        if was_ignored:
            ignored.append('subdomain')

        if spec_name:
            ignored.append('spec_name')
        if group_id:
            ignored.append('group_id')

        current = gather_settings(sdk_dir, found)
        changed = compute_changed_settings(previous, current)
    else:
        missing_params = []
        if not spec_name:
            missing_params.append('--spec-name')

        if not version:
            missing_params.append('--version')

        if not group_id:
            missing_params.append('--group-id')

        if not artifact_id:
            missing_params.append('--artifact-id')

        if not relative_spec_path:
            missing_params.append('--relative-spec-path')

        # not checking 'endpoint' anymore; can be specified either in ticket
        # or in spec using 'x-obmcs-endpoint-template'. If neither is specified,
        # this fails in the generator

        if missing_params:
            raise MissingParameter('The following options must be specified for a new spec: {}'.format(', '.join(missing_params)))

        if endpoint and subdomain:
            raise UsageError('Cannot specify both --endpoint and --subdomain')

        if endpoint:
            print('Ignoring endpoint setting for new services; new services have to get it from the spec')
            ignored.append('endpoint')

        if subdomain:
            print('Ignoring subdomain setting for new services; new services have to get it from the spec')
            ignored.append('subdomain')

        if spec_generation_type:
            print('Note: --spec-generation-type is ignored for the .NET SDK, since it is set in the codegen/pom.xml file for all services')

        service_name = spec_name.lower().replace('_', '')                       # service_name is "newservice"
        module_name = service_name.title()                                      # module_name is "Newservice"
        service_friendly_name = string.capwords(spec_name.replace('_', ' '))    # service_friendly_name is "New Service"

        # Find OCI DotNet SDK Version
        pom = parse_pom(pom_location)
        xpath = ".//ns:version"
        property = pom.findall(xpath, ns)[0]
        sdk_version = property.text

        guid = str(uuid.uuid4()).upper()
        print('Project GUID: {}'.format(guid))

        print('Module {} does not exist in pom.xml. Adding it...'.format(module_name))
        generate_child_codegen_pom(sdk_dir, service_name, module_name, service_friendly_name, sdk_version, group_id, artifact_id, version, relative_spec_path, spec_generation_type, generate_waiters, generate_paginators)
        add_child_codegen_module(sdk_dir, service_name)
        generate_child_project_and_readme(sdk_dir, module_name, service_friendly_name, guid)
        update_sln_file(sdk_dir, module_name, guid)
        add_spec_module_to_github_whitelist(module_name, github_whitelist_location, '\n^{}/.*')

    print('Success!')

    return AddOrUpdateSpecResult(
        updated=(not found) or changed != [],  # not found means it's a new spec, or if it is an existing spec, changed needs to be non-empty
        existing=found is not None,
        ignored=ignored,
        previous=previous,
        changed=changed)


def find_existing_module(sdk_dir, artifact_id):
    codegen_dir = os.path.join(sdk_dir, DOTNETSDK_CODEGEN_LOCATION)
    pom_files = [y for x in os.walk(codegen_dir) for y in glob(os.path.join(x[0], 'pom.xml'))]
    for ldr_path in pom_files:
        pom = parse_pom(ldr_path)
        xpath = ".//ns:properties//ns:codegen.artifactory.artifactId"
        properties = pom.findall(xpath, ns)
        if len(properties) > 0 and artifact_id == properties[0].text:
            codegen_artifact_id = pom.findall("./ns:artifactId", ns)[0].text
            m = re.match("oci-dotnet-sdk-([^-]*)-codegen", codegen_artifact_id)
            if m:
                return m.group(1)

    return None


def gather_settings(sdk_dir, spec_name):
    settings = {}

    file_name = os.path.join(sdk_dir, DOTNETSDK_CODEGEN_LOCATION, spec_name, "pom.xml")
    pom = parse_pom(file_name)

    xpath = ".//ns:properties//ns:codegen.artifactory.groupId"
    property = pom.findall(xpath, ns)[0]
    settings["group_id"] = property.text

    xpath = ".//ns:properties//ns:codegen.artifactory.artifactId"
    property = pom.findall(xpath, ns)[0]
    settings["artifact_id"] = property.text

    xpath = ".//ns:properties//ns:codegen.artifactory.version"
    property = pom.findall(xpath, ns)[0]
    settings["version"] = property.text

    settings["module_name"] = spec_name

    xpath = ".//ns:properties//ns:codegen.spec.name"
    property = pom.findall(xpath, ns)[0]
    settings["relative_spec_path"] = property.text

    xpath = ".//ns:properties//ns:codegen.endpoint"
    results = pom.findall(xpath, ns)
    if results:
        property = results[0]
        settings["subdomain"] = property.text

    return settings


def update_version_of_existing_spec(sdk_dir, spec_name, version):
    file_name = os.path.join(sdk_dir, DOTNETSDK_CODEGEN_LOCATION, spec_name, "pom.xml")
    pom = parse_pom(file_name)

    xpath = ".//ns:properties//ns:codegen.artifactory.version"
    property = pom.findall(xpath, ns)[0]
    old_version = property.text

    if not is_version_increasing(old_version, version):
        return old_version

    property.text = version
    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)
    return None  # the old version was lower


def update_relative_spec_path_of_existing_spec(sdk_dir, spec_name, relative_spec_path):
    file_name = os.path.join(sdk_dir, DOTNETSDK_CODEGEN_LOCATION, spec_name, "pom.xml")
    pom = parse_pom(file_name)

    xpath = ".//ns:properties//ns:codegen.spec.name"
    property = pom.findall(xpath, ns)[0]
    property.text = relative_spec_path
    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


# Returns was_changed, was_ignored
def update_subdomain(sdk_dir, spec_name, subdomain):
    file_name = os.path.join(sdk_dir, DOTNETSDK_CODEGEN_LOCATION, spec_name, "pom.xml")
    pom = parse_pom(file_name)

    xpath = ".//ns:properties//ns:codegen.endpoint"
    results = pom.findall(xpath, ns)
    if results:
        property = results[0]
        was_changed = property.text != subdomain
        property.text = subdomain
        indent(pom.getroot())
        pom.write(file_name, encoding="UTF-8", xml_declaration=True)
        return was_changed, False
    else:
        print('Ignored update to subdomain/endpoint -- service had no setting for this in pom.xml file and is getting that information from spec')
        return False, True


def generate_child_codegen_pom(sdk_dir, service_name, module_name, service_friendly_name, sdk_version, group_id, artifact_id, artifact_version, spec_path_relative_to_jar, spec_generation_type, generate_waiters, generate_paginators):
    dir_name = os.path.join(sdk_dir, DOTNETSDK_CODEGEN_LOCATION, service_name)
    os.mkdir(dir_name)
    file_name = os.path.join(dir_name, "pom.xml")

    content = DOTNETSDK_POM_FILE_TEMPLATE.format(
        group_id=group_id,
        artifact_id=artifact_id,
        artifact_version=artifact_version,
        sdk_version=sdk_version,
        service_name=service_name,
        service_friendly_name=service_friendly_name,
        module_name=module_name,
        spec_path_relative_to_jar=spec_path_relative_to_jar,
        spec_generation_type=spec_generation_type,
        generate_waiters=str(generate_waiters).lower(),
        generate_paginators=str(generate_paginators).lower())

    root = ET.fromstring(content, parser=ET.XMLParser(target=CommentedTreeBuilder()))
    pom = ET.ElementTree(element=root)
    write_xml(file_name, pom)


def add_child_codegen_module(sdk_dir, service_name):
    file_name = os.path.join(sdk_dir, DOTNETSDK_CODEGEN_LOCATION, "pom.xml")
    pom = parse_pom(file_name)

    content = MODULE_TEMPLATE.format(name=service_name)
    module_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))

    # find modules
    modules = pom.findall("./ns:modules", ns)[0]
    modules.append(module_element)

    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


def generate_child_project_and_readme(sdk_dir, module_name, service_friendly_name, guid):
    dir_name = os.path.join(sdk_dir, module_name)
    os.mkdir(dir_name)
    file_name = os.path.join(dir_name, "OCI.DotNetSDK." + module_name + ".csproj")
    readme = os.path.join(dir_name, "README.md")

    content = MODULE_PROJECT_FILE_TEMPLATE.format(
        guid=guid,
        module_name=module_name,
        service_friendly_name=service_friendly_name)

    readme_content = MODULE_README_FORMAT.format(service_friendly_name=service_friendly_name,
                                                 module_name=module_name)

    project_file = open(file_name, "w")
    project_file.write(content)
    readme_file = open(readme, "w")
    readme_file.write(readme_content)


def update_sln_file(sdk_dir, module_name, guid):
    project_added = False
    project_platform_added = False

    sln_file = os.path.join(sdk_dir, "oci-dotnet-sdk.sln")
    with open(sln_file) as sln_file_read:
        lines = sln_file_read.readlines()

    locs = [i for i, val in enumerate(lines) if 'EndProject' in val]

    # Ignoring case when a solution file does not contain any existing projects.This should never happen since we already have projects.
    if len(locs) > 1:
        print("Will add project")
        new_project_entry = PROJECT_ENTRY_TEMPLATE.format(guid=guid, module_name=module_name)
        lines.insert(locs[-1] + 1, new_project_entry)
        project_added = True

        project_platform_section_found = False
        # First find the start of ProjectConfigurationPlatforms section
        for loc, line in enumerate(lines):
            if 'ProjectConfigurationPlatforms' in line:
                project_platform_section_found = True
                continue
            # Find the matching closing line for ProjectConfigurationPlatforms section
            if project_platform_section_found and 'EndGlobalSection' in line:
                print("Will update plateform")
                new_project_platform = PROJECT_PLATEFORM_TEMPLATE.format(guid=guid)
                lines.insert(loc, new_project_platform)
                project_platform_added = True
                break
        if not project_platform_section_found:
            print("Cannot find project platform section.")
    else:
        print("Cannot find existing projects.")

    if project_added and project_platform_added:
        with open(sln_file, 'w') as sln_file_write:
            sln_file_write.write(''.join(lines))
    else:
        print("Unable to update sln file.")


@click.command()
@click.option('--artifact-id', required=True, help='The artifact id for the spec artifact (e.g. coreservices-api-spec')
@click.option('--group-id', help='The group id for the spec artifact (e.g. com.oracle.pic.commons)')
@click.option('--spec-name', help='The name of the spec. This will be used (e.g. core, identity, objectstorage). This is also used as the module name (\'bmc-servicename\') and base package (\'com.oracle.bmc.servicename\'). Underscores are removed, everything is lower-cased.')
@click.option('--relative-spec-path', help='The relative path of the spec within the artifact (e.g. coreservices-api-spec-20160918-external.yaml)')
@click.option('--endpoint', help='The base endpoint for the service (e.g. https://iaas.{domain}/20160918)')
@click.option('--subdomain', help='The subdomain for the service (e.g. if the endpoint is https://iaas.{domain}/20160918), the subdomain is "iaas"')
@click.option('--version', help='The version of the spec artifact (e.g. 0.0.1-SNAPSHOT')
@click.option('--spec-generation-type', help='The generation type: PUBLIC or PREVIEW')
@click.option('--generate-waiters/--no-generate-waiters', default=True, help='Generate waiters')
@click.option('--generate-paginators/--no-generate-paginators', default=True, help='Generate paginators')
@click.option('--pom-location', type=click.Path(exists=True), default=DEFAULT_PARENT_POM_LOCATION, help='Location of the pom.xml file in the root directory of the OCI .NET SDK')
@click.option('--github-whitelist-location', type=click.Path(exists=True), default=DEFAULT_GITHUB_WHITELIST_LOCATION, help='Location of the github.whitelist file to update')
def add_or_update_spec_command(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version, spec_generation_type, generate_waiters, generate_paginators, pom_location, github_whitelist_location):
    print(add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version, spec_generation_type, generate_waiters, generate_paginators, pom_location, github_whitelist_location))


if __name__ == '__main__':
    add_or_update_spec_command()
