# \NcosLevelApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NcoslevelsGet**](NcosLevelApi.md#NcoslevelsGet) | **Get** /ncoslevels/ | Get NcosLevel items
[**NcoslevelsIdDelete**](NcosLevelApi.md#NcoslevelsIdDelete) | **Delete** /ncoslevels/{id} | Delete a specific NcosLevel
[**NcoslevelsIdGet**](NcosLevelApi.md#NcoslevelsIdGet) | **Get** /ncoslevels/{id} | Get a specific NcosLevel
[**NcoslevelsIdPatch**](NcosLevelApi.md#NcoslevelsIdPatch) | **Patch** /ncoslevels/{id} | Change a specific NcosLevel
[**NcoslevelsIdPut**](NcosLevelApi.md#NcoslevelsIdPut) | **Put** /ncoslevels/{id} | Replace/change a specific NcosLevel
[**NcoslevelsPost**](NcosLevelApi.md#NcoslevelsPost) | **Post** /ncoslevels/ | Create a new NcosLevel



## NcoslevelsGet

> []NcosLevel NcoslevelsGet(ctx, optional)

Get NcosLevel items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NcoslevelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NcoslevelsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for ncos levels belonging to a specific reseller | 
 **level** | **optional.String**| Filter for levels matching the given pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]NcosLevel**](NcosLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslevelsIdDelete

> NcoslevelsIdDelete(ctx, id)

Delete a specific NcosLevel

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


## NcoslevelsIdGet

> NcosLevel NcoslevelsIdGet(ctx, id)

Get a specific NcosLevel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**NcosLevel**](NcosLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslevelsIdPatch

> NcosLevel NcoslevelsIdPatch(ctx, id, operation)

Change a specific NcosLevel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**NcosLevel**](NcosLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslevelsIdPut

> NcosLevel NcoslevelsIdPut(ctx, id, ncosLevel)

Replace/change a specific NcosLevel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**ncosLevel** | [**NcosLevel**](NcosLevel.md)|  | 

### Return type

[**NcosLevel**](NcosLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslevelsPost

> []NcosLevel NcoslevelsPost(ctx, ncosLevel)

Create a new NcosLevel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ncosLevel** | [**NcosLevel**](NcosLevel.md)|  | 

### Return type

[**[]NcosLevel**](NcosLevel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

