import os
from confluence_api_utils import utils

username = os.environ.get("CONFLUENCE_USERNAME")
password = os.environ.get("CONFLUENCE_PASSWORD")
personal_access_token = os.environ.get("CONFLUENCE_TOKEN")
page_id = "2461340038"  #used page id for  jira phonebook lookup tool page
file_name='sample.csv'
host = 'https://confluence.oci.oraclecorp.com'

# below 2 are required if you wanted to access confluence page using space and title
page_space_name = ''
page_title = ''

# auth 0 means authentication using username-password and 1 means using PAT
auth=1

if auth == 0:
    conf = utils.auth_using_username_password(host,username,password)

# If you want to login through Personal acess token
if auth == 1:
    conf = utils.auth_using_personal_access_token(host,personal_access_token)

# If you want page content by page id
#content1 = utils.get_content_of_confluence_page_by_page_id(conf,page_id)


#If you want page content by space and title --- uncomment to use
#content1 = utils.get_content_of_confluence_page_by_page_space_and_title(conf,page_space_name, page_title)

# upload a csv file
#utils.upload_with_no_replacement_a_csv_file_to_confluence_page(conf,page_id,file_name)

# remove a csv file
#utils.remove_a_csv_file_from_confluence_page(conf,page_id,file_name)

utils.upload_with_replacement_a_csv_file_to_confluence_page(conf,page_id,file_name)
