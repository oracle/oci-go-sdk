import logging
import os 
from .region_updater_utils import get_regions_by_realms, update_regions_json_file, get_regions_from_json_file


class GoSDKRegionsUpdater:
    script_dir = os.path.dirname(os.path.realpath(__file__))
    regions_json_file_path = os.path.join("common", "regions.json")
    regions_source_file_path = os.path.join("common", "regions.go")
    regions_template_path = os.path.join(script_dir, "templates", "go-sdk-regions.tpl")
    region_realm_mapping_entry_template = '    Region{}: "{}",\n'
    short_name_region_mapping_entry_template = '    "{}": Region{},\n'
    region_definition_comment_template = '    //Region{} region {}\n'
    gov_region_definition_comment_template = '    //Region{} gov region {}\n'
    region_definition_template = '    Region{} Region = "{}"\n'
    realm_definition_template = '    "{}": "{}",\n'
    # Some earlier regions in GoSDK were named in a way that is different from the current naming conventions.
    # These legacy names need to be handled.
    special_region_names = {
        'us-phoenix-1': 'PHX',
        'us-ashburn-1': 'IAD',
        'us-sanjose-1': 'SJC1',
        'eu-frankfurt-1': 'FRA',
        'uk-london-1': 'LHR'
    }

    def __init__(self, sdk_path):
        self.logger = logging.getLogger("OCI-GoSDK-Regions-Updater")
        logging.basicConfig(level=logging.INFO)
        self.go_sdk_path = sdk_path

    def _build_region_name(self, region_id):
        if region_id.lower() in self.special_region_names:
            return self.special_region_names[region_id.lower()]
        else:
            region_name_parts = region_id.lower().split('-')
            region_name = region_name_parts[0].upper()
            for part in region_name_parts[1:]:
                region_name += part.title()
            return region_name

    def _build_region_comment(self, region_id):
        region_name = self._build_region_name(region_id)
        parts = region_id.lower().split('-')
        if len(parts) == 4 and parts[1] == 'gov':
            return self.gov_region_definition_comment_template.format(region_name, parts[2].title())
        else:
            return self.region_definition_comment_template.format(region_name, parts[-2].title())

    def _update_regions_file(self, regions):
        regions_source_file_full_path = os.path.join(self.go_sdk_path, self.regions_source_file_path)
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
                    realms_definitions += self.realm_definition_template.format(region['realmKey'].lower(), region['realmDomainComponent'].lower())
                    processed_realms.append(region['realmKey'].lower())
                if region['regionIdentifier'] not in processed_regions:
                    region_name = self._build_region_name(region['regionIdentifier'])
                    # Region key (short code) is optional. 
                    if 'regionKey' in region:
                        short_name_regions += self.short_name_region_mapping_entry_template.format(region['regionKey'].lower(), region_name)
                    region_realm_mappings += self.region_realm_mapping_entry_template.format(region_name, region['realmKey'].lower())
                    regions_definitions += self._build_region_comment(region['regionIdentifier'])
                    regions_definitions += self.region_definition_template.format(region_name, region['regionIdentifier'].lower())
                    processed_regions.append(region['regionIdentifier'])
            region_realm_mappings += '\n'
        short_name_regions = short_name_regions.rstrip()
        region_realm_mappings = region_realm_mappings.rstrip()
        regions_definitions = regions_definitions.rstrip()
        realms_definitions = realms_definitions.rstrip()
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
        regions_json_file_path = os.path.join(self.go_sdk_path, self.regions_json_file_path)
        update_regions_json_file(new_regions, regions_json_file_path)
        regions = get_regions_from_json_file(regions_json_file_path)
        self._update_regions_file(regions)
