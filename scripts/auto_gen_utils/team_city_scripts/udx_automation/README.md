# UDX intake Process Automation

## Introduction
Today when a Service teams wants to GA a new entire service or a new feature in their existing service, they need to file an ORM ticket which automatically creates a UDX ticket for tracking the UDX approvals needed this GA. Currently, the TPM manually does all the in-take process of such new UDX tickets which involves some verification of the ticket and collecting various information from related tickets.

This approach has a lot of disadvantages as detailed below:-

* This work though trivial takes a lot of effort and time of the TPM to manage. This time can be better utilized by the TPM to other higher priorities task.
* As this is a manual work, it is prone to human errors ,and we should automate wherever we can.
* As the number of services keep on increasing, this workflow will be unsustainable as it does not scale well.
* If the TPM is out of office or is busy in other higher priority tasks, the review for such tickets will pile up and the GA for services might get delayed.
* Makes it harder to hand off this work to a new TPM

## Automation Goal

* When a UDX tickets gets created for review, automation should be the first one to run and check if the ticket is ready.
* Transition to **In Design** state if the ticket meets the bar. Otherwise, move to **Closed** if it is manually created or move to **More information needed** state in case of issues
* Job will run on a regular frequency and process all the tickets needing triage
* Make the job extensible to allow for adding possibilities of modifying any requirements or add new ones.

## Automation Steps
The automation would do the following steps:-

1. Gets ticket in **Needs Triage** status in the **UDX** project which do not have the Bypass Label **Bypass-UDX-Automation**. If you want to modify the Jira Query the automation uses, then you can set the environment variable **TICKETS_IN_TRIAGE_QUERY** which has the value of the query you want to do instead.
2. Check if the ticket has a Root ORM ticket attached to it. If the ticket does not have a Root ORM then the UDX ticket is **closed** with a comment and link to create UDX ticket via the ORM process.
3. If the root ORM ticket has ***Customer facing changes*** set to **No**, then transition the UDX ticket to **closed** state with a comment that UDX ticket is not needed when there are no such changes.
4. Start the intake process for the UDX ticket
   1. If the UDX ticket **reporter** is set to ***jira-automation-bot*** then set reporter to match the reporter set in the Root ORM.
   2. If the **Service Team Project/Queue** is missing, then update this with the value from the Root ORM ticket.
   3. Find the TC ticket linked to UDX by looking at the Root ORM ticket and finding the child Technical content ORM ticket to get the TC ticket and create a link to it.
   4. Check if the **Public facing API changes**.
      1. If it is set to **Yes**, then:- 
         - set **Surface Exempt from Feature Impact = No**
         - Check if **GA date** is either a *Tuesday or Wednesday* and if it is not then add a comment asking the *Reporter* to set the **GA Date** to a *Tuesday/Wednesday* on the Root ORM.
      2. If it is set to **No**, then
         - Check if there is either an *SDK/CLI* or *Terraform* as **surface exemptions**, if not move the ticket to **More Information Needed** status with a comment asking the *Reporter* to check if they forgot to add the exemptions.
         - If the necessary **surface exemptions** are present then check if the **GA Date** is a weekday, if it is not then ask the *Reporter* to set the **GA Date** to a weekday on the Root ORM. 
   
   5. Transition tickets still not transitioned to **In Design** with a comment asking the *Reporter* to provide the VPAT documentation information.

