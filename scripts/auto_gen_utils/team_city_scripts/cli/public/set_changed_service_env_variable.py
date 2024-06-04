import sys
import os
import re

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '../../..'))
sys.path.append(os.path.join(dir_path, '../../'))

import config # noqa: ignore=F402
import util # noqa: ignore=F402
from add_or_update_scripts.cli_add_or_update_spec import determine_pom_location # noqa: ignore=F402

CHANGED_SERVICE = 'env.CHANGED_SERVICE'
SERVICES_DIR = "services"
POM_LOCATION_PATTERN = "services/(.*)/pom.xml"


def _get_service_name_from_path(file_path):
    result = re.search(POM_LOCATION_PATTERN, file_path)
    return result.group(1)  # group(1) would return the first capture group i.e service name


def get_service_name_from_issue(dexreq_issue):
    jira_obj = util.get_dexreq_issue(dexreq_issue)
    spec_name = getattr(jira_obj.fields, config.CUSTOM_FIELD_ID_SPEC_FRIENDLY_NAME)
    artifact_id = getattr(jira_obj.fields, config.CUSTOM_FIELD_ID_ARTIFACT_ID)
    services_root_dir = os.path.join(config.CLI_REPO_RELATIVE_LOCATION, SERVICES_DIR)

    path = determine_pom_location(artifact_id, spec_name, services_root_dir)
    return _get_service_name_from_path(path)


if __name__ == '__main__':
    """
    This script will be used in Preview/Public CLI pipeline to determine the changed service directory.
    We use this information to run make gen and make docs for the changed service.
    """
    tool_name = config.CLI_NAME
    last_commit_message = util.get_last_commit_message(tool_name)
    issue_keys = util.parse_issue_keys_from_commit_message(last_commit_message)
    if len(issue_keys) != 1:
        print('More than one DEXReq issues found {}.', issue_keys)
        sys.exit(0)

    service_name = get_service_name_from_issue(issue_keys[0])
    print("Changed service:" + service_name)
    print("##teamcity[setParameter name='{}' value='{}']".format(CHANGED_SERVICE, service_name))
