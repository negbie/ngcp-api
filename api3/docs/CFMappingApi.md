# \CFMappingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfmappingsGet**](CFMappingApi.md#CfmappingsGet) | **Get** /cfmappings/ | Get CFMapping items
[**CfmappingsIdGet**](CFMappingApi.md#CfmappingsIdGet) | **Get** /cfmappings/{id} | Get a specific CFMapping
[**CfmappingsIdPatch**](CFMappingApi.md#CfmappingsIdPatch) | **Patch** /cfmappings/{id} | Change a specific CFMapping
[**CfmappingsIdPut**](CFMappingApi.md#CfmappingsIdPut) | **Put** /cfmappings/{id} | Replace/change a specific CFMapping



## CfmappingsGet

> []CfMapping CfmappingsGet(ctx, optional)

Get CFMapping items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CfmappingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CfmappingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CfMapping**](CFMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfmappingsIdGet

> CfMapping CfmappingsIdGet(ctx, id)

Get a specific CFMapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CfMapping**](CFMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfmappingsIdPatch

> CfMapping CfmappingsIdPatch(ctx, id, operation)

Change a specific CFMapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CfMapping**](CFMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfmappingsIdPut

> CfMapping CfmappingsIdPut(ctx, id, cfMapping)

Replace/change a specific CFMapping

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**cfMapping** | [**CfMapping**](CfMapping.md)|  | 

### Return type

[**CfMapping**](CFMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

