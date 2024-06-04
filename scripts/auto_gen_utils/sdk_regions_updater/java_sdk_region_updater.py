import logging
import os 
from .region_updater_utils import get_regions_by_realms, update_regions_json_file


class JavaSDKRegionsUpdater:
    realms_file_path = os.path.join("bmc-common", "src", "main", "java", "com", "oracle", "bmc", "Realm.java")
    regions_file_path = os.path.join("bmc-common", "src", "main", "java", "com", "oracle", "bmc", "Region.java")
    regions_json_file_path = os.path.join("bmc-common", "src", "test", "resources", "regions.json")
    realm_line_prefix = "public static final Realm "
    region_line_prefix = "public static final Region "
    realm_line_template = '    public static final Realm {} = new Realm("{}", "{}");\n'
    region_line_template = '    public static final Region {} = register("{}", Realm.{}, "{}");\n'
    region_no_short_code_line_template = '    public static final Region {} = register("{}", Realm.{};\n'

    def __init__(self, sdk_path):
        self.logger = logging.getLogger("OCI-JavaSDK-Regions-Updater")
        logging.basicConfig(level=logging.DEBUG)
        self.java_sdk_path = sdk_path

    # We are assuming that realms are added in the right orders. 
    # New realm numbers are always larger than existing ones. 
    # If the assumption is not correct, realms will be organized out of order. 
    # This does not affect code functionalities though.
    def _update_realms_file(self, new_regions_by_realm):
        realms_file_full_path = os.path.join(self.java_sdk_path, self.realms_file_path)
        existing_realms = []
        should_inject_realms = False
        new_realms_file_content = ''
        with open(realms_file_full_path) as f:
            for line in f.readlines():
                if line.strip().startswith(self.realm_line_prefix):
                    should_inject_realms = True
                    realm_id = line.strip()[len(self.realm_line_prefix):].split('=')[0].strip()
                    existing_realms.append(realm_id)
                else:
                    if should_inject_realms:
                        # need to inject new realm(s) here
                        realm_numbers = list(new_regions_by_realm.keys())
                        realm_numbers.sort()
                        for number in realm_numbers:
                            new_region = new_regions_by_realm[number][0]
                            if new_region['realmKey'].upper() in existing_realms:
                                self.logger.log(logging.INFO, "realm {} already exists. Will not add again.".format(new_region['realmKey']))
                            else:
                                self.logger.log(logging.INFO, "Will add new realm {} .".format(new_region['realmKey']))
                                new_line = self.realm_line_template.format(new_region['realmKey'].upper(), new_region['realmKey'].lower(), new_region['realmDomainComponent'])
                                new_realms_file_content += new_line
                        should_inject_realms = False
                if (new_realms_file_content == ''):
                    new_realms_file_content = line
                else:
                    new_realms_file_content += line
        with open(realms_file_full_path, 'w') as fw:
            fw.write(new_realms_file_content)
        
    def _update_regions_file(self, new_regions_by_realm):
        regions_file_full_path = os.path.join(self.java_sdk_path, self.regions_file_path)
        existing_regions = []
        should_inject_regions = False
        new_regions_file_content = ''
        current_realm = 0
        is_region_definition_complete = True
        with open(regions_file_full_path) as f:
            for line in f.readlines():
                line_stripped = line.strip()
                if line_stripped.startswith('// OC'):
                    # This is the beginning of the code block defining regions for a realm.
                    current_realm = int(line_stripped.split(' ')[1][2:])
                elif line_stripped.startswith(self.region_line_prefix):
                    # This line is region definition.
                    should_inject_regions = True
                    region_name = line_stripped[len(self.region_line_prefix):].split('=')[0].strip()
                    existing_regions.append(region_name)
                    # The region definition can be broken into multiple lines. If the current line
                    # does not end with ";", then this is not a complete region definition statement.
                    if not line_stripped.endswith(';'):
                        is_region_definition_complete = False
                else:
                    # Check if the we have a region definition statement that is not complete. 
                    # The first ";" should complete the defintion statement.
                    if not is_region_definition_complete:
                        if line_stripped.endswith(';'):
                            is_region_definition_complete = True
                    # If line_stripped != '' and not line_stripped.startswith('// OC') and should_inject_regions:
                    elif should_inject_regions:
                        if line_stripped == '':
                            # At an empty line, inject new region definitions for this realm, if any
                            if current_realm in new_regions_by_realm:
                                # need to inject new region(s) here
                                for region in new_regions_by_realm[current_realm]:
                                    region_name = region['regionIdentifier'].upper().replace('-', '_')
                                    if region_name in existing_regions:
                                        self.logger.log(logging.INFO, "Region {} already exists. Will not add again.".format(region['regionIdentifier']))
                                    else:
                                        self.logger.log(logging.INFO, "Will add new region {} under realm OC{}.".format(region['regionIdentifier'], current_realm))
                                        new_line = self.region_line_template.format(region_name, region['regionIdentifier'].lower(), region['realmKey'].upper(), region['regionKey'].lower())
                                        new_regions_file_content += new_line
                                del new_regions_by_realm[current_realm]
                            # Check if there are additional regions left in new_regions_by_realm.
                            if len(new_regions_by_realm) == 0:
                                should_inject_regions = False
                        else:
                            # This is not a region defintion, and not an empty space, and not a region block comment. 
                            # This must be the first line of code after the whole region definiton block. 
                            # If there are new regions under new realms, add them here.
                            realm_numbers = list(new_regions_by_realm.keys())
                            realm_numbers.sort()
                            for number in realm_numbers:
                                comment_line = '    // OC{}\n'.format(number)
                                new_regions_file_content += comment_line
                                for region in new_regions_by_realm[number]:
                                    self.logger.log(logging.INFO, "Will add new region {} for new realm OC{}.".format(region['regionIdentifier'], number))
                                    new_line = self.region_line_template.format(region['regionIdentifier'].upper().replace('-', '_'), region['regionIdentifier'].lower(), region['realmKey'].upper(), region['regionKey'].lower())
                                    new_regions_file_content += new_line
                                new_regions_file_content += '\n'
                            should_inject_regions = False

                if (new_regions_file_content == ''):
                    new_regions_file_content = line
                else:
                    new_regions_file_content += line
        with open(regions_file_full_path, 'w') as fw:
            fw.write(new_regions_file_content)

    def update_regions(self, new_regions):
        regions_json_file_path = os.path.join(self.java_sdk_path, self.regions_json_file_path)
        update_regions_json_file(new_regions, regions_json_file_path)
        new_regions_by_realm = get_regions_by_realms(new_regions)
        self._update_realms_file(new_regions_by_realm)
        self._update_regions_file(new_regions_by_realm)
