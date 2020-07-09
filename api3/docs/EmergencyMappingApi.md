# \EmergencyMappingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmergencymappingsGet**](EmergencyMappingApi.md#EmergencymappingsGet) | **Get** /emergencymappings/ | Get EmergencyMapping items
[**EmergencymappingsIdDelete**](EmergencyMappingApi.md#EmergencymappingsIdDelete) | **Delete** /emergencymappings/{id} | Delete a specific EmergencyMapping
[**EmergencymappingsIdGet**](EmergencyMappingApi.md#EmergencymappingsIdGet) | **Get** /emergencymappings/{id} | Get a specific EmergencyMapping
[**EmergencymappingsIdPatch**](EmergencyMappingApi.md#EmergencymappingsIdPatch) | **Patch** /emergencymappings/{id} | Change a specific EmergencyMapping
[**EmergencymappingsIdPut**](EmergencyMappingApi.md#EmergencymappingsIdPut) | **Put** /emergencymappings/{id} | Replace/change a specific EmergencyMapping
[**EmergencymappingsPost**](EmergencyMappingApi.md#EmergencymappingsPost) | **Post** /emergencymappings/ | Create a new EmergencyMapping



## EmergencymappingsGet

> []EmergencyMapping EmergencymappingsGet(ctx, optional)

Get EmergencyMapping items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmergencymappingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmergencymappingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emergencyContainerId** | **optional.String**| Filter for emergency mappings belonging to a specific emergency mapping container | 
 **resellerId** | **optional.String**| Filter for emergency mappings belonging to a specific reseller | 
 **code** | **optional.String**| Filter for mappings with a specific code (wildcards possible) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]EmergencyMapping**](EmergencyMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingsIdDelete

> EmergencymappingsIdDelete(ctx, id)

Delete a specific EmergencyMapping

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


## EmergencymappingsIdGet

> EmergencyMapping EmergencymappingsIdGet(ctx, id)

Get a specific EmergencyMapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**EmergencyMapping**](EmergencyMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingsIdPatch

> EmergencyMapping EmergencymappingsIdPatch(ctx, id, operation)

Change a specific EmergencyMapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**EmergencyMapping**](EmergencyMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingsIdPut

> EmergencyMapping EmergencymappingsIdPut(ctx, id, emergencyMapping)

Replace/change a specific EmergencyMapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**emergencyMapping** | [**EmergencyMapping**](EmergencyMapping.md)|  | 

### Return type

[**EmergencyMapping**](EmergencyMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmergencymappingsPost

> []EmergencyMapping EmergencymappingsPost(ctx, emergencyMapping)

Create a new EmergencyMapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emergencyMapping** | [**EmergencyMapping**](EmergencyMapping.md)|  | 

### Return type

[**[]EmergencyMapping**](EmergencyMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

