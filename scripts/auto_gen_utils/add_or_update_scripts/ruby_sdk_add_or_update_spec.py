#!/usr/bin/env python

#
# This script manipulates the pom.xml tree to either add new specs or update the versions of existing specs.
#

import os
import click

from datetime import datetime
from .module_pom_file_add_or_update_spec import DEFAULT_POM_LOCATION
from .spec_updater_base import SpecUpdaterBase


# The path for the modules is under "ruby-sdk/codegen"
RUBY_MODULE_LOCATION = "codegen"

# The service pom.xml template for the ruby sdk
RUBY_POM_FILE_TEMPLATE = """<?xml version='1.0' encoding='UTF-8'?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <parent>
    <groupId>com.oci.sdk</groupId>
    <artifactId>ruby-sdk-codegen-template</artifactId>
    <version>{sdk_version}</version>
    <relativePath>..</relativePath>
  </parent>
  <artifactId>ruby-sdk-{module_name}-codegen</artifactId>
  <name>OCI Ruby SDK - {service_friendly_name} Service Codegen</name>
  <packaging>pom</packaging>
  <properties>
    <codegen.artifactory.groupId>{group_id}</codegen.artifactory.groupId>
    <codegen.artifactory.artifactId>{artifact_id}</codegen.artifactory.artifactId>
    <codegen.artifactory.version>{artifact_version}</codegen.artifactory.version>
    <codegen.artifact.spec.path>{spec_path_relative_to_jar}</codegen.artifact.spec.path>
    <codegen.service.name>{module_name}</codegen.service.name>
    <codegen.endpoint.prefix>{subdomain}</codegen.endpoint.prefix>
    <codegen.generate.waiters>true</codegen.generate.waiters>
    <codegen.generate.paginators>true</codegen.generate.paginators>
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
        <artifactId>exec-maven-plugin</artifactId>
        <groupId>org.codehaus.mojo</groupId>
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
RUBY_SPEC_PARAMS_XML_PATH_DICT = {
    'group_id': ".//ns:properties//ns:codegen.artifactory.groupId",
    'artifact_id': ".//ns:properties//ns:codegen.artifactory.artifactId",
    'version':  ".//ns:properties//ns:codegen.artifactory.version",
    'relative_spec_path': ".//ns:properties//ns:codegen.artifact.spec.path",
    'service_name': ".//ns:properties//ns:codegen.service.name",
    'subdomain': ".//ns:properties//ns:codegen.endpoint.prefix"
}


class RubySpecUpdater(SpecUpdaterBase):
    # Required for new services to reference from the generated composite operations
    # Generates an empty util.rb file as ruby-sdk/lib/oci/module_name/util.rb
    def generate_util_rb(self, sdk_dir, module_name):
        gen_path = os.path.join(sdk_dir, 'lib', 'oci', module_name)
        util_file_path = os.path.join(gen_path, 'util.rb')

        if os.path.isfile(util_file_path):
            print("{} already exists.  Not overwriting".format(util_file_path))
            return

        if not os.path.exists(gen_path):
            os.makedirs(gen_path)

        with open(util_file_path, 'w') as util_file:
            util_contents = "# Copyright (c) 2016, {}, Oracle and/or its affiliates. All rights reserved.\n\n" \
                .format(datetime.now().year)
            util_file.write(util_contents)
        print("Created {}".format(util_file_path))

    # Override add_spec to generate the util.rb file that is required for the ruby sdk
    def add_spec(self,
                 sdk_dir,
                 module_location,
                 spec_name,
                 group_id,
                 artifact_id,
                 version,
                 relative_spec_path,
                 subdomain):
        result = SpecUpdaterBase.add_spec(self,
                                          sdk_dir,
                                          module_location,
                                          spec_name,
                                          group_id,
                                          artifact_id,
                                          version,
                                          relative_spec_path,
                                          subdomain)
        self.generate_util_rb(sdk_dir, spec_name)
        return result


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
    spec_updater = RubySpecUpdater(RUBY_MODULE_LOCATION, RUBY_POM_FILE_TEMPLATE, RUBY_SPEC_PARAMS_XML_PATH_DICT)
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
