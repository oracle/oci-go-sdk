import requests
import os

# PAR that allows overwriting reports_homepage so we can publish new versions of it
reports_homepage_par = 'https://objectstorage.us-phoenix-1.oraclecloud.com/p/3AARHQXKrCbO-S7PfKc2nJG0Ad5ZCeg9KEU4K4Y8RGs/n/dex-us-phx-cli-1/b/dexreq_reports/o/reports_homepage.html'

homepage_html_file_path = os.path.join(os.path.dirname(os.path.realpath(__file__)), 'reports_homepage.html')

with open(homepage_html_file_path, 'r') as f:
    content = f.read()
    result = requests.put(
        reports_homepage_par,
        data=content,
        headers={'Content-type': 'text/html'}
    )

    if result.status_code == 200:
        print('Successful upload')
    else:
        print('Failed uploading')
        print(str(result.content))
