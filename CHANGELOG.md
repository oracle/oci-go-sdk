# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)

## 24.3.0 - 2020-09-08
### Added
- Support for Logging Service
- Support for Logging Analytics Service
- Support for Logging Search Service
- Support for Logging Ingestion Service
- Support for Management Agent Cloud Service
- Support for Management Dashboard Service
- Support for Service Connector Hub service
- Support for Policy based Request/Response transformation in the API Gateway Service
- Support for sending diagnostic interrupt to a VM instance in the Compute Service
- Support for custom Database Software Images in the Database Service
- Support for getting and listing container database patches for Autonomous Container Database resources in the Database Service
- Support for updating patch id on maintenance run for Autonomous Container Database resources in the Database Service
- Support for searching Oracle Cloud resources across tenancies in the Search Service
- Documentation update for Logging Policies in the API Gateway service

## 24.2.0 - 2020-09-01
### Added
- Support for calling Oracle Cloud Infrastructure services in the ap-chiyoda-1 region
- Support for VM database cloning in the Database service
- Support for the MAINTENANCE_IN_PROGRESS lifecycle state on database systems, VM clusters, and Cloud Exadata in the Database service
- Support for provisioning refreshable clones in the Database service
- Support for new options on listeners and backend sets for specifying SSL protocols, SSL cipher suites, and server ordering preferences in the Load Balancing service
- Support for AMD flexible shapes with configurable CPU in the Container Engine for Kubernetes service
- Support for network sources in authentication policies in the Identity service

## 24.1.0 - 2020-08-18
### Added
- Support for custom boot volume size and other node pool updates in the Container Engine for Kubernetes service
- Support for Data Guard on Exadata Cloud at Customer VM clusters in the Database service
- Support for stopping VM instances after scheduled maintenance or hypervisor reboots in the Compute service
- Support for creating and managing private endpoints in the Data Flow service
- Fix upload manager upload extreme large file out of memory issue
- Temporarily remove go mod feature

## 24.0.0 - 2020-08-11
### Added
- Support for autonomous json databases in the Database service
- Support for cleaning up uncommitted multipart uploads in the Object Storage service
- Support for additional list API filters in the Data Catalog service
- Support for Go SDK logging to file 
- Support for Go Modules

### Breaking changes
- Some unusable region enums were removed from the Support Management service
- `CreateIncidentRequest` parameter `OpcRetryToken` was removed from the Support Management service

## 23.0.0 - 2020-08-04
### Added
- Support for calling Oracle Cloud Infrastructure services in the uk-gov-cardiff-1 region
- Support for creating and managing private endpoints in the Data Flow service
- Support for changing instance shapes and restarting nodes in the Big Data service
- Support for additional versions (for example CSQL) in the Big Data service
- Support for creating stacks from compartments in the Resource Manager service

### Breaking changes
- Updated the property of `LifeCycleDetails` to `LifecycleDetails` from the model of `BlockchainPlatformSummary` and `BlockchainPlatformByHostname` in the blockchain service

## 22.0.0 - 2020-07-28
### Added
- Support for calling Oracle Cloud Infrastructure services in the us-sanjose-1 region
- Support for PKCS#8 format API Keys
- Support for updating the fault domain and launch options of VM instances in the Compute service
- Support for image capability schemas and schema versions in the Compute service
- Support for 'Patch Now' maintenance runs for autonomous Exadata infrastructure and autonomous container database resources in the Database service
- Support for automatic performance and cost tuning on volumes in the Block Storage service

### Breaking changes
- Removed the accessToken field from the GitlabAccessTokenConfigurationSourceProvider model in the Resource Manager service

## 21.4.0 - 2020-07-21
### Added
- Support for license types on instances in the Content and Experience service

## 21.3.0 - 2020-07-14
### Added
- Support for the Blockchain service
- Support for failing over an autonomous database that has Data Guard enabled in the Database service
- Support for switching over an autonomous database that has Data Guard enabled in the Database service
- Support for git configuration sources in the Resource Manager service
- Support for optionally specifying a VCN id on list operations of DHCP options, subnets, security lists, route tables, internet gateways, and local peering gateways in the Networking service

## 21.2.0 - 2020-07-07
### Added
- Support for registering and deregistering autonomous dedicated databases with Data Safe in the Database service
- Support for switching between non-private-endpoints and private endpoints on autonomous databases in the Database service
- Support for returning group names when listing identity provider groups in the Identity service
- Support for server-side object re-encryption in the Object Storage service
- Support for private endpoint (ingress) and public endpoint whitelisting in the Analytics Cloud service

### Fixed
- Fixed a bug where setting a Content-Type header twice in the same request

## 21.1.0 - 2020-06-30
### Added
- Support for the Usage service
- Support for the VMware Provisioning service
- Support for applying one-off patches to databases in the Database service
- Support for layer-2 virtualization features on vlans in the Networking service
- Support for all AttachVolumeDetails and ParavirtualizedAttachVolumeDetails properties on instance configurations in the Compute Management service
- Support for setting HTTP header size and allowing invalid characters in HTTP request headers in the Load Balancing service

## 21.0.0 - 2020-06-23
### Added
- Support for the Data Integration service
- Support for updating database home IDs on databases in the Database service
- Support for backing up autonomous databases on Cloud at Customer in the Database service
- Support for managing autonomous VM clusters on Cloud at Customer in the Database service
- Support for accessing data assets via private endpoints in the Data Catalog service
- Support for dependency archive zip files to be specified for use by applications in the Data Flow service

### Breaking changes
- Property `LifecycleState` type changed from `JobLifecycleStateEnum` to `ListJobsLifecycleStateEnum` in the Data Catalog service
- Property `JobType` type changed from `JobTypeEnum` to `ListJobsJobTypeEnum` in the Data Catalog service

