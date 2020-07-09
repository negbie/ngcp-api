# \CCMapEntryApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CcmapentriesGet**](CCMapEntryApi.md#CcmapentriesGet) | **Get** /ccmapentries/ | Get CCMapEntry items
[**CcmapentriesIdDelete**](CCMapEntryApi.md#CcmapentriesIdDelete) | **Delete** /ccmapentries/{id} | Delete a specific CCMapEntry
[**CcmapentriesIdGet**](CCMapEntryApi.md#CcmapentriesIdGet) | **Get** /ccmapentries/{id} | Get a specific CCMapEntry
[**CcmapentriesIdPatch**](CCMapEntryApi.md#CcmapentriesIdPatch) | **Patch** /ccmapentries/{id} | Change a specific CCMapEntry
[**CcmapentriesIdPut**](CCMapEntryApi.md#CcmapentriesIdPut) | **Put** /ccmapentries/{id} | Replace/change a specific CCMapEntry



## CcmapentriesGet

> []CcMapEntry CcmapentriesGet(ctx, optional)

Get CCMapEntry items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CcmapentriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CcmapentriesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CcMapEntry**](CCMapEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CcmapentriesIdDelete

> CcmapentriesIdDelete(ctx, id)

Delete a specific CCMapEntry

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


## CcmapentriesIdGet

> CcMapEntry CcmapentriesIdGet(ctx, id)

Get a specific CCMapEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CcMapEntry**](CCMapEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CcmapentriesIdPatch

> CcMapEntry CcmapentriesIdPatch(ctx, id, operation)

Change a specific CCMapEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CcMapEntry**](CCMapEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CcmapentriesIdPut

> CcMapEntry CcmapentriesIdPut(ctx, id, ccMapEntry)

Replace/change a specific CCMapEntry

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**ccMapEntry** | [**CcMapEntry**](CcMapEntry.md)|  | 

### Return type

[**CcMapEntry**](CCMapEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

