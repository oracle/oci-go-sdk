#!/usr/bin/env python

#
# This script manipulates the pom.xml tree to either add new specs or update the versions of existing specs.
#

import click

from .module_pom_file_add_or_update_spec import DEFAULT_POM_LOCATION
from .spec_updater_base import SpecUpdaterBase


# The path for the modules is under "sdk-client-test-data/codegen"
TEST_DATA_GEN_MODULE_LOCATION = "codegen"

# Pom.xml template specific to sdk-client-test-data
TEST_DATA_GEN_POM_FILE_TEMPLATE = """<?xml version='1.0' encoding='UTF-8'?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <parent>
    <groupId>com.oci.sdk</groupId>
    <artifactId>sdk-test-data-codegen-template</artifactId>
    <version>{sdk_version}</version>
    <relativePath>..</relativePath>
  </parent>
  <artifactId>sdk-test-data-{module_name}-codegen</artifactId>
  <name>OCI Generated Test Data - {service_friendly_name} Service Codegen</name>
  <packaging>pom</packaging>
  <properties>
    <codegen.artifactory.groupId>{group_id}</codegen.artifactory.groupId>
    <codegen.artifactory.artifactId>{artifact_id}</codegen.artifactory.artifactId>
    <codegen.artifactory.version>{artifact_version}</codegen.artifactory.version>
    <codegen.artifact.spec.path>{spec_path_relative_to_jar}</codegen.artifact.spec.path>
    <codegen.service.name>{module_name}</codegen.service.name>
    <codegen.service.group.endpoint>{subdomain}</codegen.service.group.endpoint>
  </properties>
  <build>
    <plugins>
      <plugin>
        <groupId>org.codehaus.mojo</groupId>
        <artifactId>build-helper-maven-plugin</artifactId>
      </plugin>
      <plugin>
        <groupId>com.oracle.oci.sdk.utilities</groupId>
        <artifactId>dex-get-spec-artifact-plugin</artifactId>
        <version>${{oci.get.spec.artifact.plugin.version}}</version>
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
        <groupId>com.oracle.oci</groupId>
        <artifactId>sdk-client-test-data-generator-maven-plugin</artifactId>
      </plugin>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-antrun-plugin</artifactId>
      </plugin>
      <plugin>
        <artifactId>maven-clean-plugin</artifactId>
      </plugin>
    </plugins>
  </build>
</project>
"""


# The ruby-specific pom.xml template parameters for XML parsing
TEST_DATA_GEN_SPEC_PARAMS_XML_PATH_DICT = {
    'group_id': ".//ns:properties//ns:codegen.artifactory.groupId",
    'artifact_id': ".//ns:properties//ns:codegen.artifactory.artifactId",
    'version':  ".//ns:properties//ns:codegen.artifactory.version",
    'relative_spec_path': ".//ns:properties//ns:codegen.artifact.spec.path",
    'service_name': ".//ns:properties//ns:codegen.service.name",
    'subdomain': ".//ns:properties//ns:codegen.service.group.endpoint"
}


class TestDataGenSpecUpdater(SpecUpdaterBase):
    # Override the spec name that is defined in the service pom.xml files as the testing service
    # references folders without snake-case (based on the java sdk formatting).
    def format_module_name_for_template(self, module_name):
        return module_name.replace("-", "").replace("_", "").replace(" ","")


##################################################
# Main
##################################################
@click.command()
@click.option('--artifact-id', required=True, help='The artifact id for the spec artifact (e.g. coreservices-api-spec')
@click.option('--group-id', help='The group id for the spec artifact (e.g. com.oracle.pic.commons)')
@click.option('--spec-name', help='The name of the spec. This will be used (e.g. core, identity, object_storage). '
                                  'This is also used as the module name.')
@click.option('--relative-spec-path', help='The relative path of the spec within the artifact '
                                           '(e.g. coreservices-api-spec-20160918-external.yaml)')
@click.option('--endpoint', help='The base endpoint for the service (e.g. https://iaas.{domain}/20160918)')
@click.option('--subdomain', help='The subdomain for the service (e.g. \'iaas\')')
@click.option('--version', help='The version of the spec artifact (e.g. 0.0.1-SNAPSHOT')
@click.option('--spec-generation-type', help='The generation type: PUBLIC or PREVIEW (is ignored for Ruby)')
@click.option('--regional-sub-service-overrides', multiple=True, help="Is ignored for Ruby SDK.")
@click.option('--non-regional-sub-service-overrides',
              multiple=True,
              help="Is ignored for Ruby SDK. Non-regional client overrides require manual update to the service module's pom.xml")
@click.option('--signing-strategy',
              help='The signing strategy to use for the client.  Is ignored for Ruby. Requires manual pom.xml update to override')
@click.option('--pom-location',
              type=click.Path(exists=True),
              default=DEFAULT_POM_LOCATION,
              help='Location of the root pom.xml file for the Ruby SDK')
@click.option('--module-location',
              type=click.Path(exists=True),
              help="Parent directory containing the module pom.xml files")
def add_or_update_command(artifact_id,
                          group_id,
                          spec_name,
                          relative_spec_path,
                          endpoint,
                          subdomain,
                          signing_strategy,
                          version,
                          spec_generation_type,
                          regional_sub_service_overrides,
                          non_regional_sub_service_overrides,
                          pom_location,
                          module_location):
    spec_updater = TestDataGenSpecUpdater(TEST_DATA_GEN_MODULE_LOCATION,
                                          TEST_DATA_GEN_POM_FILE_TEMPLATE,
                                          TEST_DATA_GEN_SPEC_PARAMS_XML_PATH_DICT)
    print(spec_updater.add_or_update_spec(
        artifact_id=artifact_id,
        group_id=group_id,
        spec_name=spec_name,
        relative_spec_path=relative_spec_path,
        endpoint=endpoint,
        subdomain=subdomain,
        version=version,
        pom_location=pom_location,
        module_location=module_location))


if __name__ == '__main__':
    add_or_update_command()
