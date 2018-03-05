// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// File Storage Service API
//
// APIs for OCI file storage service.
//

package ffsw

import(
    "github.com/oracle/oci-go-sdk/common"
    "context"
    "fmt"
    "net/http"
)

//FileStorageClientData a client for FileStorage
type FileStorageClientData struct {
    common.BaseClient
    config *common.ConfigurationProvider
}


// NewFileStorageClientWithConfigurationProvider Creates a new default FileStorage client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewFileStorageClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (FileStorageClient, error){
    baseClient, err := common.NewClientWithConfig(configProvider)
    if err != nil {
        return nil, err
    }

    client := FileStorageClientData{BaseClient: baseClient}
    client.BasePath = "20171215"
    err = client.setConfigurationProvider(configProvider)
    if err != nil {
        return nil, err
    }
    return &client, err
}


// GetBaseClient get the BaseClient object of this client
func (client *FileStorageClientData) GetBaseClient() (*common.BaseClient)  {
     return &client.BaseClient
}

// SetRegion overrides the region of this client.
func (client *FileStorageClientData) SetRegion(region string)  {
    client.Host = fmt.Sprintf(common.DefaultHostURLTemplate, "ffsw", region)
}


// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *FileStorageClientData) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
    if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
        return err
    }

    // Error has been checked already
    region, _ := configProvider.Region()
    client.config = &configProvider
    client.SetRegion(region)
    return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *FileStorageClientData) ConfigurationProvider() *common.ConfigurationProvider {
    return client.config
}






 // CreateExport Creates a new export in the specifed export set, path, and
 // file system.
