# \NcosPatternApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NcospatternsGet**](NcosPatternApi.md#NcospatternsGet) | **Get** /ncospatterns/ | Get NcosPattern items
[**NcospatternsIdDelete**](NcosPatternApi.md#NcospatternsIdDelete) | **Delete** /ncospatterns/{id} | Delete a specific NcosPattern
[**NcospatternsIdGet**](NcosPatternApi.md#NcospatternsIdGet) | **Get** /ncospatterns/{id} | Get a specific NcosPattern
[**NcospatternsIdPatch**](NcosPatternApi.md#NcospatternsIdPatch) | **Patch** /ncospatterns/{id} | Change a specific NcosPattern
[**NcospatternsIdPut**](NcosPatternApi.md#NcospatternsIdPut) | **Put** /ncospatterns/{id} | Replace/change a specific NcosPattern
[**NcospatternsPost**](NcosPatternApi.md#NcospatternsPost) | **Post** /ncospatterns/ | Create a new NcosPattern



## NcospatternsGet

> []NcosPattern NcospatternsGet(ctx, optional)

Get NcosPattern items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NcospatternsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NcospatternsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ncosLevelId** | **optional.String**| Filter for NCOS patterns belonging to a specific NCOS level. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]NcosPattern**](NcosPattern.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcospatternsIdDelete

> NcospatternsIdDelete(ctx, id)

Delete a specific NcosPattern

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


## NcospatternsIdGet

> NcosPattern NcospatternsIdGet(ctx, id)

Get a specific NcosPattern

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**NcosPattern**](NcosPattern.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcospatternsIdPatch

> NcosPattern NcospatternsIdPatch(ctx, id, operation)

Change a specific NcosPattern

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**NcosPattern**](NcosPattern.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcospatternsIdPut

> NcosPattern NcospatternsIdPut(ctx, id, ncosPattern)

Replace/change a specific NcosPattern

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**ncosPattern** | [**NcosPattern**](NcosPattern.md)|  | 

### Return type

[**NcosPattern**](NcosPattern.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcospatternsPost

> []NcosPattern NcospatternsPost(ctx, ncosPattern)

Create a new NcosPattern

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ncosPattern** | [**NcosPattern**](NcosPattern.md)|  | 

### Return type

[**[]NcosPattern**](NcosPattern.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

