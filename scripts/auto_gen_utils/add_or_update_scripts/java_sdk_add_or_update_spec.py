#!/usr/bin/env python

#
# This script manipulates pom.xml to either add new specs or update the versions of existing specs.
#

import datetime
import xml.etree.ElementTree as ET
import re
import click
import os
import parse
import string
from click.exceptions import UsageError, MissingParameter
from glob import glob

from shared.version_utils import is_version_increasing
from .add_or_update_spec_utils import AddOrUpdateSpecResult, compute_changed_settings, indent
from .add_or_update_spec_utils import add_spec_module_to_github_whitelist, write_xml
from .add_or_update_spec_utils import CommentedTreeBuilder
from .add_or_update_spec_utils import parse_pom


DEFAULT_PARENT_POM_LOCATION = "pom.xml"
DEFAULT_GITHUB_WHITELIST_LOCATION = "github.whitelist"

CODEGEN_POM_FILE_ARTIFACTORY_VERSION_COMMENT_TEMPLATE = "DEXREQ ticket requested version '{}', but version was already '{}' ('{}')."

CODEGEN_POM_FILE_TEMPLATE = """
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <parent>
    <groupId>com.oracle.oci.sdk</groupId>
    <artifactId>oci-java-sdk-codegen</artifactId>
    <version>{sdk_version}</version>
    <relativePath>../pom.xml</relativePath>
  </parent>

  <artifactId>oci-java-sdk-{module_name}-codegen</artifactId>
  <name>Oracle Cloud Infrastructure SDK - {service_friendly_name} Codegen</name>
  <description>This project contains the code generation spec and configuration for the {service_friendly_name} Service</description>
  <url>https://docs.cloud.oracle.com/Content/API/SDKDocs/javasdk.htm</url>

  <properties>
    <codegen.artifactory.groupId>{group_id}</codegen.artifactory.groupId>
    <codegen.artifactory.artifactId>{artifact_id}</codegen.artifactory.artifactId>
    <codegen.artifactory.version>{artifact_version}</codegen.artifactory.version>
    <codegen.spec.name>{spec_path_relative_to_jar}</codegen.spec.name>

    <codegen.package.name>{module_name}</codegen.package.name>
    <codegen.generate.waiters>{generate_waiters}</codegen.generate.waiters>
    <codegen.generate.paginators>{generate_paginators}</codegen.generate.paginators>
  </properties>

  <profiles>
    <profile>
      <id>codegen</id>
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
            <groupId>com.theoryinpractise</groupId>
            <artifactId>googleformatter-maven-plugin</artifactId>
          </plugin>
          <plugin>
            <artifactId>maven-antrun-plugin</artifactId>
          </plugin>
        </plugins>
      </build>
    </profile>
    <profile>
      <id>specs</id>
      <build>
        <plugins>
          <plugin>
            <groupId>org.apache.maven.plugins</groupId>
            <artifactId>maven-assembly-plugin</artifactId>
            <configuration>
              <skipAssembly>true</skipAssembly>
            </configuration>
          </plugin>
        </plugins>
      </build>
    </profile>
  </profiles>

</project>
"""

