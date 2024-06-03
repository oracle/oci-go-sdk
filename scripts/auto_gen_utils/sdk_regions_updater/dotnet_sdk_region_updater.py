import logging
import os 
from .region_updater_utils import get_regions_by_realms, update_regions_json_file, get_regions_from_json_file


class DotNetSDKRegionsUpdater:
    script_dir = os.path.dirname(os.path.realpath(__file__))
    regions_json_file_path = os.path.join("Commontests", "Regions.json")
    realms_file_path = os.path.join("Common", "Src", "RealmDefinitions.cs")
    regions_file_path = os.path.join("Common", "Src", "RegionDefinitions.cs")
    realms_template_path = os.path.join(script_dir, "templates", "dotnet-sdk-realms.tpl")
    regions_template_path = os.path.join(script_dir, "templates", "dotnet-sdk-regions.tpl")
    realm_line_template = '        public static readonly Realm {} = new Realm("{}", "{}");\n'
    region_line_template = '        public static readonly Region {} = Register("{}", Realm.{}, "{}");\n'
    region_no_short_code_line_template = '        public static readonly Region {} = Register("{}", Realm.{});\n'

    def __init__(self, sdk_path):
        self.logger = logging.getLogger("OCI-DotNetSDK-Regions-Updater")
        logging.basicConfig(level=logging.INFO)
        self.dotnet_sdk_path = sdk_path

    def _update_realms_file(self, regions):
        realms_source_file_full_path = os.path.join(self.dotnet_sdk_path, self.realms_file_path)
        processed_realms = []
        new_realms_file_content = ''
        for region in regions:
            if region['realmKey'].lower() not in processed_realms:
                processed_realms.append(region['realmKey'].lower())
                new_line = self.realm_line_template.format(region['realmKey'].upper(), region['realmKey'].lower(), region['realmDomainComponent'])
                new_realms_file_content += new_line
        new_realms_file_content = new_realms_file_content.rstrip()
        with open(self.realms_template_path, 'r') as ft, open(realms_source_file_full_path, 'w') as fw:
            realms_template = ft.read()
            realms_code = realms_template.replace('%REALMS_DEFINITIONS%', new_realms_file_content)
            self.logger.log(logging.INFO, 'Writing realms source code to {}'.format(realms_source_file_full_path))
            fw.write(realms_code)

    def _update_regions_file(self, regions):
        regions_source_file_full_path = os.path.join(self.dotnet_sdk_path, self.regions_file_path)
        processed_regions = []
        new_regions_file_content = ''
        regions_by_realm_number = get_regions_by_realms(regions)
        # Sort realms so that the regions can be added by realm
        realm_numbers = list(regions_by_realm_number.keys())
        realm_numbers.sort()
        for number in realm_numbers:
            comment_line = '        // OC{}\n'.format(number)
            new_regions_file_content += comment_line
            for region in regions_by_realm_number[number]:
                if region['regionIdentifier'] not in processed_regions:
                    processed_regions.append(region['regionIdentifier'])
                    # Region key (short code) is optional. 
                    if 'regionKey' in region:
                        new_line = self.region_line_template.format(region['regionIdentifier'].upper().replace('-', '_'), region['regionIdentifier'].lower(), region['realmKey'].upper(), region['regionKey'].lower())
                    else:
                        new_line = self.region_no_short_code_line_template.format(region['regionIdentifier'].upper().replace('-', '_'), region['regionIdentifier'].lower(), region['realmKey'].upper())
                    new_regions_file_content += new_line
            new_regions_file_content += '\n'
        new_regions_file_content = new_regions_file_content.rstrip()
        with open(self.regions_template_path, 'r') as ft, open(regions_source_file_full_path, 'w') as fw:
            regions_template = ft.read()
            regions_code = regions_template.replace('%REGIONS_DEFINITIONS%', new_regions_file_content)
            self.logger.log(logging.INFO, 'Writing regions source code to {}'.format(regions_source_file_full_path))
            fw.write(regions_code)

    def update_regions(self, new_regions):
        regions_json_file_path = os.path.join(self.dotnet_sdk_path, self.regions_json_file_path)
        update_regions_json_file(new_regions, regions_json_file_path)
        regions = get_regions_from_json_file(regions_json_file_path)
        self._update_realms_file(regions)
        self._update_regions_file(regions)
