import subprocess
import sys
subprocess.check_call([sys.executable, "-m", "pip", "install", "atlassian-python-api"])
from atlassian import Confluence

def auth_using_username_password(host,username,password):
    conf = Confluence(url=host, username=username, password=password)
    return conf

def auth_using_personal_access_token(host,personal_access_token):
    conf = Confluence(url=host, token=personal_access_token)
    return conf

def get_content_of_confluence_page_by_page_id(conf,page_id):
    content = conf.get_page_by_id(page_id=page_id, expand='body.storage')
    print(content['title'])
    print(content['body']['storage']['value'])

def get_content_of_confluence_page_by_page_space_and_title(conf,space,title):
    content = conf.get_page_by_title(space=space, title=title,expand='body.storage')
    print(content['title'])
    print(content['body']['storage']['value'])

def get_all_attachment_name(page_id,conf):
    t = conf.get_attachments_from_content(page_id=page_id )
    return [t['results'][i]['title'] for i in range(len(t['results']))]

def upload_with_no_replacement_a_csv_file_to_confluence_page(conf,page_id,file_to_be_uploaded):
    # get all files at that page
    temp = get_all_attachment_name(page_id,conf)
    # check if same name file already exist
    if file_to_be_uploaded in temp:
        print('file already exist')
        return

    conf.attach_file(filename=file_to_be_uploaded, name=file_to_be_uploaded, page_id=page_id, comment='new file uploaded')
    print("file attached to confluence page")

def remove_a_csv_file_from_confluence_page(conf,page_id,file_to_be_removed):
    #get all files at that page
    temp=get_all_attachment_name(page_id,conf)
    #check if file exist
    if file_to_be_removed not in temp:
        print('file do not exist')
        return
    conf.delete_attachment(page_id=page_id,filename=file_to_be_removed)
    print("file removed from confluence page")

def upload_with_replacement_a_csv_file_to_confluence_page(conf,page_id,file_to_be_uploaded):
    conf.attach_file(filename=file_to_be_uploaded, name=file_to_be_uploaded, page_id=page_id, comment=' ')
    print("file attached to confluence page")