## 20.1.0 - 2020-06-16
### Added
- Support for creating a new database from an existing database based on a given timestamp in the Database service
- Support for enabling archive log backups of databases in the Database service
- Support for returning the database version on autonomous container databases in the Database service
- Support for the new DNS format of the Data Transfer service
- Support for scheduled autoscaling, which allows for scaling actions triggered at particular times based on CRON expressions, in the Compute Autoscaling service
- Support for filtering of list APIs for groups, identity providers, identity provider groups, compartments, dynamic groups, network sources, policies, and users by name or lifecycle state in the Identity Service

## 20.0.0 - 2020-06-09
### Added
- Support for returning the database version of backups in the Database service
- Support for patching on Exadata Cloud at Customer resources in the Database service
- Support for new lifecycle substates on instances in the Digital Assistant service
- Support for file servers in the Integration service
- Support for deleting non-empty tag namespaces and bulk deleting tags in the Identity service
- Support for bulk move and bulk delete of resources by compartment in the Identity service
- Support for allowing config file location to be set via env var

### Breaking changes
- Updated property DataStorageSizeInTBs type from *int to *float64 in the database service
- Removed state 'OFFLINE' and added 'DISCONNECTED' for property ExadataInfrastructureLifecycleStateEnum in database service


## 19.4.0 - 2020-06-02
### Added
- Support for optionally supplying a signature when deleting an agreement in the Marketplace service
- Support for launching paid listings in non-US regions in the Marketplace service
- Support for returning the image id of packages in the Marketplace service
- Support for calling Oracle Cloud Infrastructure services in the ap-chuncheon-1 region
- Support security-token based authentication for all services


## 19.3.0 - 2020-05-18
### Added
- Support for returning the private IP of a private endpoint database in the Database service
- Support for native JWT validation in the API Gateway service


## 19.2.0 - 2020-05-12
### Added
- Support for drift detection in the Resource Manager service


## 19.1.0 - 2020-05-05
### Added
- Support for updating the license type of database systems in the Database service
- Support for updating the version of 19c autonomous databases in the Database service
- Support for backup and restore functionality in the Key Management service
- Support for reports in the Marketplace service
- Support for calling Oracle Cloud Infrastructure services in the ap-hyderabad-1 region

## 19.0.0 - 2020-04-28
### Added
- Support for the MySQL Database service
- Support for updating the database home of a database in the Database service
- Support for government regions in the Marketplace service
- Support for starting and stopping instances in the Integration service
- Support for installing Windows updates in the OS Management service

### Breaking changes
- Removed the models of `UpdatablePackageSummary`, `ManagedInstanceUpdateDetails` and the header parameter `etag` from `WorkRequestErrorsResponse` and `WorkRequestLogsResponse` in the osmanagement service

## 18.0.0 - 2020-04-21
### Added
- Support for the Data Safe service
- Support for the Incident Management service
- Support for showing which database versions support always-free in the Database service
- Support in instance configurations for flex shapes, dedicated VM hosts, encryption in transit, and KMS keys in the Compute Autoscaling service
- Support for server-side object encryption using a customer-provided encryption key in the Object Storage service
- Support for specifying maintenance preferences while launching and updating Exadata Database systems in the Database service
- Support for flexible-shaped VM instances in the Compute service
- Support for scheduled cross-region backups in the Block Volume service
- Support for object versioning in the Object Storage service

### Breaking changes
- Removed the models of `Archiver`, `CreateArchiverDetails` and `UpdateArchiverDetails`, operations of `CreateArchiver`, `GetArchiver`, `StartArchiver`, `StopArchiver` and `UpdateArchiver` in the streaming service

## 17.4.0 - 2020-04-14
### Added
- Support for access types on instances in the Content and Experience service
- Support for identity contexts in the Search service

## 17.3.0 - 2020-04-07
### Added
- Support for changing compartments of runs and applications in the Data Flow service
- Support for getting usage information in the Key Management Vault service
- Support for custom Key Management service endpoints and private endpoints on stream pools in the Streaming service

## 17.2.0 - 2020-03-31
### Added
- Support for the Secrets Management service 
- Support for the Big Data service
- Support for updating class name, file URI, language, and spark version of applications in the Data Flow service
- Support for cross-region replication in the Object Storage service
- Support for retention rules in the Object Storage service
- Support for enabling and disabling pod security policy admission controllers in the Container Engine for Kubernetes service

## 17.1.0 - 2020-03-24
### Added
- Support for Web Application Acceleration and Security configurations on instances in the Content and Experience service
- Support for shared database homes on Exadata Cloud at Customer resources in the Database service
- Support for Exadata database creation from backup in the Database service
- Support for conditions on JavaScript challenges, new action types on access rules, new policy configuration settings, exclusions on custom protection rules, and IP address lists on IP whitelists in the Web Application Acceleration and Security service

## 17.0.0 - 2020-03-17
### Added
- Support for serial console connections in the Database service
- Support for preview database versions in the Database service
- Support for node reboot migration maintenance status and maintenance windows in the Database service
- Support for using instance metadata API v2 for instance principals authentication


### Breaking changes
- Removed the model of `AutonomousExadataInfrastructureMaintenanceWindow` from Database service

## 16.0.0 - 2020-03-10
### Added
- Support for Events service integration with alerts in the Budgets service
- Support delegation-token auth for all services

### Breaking changes
- The parameters sort_by and lifecycle_state type from Budget service are changed from str to enum

## 15.8.0 - 2020-03-03
### Added
- Support for updating the shape of a Database System in the Database service
- Support for generating CPE configurations for download in the Networking service
- Support for private IPs and fault domains of cluster nodes in the Container Engine for Kubernetes service
- Support for calling Oracle Cloud Infrastructure services in the ca-montreal-1 region
- Support for exposing error after retrying failed in all services.

## 15.7.0 - 2020-02-25
### Added
- Support for restarting autonomous databases in the Database service
- Support for private endpoints on autonomous databases in the Database service
- Support for IP-based policies in the Identity service
- Support for management of OAuth 2.0 client credentials in the Identity service
- Support for OCI Functions as a subscription protocol in the Notifications service

## 15.6.0 - 2020-02-18
### Added
- Support for the NoSQL Database service
- Support for filtering database versions by storage management type in the Database service
- Support for specifying paid listing types within pricing models in the Marketplace service
- Support for primary and non-primary instance types in the Content and Experience service