MODULE_POM_FILE_TEMPLATE = """
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <parent>
    <groupId>com.oracle.oci.sdk</groupId>
    <artifactId>oci-java-sdk</artifactId>
    <version>{sdk_version}</version>
    <relativePath>../pom.xml</relativePath>
  </parent>

  <artifactId>oci-java-sdk-{module_name}</artifactId>
  <name>Oracle Cloud Infrastructure SDK - {service_friendly_name}</name>
  <description>This project contains the SDK used for Oracle Cloud Infrastructure {service_friendly_name}</description>
  <url>https://docs.cloud.oracle.com/Content/API/SDKDocs/javasdk.htm</url>

  <!-- START: Delete Public GitHub -->
  <profiles>
    <profile>
      <id>sign-individual</id>
      <build>
        <plugins>
          <plugin>
            <groupId>com.oracle.ccss</groupId>
            <artifactId>codesign-maven-plugin</artifactId>
          </plugin>
        </plugins>
      </build>
    </profile>
  </profiles>
  <build>
    <plugins>
      <plugin>
        <groupId>org.codehaus.mojo</groupId>
        <artifactId>clirr-maven-plugin</artifactId>
        <version>2.9.1-oracle-SNAPSHOT</version>
        <executions>
          <execution>
            <id>clirr</id>
            <phase>verify</phase>
            <goals>
              <goal>check</goal>
            </goals>
            <configuration>
              <failOnError>${{clirr.fail.on.error}}</failOnError>
              <textOutputFile>${{project.build.directory}}/clirr-{module_name}.txt</textOutputFile>
              <logResults>true</logResults>
              <excludeVersionRegexes>
                <param>.*preview.*</param>
                <param>1.23.4.*stream.*</param>
                <param>.*experimental.*</param>
              </excludeVersionRegexes>
              <comparisonVersion>${{clirr.comparison.version}}</comparisonVersion>
              <ignored>
                <!-- Ignored differences by design -->
                <difference>
                  <!-- Ignore 'internal' packages -->
                  <className>**/internal/**/*</className>
                  <field>*</field>
                  <method>*</method>
                </difference>
                <difference>
                  <!-- Ignore 'Async' and 'Client' classes, just look at the sync interface -->
                  <className>com/oracle/bmc/{module_name}/*Async*</className>
                  <field>*</field>
                  <method>*</method>
                </difference>
                <difference>
                  <!-- Ignore 'Async' and 'Client' classes, just look at the sync interface -->
                  <className>com/oracle/bmc/{module_name}/*Client</className>
                  <field>*</field>
                  <method>*</method>
                </difference>
                <difference>
                  <!-- Ignore added methods to clients -->
                  <className>com/oracle/bmc/{module_name}/*</className>
                  <field>*</field>
                  <method>*</method>
                  <differenceType>7012</differenceType>
                </difference>
                <difference>
                  <!-- Ignore added parameters in models, requests, responses -->
                  <className>com/oracle/bmc/{module_name}/*/*</className>
                  <field>*</field>
                  <method>*</method>
                  <differenceType>7004</differenceType>
                </difference>
                <difference>
                  <!-- Ignore parameter type changes in models, requests, responses -->
                  <!-- These will be caught by return type changes. -->
                  <className>com/oracle/bmc/{module_name}/*/*</className>
                  <field>*</field>
                  <method>*</method>
                  <differenceType>7005</differenceType>
                  <to>*</to>
                </difference>
                <difference>
                  <!-- Ignore builders. -->
                  <className>com/oracle/bmc/{module_name}/*/*Builder</className>
                  <field>*</field>
                  <method>*</method>
                </difference>
                <difference>
                  <!-- Ignore canEqual method. -->
                  <className>com/oracle/bmc/{module_name}/*/*</className>
                  <method>boolean canEqual(java.lang.Object)</method>
                  <differenceType>7002</differenceType>
                </difference>
                <!-- Add backward-incompatible differences below -->
              </ignored>
            </configuration>
          </execution>
        </executions>
      </plugin>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-checkstyle-plugin</artifactId>
        <version>3.1.1</version>
        <configuration>
          <configLocation>${{checkstyle-rules-file}}</configLocation>
          <encoding>UTF-8</encoding>
          <consoleOutput>true</consoleOutput>
          <failsOnError>true</failsOnError>
          <linkXRef>false</linkXRef>
        </configuration>
        <executions>
          <execution>
            <phase>verify</phase>
            <goals>
              <goal>check</goal>
            </goals>
          </execution>
        </executions>
      </plugin>
    </plugins>
  </build>
  <!-- END: Delete Public GitHub -->

  <dependencies>
    <dependency>
      <groupId>com.oracle.oci.sdk</groupId>
      <artifactId>oci-java-sdk-common</artifactId>
      <version>{sdk_version}</version>
    </dependency>
  </dependencies>

</project>
"""

MODULE_TEMPLATE = "<module>{name}</module>"

BOM_DEPENDENCY_TEMPLATE = """
<dependency>
  <groupId>com.oracle.oci.sdk</groupId>
  <artifactId>oci-java-sdk-{module_name}</artifactId>
  <version>{sdk_version}</version>
  <optional>false</optional>
</dependency>
"""

FULL_JAVADOC_SOURCEPATH_TEMPLATE = os.linesep + "            ../bmc-{module_name}/src/main/java;"

FULL_JAVADOC_GROUP_TEMPLATE = """
<group>
  <title>Oracle Cloud Infrastructure {service_friendly_name}</title>
  <packages>com.oracle.bmc.{module_name}*</packages>
</group>
"""

