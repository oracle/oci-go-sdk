// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity


type Identity interface {

        // AddUserToGroup
        // Adds the specified user to the specified group and returns a `UserGroupMembership` object with its own OCID.\n\nAfter you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the\nobject, first make sure its `lifecycleState` has changed to ACTIVE.\n
        AddUserToGroup(AddUserToGroupDetails AddUserToGroupDetails, OpcRetryToken string) (*UserGroupMembership, *APIResponse, error)


        // CreateCompartment
        // Creates a new compartment in your tenancy.\n\n**Important:** Compartments cannot be renamed or deleted.\n\nYou must specify your tenancy's OCID as the compartment ID in the request object. Remember that the tenancy\nis simply the root compartment. For information about OCIDs, see\n[Resource Identifiers](/Content/General/Concepts/identifiers.htm).\n\nYou must also specify a *name* for the compartment, which must be unique across all compartments in\nyour tenancy and cannot be changed. You can use this name or the OCID when writing policies that apply\nto the compartment. For more information about policies, see\n[How Policies Work](/Content/Identity/Concepts/policies.htm).\n\nYou must also specify a *description* for the compartment (although it can be an empty string). It does\nnot have to be unique, and you can change it anytime with\n[UpdateCompartment](#/en/identity/20160918/Compartment/UpdateCompartment).\n\nAfter you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the\nobject, first make sure its `lifecycleState` has changed to ACTIVE.\n
        CreateCompartment(CreateCompartmentDetails CreateCompartmentDetails, OpcRetryToken string) (*Compartment, *APIResponse, error)


        // CreateGroup
        // Creates a new group in your tenancy.\n\nYou must specify your tenancy's OCID as the compartment ID in the request object (remember that the tenancy\nis simply the root compartment). Notice that IAM resources (users, groups, compartments, and some policies)\nreside within the tenancy itself, unlike cloud resources such as compute instances, which typically\nreside within compartments inside the tenancy. For information about OCIDs, see\n[Resource Identifiers](/Content/General/Concepts/identifiers.htm).\n\nYou must also specify a *name* for the group, which must be unique across all groups in your tenancy and\ncannot be changed. You can use this name or the OCID when writing policies that apply to the group. For more\ninformation about policies, see [How Policies Work](/Content/Identity/Concepts/policies.htm).\n\nYou must also specify a *description* for the group (although it can be an empty string). It does not\nhave to be unique, and you can change it anytime with [UpdateGroup](#/en/identity/20160918/Group/UpdateGroup).\n\nAfter you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the\nobject, first make sure its `lifecycleState` has changed to ACTIVE.\n\nAfter creating the group, you need to put users in it and write policies for it.\nSee [AddUserToGroup](#/en/identity/20160918/UserGroupMembership/AddUserToGroup) and\n[CreatePolicy](#/en/identity/20160918/Policy/CreatePolicy).\n
        CreateGroup(CreateGroupDetails CreateGroupDetails, OpcRetryToken string) (*Group, *APIResponse, error)


        // CreateIdentityProvider
        // Creates a new identity provider in your tenancy. For more information, see\n[Identity Providers and Federation](/Content/Identity/Concepts/federation.htm).\n\nYou must specify your tenancy's OCID as the compartment ID in the request object.\nRemember that the tenancy is simply the root compartment. For information about\nOCIDs, see [Resource Identifiers](/Content/General/Concepts/identifiers.htm).\n\nYou must also specify a *name* for the `IdentityProvider`, which must be unique\nacross all `IdentityProvider` objects in your tenancy and cannot be changed.\n\nYou must also specify a *description* for the `IdentityProvider` (although\nit can be an empty string). It does not have to be unique, and you can change\nit anytime with\n[UpdateIdentityProvider](#/en/identity/20160918/IdentityProvider/UpdateIdentityProvider).\n\nAfter you send your request, the new object's `lifecycleState` will temporarily\nbe CREATING. Before using the object, first make sure its `lifecycleState` has\nchanged to ACTIVE.\n
        CreateIdentityProvider(CreateIdentityProviderDetails CreateIdentityProviderDetails, OpcRetryToken string) (*IdentityProvider, *APIResponse, error)


        // CreateIdpGroupMapping
        // Creates a single mapping between an IdP group and an IAM Service\n[group](#/en/identity/20160918/Group/).\n
        CreateIdpGroupMapping(CreateIdpGroupMappingDetails CreateIdpGroupMappingDetails, IdentityProviderID string, OpcRetryToken string) (*IdpGroupMapping, *APIResponse, error)


        // CreateOrResetUIPassword
        // Creates a new Console one-time password for the specified user. For more information about user\ncredentials, see [User Credentials](/Content/Identity/Concepts/usercredentials.htm).\n\nUse this operation after creating a new user, or if a user forgets their password. The new one-time\npassword is returned to you in the response, and you must securely deliver it to the user. They'll\nbe prompted to change this password the next time they sign in to the Console. If they don't change\nit within 7 days, the password will expire and you'll need to create a new one-time password for the\nuser.\n\n**Note:** The user's Console login is the unique name you specified when you created the user\n(see [CreateUser](#/en/identity/20160918/User/CreateUser)).\n
        CreateOrResetUIPassword(UserID string, OpcRetryToken string) (*UiPassword, *APIResponse, error)


        // CreatePolicy
        // Creates a new policy in the specified compartment (either the tenancy or another of your compartments).\nIf you're new to policies, see [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).\n\nYou must specify a *name* for the policy, which must be unique across all policies in your tenancy\nand cannot be changed.\n\nYou must also specify a *description* for the policy (although it can be an empty string). It does not\nhave to be unique, and you can change it anytime with [UpdatePolicy](#/en/identity/20160918/Policy/UpdatePolicy).\n\nYou must specify one or more policy statements in the statements array. For information about writing\npolicies, see [How Policies Work](/Content/Identity/Concepts/policies.htm) and\n[Common Policies](/Content/Identity/Concepts/commonpolicies.htm).\n\nAfter you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using the\nobject, first make sure its `lifecycleState` has changed to ACTIVE.\n\nNew policies take effect typically within 10 seconds.\n
        CreatePolicy(CreatePolicyDetails CreatePolicyDetails, OpcRetryToken string) (*Policy, *APIResponse, error)


        // CreateRegionSubscription
        // Creates a subscription to a region for a tenancy.\n
        CreateRegionSubscription(CreateRegionSubscriptionDetails CreateRegionSubscriptionDetails, TenancyID string, OpcRetryToken string) (*RegionSubscription, *APIResponse, error)


        // CreateSwiftPassword
        // Creates a new Swift password for the specified user. For information about what Swift passwords are for, see\n[Managing User Credentials](/Content/Identity/Tasks/managingcredentials.htm).\n\nYou must specify a *description* for the Swift password (although it can be an empty string). It does not\nhave to be unique, and you can change it anytime with\n[UpdateSwiftPassword](#/en/identity/20160918/SwiftPassword/UpdateSwiftPassword).\n\nEvery user has permission to create a Swift password for *their own user ID*. An administrator in your organization\ndoes not need to write a policy to give users this ability. To compare, administrators who have permission to the\ntenancy can use this operation to create a Swift password for any user, including themselves.\n
        CreateSwiftPassword(CreateSwiftPasswordDetails CreateSwiftPasswordDetails, UserID string, OpcRetryToken string) (*SwiftPassword, *APIResponse, error)


        // CreateUser
        // Creates a new user in your tenancy. For conceptual information about users, your tenancy, and other\nIAM Service components, see [Overview of the IAM Service](/Content/Identity/Concepts/overview.htm).\n\nYou must specify your tenancy's OCID as the compartment ID in the request object (remember that the\ntenancy is simply the root compartment). Notice that IAM resources (users, groups, compartments, and\nsome policies) reside within the tenancy itself, unlike cloud resources such as compute instances,\nwhich typically reside within compartments inside the tenancy. For information about OCIDs, see\n[Resource Identifiers](/Content/General/Concepts/identifiers.htm).\n\nYou must also specify a *name* for the user, which must be unique across all users in your tenancy\nand cannot be changed. Allowed characters: No spaces. Only letters, numerals, hyphens, periods,\nunderscores, +, and @. If you specify a name that's already in use, you'll get a 409 error.\nThis name will be the user's login to the Console. You might want to pick a\nname that your company's own identity system (e.g., Active Directory, LDAP, etc.) already uses.\nIf you delete a user and then create a new user with the same name, they'll be considered different\nusers because they have different OCIDs.\n\nYou must also specify a *description* for the user (although it can be an empty string).\nIt does not have to be unique, and you can change it anytime with\n[UpdateUser](#/en/identity/20160918/User/UpdateUser). You can use the field to provide the user's\nfull name, a description, a nickname, or other information to generally identify the user.\n\nAfter you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before\nusing the object, first make sure its `lifecycleState` has changed to ACTIVE.\n\nA new user has no permissions until you place the user in one or more groups (see\n[AddUserToGroup](#/en/identity/20160918/UserGroupMembership/AddUserToGroup)). If the user needs to\naccess the Console, you need to provide the user a password (see\n[CreateOrResetUIPassword](#/en/identity/20160918/UIPassword/CreateOrResetUIPassword)).\nIf the user needs to access the Oracle Bare Metal Cloud Services REST API, you need to upload a\npublic API signing key for that user (see\n[Required Keys and OCIDs](/Content/API/Concepts/apisigningkey.htm) and also\n[UploadApiKey](#/en/identity/20160918/ApiKey/UploadApiKey)).\n\n**Important:** Make sure to inform the new user which compartment(s) they have access to.\n
        CreateUser(CreateUserDetails CreateUserDetails, OpcRetryToken string) (*User, *APIResponse, error)


        // DeleteApiKey
        // Deletes the specified API signing key for the specified user.\n\nEvery user has permission to use this operation to delete a key for *their own user ID*. An\nadministrator in your organization does not need to write a policy to give users this ability.\nTo compare, administrators who have permission to the tenancy can use this operation to delete\na key for any user, including themselves.\n
        DeleteApiKey(UserID string, Fingerprint string, IfMatch string) (*APIResponse, error)


        // DeleteGroup
        // Deletes the specified group. The group must be empty.\n
        DeleteGroup(GroupID string, IfMatch string) (*APIResponse, error)


        // DeleteIdentityProvider
        // Deletes the specified identity provider. The identity provider must not have\nany group mappings (see [IdpGroupMapping](#/en/identity/20160918/IdpGroupMapping/)).\n
        DeleteIdentityProvider(IdentityProviderID string, IfMatch string) (*APIResponse, error)


        // DeleteIdpGroupMapping
        // Deletes the specified group mapping.
        DeleteIdpGroupMapping(IdentityProviderID string, MappingID string, IfMatch string) (*APIResponse, error)


        // DeletePolicy
        // Deletes the specified policy. The deletion takes effect typically within 10 seconds.
        DeletePolicy(PolicyID string, IfMatch string) (*APIResponse, error)


        // DeleteSwiftPassword
        // Deletes the specified Swift password for the specified user.\n
        DeleteSwiftPassword(UserID string, SwiftPasswordID string, IfMatch string) (*APIResponse, error)


        // DeleteUser
        // Deletes the specified user. The user must not be in any groups.
        DeleteUser(UserID string, IfMatch string) (*APIResponse, error)


        // GetCompartment
        // Gets the specified compartment's information.\n\nThis operation does not return a list of all the resources inside the compartment. There is no single\nAPI operation that does that. Compartments can contain multiple types of resources (instances, block\nstorage volumes, etc.). To find out what's in a compartment, you must call the \"List\" operation for\neach resource type and specify the compartment's OCID as a query parameter in the request. For example,\ncall the [ListInstances](#/en/iaas/20160918/Instance/ListInstances) operation in the Cloud Compute\nService or the [ListVolumes](#/en/iaas/20160918/Volume/ListVolumes) operation in Cloud Block Storage.\n
        GetCompartment(CompartmentID string) (*Compartment, *APIResponse, error)


        // GetGroup
        // Gets the specified group's information.\n\nThis operation does not return a list of all the users in the group. To do that, use\n[ListUserGroupMemberships](#/en/identity/20160918/UserGroupMembership/ListUserGroupMemberships) and\nprovide the group's OCID as a query parameter in the request.\n
        GetGroup(GroupID string) (*Group, *APIResponse, error)


        // GetIdentityProvider
        // Gets the specified identity provider's information.
        GetIdentityProvider(IdentityProviderID string) (*IdentityProvider, *APIResponse, error)


        // GetIdpGroupMapping
        // Gets the specified group mapping.
        GetIdpGroupMapping(IdentityProviderID string, MappingID string) (*IdpGroupMapping, *APIResponse, error)


        // GetPolicy
        // Gets the specified policy's information.
        GetPolicy(PolicyID string) (*Policy, *APIResponse, error)


        // GetTenancy
        // Get the specified tenancy's information.
        GetTenancy(TenancyID string) (*Tenancy, *APIResponse, error)


        // GetUser
        // Gets the specified user's information.
        GetUser(UserID string) (*User, *APIResponse, error)


        // GetUserGroupMembership
        // Gets the specified UserGroupMembership's information.
        GetUserGroupMembership(UserGroupMembershipID string) (*UserGroupMembership, *APIResponse, error)


        // ListApiKeys
        // Lists the API signing keys for the specified user. A user can have a maximum of three keys.\n\nEvery user has permission to use this API call for *their own user ID*.  An administrator in your\norganization does not need to write a policy to give users this ability.\n
        ListApiKeys(UserID string) ([]ApiKey, *APIResponse, error)


        // ListAvailabilityDomains
        // Lists the Availability Domains in your tenancy. Specify the OCID of either the tenancy or another\nof your compartments as the value for the compartment ID (remember that the tenancy is simply the root compartment).\nSee [Where to Get the Tenancy's OCID and User's OCID](/Content/API/Concepts/apisigningkey.htm#five).\n
        ListAvailabilityDomains(CompartmentID string) ([]AvailabilityDomain, *APIResponse, error)


        // ListCompartments
        // Lists the compartments in your tenancy. You must specify your tenancy's OCID as the value\nfor the compartment ID (remember that the tenancy is simply the root compartment).\nSee [Where to Get the Tenancy's OCID and User's OCID](/Content/API/Concepts/apisigningkey.htm#five).\n
        ListCompartments(CompartmentID string, Page string, Limit int32) ([]Compartment, *APIResponse, error)


        // ListGroups
        // Lists the groups in your tenancy. You must specify your tenancy's OCID as the value for\nthe compartment ID (remember that the tenancy is simply the root compartment).\nSee [Where to Get the Tenancy's OCID and User's OCID](/Content/API/Concepts/apisigningkey.htm#five).\n
        ListGroups(CompartmentID string, Page string, Limit int32) ([]Group, *APIResponse, error)


        // ListIdentityProviders
        // Lists all the identity providers in your tenancy. You must specify the identity provider type (e.g., `SAML2` for\nidentity providers using the SAML2.0 protocol). You must specify your tenancy's OCID as the value for the\ncompartment ID (remember that the tenancy is simply the root compartment).\nSee [Where to Get the Tenancy's OCID and User's OCID](/Content/API/Concepts/apisigningkey.htm#five).\n
        ListIdentityProviders(Protocol string, CompartmentID string, Page string, Limit int32) ([]IdentityProvider, *APIResponse, error)


        // ListIdpGroupMappings
        // Lists the group mappings for the specified identity provider.\n
        ListIdpGroupMappings(IdentityProviderID string, Page string, Limit int32) ([]IdpGroupMapping, *APIResponse, error)


        // ListPolicies
        // Lists the policies in the specified compartment (either the tenancy or another of your compartments).\nSee [Where to Get the Tenancy's OCID and User's OCID](/Content/API/Concepts/apisigningkey.htm#five).\n\nTo determine which policies apply to a particular group or compartment, you must view the individual\nstatements inside all your policies. There isn't a way to automatically obtain that information via the API.\n
        ListPolicies(CompartmentID string, Page string, Limit int32) ([]Policy, *APIResponse, error)


        // ListRegionSubscriptions
        // Lists the region subscriptions for the specified tenancy.
        ListRegionSubscriptions(TenancyID string) ([]RegionSubscription, *APIResponse, error)


        // ListRegions
        // Lists all the regions offered by Oracle Bare Metal Cloud Services.
        ListRegions() ([]Region, *APIResponse, error)


        // ListSwiftPasswords
        // Lists the Swift passwords for the specified user. The returned object contains the password's OCID, but not\nthe password itself. The actual password is returned only upon creation.\n
        ListSwiftPasswords(UserID string) ([]SwiftPassword, *APIResponse, error)


        // ListUserGroupMemberships
        // Lists the `UserGroupMembership` objects in your tenancy. You must specify your tenancy's OCID\nas the value for the compartment ID\n(see [Where to Get the Tenancy's OCID and User's OCID](/Content/API/Concepts/apisigningkey.htm#five)).\nYou must also then filter the list in one of these ways:\n\n- You can limit the results to just the memberships for a given user by specifying a `userId`.\n- Similarly, you can limit the results to just the memberships for a given group by specifying a `groupId`.\n- You can set both the `userId` and `groupId` to determine if the specified user is in the specified group.\nIf the answer is no, the response is an empty list.\n
        ListUserGroupMemberships(CompartmentID string, UserID string, GroupID string, Page string, Limit int32) ([]UserGroupMembership, *APIResponse, error)


        // ListUsers
        // Lists the users in your tenancy. You must specify your tenancy's OCID as the value for the\ncompartment ID (remember that the tenancy is simply the root compartment).\nSee [Where to Get the Tenancy's OCID and User's OCID](/Content/API/Concepts/apisigningkey.htm#five).\n
        ListUsers(CompartmentID string, Page string, Limit int32) ([]User, *APIResponse, error)


        // RemoveUserFromGroup
        // Removes a user from a group by deleting the corresponding `UserGroupMembership`.
        RemoveUserFromGroup(UserGroupMembershipID string, IfMatch string) (*APIResponse, error)


        // UpdateCompartment
        // Updates the specified compartment's description.
        UpdateCompartment(CompartmentID string, UpdateCompartmentDetails UpdateCompartmentDetails, IfMatch string) (*Compartment, *APIResponse, error)


        // UpdateGroup
        // Updates the specified group.
        UpdateGroup(GroupID string, UpdateGroupDetails UpdateGroupDetails, IfMatch string) (*Group, *APIResponse, error)


        // UpdateIdentityProvider
        // Updates the specified identity provider.
        UpdateIdentityProvider(IdentityProviderID string, UpdateIdentityProviderDetails UpdateIdentityProviderDetails, IfMatch string) (*IdentityProvider, *APIResponse, error)


        // UpdateIdpGroupMapping
        // Updates the specified group mapping.
        UpdateIdpGroupMapping(IdentityProviderID string, MappingID string, UpdateIdpGroupMappingDetails UpdateIdpGroupMappingDetails, IfMatch string) (*IdpGroupMapping, *APIResponse, error)


        // UpdatePolicy
        // Updates the specified policy. You can update the description or the policy statements themselves.\n\nPolicy changes take effect typically within 10 seconds.\n
        UpdatePolicy(PolicyID string, UpdatePolicyDetails UpdatePolicyDetails, IfMatch string) (*Policy, *APIResponse, error)


        // UpdateSwiftPassword
        // Updates the specified Swift password's description.\n
        UpdateSwiftPassword(UserID string, SwiftPasswordID string, UpdateSwiftPasswordDetails UpdateSwiftPasswordDetails, IfMatch string) (*SwiftPassword, *APIResponse, error)


        // UpdateUser
        // Updates the description of the specified user.
        UpdateUser(UserID string, UpdateUserDetails UpdateUserDetails, IfMatch string) (*User, *APIResponse, error)


        // UpdateUserState
        // Updates the state of the specified user.\n
        UpdateUserState(UserID string, UpdateStateDetails UpdateStateDetails, IfMatch string) (*User, *APIResponse, error)


        // UploadApiKey
        // Uploads an API signing key for the specified user.\n\nEvery user has permission to use this operation to upload a key for *their own user ID*. An\nadministrator in your organization does not need to write a policy to give users this ability.\nTo compare, administrators who have permission to the tenancy can use this operation to upload a\nkey for any user, including themselves.\n\n**Important:** Even though you have permission to upload an API key, you might not yet\nhave permission to do much else. If you try calling an operation unrelated to your own credential\nmanagement (e.g., `ListUsers`, `LaunchInstance`) and receive an \"unauthorized\" error,\ncheck with an administrator to confirm which IAM Service group(s) you're in and what access\nyou have. Also confirm you're working in the correct compartment.\n\nAfter you send your request, the new object's `lifecycleState` will temporarily be CREATING. Before using\nthe object, first make sure its `lifecycleState` has changed to ACTIVE.\n
        UploadApiKey(UserID string, CreateApiKeyDetails CreateApiKeyDetails, OpcRetryToken string) (*ApiKey, *APIResponse, error)

}