## 15.5.0 - 2020-02-11
### Added
- Support for listing supported database versions for Autonomous Database Serverless, and selecting a version at provisioning time in the Database service
- Support for TCP proxy protocol versions on listener connection configurations in the Load Balancer service
- Support for calling the Notifications service in alternate realms
- Support for calling Oracle Cloud Infrastructure services in the eu-amsterdam-1 and me-jeddah-1 regions
- Support for non-default profiles for credentials

## 15.4.0 - 2020-02-04
## Added
- Support for the Data Science service
- Support for calling Oracle Cloud Infrastructure services in the ap-osaka-1 and ap-melbourne-1 regions

## 15.3.0 - 2020-01-28
## Added
- Support for the Application Migration service
- Support for the Data Flow service
- Support for the Data Catalog service
- Support for cross-shape Data Guard in the Database service
- Support for offline data export in the Data Transfer service

## 15.2.0 - 2020-01-21
## Added
- Support for getting DRG redundancy status in the Networking service
- Support for cloning autonomous databases from backups in the Database service

## 15.1.0 - 2020-01-14
### Added
- Support for a description field on route rules and security rules in the Networking service
- Support for starting and stopping Digital Assistant instances in the Digital Assistant service
- Support for shared database homes on Exadata, bare metal, and virtual machine instances in the Database service
- Support for tracking a number of Database service operations through the Work Requests service

## 15.0.0 - 2020-01-07
### Added
- Support for optionally specifying the corporate proxy field when creating Exadata infrastructure in the Database service
- Support for maintenance windows, and rescheduling maintenance runs, on autonomous container databases in the Database service

### Breaking changes
- Field `hostname` in `NodeDetails` from Database service is changed to mandatory

## 14.0.0 - 2019-12-17
### Added
- Support for the API Gateway service
- Support for the OS Management service
- Support for the Marketplace service
- Support for "default"-type vaults in the Key Management service
- Support for bringing your own keys in the Key Management service 
- Support for cross-region backups of boot volumes in the Block Storage service
- Support for top-level TSIG keys in the DNS service
- Support for resizing virtual machine instances to different shapes in the Compute service
- Support for management configuration of cloud agents in the Compute service
- Support for launching node pools using image IDs in the Container Engine for Kubernetes service

### Breaking changes
- Removed support for v1 auth tokens in kubeconfig files in the `CreateClusterKubeconfigContentDetails` class of the Container Engine for Kubernetes service
- Removed the IDCS access token requirement on the delete deleteOceInstance operation in the Content and Experience service, which is why the `DeleteOceInstanceDetails` class was removed
- Parameter `compartment_id` in `list_stream_pools` API from Streaming service is changed to required parameter

## 13.1.0 - 2019-12-10
### Added
- Support for etags on results of the List Objects API in the Object Storage service
- Support for OCIDs on buckets in the Object Storage service
- Support for content-disposition and cache-control headers on objects in the Object Storage service
- Support for recovering deleted compartments in the Identity service
- Support for sharing volumes across multiple instances in the Block Storage service
- Support for connect harnesses and stream pools in the Streaming service
- Support for associating file storage mount targets with network security groups in the File Storage service 
- Support for calling Oracle Cloud Infrastructure services in the uk-gov-london-1 region

## 13.0.0 - 2019-11-26
### Added
- Support for maintenance windows on autonomous databases in the Database service
- Support for getting the compute units (OCPUs) of an Exadata autonomous transaction processing - dedicated resource in the Database service

### Breaking changes
- Create database home from VM_CLUSTER_BACKUP is removed from Database Service
- Response type is changed for following two APIs in Virtual Network Service 
    - Before

    ```golang
    BulkAddVirtualCircuitPublicPrefixes (err error)

    BulkDeleteVirtualCircuitPublicPrefixes (err error)
    ```

    - After

    ```golang
    BulkAddVirtualCircuitPublicPrefixes (response BulkAddVirtualCircuitPublicPrefixesResponse, err error)

    BulkDeleteVirtualCircuitPublicPrefixes (response BulkDeleteVirtualCircuitPublicPrefixesResponse, err error)
    ```

## 12.5.0 - 2019-11-19
### Added
- Support for four-byte autonomous system numbers (ASNs) on FastConnect resources in the Networking service
- Support for choosing fault domains when creating instance pools in the Compute service
- Support for allowing connections from only specific VCNs to autonomous data warehouse and autonomous transaction processing instances in the Database service
- Support for Streaming Client Non-Regional