FULL_DEPENDENCY_TEMPLATE = """
    <dependency>
      <groupId>com.oracle.oci.sdk</groupId>
      <artifactId>oci-java-sdk-{module_name}</artifactId>
    </dependency>
"""

SHADED_DEPENDENCY_TEMPLATE = """
    <dependency>
      <groupId>com.oracle.oci.sdk</groupId>
      <artifactId>oci-java-sdk-{module_name}</artifactId>
      <scope>provided</scope>
    </dependency>
"""

INTEGTESTS_DEPENDENCY_TEMPLATE = """
    <dependency>
      <groupId>com.oracle.oci.sdk</groupId>
      <artifactId>oci-java-sdk-{module_name}</artifactId>{scope_text}
    </dependency>
"""

ns = {"ns":"http://maven.apache.org/POM/4.0.0"}

# allow default namespace for output, dont print ns0: prefixes everywhere
ET.register_namespace('',"http://maven.apache.org/POM/4.0.0")


def generate_child_codegen_pom(sdk_dir, module_name, service_friendly_name, sdk_version, group_id, artifact_id, artifact_version, spec_path_relative_to_jar, spec_generation_type, generate_waiters, generate_paginators):
    dir_name = os.path.join(sdk_dir, "bmc-codegen", "bmc-" + module_name + "-codegen")
    os.mkdir(dir_name)
    file_name = os.path.join(dir_name, "pom.xml")

    content = CODEGEN_POM_FILE_TEMPLATE.format(
        group_id=group_id,
        artifact_id=artifact_id,
        artifact_version=artifact_version,
        sdk_version=sdk_version,
        service_friendly_name=service_friendly_name,
        module_name=module_name,
        spec_path_relative_to_jar=spec_path_relative_to_jar,
        spec_generation_type=spec_generation_type,
        generate_waiters=str(generate_waiters).lower(),
        generate_paginators=str(generate_paginators).lower())

    root = ET.fromstring(content, parser=ET.XMLParser(target=CommentedTreeBuilder()))
    pom = ET.ElementTree(element=root)
    write_xml(file_name, pom)


def add_child_codegen_module(sdk_dir, module_name):
    file_name = os.path.join(sdk_dir, "bmc-codegen", "pom.xml")
    pom = parse_pom(file_name)

    content = MODULE_TEMPLATE.format(name="bmc-" + module_name + "-codegen")
    module_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))

    # find modules
    modules = pom.findall("./ns:modules", ns)[0]
    modules.append(module_element)

    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


def generate_child_module(sdk_dir, module_name, service_friendly_name, sdk_version):
    dir_name = os.path.join(sdk_dir, "bmc-" + module_name)
    os.mkdir(dir_name)
    file_name = os.path.join(dir_name, "pom.xml")

    content = MODULE_POM_FILE_TEMPLATE.format(
        sdk_version=sdk_version,
        service_friendly_name=service_friendly_name,
        module_name=module_name)

    root = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))
    pom = ET.ElementTree(element=root)
    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)

    os.makedirs(os.path.join(dir_name, "src", "main", "java"))


def add_child_module(parent_pom_location, module_name):
    pom = parse_pom(parent_pom_location)

    content = MODULE_TEMPLATE.format(name="bmc-" + module_name)
    module_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))

    # find <modules>
    modules = pom.findall("./ns:modules", ns)[0]

    # we want to insert this just before the first "START: Delete Public GitHub"
    index_to_insert = 0
    for index, child in enumerate(list(modules)):
        index_to_insert = index
        if "START: Delete Public GitHub" in child.text:
            break

    modules.insert(index_to_insert, module_element)

    indent(pom.getroot())
    pom.write(parent_pom_location, encoding="UTF-8", xml_declaration=True)


def add_bom_dependency(sdk_dir, module_name, sdk_version):
    file_name = os.path.join(sdk_dir, "bmc-bom", "pom.xml")
    pom = parse_pom(file_name)

    content = BOM_DEPENDENCY_TEMPLATE.format(module_name=module_name, sdk_version=sdk_version)
    dependency_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))

    # find dependencies
    dependencies = pom.findall("./ns:dependencyManagement/ns:dependencies", ns)[0]
    dependencies.append(dependency_element)

    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


