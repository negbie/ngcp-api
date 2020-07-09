# \EmergencyMappingContainerApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmergencymappingcontainersGet**](EmergencyMappingContainerApi.md#EmergencymappingcontainersGet) | **Get** /emergencymappingcontainers/ | Get EmergencyMappingContainer items
[**EmergencymappingcontainersIdDelete**](EmergencyMappingContainerApi.md#EmergencymappingcontainersIdDelete) | **Delete** /emergencymappingcontainers/{id} | Delete a specific EmergencyMappingContainer
[**EmergencymappingcontainersIdGet**](EmergencyMappingContainerApi.md#EmergencymappingcontainersIdGet) | **Get** /emergencymappingcontainers/{id} | Get a specific EmergencyMappingContainer
[**EmergencymappingcontainersIdPatch**](EmergencyMappingContainerApi.md#EmergencymappingcontainersIdPatch) | **Patch** /emergencymappingcontainers/{id} | Change a specific EmergencyMappingContainer
[**EmergencymappingcontainersIdPut**](EmergencyMappingContainerApi.md#EmergencymappingcontainersIdPut) | **Put** /emergencymappingcontainers/{id} | Replace/change a specific EmergencyMappingContainer
[**EmergencymappingcontainersPost**](EmergencyMappingContainerApi.md#EmergencymappingcontainersPost) | **Post** /emergencymappingcontainers/ | Create a new EmergencyMappingContainer



## EmergencymappingcontainersGet

> []EmergencyMappingContainer EmergencymappingcontainersGet(ctx, optional)

Get EmergencyMappingContainer items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmergencymappingcontainersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmergencymappingcontainersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter for emergency mapping containers with a specific name (wildcards possible) | 
 **resellerId** | **optional.String**| Filter for emergency mapping containers for a specific reseller | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]EmergencyMappingContainer**](EmergencyMappingContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingcontainersIdDelete

> EmergencymappingcontainersIdDelete(ctx, id)

Delete a specific EmergencyMappingContainer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingcontainersIdGet

> EmergencyMappingContainer EmergencymappingcontainersIdGet(ctx, id)

Get a specific EmergencyMappingContainer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**EmergencyMappingContainer**](EmergencyMappingContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingcontainersIdPatch

> EmergencyMappingContainer EmergencymappingcontainersIdPatch(ctx, id, operation)

Change a specific EmergencyMappingContainer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**EmergencyMappingContainer**](EmergencyMappingContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingcontainersIdPut

> EmergencyMappingContainer EmergencymappingcontainersIdPut(ctx, id, emergencyMappingContainer)

Replace/change a specific EmergencyMappingContainer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**emergencyMappingContainer** | [**EmergencyMappingContainer**](EmergencyMappingContainer.md)|  | 

### Return type

[**EmergencyMappingContainer**](EmergencyMappingContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingcontainersPost

> []EmergencyMappingContainer EmergencymappingcontainersPost(ctx, emergencyMappingContainer)

Create a new EmergencyMappingContainer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emergencyMappingContainer** | [**EmergencyMappingContainer**](EmergencyMappingContainer.md)|  | 

### Return type

[**[]EmergencyMappingContainer**](EmergencyMappingContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