## 12.4.0 - 2019-11-12
### Added
- Support for access to APEX and SQL Dev features on autonomous transaction processing and autonomous data warehouse resources in the Database service
- Support for registering / deregistering autonomous transaction processing and autonomous data warehouse resources with Data Safe in the Database service
- Support for redirecting HTTP / HTTPS request URIs to different URIs in the Load Balancing service
- Support for specifying compartments on options APIs in the Container Engine for Kubernetes service
- Support for volume performance units on block volumes in the Block Storage service
- Support for opc-multipart-md5 verification for UploadManager. Example can be found on [Github](https://github.com/oracle/oci-go-sdk/blob/v8.0.0/example/example_objectstorage_test.go#L57)

## 12.3.0 - 2019-11-05
### Added
- Support for the Analytics Cloud service
- Support for the Integration Cloud service
- Support for IKE versions in IPSec connections in the Virtual Networking service
- Support for getting a stack's Terraform state in the Resource Manager service

## 12.2.0 - 2019-10-29
### Added
- Support for wallet rotation operations on Autonomous Databases in the Database service
- Support for adding and removing image shape compatibility entries in the Compute service
- Support for managing redirects in the Web Application Acceleration and Security service
- Support for migrating zones from the Dyn HTTP Redirect Service to Oracle Cloud Infrastructure in the DNS service

## 12.1.0 - 2019-10-15
### Added
- Support for the Digital Assistant service
- Support for work requests on Instance Pool operations in the Compute service

## 12.0.0 - 2019-10-08
### Added
- Support for the new schema for events in the Audit service
- Support for entitlements in the Data Transfer service
- Support for custom scheduled backup policies on volumes in the Block Storage service
- Support for specifying the network type when launching virtual machine instances in the Compute service
- Support for Monitoring service integration in the Health Checks service

### Fixed
- OCI Golang SDK hook/callback to display progress bar for uploads [Github issue 187](https://github.com/oracle/oci-go-sdk/issues/187)

### Breaking changes
* The TenantId parameter is now Id (Id of the Transfer Application Entitlement) for GetTransferApplianceEntitlementRequest in TransferApplianceEntitlementClient
* The Audit service version was bumped to 20190901, use older version of Go SDK for Audit service version 20160918 

## 11.0.0 - 2019-10-01
### Added
- Support for required tags in the Identity service
- Support for work requests on tagging operations in the Identity service
- Support for enumerated tag values in the Identity service
- Support for moving dynamic routing gateway resources across compartments in the Networking service
- Support for migrating zones from Dyn managed DNS to OCI in the DNS service
- Support for fast provisioning for virtual machine databases in the Database service

### Breaking changes
- The field``CreateZoneDetails`` is no longer an anonymous field and the type changed from struct to interface in struct ``CreateZoneRequest``. Here is sample code that shows how to update your code to incorporate this change. 


    - Before

    ```golang
    // There were two ways to initialize the CreateZoneRequest struct.
    // This breaking change only impact option #2
    request := dns.CreateZoneRequest{}

    // #1. Instantiate CreateZoneDetails directly: no impact
    details := dns.CreateZoneDetails{}
    details.Name = common.String('some name')
    // ... other properties
    // Set it to the request class
    request.CreateZoneDetails = details

    // #2. Instantiate CreateZoneDetails through anonymous fields: will break
    request.Name = common.String('some name')
    // ... other properties
    ```

    - After

    ```golang
    // #2 no longer supported. Create CreateZoneDetails directly
    details := dns.CreateZoneDetails{}
    details.Name = common.String('some name')
    // ... other properties

    request := dns.CreateZoneRequest{
        CreateZoneDetails: details
    }
    // ...
    ```

## 10.1.0 - 2019-09-24
### Added
- Support for selecting the Terraform version to use in the Resource Manager service
- Support for bucket re-encryption in the Object Storage service
- Support for enabling / disabling bucket-level events in the Object Storage service

## 10.0.0 - 2019-09-17
### Added
- Support for importing state files in the Resource Manager service
- Support for Exadata Cloud at Customer in the Database service
- Support for free tier resources and system tags in the Load Balancing service
- Support for free tier resources and system tags in the Compute service
- Support for free tier resources and system tags in the Block Storage service
- Support for free tier and system tags on autonomous databases in the Database service

### Breaking
- Interface CreateDbHomeWithDbSystemIdBase is renamed to CreateDbHomeBase and dbSystemId property is removed from it
- CreateDbHomeWithDbSystemIdBase in CreateDbHomeRequest is replaced with CreateDbHomeWithDbSystemIdDetails

## 9.0.0 - 2019-09-10
### Added
- Support for specifying the autoBackupWindow field for scheduling backups in the Database service
- Support for network security groups on autonomous Exadata infrastructure in the Database service
- Support for Kubernetes secrets encryption in customer clusters, regional subnets, and cluster authentication for instance principals in the Container Engine for Kubernetes service
- Support for the Oracle Content and Experience service

### Breaking
- The etag field has been removed from the ChangeSubscriptionCompartmentResponse and ChangeTopicCompartmentResponse structs of the Notifications service

## 8.1.0 - 2019-09-03
### Added
- Support for the Sydney (SYD) region
- Support for managing cluster networks in the Compute Autoscaling service
- Support for tracking asynchronous operations via work requests in the Database service

## 8.0.0 - 2019-08-27
### Added
- Support for the Sao Paulo (GRU) region
- Support for dedicated virtual machine hosts in the Compute service
- Support for resource groups in metrics and alarms in the Monitoring service
- Support for resource principle auth. Example can be found on [Github](https://github.com/oracle/oci-go-sdk/tree/master/example/example_resource_principal_function/README.md)

### Breaking changes
- Breaking changes were made for following enum values
    - Before
    ```golang
    autoscaling.ActionTypeEnum.ActionTypeBy
    keymanagement.CreateVaultDetailsVaultTypeEnum.CreateVaultDetailsVaultTypePrivate
    keymanagement.VaultSummaryVaultTypeEnum.VaultSummaryVaultTypePrivate
    keymanagement.VaultVaultTypeEnum.VaultVaultTypePrivate
    objectstorage.WorkRequestSummaryOperationTypeEnum.WorkRequestSummaryOperationTypeObject
    objectstorage.WorkRequestOperationTypeEnum.WorkRequestOperationTypeObject
    resourcemanager.LogEntryTypeEnum.LogEntryTypeConsole
    resourcemanager.WorkRequestOperationTypeEnum.WorkRequestOperationTypeCompartment
    ```

    - After
    ```golang
    autoscaling.ActionTypeEnum.ActionTypeChangeCountBy
    keymanagement.CreateVaultDetailsVaultTypeEnum.CreateVaultDetailsVaultTypeVirtualPrivate
    keymanagement.VaultSummaryVaultTypeEnum.VaultSummaryVaultTypeVirtualPrivate
    keymanagement.VaultVaultTypeEnum.VaultVaultTypeVirtualPrivate
    objectstorage.WorkRequestSummaryOperationTypeEnum.WorkRequestSummaryOperationTypeCopyObject
    objectstorage.WorkRequestOperationTypeEnum.WorkRequestOperationTypeCopyObject
    resourcemanager.LogEntryTypeEnum.LogEntryTypeTerraformConsole
    resourcemanager.WorkRequestOperationTypeEnum.WorkRequestOperationTypeChangeStackCompartment
    ```

## 7.1.0 - 2019-08-20
### Added
- Support for the Limits service
- Support for archiving to Object Storage in the Streaming service
- Support for etags on resources in the Streaming service
- Support for Key Management service (KMS) encryption of file systems in the File Storage service
- Support for moving public IP, DHCP, local peering gateway, internet gateway, network security group, and DRG attachment resources across compartments in the Networking service
- Support for multi-origin, basic cache, certificate mapping, and OCI Monitoring service integration in the Web Application Acceleration and Security service

## 7.0.0 - 2019-08-13
### Added
- Support for the Data Transfer service
- Support for the Zurich (ZRH) region

### Breaking changes
- Breaking changes were made in the Web Application Acceleration and Security (WAAS) service
  - `WorkRequestSummaryOperationTypePurgeWaasPolicy` const removed from `waas/work_request_summary.go`
  - `WorkRequestOperationTypesPurgeWaasPolicy` const removed from `waas/work_request_operation_types.go`
  - `WorkRequestOperationTypesPurgeWaasPolicy` const removed from `waas/work_request.go`
  - `IssuerName` in `Certificate` struct changed type from `*CertificateSubjectName` to `*CertificateIssuerName`
  - `LifecycleState` changed from array of string to array of `ListCertificateLifeCycleStateEnum` in `waas/list_certificates_request_response.go` and `waas/list_waas_policies_request_response.go`
  - `Etag` was removed from the following structs:
     - `AcceptRecommendationsResponse`
     - `DeleteWaasPolicyResponse`
     - `UpdateAccessRulesResponse`
     - `UpdateCaptchasResponse`
     - `UpdateDeviceFingerprintChallengeResponse`
     - `UpdateGoodBotsResponse`
     - `UpdateHumanInteractionChallengeResponse`
     - `UpdateJsChallengeResponse`
     - `UpdatePolicyConfigResponse`
     - `UpdateProtectionRulesResponse`
     - `UpdateProtectionSettingsResponse`
     - `UpdateThreatFeedsResponse`
     - `UpdateWaasPolicyResponse`
     - `UpdateWafAddressRateLimitingResponse`
     - `UpdateWafConfigResponse`
     - `UpdateWhitelistsResponse`

## 6.2.0 - 2019-08-06
### Added
- Support for IPv6 load balancers in the Load Balancing service
- Support for IPv6 on VCN and FastConnect resources in the Networking service

## 6.1.0 - 2019-07-30
### Added
- Support for the Mumbai (BOM) region
- Support for the Events service
- Support for moving streams across compartments in the Streaming service
- Support for moving FastConnect resources across compartments in the Networking service
- Support for moving policies across compartments in the Web Application Acceleration and Security service
- Support for tagging FastConnect resources in the Networking service

## 6.0.0 - 2019-07-23
### Added
- Support for moving resources across compartments in the Database service
- Support for moving resources across compartments in the Health Checks service
- Support for moving alarms across compartments in the Monitoring service
- Support for creating instance configurations from running instances in the Compute service
- Support for setting up budget alerts for cost tracking tags in the Budgets service

## 5.15.0 - 2019-07-16
### Added
- Support for the Functions service
- Support for the Quotas service
- Support for moving resources across compartments in the DNS service
- Support for moving instances across compartments in the Compute service
- Support for moving keys and vaults across compartments in the Key Management service
- Support for moving topics and subscriptions across compartments in the Notifications service
- Support for moving load balancers across compartments in the Load Balancing service
- Support for specifying permitted REST methods in load balancer rule sets in the Load Balancing service
- Support for configuring cookie session persistence in backend sets in the Load Balancing service
- Support for ACL rules in rule sets in the Load Balancing service
- Support for move compartment tree in the Identity service
- Support for specifying and returning a KMS key in backup operations in the Block Storage service
- Support for transit routing in the Networking service

## 5.14.0 - 2019-07-09
### Added
- Support for network security groups in the Load Balancing service
- Support for network security groups in Core Services
- Support for network security groups on database systems in the Database service
- Support for creating autonomous transaction processing and autonomous data warehouse previews in the Database service
- Support for getting the load balancer attachments of instance pools in the Compute service
- Support for moving resources across compartments in the Resource Manager service
- Support for moving VCN resources across compartments in the Networking service

## 5.13.0 - 2019-07-02
### Added
- Support for moving images, instance configurations, and instance pools across compartments in Core Services
- Support for moving autoscaling configurations across compartments in the Compute Autoscaling service

### Fixed
- Fixed a bug where the Streaming service's endpoints in Tokyo, Seoul, and future regions were not reachable from the SDK

## 5.12.0 - 2019-06-25
### Added
- Support for moving senders across compartments in the Email service
- Support for moving NAT gateway resources across compartments in Core Services

## 5.11.0 - 2019-06-18
### Added
- Support for moving service gateway resources across compartments in Core Services
- Support for moving block storage resources across compartments in Core Services
- Support for key deletion in the Key Management service

## 5.10.0 - 2019-06-11
### Added
- Support for specifying custom boot volume sizes on instance configurations in the Compute Autoscaling service
- Support for 'Autonomous Transaction Processing - Dedicated' features, as well as maintenance run and backup operations on autonomous databases, autonomous container databases, and autonomous Exadata infrastructure in the Database service

## 5.9.0 - 2019-06-04
### Added
- Support for autoscaling autonomous databases and autonomous data warehouses in the Database service
- Support for specifying fault domains as part of instance configurations in the Compute Autoscaling service
- Support for deleting tag definitions and tag namespaces in the Identity service

### Fixed
- Support for regions in realms other than oraclecloud.com in the Load Balancing service

## 5.8.0 - 2019-05-28
### Added
- Support for the Work Requests service, and tracking of a number of Core Services operations through work requests
- Support for emulated volume attachments in Core Services
- Support for changing the compartment of resources in the File Storage service
- Support for tags in list operations in the File Storage service
- Support for returning UI password creation dates in the Identity service

## 5.7.0 - 2019-05-21
### Added
- Support for returning tags when listing instance configurations, instance pools, or autoscaling configurations in the Compute Autoscaling service
- Support for getting the namespace of another tenancy than the caller's tenancy in the Object Storage service
- Support for BGP dynamic routing and providing pre-shared secrets (PSKs) when establishing tunnels in the Networking service

## 5.6.0 - 2019-05-14
### Added
- Support for the Seoul (ICN) region
- Support for logging context fields on data-plane APIs of the Key Management Service
- Support for reverse pagination on list operations of the Email service
- Support for configuring backup retention windows on database backups in the Database service

## 5.5.0 - 2019-05-07
### Added
- Support for the Tokyo (NRT) region

- Support UploadManager for uploading large objects. Sample is available on [Github](https://github.com/oracle/oci-go-sdk/tree/master/example/example_objectstorage_test.go)

## 5.4.0 - 2019-04-16
### Added
- Support for tagging dynamic groups in the Identity service
- Support for updating network ACLs and license types for autonomous databases and autonomous data warehouses in the Database service
- Support for editing static routes and IPSec remote IDs in the Virtual Networking service

## 5.3.0 - 2019-04-09
### Added
- Support for etag and if-match headers (for optimistic concurrency control) in the Email service

## 5.2.0 - 2019-04-02
### Added
- Support for provider service key names on virtual circuits in the FastConnect service
- Support for customer reference names on cross connects and cross connect groups in the FastConnect service

## 5.1.0 - 2019-03-26
### Added
- Support for glob patterns and exclusions for object lifecycle management in the Object Storage service
- Documentation enhancements and corrections for traffic management in the DNS service

### Fixed
- The 'tag' info is always ignored in the returned string of Version() function [Github issue 157](https://github.com/oracle/oci-go-sdk/issues/157)

## 5.0.0 - 2019-03-19
### Added

- Support for specifying metadata on node pools in the Container Engine for Kubernetes service
- Support for provisioning a new autonomous database or autonomous data warehouse as a clone of another in the Database service
### Breaking changes
- The field``CreateAutonomousDatabaseDetails`` is no longer an anonymous field and the type changed from struct to interface in struct ``CreateAutonomousDatabaseRequest``. Here is sample code that shows how to update your code to incorporate this change. 

    - Before

    ```golang
    // create a CreateAutonomousDatabaseRequest
    // There were two ways to initialize the CreateAutonomousDatabaseRequest struct.
    // This breaking change only impact option #2
    request := database.CreateAutonomousDatabaseRequest{}

    // #1. Instantiate CreateAutonomousDatabaseDetails directly: no impact
    details := database.CreateAutonomousDatabaseDetails{}
    details.CompartmentId = common.String(getCompartmentID())
    // ... other properties

    // Set it to the request class
    request.CreateAutonomousDatabaseDetails = details

    // #2. Instantiate CreateAutnomousDatabaseDetails through  anonymous fields: will break
    request.CompartmentId = common.String(getCompartmentID())
    // ... other properties
    ```

    - After

    ```golang
    // #2 no longer supported. Create CreateAutonomousDatabaseDetails directly
    details := database.CreateAutonomousDatabaseDetails{}
    details.CompartmentId = common.String(getCompartmentID())
    // ... other properties

    // and set the details to CreateAutonomousDatabaseBase
    request := database.CreateAutonomousDatabaseRequest{}
    request.CreateAutonomousDatabaseDetails = details
    // ...
    ```


## 4.2.0 - 2019-03-12
### Added
- Support for the Budgets service
- Support for managing multifactor authentication in the Identity service
- Support for managing default tags in the Identity service
- Support for account recovery in the Identity service
- Support for authentication policies in the Identity service
- Support for specifying the workload type when creating autonomous databases in the Database service
- Support for I/O resource management for Exadata database systems in the Database service
- Support for customer-specified timezones on database systems in the Database service

## 4.1.0 - 2019-02-28
### Added
- Support for the Monitoring service
- Support for the Notification service
- Support for the Resource Manager service
- Support for the Compute Autoscaling service
- Support for changing the compartment of a tag namespace in the Identity service
- Support for specifying fault domains in the Database service
- Support for managing instance monitoring in the Compute service
- Support for attaching/detaching load balancers to instance pools in the Compute service

## 4.0.0 - 2019-02-21
### Added
- Support for government-realm regions
- Support for the Streaming service
- Support for tags in the Key Management service
- Support for regional subnets in the Virtual Networking service

### Fixed
- Removed unused Announcements service 'NotificationFollowupDetails' struct and 'GetFollowups' operation
- InstancePrincipals now invalidates a token shortly before its expiration time to avoid making  a service call with an expired token
- Requests with binary bodies that require its body to be included in the signature are now being signed correctly

## 3.7.0 - 2019-02-07
### Added
- Support for the Web Application Acceleration and Security (WAAS) service
- Support for the Health Checks service
- Support for connection strings on Database resources in the Database service
- Support for traffic management in the DNS service
- Support for tagging in the Email service
### Fixed
- Retry context in now cancelable during wait for new retry

## 3.6.0 - 2019-01-31
### Added
- Support for the Announcements service

## 3.5.0 - 2019-01-24
### Added

- Support for renaming databases during restore-from-backup operations in the Database service
- Built-in logging now supports log levels. More information about the changes can be found in the [go-docs page](https://godoc.org/github.com/oracle/oci-go-sdk#hdr-Logging_and_Debugging)
- Support for calling Oracle Cloud Infrastructure services in the ca-toronto-1 region

## 3.4.0 - 2019-01-10
### Added 
- Support for device attributes on volume attachments in the Compute service
- Support for custom header rulesets in the Load Balancing service


## 3.3.0 - 2018-12-13
### Added 
- Support for Data Guard for VM shapes in the Database service
- Support for sparse disk groups for Exadata shapes in the Database service
- Support for a new field, isLatestForMajorVersion, when listing DB versions in the Database service
- Support for in-transit encryption for paravirtualized boot volume and data volume attachments in the Block Storage service
- Support for tagging DNS Zones in the DNS service
- Support for resetting credentials for SCIM clients associated with an Identity provider and updating user capabilities in the Identity service

## 3.2.0 - 2018-11-29
### Added 
- Support for getting bucket statistics in the Object Storage service

### Fixed
- Block Storage service for copying volume backups across regions is now enabled 
- Object storage `PutObject` and `UploadPart` operations now do not override the client's signer

## 3.1.0 - 2018-11-15
### Added
- Support for VCN transit routing in the Networking service 

## 3.0.0 - 2018-11-01
### Added
- Support for modifying the route table, DHCP options and security lists associated with a subnet in the Networking service.
- Support for tagging of File Systems, Mount Targets and Snapshots in the File Storage service.
- Support for nested compartments in the Identity service

### Notes
- The version is bumped due to breaking changes in previous release.

## 2.7.0 - 2018-10-18
### Added
- Support for cost tracking tags in the Identity service
- Support for generating and downloading wallets in the Database service
- Support for creating a standalone backup from an on-premises database in the Database service
- Support for db version and additional connection strings in the Autonomous Transaction Processing and Autonomous Data Warehouse resources of the Database service
- Support for copying volume backups across regions in the Block Storage service
- Support for deleting compartments in the Identity service
- Support for reboot migration for virtual machines in the Compute service
- Support for Instance Pools and Instance Configurations in the Compute service

### Fixed
- The signing algorithm does not lower case the header fields [Github issue 132](https://github.com/oracle/oci-go-sdk/issues/132)
- Raw configuration provider does not check for empty strings [Github issue 134](https://github.com/oracle/oci-go-sdk/issues/134)

### Breaking change
- DbDataSizeInMBs field in Backup and BackupSummary struct was renamed to DatabaseSizeInGBs and type changed from *int to *float64 
    - Before
    ```golang
    // Size of the database in megabytes (MB) at the time the backup was taken.
    DbDataSizeInMBs *int `mandatory:"false" json:"dbDataSizeInMBs"`
    ```

    - After

    ```golang
    // The size of the database in gigabytes at the time the backup was taken.
    DatabaseSizeInGBs *float64 `mandatory:"false" json:"databaseSizeInGBs"`
    ```
- Data type for DatabaseEdition in Backup and BackupSummary struct was changed from *string to BackupDatabaseEditionEnum
    - Before

    ```golang
    // The Oracle Database edition of the DB system from which the database backup was taken.
    DatabaseEdition *string `mandatory:"false" json:"databaseEdition"`
    ```

    - After

    ```golang
     // The Oracle Database edition of the DB system from which the database backup was taken.
     DatabaseEdition BackupDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`
    ```

## 2.6.0 - 2018-10-04
### Added
- Support for trusted partner images through application listings and subscriptions in the Compute service
- Support for object lifecycle policies in the Object Storage service
- Support for copying objects across regions in the Object Storage service
- Support for network address translation (NAT) gateways in the Networking service

## 2.5.0 - 2018-09-27
### Added
- Support for paravirtualized launch mode when importing images in the Compute service
- Support for Key Management service
- Support for encrypting the contents of an Object Storage bucket using a Key Management service key
- Support for specifying a Key Management service key when launching a compute instance in the Compute service
- Support for specifying a Key Management service key when backing up or restoring a block storage volume in the Block Volume service

## 2.4.0 - 2018-09-06
### Added
- Added support for updating metadata fields on an instance in the Compute service

## 2.3.0 - 2018-08-23
### Added
- Support for fault domain in the Identity Service
- Support for Autonomous Data Warehouse and Autonomous Transaction Processing in the Database service
- Support for resizing an offline volume in the Block Storage service
- Nil interface when polymorphic json response object is null

## 2.2.0 - 2018-08-09
### Added
- Support for fault domains in the Compute service
- A sample showing how to use Search service from the SDK is available on [Github](https://github.com/oracle/oci-go-sdk/tree/master/example/example_resourcesearch_test.go)

## 2.1.0 - 2018-07-26
### Added
- Support for the Search service
- Support for specifying a backup policy when creating a boot volume in the Block Storage service

### Fixed
- OCI error is missing opc-request-id value [Github Issue 120](https://github.com/oracle/oci-go-sdk/issues/120)
- Include raw http response when service error occurred

## 2.0.0 - 2018-07-12
### Added
- Support for tagging Load Balancers in the Load Balancing service
- Support for export options in the File Storage service
- Support for retrieving compartment name and user name as part of events in the Audit service

### Fixed
- CreateKubeconfig function should not close http reponse body [Github Issue 116](https://github.com/oracle/oci-go-sdk/issues/116)

### Breaking changes
- Datatype changed from *int to *int64 for several request/response structs. Here is sample code that shows how to update your code to incorporate this change. 

    - Before

    ```golang
    // Update the impacted properties from common.Int to common.Int64.
    // Here is the updates for CreateBootVolumeDetails
    details := core.CreateBootVolumeDetails{
        SizeInGBs: common.Int(10),
    }
    ```

    - After

    ```golang
    details := core.CreateBootVolumeDetails{
        SizeInGBs: common.Int64(10),
    }
    ```

- Impacted packages and structs
    - core
        - BootVolume.(SizeInGBs, SizeInMBs)
        - BootVolumeBackup.(SizeInGBs, UniqueSizeInGBs)
        - CreateBootVolumeDetails.SizeInGBs
        - CreateVolumeDetails.(SizeInGBs, SizeInMBs)
        - Image.SizeInMBs
        - InstanceSourceViaImageDetails.BootVolumeSizeInGBs
        - Volume.(SizeInGBs, SizeInMBs)
        - VolumeBackup.(SizeInGBs, SizeInMBs, UniqueSizeInGBs, UniqueSizeInMbs)
        - VolumeGroup.(SizeInMBs, SizeInGBs)
        - VolumeGroupBackup.(SizeInMBs, SizeInGBs, UniqueSizeInMbs, UniqueSizeInGbs)
    - dns
        - GetDomainRecordsRequest.Limit
        - GetRRSetRequest.Limit
        - GetZoneRecordsRequest.Limit
        - ListZonesRequest.Limit
        - Zone.Serial
        - ZoneSummary.Serial
    - filestorage
        - ExportSet.(MaxFsStatBytes, MaxFsStatFiles)
        - FileSystem.MeteredBytes
        - FileSystemSummary.MeteredBytes
        - UpdateExportSetDetails.(MaxFsStatBytes, MaxFsStatFiles)
    - identity
        - ApiKey.InactiveStatus
        - AuthToken.InactiveStatus
        - Compartment.InactiveStatus
        - CustomerSecretKey.InactiveStatus
        - CustomerSecretKeySummary.InactiveStatus
        - DynamicGroup.InactiveStatus
        - Group.InactiveStatus
        - IdentityProvider.InactiveStatus
        - IdpGroupMapping.InactiveStatus
        - Policy.InactiveStatus
        - Saml2IdentityProvider.InactiveStatus
        - SmtpCredential.InactiveStatus
        - SmtpCredentialSummary.InactiveStatus
        - SwiftPassword.InactiveStatus
        - UiPassword.InactiveStatus
        - User.InactiveStatus
        - UserGroupMembership.InactiveStatus
    - loadbalancer
        - ConnectionConfiguration.IdleTimeout
        - ListLoadBalancerHealthsRequest.Limit
        - ListLoadBalancersRequest.Limit
        - ListPoliciesRequest 
        - ListProtocolsRequest.Limit
        - ListShapesRequest.Limit
        - ListWorkRequestsRequest.Limit
    - objectstorage
        - GetObjectResponse.ContentLength
        - HeadObjectResponse.ContentLength
        - MultipartUploadPartSummary.Size
        - ObjectSummary.Size
        - PutObjectRequest.ContentLength
        - UploadPartRequest.ContentLength

## 1.8.0 - 2018-06-28
### Added
- Support for service gateway management in the Networking service
- Support for backup and clone of boot volumes in the Block Storage service

## 1.7.0 - 2018-06-14
### Added
- Support for the Container Engine service. A sample showing how to use this service from the SDK is available [Github](https://github.com/oracle/oci-go-sdk/tree/master/example/example_containerengine_test.go)

### Fixed
- Empty string was send to backend service for optional enum if it's not set

## 1.6.0 - 2018-05-31
### Added
- Support for the "soft shutdown" instance action in the Compute service
- Support for Auth Token management in the Identity service
- Support for backup or clone of multiple volumes at once using volume groups in the Block Storage service
- Support for launching a database system from a backup in the Database service

### Breaking changes
- ``LaunchDbSystemDetails`` is renamed to ``LaunchDbSystemBase`` and the type changed from struct to interface in ``LaunchDbSystemRequest``. Here is sample code that shows how to update your code to incorporate this change. 

    - Before

    ```golang
    // create a LaunchDbSystemRequest
    // There were two ways to initialize the LaunchDbSystemRequest struct.
    // This breaking change only impact option #2
    request := database.LaunchDbSystemRequest{}

    // #1. explicity create LaunchDbSystemDetails struct (No impact)
    details := database.LaunchDbSystemDetails{}
    details.AvailabilityDomain = common.String(validAD())
    details.CompartmentId = common.String(getCompartmentID())
    // ... other properties
    request.LaunchDbSystemDetails = details

    // #2. use anonymous fields (Will break)
    request.AvailabilityDomain = common.String(validAD())
    request.CompartmentId = common.String(getCompartmentID())
    // ...
    ```

    - After

    ```golang
    // create a LaunchDbSystemRequest
    request := database.LaunchDbSystemRequest{}
    details := database.LaunchDbSystemDetails{}
    details.AvailabilityDomain = common.String(validAD())
    details.CompartmentId = common.String(getCompartmentID())
    // ... other properties

    // set the details to LaunchDbSystemBase
    request.LaunchDbSystemBase = details
    // ...
    ```

## 1.5.0 - 2018-05-17
### Added
- ~~Support for backup or clone of multiple volumes at once using volume groups in the Block Storage service~~
- Support for the ability to optionally specify a compartment filter when listing exports in the File Storage service
- Support for tagging virtual cloud network resources in the Networking service
- Support for specifying the PARAVIRTUALIZED remote volume type when creating a virtual image or launching a new instance in the Compute service
- Support for tilde in private key path in configuration files

## 1.4.0 - 2018-05-03
### Added
- Support for ``event_name`` in Audit Service
- Support for multiple ``hostnames`` for loadbalancer listener in LoadBalance service
- Support for auto-generating opc-request-id for all operations
- Add opc-request-id property for all requests except for Object Storage which use opc-client-request-id

## 1.3.0 - 2018-04-19
### Added
- Support for retry on Oracle Cloud Infrastructure service APIs. Example can be found on [Github](https://github.com/oracle/oci-go-sdk/tree/master/example/example_retry_test.go)
- Support for tagging DbSystem and Database resources in the Database Service
- Support for filtering by DbSystemId in ListDbVersions operation in Database Service

### Fixed
- Fixed a request signing bug for PatchZoneRecords API
- Fixed a bug in DebugLn

## 1.2.0 - 2018-04-05
### Added
- Support for Email Delivery Service. Example can be found on [Github](https://github.com/oracle/oci-go-sdk/tree/master/example/example_email_test.go)
- Support for paravirtualized volume attachments in Core Services
- Support for remote VCN peering across regions
- Support for variable size boot volumes in Core Services
- Support for SMTP credentials in the Identity Service
- Support for tagging Bucket resources in the Object Storage Service

## 1.1.0 - 2018-03-27
### Added
- Support for DNS service
- Support for File Storage service
- Support for PathRouteSets and Listeners in Load Balancing service
- Support for Public IPs in Core Services
- Support for Dynamic Groups in Identity service
- Support for tagging in Core Services and Identity service. Example can be found on [Github](https://github.com/oracle/oci-go-sdk/tree/master/example/example_tagging_test.go)
- Fix ComposingConfigurationProvider to not accept a nil ConfigurationProvider
- Support for passphrase configuration to FileConfiguration provider

## 1.0.0 - 2018-02-28 Initial Release
### Added
- Support for Audit service
- Support for Core Services (Networking, Compute, Block Volume)
- Support for Database service
- Support for IAM service
- Support for Load Balancing service
- Support for Object Storage service