def add_full_dependency(sdk_dir, module_name, service_friendly_name, sdk_version):
    file_name = os.path.join(sdk_dir, "bmc-full", "pom.xml")
    pom = parse_pom(file_name)

    content = FULL_DEPENDENCY_TEMPLATE.format(module_name=module_name)
    dependency_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))

    # find dependencies
    dependencies = pom.findall("./ns:dependencies", ns)[0]
    dependencies.append(dependency_element)

    # find Javadoc sourcepath
    sourcepath_node = pom.findall("./ns:build/ns:plugins/ns:plugin[ns:artifactId='maven-javadoc-plugin']/ns:configuration/ns:sourcepath", ns)[0]
    sourcepath_node.text = sourcepath_node.text + FULL_JAVADOC_SOURCEPATH_TEMPLATE.format(module_name=module_name)

    # find Javadoc groups
    groups = pom.findall("./ns:build/ns:plugins/ns:plugin[ns:artifactId='maven-javadoc-plugin']/ns:configuration/ns:groups", ns)[0]

    content = FULL_JAVADOC_GROUP_TEMPLATE.format(module_name=module_name, service_friendly_name=service_friendly_name)
    group_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))
    groups.append(group_element)

    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


def add_dependency_to_other_pom_file(file_name, module_name, scope):
    pom = parse_pom(file_name)

    content = INTEGTESTS_DEPENDENCY_TEMPLATE.format(
        module_name=module_name,
        scope_text="\n      <scope>{}</scope>".format(scope) if scope else "")
    dependency_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))

    # find dependencies
    dependencies = pom.findall("./ns:dependencies", ns)[0]
    dependencies.append(dependency_element)

    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


def update_version_of_existing_spec(sdk_dir, spec_name, version):
    file_name = os.path.join(sdk_dir, "bmc-codegen", "bmc-" + spec_name + "-codegen", "pom.xml")
    pom = parse_pom(file_name)

    xpath = ".//ns:properties//ns:codegen.artifactory.version"
    property = pom.findall(xpath, ns)[0]
    old_version = property.text

    version_increased = is_version_increasing(old_version, version)
    if version_increased:
        property.text = version

    # remove old comments and find the place to insert the version comment, if necessary
    parent = pom.findall(".//ns:properties", ns)[0]
    comments_to_remove = []
    comment_insertion_index = 0
    for index, el in enumerate(parent.iter()):
        if el.tag is ET.Comment and el.text:
            result = parse.search(CODEGEN_POM_FILE_ARTIFACTORY_VERSION_COMMENT_TEMPLATE, el.text)
            if result:
                # matched the formats
                comments_to_remove.append(el)
        else:
            if el.tag.endswith("codegen.artifactory.version"):
                if index > 0:
                    comment_insertion_index = index - len(comments_to_remove) - 1

    for el in comments_to_remove:
        parent.remove(el)

    if not version_increased:
        # insert version comment
        comment = ET.Comment(CODEGEN_POM_FILE_ARTIFACTORY_VERSION_COMMENT_TEMPLATE.format(version, old_version, datetime.datetime.now()))
        parent.insert(comment_insertion_index, comment)

    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)
    
    if not version_increased:
        return old_version

    return None  # the old version was lower


def update_relative_spec_path_of_existing_spec(sdk_dir, spec_name, relative_spec_path):
    file_name = os.path.join(sdk_dir, "bmc-codegen", "bmc-" + spec_name + "-codegen", "pom.xml")
    pom = parse_pom(file_name)

    xpath = ".//ns:properties//ns:codegen.spec.name"
    property = pom.findall(xpath, ns)[0]
    property.text = relative_spec_path
    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


# Returns was_changed, was_ignored
def update_subdomain(sdk_dir, spec_name, subdomain):
    file_name = os.path.join(sdk_dir, "bmc-codegen", "bmc-" + spec_name + "-codegen", "pom.xml")
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


def add_dependency_for_shaded_fatjar(file_name, module_name, template):
    pom = parse_pom(file_name)

    content = template.format(module_name=module_name)
    dependency_element = ET.fromstring(content, ET.XMLParser(target=CommentedTreeBuilder()))

    # find <modules>
    dependencies = pom.findall("./ns:dependencies", ns)[0]

    # we want to insert this just before the first "START: Delete Public GitHub"
    index_to_insert = 0
    for index, child in enumerate(list(dependencies)):
        index_to_insert = index
        if "END OF OCI SDK DEPENDENCIES" in child.text:
            break

    dependencies.insert(index_to_insert, dependency_element)

    indent(pom.getroot())
    pom.write(file_name, encoding="UTF-8", xml_declaration=True)


