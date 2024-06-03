import json
import os
import skclient
import sys
import time

from requests.exceptions import ConnectionError
from skclient.rest import ApiException
from functools import wraps

# Add the root of the package, two directories up, to the sys.path
dir_path = os.path.dirname(os.path.realpath(__file__))
sys.path.append(os.path.join(dir_path, '..'))
import config    # noqa: ignore=F402
import util      # noqa: ignore=F402

_STOREKEEPER_CLIENT = None


def setup_endpoint(endpoint):
    if not endpoint:
        raise Exception('Endpoint is not configured: %s' % endpoint)
    skclient.Configuration().host = endpoint


def setup_skclient(r0=False):
    if not r0:
        sk_endpoint = os.environ.get('SK_ENDPOINT')
    else:
        sk_endpoint = 'https://storekeeper.oci.oraclecorp.com/v1'

    setup_endpoint(sk_endpoint)
    print('storekeeper endpoint:', skclient.Configuration().host)
    setup_skuser(r0)


def setup_skuser(r0=False):
    tenant_id = os.environ.get('TENANT_ID')
    user_id = os.environ.get('USER_ID')
    user_fingerprint = os.environ.get('USER_FINGERPRINT')
    private_key_filename = os.environ.get('PRIVATE_KEY_FILENAME')
    print('tenant_id: {}'.format(tenant_id))
    print('user_id: {}'.format(user_id))
    print('user_fingerprint: {}'.format(user_fingerprint))
    print('private_key_filename: {}'.format(private_key_filename))

    if tenant_id and user_id and user_fingerprint and private_key_filename:
        skclient.setup_user(tenant_id, user_id, user_fingerprint,
                            private_key_filename)
    else:
        raise ValueError('Please provide values for environment variables TENANT_ID, USER_ID, USER_FINGERPRINT, and PRIVATE_KEY_FILENAME.')


def retry(exceptions, tries=3, delay=5, backoff=2, logger=None):
    def deco_retry(f):
        @wraps(f)
        def f_retry(*args, **kwargs):
            nextretry, nextdelay = tries, delay
            while nextretry > 0:
                try:
                    return f(*args, **kwargs)
                except exceptions as e:
                    print("Retrying in {} seconds because of '{}' ".format(
                        nextdelay, e))
                    if logger:
                        msg = "Retrying in {} seconds because of '{}' " \
                              "...".format(nextdelay, e)
                        logger.warning(msg)
                    time.sleep(nextdelay)
                    nextretry -= 1
                    nextdelay *= backoff
            return f(*args, **kwargs)
        return f_retry
    return deco_retry


@retry(ConnectionError, tries=3, delay=30)
def sk_api_retry(func, *args, **kwargs):
    try:
        return func(*args, **kwargs)
    except ApiException as ex:
        if 599 > ex.status >= 500:
            raise ConnectionError("SK api failed due to {}".format(ex))
        raise


def STOREKEEPER_CLIENT():
    global _STOREKEEPER_CLIENT

    if _STOREKEEPER_CLIENT:
        return _STOREKEEPER_CLIENT

    setup_skclient(True)
    _STOREKEEPER_CLIENT = skclient.StoreKeeperApi()
    return _STOREKEEPER_CLIENT


def list_regions_from_storekeeper():
    page_token = 0
    regions = []
    while (page_token is not None and int(page_token) >= 0):
        response = sk_api_retry(
            STOREKEEPER_CLIENT().list_regions,
            num_results=50, page_token=page_token)
        page_token = response.next_page_token
        print("page_token is {}".format(page_token))
        for region in response.regions:
            if region.realm_name is not None and region.realm_name != "OC0" and region.region_state == "active":
                print("canonical_name: {}; name: {}; realm: {}; status: {}".format(region.canonical_name, region.name, region.realm_name, region.region_state))
                regions.append(region)
    return regions


def is_region_type_in_special_list(region, special_region_types):
    print("Region: {} has region_service_types: {}".format(region.canonical_name, region.region_service_types))
    for rtype in region.region_service_types:
        for special_region_type in special_region_types:
            if special_region_type in rtype.region_service_type:
                print('Region has special type {}'.format(special_region_type))
                return True
    return False


def get_region_from_storekeeper(region_id, region_types_to_ignore=[]):
    try:
        region = STOREKEEPER_CLIENT().get_region(region_id)
    except ApiException as e:
        print("Caught exception {}".format(e))
        region = None
    # There must be some error either in the DEX ticket (maybe mis-typed region id)
    # or Storekeeper is not behaving correctly. We cannot proceed.
    if region is None:
        return ""
    # Before returning the region information, check if the region is supposed to be
    # processed. region_types_to_ignore contains a list of types that should not be
    # included in SDK. Return None for those types. The caller should handle None
    # properly.
    if region.region_service_types is not None:
        if is_region_type_in_special_list(region, region_types_to_ignore):
            print('region {} with special type will be ignored'.format(region_id))
            return None
    else:
        # The absense of region_service_types in the region returned from Storekeeper
        # will be treated as normal region and it will be processed.
        print('region_service_types not found')
    region_json = {}
    region_json['regionKey'] = region.canonical_short_code
    region_json['realmKey'] = region.realm_name.lower()
    region_json['regionIdentifier'] = region.canonical_name
    region_json['realmDomainComponent'] = region.realm_domain_name
    return region_json


def get_regions_by_realms(regions):
    regions_by_realm_number = {}
    for region in regions:
        realm_number = int(region['realmKey'][2:])
        if realm_number not in regions_by_realm_number:
            regions_by_realm_number[realm_number] = [region]
        else:
            regions_by_realm_number[realm_number].append(region)
    return regions_by_realm_number


def get_regions_from_json_file(regions_json_file_full_path):
    regions = None
    if not os.path.isfile(regions_json_file_full_path):
        raise ValueError('{} is not a valid file.'.format(regions_json_file_full_path))
    with open(regions_json_file_full_path, 'r') as f:
        regions = json.load(f)
    return regions


def update_regions_json_file(new_regions, regions_json_file_full_path, json_indent=4):
    data = get_regions_from_json_file(regions_json_file_full_path)
    existing_region_ids = []
    for item in data:
        existing_region_ids.append(item['regionIdentifier'])
    for new_region in new_regions:
        if new_region['regionIdentifier'] not in existing_region_ids:
            data.append(new_region)
    with open(regions_json_file_full_path, 'w') as f:
        json.dump(data, f, indent=json_indent)


def get_new_regions_info_for_branch(branch='preview'):
    issues = util.get_region_support_tickets_to_process(branch)
    return get_new_regions_info_from_issues(issues)


def get_new_regions_info_from_issues(issues):
    ids = util.get_region_id_from_dex_tickets(issues)
    new_regions = []
    for id in ids:
        print('Getting region details for {}'.format(id))
        new_regions.append(get_region_from_storekeeper(id))
    return new_regions


def get_issues_with_special_regions_to_be_ignored(issues, region_types_to_ignore):
    issues_to_ignore = []
    issues_with_invalid_regions = []
    for issue in issues:
        region_id = issue.raw['fields']['summary'].split()[-1]
        print('Getting region details for {}'.format(region_id))
        new_region = get_region_from_storekeeper(region_id, region_types_to_ignore)
        if new_region is None:
            issues_to_ignore.append(issue)
        elif new_region == "":
            issues_with_invalid_regions.append(issue)
    return issues_to_ignore, issues_with_invalid_regions