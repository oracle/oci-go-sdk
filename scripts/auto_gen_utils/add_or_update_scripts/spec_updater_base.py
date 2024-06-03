import os
import re
import string
import xml.etree.ElementTree as ET

from click.exceptions import MissingParameter
from shared.version_utils import is_version_increasing
from .add_or_update_spec_utils import CommentedTreeBuilder, ns, indent, write_xml, parse_pom, find_pom_version
from .add_or_update_spec_utils import AddOrUpdateSpecResult, compute_changed_settings, convert_camel_to_snake_case
from .java_sdk_add_or_update_spec import MODULE_TEMPLATE

# Example dictionary for defining the paths in the XML for the templates:
# SPEC_PARAMS_XML_PATH_DICT = {
#     'group_id': ".//ns:properties//ns:codegen.artifactory.groupId",
#     'artifact_id': ".//ns:properties//ns:codegen.artifactory.artifactId",
#     'version':  ".//ns:properties//ns:codegen.artifactory.version",
#     'relative_spec_path': ".//ns:properties//ns:codegen.artifact.spec.path",
#     'service_name': ".//ns:properties//ns:codegen.service.name",
#     'subdomain': ".//ns:properties//ns:codegen.endpoint.prefix"
# }


class SpecUpdaterBase:
    ET.register_namespace('', "http://maven.apache.org/POM/4.0.0")

    def __init__(self, pom_module_location, service_pom_file_template, spec_params_xml_path_dict):
        self.pom_module_location = pom_module_location
        self.service_pom_file_template = service_pom_file_template
        self.spec_params_xml_path_dict = spec_params_xml_path_dict

    # Creates a new pom file as ruby-sdk/codegen/module_name/pom.xml
    def generate_child_codegen_pom(self,
                                   service_pom_path,
                                   module_name,
                                   service_friendly_name,
                                   sdk_version,
                                   group_id,
                                   artifact_id,
                                   artifact_version,
                                   spec_path_relative_to_jar,
                                   subdomain):
        if not subdomain:
            subdomain = "None"

        content = self.service_pom_file_template.format(
            group_id=group_id,
            artifact_id=artifact_id,
            artifact_version=artifact_version,
            sdk_version=sdk_version,
            service_friendly_name=service_friendly_name,
            module_name=self.format_module_name_for_template(module_name),
            spec_path_relative_to_jar=spec_path_relative_to_jar,
            subdomain=subdomain)

        root = ET.fromstring(content, parser=ET.XMLParser(target=CommentedTreeBuilder()))
        pom = ET.ElementTree(element=root)
        write_xml(service_pom_path, pom)

    def format_module_name_for_template(self, module_name):
        return module_name

    # Updates the submodule reference under ruby-sdk/codegen/pom.xml
    def add_child_codegen_module(self, sdk_dir, module_location, module_name):
        file_name = os.path.join(sdk_dir, module_location, "pom.xml")
        pom = parse_pom(file_name)

        content = MODULE_TEMPLATE.format(name=module_name)
        module_element = ET.fromstring(content, parser=ET.XMLParser(target=CommentedTreeBuilder()))

        # find modules
        modules = pom.findall("./ns:modules", ns)[0]
        modules.append(module_element)

        indent(pom.getroot())
        pom.write(file_name, encoding="UTF-8", xml_declaration=True)

    # Verifies required parameters for adding a new service spec.
    def verify_params_for_add(self, spec_name, group_id, artifact_id, version, relative_spec_path):
        # Check input parameters
        missing_params = []
        if not spec_name:
            missing_params.append('--spec-name')

        if not group_id:
            missing_params.append('--group-id')

        if not artifact_id:
            missing_params.append('--artifact-id')

        if not version:
            missing_params.append('--version')

        if not relative_spec_path:
            missing_params.append('--relative-spec-path')

        if missing_params:
            raise MissingParameter(
                'The following options must be specified for a new spec: {}'.format(', '.join(missing_params)))

    # Main business logic for adding a new service spec
    def add_spec(self,
                 sdk_dir,
                 module_location,
                 spec_name,
                 group_id,
                 artifact_id,
                 version,
                 relative_spec_path,
                 subdomain):
        self.verify_params_for_add(spec_name, group_id, artifact_id, version, relative_spec_path)

        # Create the pom.xml file for writing (e.g., ruby-sdk/codegen/new_service/pom.xml)
        service_pom_file = os.path.join(module_location, spec_name, "pom.xml")
        service_dir = os.path.join(module_location, spec_name)
        if not os.path.exists(service_dir):
            os.makedirs(service_dir)
        print('Module {} does not exist. Adding it as {}...'.format(spec_name, service_pom_file))

        sdk_version = find_pom_version(os.path.join(sdk_dir, "pom.xml"))
        service_friendly_name = string.capwords(spec_name.replace('_', ' '))

        print("SDK VERSION: {}, Friendly Name: {}".format(sdk_version, service_friendly_name))
        self.generate_child_codegen_pom(
            service_pom_file,
            spec_name,  # module_name == spec_name for Ruby
            service_friendly_name,
            sdk_version,
            group_id,
            artifact_id,
            version,
            relative_spec_path,
            subdomain)
        self.add_child_codegen_module(sdk_dir, module_location, spec_name)

        return AddOrUpdateSpecResult(
            updated=False,
            existing=False,
            ignored=[],
            previous={},
            changed=[])

    ##################################################
    # Update
    ##################################################

    def gather_settings(self, module_name, service_pom_file):
        settings = {}

        pom = parse_pom(service_pom_file)

        xpath = self.spec_params_xml_path_dict['group_id']
        xml_property = pom.findall(xpath, ns)[0]
        settings["group_id"] = xml_property.text

        xpath = self.spec_params_xml_path_dict['artifact_id']
        xml_property = pom.findall(xpath, ns)[0]
        settings["artifact_id"] = xml_property.text

        xpath = self.spec_params_xml_path_dict['version']
        xml_property = pom.findall(xpath, ns)[0]
        settings["version"] = xml_property.text

        settings["module_name"] = module_name

        xpath = self.spec_params_xml_path_dict['relative_spec_path']
        xml_property = pom.findall(xpath, ns)[0]
        settings["relative_spec_path"] = xml_property.text

        xpath = self.spec_params_xml_path_dict['subdomain']
        xml_property = pom.findall(xpath, ns)
        if xml_property:
            settings["subdomain"] = xml_property[0].text

        return settings

    def resolve_sub_domain(self, endpoint, subdomain):
        if subdomain:
            return subdomain

        if endpoint and not subdomain:
            subdomain = endpoint
            subdomain = re.sub('^.*://', '', subdomain)  # remove protocol and '://'
            subdomain = re.sub(r'\.{domain}.*$', '', subdomain)  # remove '.{domain}' and everything after it
        else:
            subdomain = None

        return subdomain

    def update_version_of_existing_spec(self, pom_file, version):
        pom = parse_pom(pom_file)

        xpath = self.spec_params_xml_path_dict['version']
        xml_property = pom.findall(xpath, ns)[0]
        old_version = xml_property.text

        if not is_version_increasing(old_version, version):
            return old_version

        xml_property.text = version
        indent(pom.getroot())
        pom.write(pom_file, encoding="UTF-8", xml_declaration=True)
        return None  # the old version was lower

    def update_relative_spec_path_of_existing_spec(self, pom_file, relative_spec_path):
        pom = parse_pom(pom_file)

        xpath = self.spec_params_xml_path_dict['relative_spec_path']
        xml_property = pom.findall(xpath, ns)[0]
        xml_property.text = relative_spec_path
        indent(pom.getroot())
        pom.write(pom_file, encoding="UTF-8", xml_declaration=True)

    def update_subdomain(self, pom_file, subdomain):
        pom = parse_pom(pom_file)

        xpath = self.spec_params_xml_path_dict['subdomain']
        xml_property = pom.findall(xpath, ns)[0]
        xml_property.text = subdomain
        indent(pom.getroot())
        pom.write(pom_file, encoding="UTF-8", xml_declaration=True)

    def update_spec(self,
                    sdk_dir,
                    service_pom_file,
                    spec_name,
                    group_id,
                    artifact_id,
                    artifact_version,
                    relative_spec_path,
                    subdomain):
        print('Artifact {} already exists in pom.xml. Updating specified fields...'.format(artifact_id))

        ignored_settings = []
        previous_settings = self.gather_settings(spec_name, service_pom_file)
        changed_settings = []

        if artifact_version:
            updated_version = self.update_version_of_existing_spec(service_pom_file, artifact_version)
            if updated_version:
                print('The version was not updated to {}, because it was already at {}.'.format(artifact_version,
                                                                                                updated_version))

        if relative_spec_path:
            self.update_relative_spec_path_of_existing_spec(service_pom_file, relative_spec_path)

        if subdomain:
            self.update_subdomain(service_pom_file, subdomain)

        if spec_name:
            ignored_settings.append('subdomain')

        if group_id:
            ignored_settings.append('group_id')

        current_settings = self.gather_settings(spec_name, service_pom_file)
        changed_settings = changed = compute_changed_settings(previous_settings, current_settings)

        return AddOrUpdateSpecResult(
            updated=changed != [],
            # not found means it's a new spec, or if it is an existing spec, changed needs to be non-empty
            existing=True,
            ignored=ignored_settings,
            previous=previous_settings,
            changed=changed_settings)

    def find_existing_spec_pom(self, source_path, spec_name, artifact_id):
        # Search for module pom.xml based on artifact_id
        pom_location_from_artifact_id = None
        for item in os.listdir(source_path):
            path = os.path.join(source_path, item)
            if not os.path.isdir(path):
                continue

            pom_location = os.path.join(path, "pom.xml")
            if not os.path.exists(pom_location):
                continue

            pom = parse_pom(pom_location)
            properties = pom.findall(".//ns:properties//ns:codegen.artifactory.artifactId", ns)
            if len(properties) > 0 and artifact_id == properties[0].text:
                pom_location_from_artifact_id = pom_location

        # Search for module pom.xml based on the spec_name
        pom_location_from_spec_name = None
        if spec_name:
            pom_location_from_spec_name = os.path.join(source_path, spec_name, "pom.xml")

        does_pom_from_spec_name_exist = pom_location_from_spec_name and os.path.exists(pom_location_from_spec_name)
        does_pom_from_artifact_id_exist = pom_location_from_artifact_id and os.path.exists(
            pom_location_from_artifact_id)

        # Service pom was found from the spec name
        if does_pom_from_spec_name_exist and not does_pom_from_artifact_id_exist:
            return pom_location_from_spec_name

        # Spec name was not provided, but the service pom was found via artifact id
        if not does_pom_from_spec_name_exist and does_pom_from_artifact_id_exist:
            return pom_location_from_artifact_id

        # This is an update to an already matching spec with the same artifact_id and spec_name.
        if does_pom_from_spec_name_exist and does_pom_from_artifact_id_exist and pom_location_from_spec_name == pom_location_from_artifact_id:
            return pom_location_from_spec_name

        # This is a new service spec.  Return None.
        if not does_pom_from_spec_name_exist and not does_pom_from_artifact_id_exist:
            return None

        # At this point, both a pom was resolved from the spec_name as well from a matching artifact_id.  Favor the spec_name.
        if does_pom_from_spec_name_exist:
            # For Ruby, favor the spec_name in order to update the artifact_id.  If the spec_name is to be renamed,
            # it will require a manual update to remove the existing service spec from the Ruby SDK.
            print(
                "artifact_id [{}] already exists under [{}].  Returning pom based on spec name [{}]".format(artifact_id,
                                                                                                            pom_location_from_artifact_id,
                                                                                                            pom_location_from_spec_name))
            return pom_location_from_spec_name

        raise EnvironmentError("Unable to determine path for service pom.xml")

    def add_or_update_spec(self,
                           artifact_id=None,
                           group_id=None,
                           spec_name=None,
                           relative_spec_path=None,
                           endpoint=None,
                           subdomain=None,
                           signing_strategy=None,  # Not used
                           version=None,
                           spec_generation_type=None,  # Not used
                           regional_sub_service_overrides=None,  # Not used
                           non_regional_sub_service_overrides=None,  # Not used
                           pom_location=None,
                           module_location=None):
        if signing_strategy:
            print("signing_strategy is ignored and requires a manual update")

        if spec_generation_type:
            print("spec_generation_type is ignored")

        if regional_sub_service_overrides:
            print("regional_sub_service_overrides is ignored")

        if non_regional_sub_service_overrides:
            print("non_regional_sub_service_overrides is ignored and requires a manual update")

        sdk_dir = os.path.dirname(pom_location)

        if spec_name:
            spec_name = convert_camel_to_snake_case(spec_name)

        if pom_location and not module_location:
            module_location = os.path.join(os.path.dirname(pom_location), self.pom_module_location)
            print("Location of pom modules: {}".format(module_location))

        service_pom_path = self.find_existing_spec_pom(module_location, spec_name, artifact_id)
        subdomain = self.resolve_sub_domain(endpoint, subdomain)
        print("Resolved subdomain: {}".format(subdomain))
        print("Service Pom.xml path:  {}".format(service_pom_path))

        if service_pom_path:
            return self.update_spec(sdk_dir,
                                    service_pom_path,
                                    spec_name,
                                    group_id,
                                    artifact_id,
                                    version,
                                    relative_spec_path,
                                    subdomain)
        else:
            return self.add_spec(sdk_dir,
                                 module_location,
                                 spec_name,
                                 group_id,
                                 artifact_id,
                                 version,
                                 relative_spec_path,
                                 subdomain)