def find_existing_module(sdk_dir, artifact_id):
    codegen_dir = os.path.join(sdk_dir, "bmc-codegen")
    pom_files = [y for x in os.walk(codegen_dir) for y in glob(os.path.join(x[0], 'pom.xml'))]
    for ldr_path in pom_files:
        pom = parse_pom(ldr_path)
        xpath = ".//ns:properties//ns:codegen.artifactory.artifactId"
        properties = pom.findall(xpath, ns)
        if len(properties) > 0 and artifact_id == properties[0].text:
            codegen_artifact_id = pom.findall("./ns:artifactId", ns)[0].text
            m = re.match("oci-java-sdk-([^-]*)-codegen", codegen_artifact_id)
            if m:
                return m.group(1)

    return None


def gather_settings(sdk_dir, spec_name):
    settings = {}

    file_name = os.path.join(sdk_dir, "bmc-codegen", "bmc-" + spec_name + "-codegen", "pom.xml")
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
            print('Note: --spec-generation-type is ignored for the Java SDK, since it is set in the bmc-codegen/pom.xml file for all modules')

        module_name = spec_name.lower().replace('_', '')  # module_name is "newservice"
        service_friendly_name = string.capwords(spec_name.replace('_', ' '))  # service_friendly_name is "New Service"

        # Find OCI Java SDK Version
        pom = parse_pom(pom_location)
        xpath = ".//ns:version"
        property = pom.findall(xpath, ns)[0]
        sdk_version = property.text

        print('Module {} does not exist in pom.xml. Adding it...'.format(module_name))
        generate_child_codegen_pom(sdk_dir, module_name, service_friendly_name, sdk_version, group_id, artifact_id, version, relative_spec_path, spec_generation_type, generate_waiters, generate_paginators)
        add_child_codegen_module(sdk_dir, module_name)
        generate_child_module(sdk_dir, module_name, service_friendly_name, sdk_version)
        add_child_module(pom_location, module_name)
        add_bom_dependency(sdk_dir, module_name, sdk_version)
        add_full_dependency(sdk_dir, module_name, service_friendly_name, sdk_version)
        add_dependency_to_other_pom_file(os.path.join(sdk_dir, "bmc-integtests", "pom.xml"), module_name, "test")
        add_dependency_to_other_pom_file(os.path.join(sdk_dir, "bmc-smoketests", "pom.xml"), module_name, "test")
        add_dependency_to_other_pom_file(os.path.join(sdk_dir, "bmc-examples", "pom.xml"), module_name, None)
        add_dependency_for_shaded_fatjar(os.path.join(sdk_dir, "bmc-shaded", "bmc-shaded-internal-fatjar", "pom.xml"), module_name, FULL_DEPENDENCY_TEMPLATE)
        add_dependency_for_shaded_fatjar(os.path.join(sdk_dir, "bmc-shaded", "bmc-shaded-full", "pom.xml"), module_name, SHADED_DEPENDENCY_TEMPLATE)
        add_spec_module_to_github_whitelist(module_name, github_whitelist_location, '\n^bmc-{}/.*')

    print('Success!')

    return AddOrUpdateSpecResult(
        updated=(not found) or changed != [],  # not found means it's a new spec, or if it is an existing spec, changed needs to be non-empty
        existing=found is not None,
        ignored=ignored,
        previous=previous,
        changed=changed)


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
@click.option('--pom-location', type=click.Path(exists=True), default=DEFAULT_PARENT_POM_LOCATION, help='Location of the pom.xml file in the root directory of the OCI Java SDK')
@click.option('--github-whitelist-location', type=click.Path(exists=True), default=DEFAULT_GITHUB_WHITELIST_LOCATION, help='Location of the github.whitelist file to update')
def add_or_update_spec_command(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version, spec_generation_type, generate_waiters, generate_paginators, pom_location, github_whitelist_location):
    print(add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, endpoint, subdomain, version, spec_generation_type, generate_waiters, generate_paginators, pom_location, github_whitelist_location))


if __name__ == '__main__':
    add_or_update_spec_command()
