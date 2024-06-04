import logging
import os 
from .region_updater_utils import get_regions_by_realms, update_regions_json_file, get_regions_from_json_file


class RubySDKRegionsUpdater:
    script_dir = os.path.dirname(os.path.realpath(__file__))
    regions_json_file_path = os.path.join("test", "resources", "regions.json")
    regions_source_file_path = os.path.join("lib", "oci", "regions_definitions.rb")
    regions_template_path = os.path.join(script_dir, "templates", "ruby-sdk-regions.tpl")
    mapping_entry_template = '      \'{}\': \'{}\'.freeze,\n'
    short_name_region_mapping_entry_template = '      \'{}\': REGION_{},\n'
    region_enum_item_template = '      REGION_{} = \'{}\'.freeze,\n'

    def __init__(self, sdk_path):
        self.logger = logging.getLogger("OCI-RubySDK-Regions-Updater")
        logging.basicConfig(level=logging.INFO)
        self.ruby_sdk_path = sdk_path

    def _update_regions_file(self, regions):
        regions_source_file_full_path = os.path.join(self.ruby_sdk_path, self.regions_source_file_path)
        processed_regions = []
        processed_realms = []
        short_name_regions = ''
        region_realm_mappings = ''
        realms_definitions = ''
        regions_definitions = ''
        regions_by_realm_number = get_regions_by_realms(regions)
        # Sort realms so that the regions can be added by realm
        realm_numbers = list(regions_by_realm_number.keys())
        realm_numbers.sort()
        for number in realm_numbers:
            for region in regions_by_realm_number[number]:
                if region['realmKey'].lower() not in processed_realms:
                    realms_definitions += self.mapping_entry_template.format(region['realmKey'].lower(), region['realmDomainComponent'].lower())
                    processed_realms.append(region['realmKey'].lower())
                if region['regionIdentifier'] not in processed_regions:
                    # Region key (short code) is optional. 
                    if 'regionKey' in region:
                        short_name_regions += self.short_name_region_mapping_entry_template.format(region['regionKey'].lower(), region['regionIdentifier'].upper().replace('-', '_'))
                    region_realm_mappings += self.mapping_entry_template.format(region['regionIdentifier'].lower(), region['realmKey'].lower())
                    regions_definitions += self.region_enum_item_template.format(region['regionIdentifier'].upper().replace('-', '_'), region['regionIdentifier'].lower())
                    processed_regions.append(region['regionIdentifier'])
            region_realm_mappings += '\n'
        short_name_regions = short_name_regions.rstrip().rstrip(',')
        region_realm_mappings = region_realm_mappings.rstrip().rstrip(',')
        regions_definitions = regions_definitions.rstrip().rstrip(',')
        realms_definitions = realms_definitions.rstrip().rstrip(',')
        with open(self.regions_template_path, 'r') as ft, open(regions_source_file_full_path, 'w') as fw:
            regions_template = ft.read()
            regions_code = (regions_template.replace('%REGIONS_DEFINITIONS%', regions_definitions)
                                            .replace('%REALMS_DEFINITIONS%', realms_definitions)
                                            .replace('%SHORT_NAME_REGIONS%', short_name_regions)
                                            .replace('%REGION_REALM_MAPPINGS%', region_realm_mappings)
            )
            self.logger.log(logging.INFO, 'Writing regions source code to {}'.format(regions_source_file_full_path))
            fw.write(regions_code)

    def update_regions(self, new_regions):
        regions_json_file_path = os.path.join(self.ruby_sdk_path, self.regions_json_file_path)
        update_regions_json_file(new_regions, regions_json_file_path)
        regions = get_regions_from_json_file(regions_json_file_path)
        self._update_regions_file(regions)
