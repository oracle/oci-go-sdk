DEXREQ PR Tracker
========================

Overview
-----------------------
The purpose of this tool is to generate a report that helps visualize the status of the various DEXREQ tickets for a given UDX feature.

The script takes in a comma separateed list of UDX ticket keys (e.g. UDX-1234,UDX-234) and generates 3 reports: one preview report showing only PRs to preview related to that UDX ticket, one public report showing only PRs to master related to that UDX ticket, and one full report showing both combined.

Note by default the script only looks back ~45 PRs in bitbucket history to limit the load on the Bitbucket APIs.  If you want to include older PRs, you can override this with the --limit parameter.

Installing dependencies
------------------------
To install dependencies for this project, execute the following command from the root of the auto-gen-utils project:

`pip install -r requirements.txt`

Generating a report
-------------------------
Below is an example invocation of the script:

`python dexreq_pr_tracker.py --issues UDX-494,UDX.497`

Below is an example invocation of the script to generate a report for all tickets in a given GA date:

`python dexreq_pr_tracker.py --sdk-cli-ga-date 2019-03-12`

Below is an example invocation of the script with more of the optional parameters:

`python dexreq_pr_tracker.py --issues UDX-494,UDX.497 --output-dir ~/reports/ --report-name release_10_18 --upload-to-object-storage`


Common Workflow
-------------------------
The goal for this tool is to allow generating reports for a specific release based on all of the UDX tickets that expect to be released.  This report would be continually updated to allow the on-call and other members of the team to track the status of all PRs for a given release.

Here is a suggested workflow for a release:
* At the beginning of the release, figure out which UDX tickets are aiming for the RELEASE_DATE
* Generate a report for all of those tickets using a standard report name  and upload to object storage (e.g. --report-name release_{RELEASE_DATE})
* In our internal tenancy (right now dex-us-phx-cli-1) create a PAR for the public, private, and full reports in object storage so you have a stable link that you can share for a given release
* Continue reviewing PRs and adding / removing UDX tickets from the report as the release progresses.  Each report generation with the same --report-name will overwrite the previous report in object storage


Running in Team City
-------------------------
This script can also be run in Team City as the following job:
https://teamcity.oci.oraclecorp.com/viewType.html?buildTypeId=Sdk_PythonCli_Misc_GenerateDexreqPrStatusReport


Persistent comments in reports
------------------------
You can leave notes in the report under the 'comments' column and when you lose focus on the text box the comments will be saved.  Comments are persisted across reports, they are tied to the corresponding UDX ticket so they will be shown in any report that includes that UDX ticket.

Persistent comment storage is done through the following object storage PAR (valid through 2020):
https://objectstorage.us-phoenix-1.oraclecloud.com/p/Sy7XXIKMDfXao-zvGfqXbMn6E1_xtz1GTCifWx6J8Cc/n/dex-us-phx-cli-1/b/dexreq_reports/o/datastore.json

This is also included in templates/report_table.html


Reports stored in object storage
-------------------------------

Read reports homepage PAR:
This is used as a permalink for all team members to find all of the reports that have been created:
https://objectstorage.us-phoenix-1.oraclecloud.com/p/_VNrbExh9IlndgYeH2fsS5NByaEL9rBTKy2NlceNJjc/n/dex-us-phx-cli-1/b/dexreq_reports/o/reports_homepage.html

Read / write reports homepage PAR:
This is used by publish_homepage.py to upload new versions of the homepage (mostly while developing)
https://objectstorage.us-phoenix-1.oraclecloud.com/p/3AARHQXKrCbO-S7PfKc2nJG0Ad5ZCeg9KEU4K4Y8RGs/n/dex-us-phx-cli-1/b/dexreq_reports/o/reports_homepage.html

Read / Write Reports Index PAR:
This is used by the reports homepage to read the index of reports and display them
It is also used by the dexreq_pr_tracker to update the index of reports when a new one is created
https://objectstorage.us-phoenix-1.oraclecloud.com/p/L8mVstLyLPSGmVO5-vFP5VKukTQGhj1W5d0y1NCVbE4/n/dex-us-phx-cli-1/b/dexreq_reports/o/reports_index.json