func(client FileStorageClientData) CreateExport(ctx context.Context, request CreateExportRequest) (response CreateExportResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/exports", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // CreateFileSystem Create a new file system in the specified compartment and
 // availability domain. File systems in one availability domain
 // can be mounted by instances in another availability domain,
 // but they may see higher latencies than instances in the same
 // availability domain as the file system.
 // Once a file system is created it can be assocated with a mount
 // target and then mounted by instances that can connect to the
 // mount target's IP address. A file system can be assocated with
 // more than one mount target at a time.
 // For information about access control and compartments, see
 // Overview of the IAM Service ({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
 // For information about Availability Domains, see Regions and
 // Availability Domains ({{DOC_SERVER_URL}}/Content/General/Concepts/regions.htm).
 // To get a list of Availability Domains, use the
 // `ListAvailabilityDomains` operation in the Identity and Access
 // Management Service API.
 // All Oracle Bare Metal Cloud Services resources, including
 // file systems, get an Oracle-assigned, unique ID called an Oracle
 // Cloud Identifier (OCID).  When you create a resource, you can
 // find its OCID in the response. You can also retrieve a
 // resource's OCID by using a List API operation on that resource
 // type, or by viewing the resource in the Console.
func(client FileStorageClientData) CreateFileSystem(ctx context.Context, request CreateFileSystemRequest) (response CreateFileSystemResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/fileSystems", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // CreateMountTarget Create a new mount target in the specified compartment and
 // subnet. A file system can only be assocated with a mount
 // target if they are both in the availablity domain. Instances
 // can connect to mount targets in another availablity domain but
 // they may see higher latencies than instances in the same
 // availability domain as the mount target.
 // Mount targets have one or more private IP addresses that can
 // be used as the host portion of remotetarget parameters in
 // client mount commands. These private IP addresses are listed
 // in privateIpIds property of the mount target and are HA. Mount
 // targets also consume additional IP addresses in their subnet.
 // For information about access control and compartments, see
 // Overview of the IAM
 // Service ({{DOC_SERVER_URL}}/Content/Identity/Concepts/overview.htm).
 // For information about Availability Domains, see Regions and
 // Availability Domains ({{DOC_SERVER_URL}}/Content/General/Concepts/regions.htm).
 // To get a list of Availability Domains, use the
 // `ListAvailabilityDomains` operation in the Identity and Access
 // Management Service API.
 // All Oracle Bare Metal Cloud Services resources, including
 // mount targets, get an Oracle-assigned, unique ID called an
 // Oracle Cloud Identifier (OCID).  When you create a resource,
 // you can find its OCID in the response. You can also retrieve a
 // resource's OCID by using a List API operation on that resource
 // type, or by viewing the resource in the Console.
func(client FileStorageClientData) CreateMountTarget(ctx context.Context, request CreateMountTargetRequest) (response CreateMountTargetResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/mountTargets", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // CreateSnapshot Creates a new snapshot of the specified file system. The
 // snapshot will be accessible at `.snapshot/<name>`.
func(client FileStorageClientData) CreateSnapshot(ctx context.Context, request CreateSnapshotRequest) (response CreateSnapshotResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPost, "/snapshots", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // DeleteExport Delete the specified export.
func(client FileStorageClientData) DeleteExport(ctx context.Context, request DeleteExportRequest) (response DeleteExportResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/exports/{exportId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // DeleteFileSystem Delete the specified file system. The file system must not be
 // referenced by any non-deleted export resources. Deleting a
 // file system also deletes all of its snapshots.
func(client FileStorageClientData) DeleteFileSystem(ctx context.Context, request DeleteFileSystemRequest) (response DeleteFileSystemResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/fileSystems/{fileSystemId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // DeleteMountTarget Delete the specified mount target. This will also delete the
 // mount target's VNICs.
func(client FileStorageClientData) DeleteMountTarget(ctx context.Context, request DeleteMountTargetRequest) (response DeleteMountTargetResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/mountTargets/{mountTargetId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // DeleteSnapshot Delete the specified snapshot.
func(client FileStorageClientData) DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest) (response DeleteSnapshotResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodDelete, "/snapshots/{snapshotId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // GetExport Gets the specified export's information.
func(client FileStorageClientData) GetExport(ctx context.Context, request GetExportRequest) (response GetExportResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exports/{exportId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // GetExportSet Gets the specified export set's information.
func(client FileStorageClientData) GetExportSet(ctx context.Context, request GetExportSetRequest) (response GetExportSetResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exportSets/{exportSetId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // GetFileSystem Gets the specified file system's information.
func(client FileStorageClientData) GetFileSystem(ctx context.Context, request GetFileSystemRequest) (response GetFileSystemResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fileSystems/{fileSystemId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // GetMountTarget Gets the specified mount targets's information.
func(client FileStorageClientData) GetMountTarget(ctx context.Context, request GetMountTargetRequest) (response GetMountTargetResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/mountTargets/{mountTargetId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // GetSnapshot Gets the specified snapshot's information.
func(client FileStorageClientData) GetSnapshot(ctx context.Context, request GetSnapshotRequest) (response GetSnapshotResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/snapshots/{snapshotId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // ListExportSets List the export set resources in the specified compartment.
func(client FileStorageClientData) ListExportSets(ctx context.Context, request ListExportSetsRequest) (response ListExportSetsResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exportSets", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // ListExports List the export resources in the specified compartment. Must
 // also specify an export set and / or a file system.
func(client FileStorageClientData) ListExports(ctx context.Context, request ListExportsRequest) (response ListExportsResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/exports", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // ListFileSystems List the file system resources in the specified compartment.
func(client FileStorageClientData) ListFileSystems(ctx context.Context, request ListFileSystemsRequest) (response ListFileSystemsResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fileSystems", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // ListLockOwners List the lock owners in a given file system.
func(client FileStorageClientData) ListLockOwners(ctx context.Context, request ListLockOwnersRequest) (response ListLockOwnersResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/fileSystems/{fileSystemId}/lockOwners", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // ListMountTargets List the mount target resources in the specified compartment.
func(client FileStorageClientData) ListMountTargets(ctx context.Context, request ListMountTargetsRequest) (response ListMountTargetsResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/mountTargets", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // ListSnapshots List the snapshots of the specified file system.
func(client FileStorageClientData) ListSnapshots(ctx context.Context, request ListSnapshotsRequest) (response ListSnapshotsResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodGet, "/snapshots", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // UpdateExportSet Update the specified export set's information.
func(client FileStorageClientData) UpdateExportSet(ctx context.Context, request UpdateExportSetRequest) (response UpdateExportSetResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/exportSets/{exportSetId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // UpdateFileSystem Update the specified file system's information.
func(client FileStorageClientData) UpdateFileSystem(ctx context.Context, request UpdateFileSystemRequest) (response UpdateFileSystemResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/fileSystems/{fileSystemId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}




 // UpdateMountTarget Update the specified mount targets's information.
func(client FileStorageClientData) UpdateMountTarget(ctx context.Context, request UpdateMountTargetRequest) (response UpdateMountTargetResponse,  err error) {
        httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStruct(http.MethodPut, "/mountTargets/{mountTargetId}", request)
        if err != nil {
            return
        }


    httpResponse, err := client.Call(ctx, &httpRequest)
    defer common.CloseBodyIfValid(httpResponse)
    response.RawResponse = httpResponse
    if err != nil {
        return
    }

        err = common.UnmarshalResponse(httpResponse, &response)
    return
}



// FileStorageClient is client interface for FileStorage
type FileStorageClient interface {
    GetBaseClient() (*common.BaseClient)
    ConfigurationProvider() (*common.ConfigurationProvider)
    CreateExport(ctx context.Context, request CreateExportRequest) (response CreateExportResponse,  err error) 
    CreateFileSystem(ctx context.Context, request CreateFileSystemRequest) (response CreateFileSystemResponse,  err error) 
    CreateMountTarget(ctx context.Context, request CreateMountTargetRequest) (response CreateMountTargetResponse,  err error) 
    CreateSnapshot(ctx context.Context, request CreateSnapshotRequest) (response CreateSnapshotResponse,  err error) 
    DeleteExport(ctx context.Context, request DeleteExportRequest) (response DeleteExportResponse,  err error) 
    DeleteFileSystem(ctx context.Context, request DeleteFileSystemRequest) (response DeleteFileSystemResponse,  err error) 
    DeleteMountTarget(ctx context.Context, request DeleteMountTargetRequest) (response DeleteMountTargetResponse,  err error) 
    DeleteSnapshot(ctx context.Context, request DeleteSnapshotRequest) (response DeleteSnapshotResponse,  err error) 
    GetExport(ctx context.Context, request GetExportRequest) (response GetExportResponse,  err error) 
    GetExportSet(ctx context.Context, request GetExportSetRequest) (response GetExportSetResponse,  err error) 
    GetFileSystem(ctx context.Context, request GetFileSystemRequest) (response GetFileSystemResponse,  err error) 
    GetMountTarget(ctx context.Context, request GetMountTargetRequest) (response GetMountTargetResponse,  err error) 
    GetSnapshot(ctx context.Context, request GetSnapshotRequest) (response GetSnapshotResponse,  err error) 
    ListExportSets(ctx context.Context, request ListExportSetsRequest) (response ListExportSetsResponse,  err error) 
    ListExports(ctx context.Context, request ListExportsRequest) (response ListExportsResponse,  err error) 
    ListFileSystems(ctx context.Context, request ListFileSystemsRequest) (response ListFileSystemsResponse,  err error) 
    ListLockOwners(ctx context.Context, request ListLockOwnersRequest) (response ListLockOwnersResponse,  err error) 
    ListMountTargets(ctx context.Context, request ListMountTargetsRequest) (response ListMountTargetsResponse,  err error) 
    ListSnapshots(ctx context.Context, request ListSnapshotsRequest) (response ListSnapshotsResponse,  err error) 
    UpdateExportSet(ctx context.Context, request UpdateExportSetRequest) (response UpdateExportSetResponse,  err error) 
    UpdateFileSystem(ctx context.Context, request UpdateFileSystemRequest) (response UpdateFileSystemResponse,  err error) 
    UpdateMountTarget(ctx context.Context, request UpdateMountTargetRequest) (response UpdateMountTargetResponse,  err error) 
}
