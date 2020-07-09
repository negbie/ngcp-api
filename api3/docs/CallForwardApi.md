# \CallForwardApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallforwardsGet**](CallForwardApi.md#CallforwardsGet) | **Get** /callforwards/ | Get CallForward items
[**CallforwardsIdDelete**](CallForwardApi.md#CallforwardsIdDelete) | **Delete** /callforwards/{id} | Delete a specific CallForward
[**CallforwardsIdGet**](CallForwardApi.md#CallforwardsIdGet) | **Get** /callforwards/{id} | Get a specific CallForward
[**CallforwardsIdPatch**](CallForwardApi.md#CallforwardsIdPatch) | **Patch** /callforwards/{id} | Change a specific CallForward
[**CallforwardsIdPut**](CallForwardApi.md#CallforwardsIdPut) | **Put** /callforwards/{id} | Replace/change a specific CallForward



## CallforwardsGet

> []CallForward CallforwardsGet(ctx, optional)

Get CallForward items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CallforwardsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CallforwardsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CallForward**](CallForward.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallforwardsIdDelete

> CallforwardsIdDelete(ctx, id)

Delete a specific CallForward

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


## CallforwardsIdGet

> CallForward CallforwardsIdGet(ctx, id)

Get a specific CallForward

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CallForward**](CallForward.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallforwardsIdPatch

> CallForward CallforwardsIdPatch(ctx, id, operation)

Change a specific CallForward

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CallForward**](CallForward.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallforwardsIdPut

> CallForward CallforwardsIdPut(ctx, id, callForward)

Replace/change a specific CallForward

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**callForward** | [**CallForward**](CallForward.md)|  | 

### Return type

[**CallForward**](CallForward.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

