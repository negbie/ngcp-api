# \UpnRewriteSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpnrewritesetsGet**](UpnRewriteSetApi.md#UpnrewritesetsGet) | **Get** /upnrewritesets/ | Get UpnRewriteSet items
[**UpnrewritesetsIdDelete**](UpnRewriteSetApi.md#UpnrewritesetsIdDelete) | **Delete** /upnrewritesets/{id} | Delete a specific UpnRewriteSet
[**UpnrewritesetsIdGet**](UpnRewriteSetApi.md#UpnrewritesetsIdGet) | **Get** /upnrewritesets/{id} | Get a specific UpnRewriteSet
[**UpnrewritesetsIdPatch**](UpnRewriteSetApi.md#UpnrewritesetsIdPatch) | **Patch** /upnrewritesets/{id} | Change a specific UpnRewriteSet
[**UpnrewritesetsIdPut**](UpnRewriteSetApi.md#UpnrewritesetsIdPut) | **Put** /upnrewritesets/{id} | Replace/change a specific UpnRewriteSet
[**UpnrewritesetsPost**](UpnRewriteSetApi.md#UpnrewritesetsPost) | **Post** /upnrewritesets/ | Create a new UpnRewriteSet



## UpnrewritesetsGet

> []UpnRewriteSet UpnrewritesetsGet(ctx, optional)

Get UpnRewriteSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpnrewritesetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpnrewritesetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for trusted sources of a specific subscriber | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]UpnRewriteSet**](UpnRewriteSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpnrewritesetsIdDelete

> UpnrewritesetsIdDelete(ctx, id)

Delete a specific UpnRewriteSet

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


## UpnrewritesetsIdGet

> UpnRewriteSet UpnrewritesetsIdGet(ctx, id)

Get a specific UpnRewriteSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**UpnRewriteSet**](UpnRewriteSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpnrewritesetsIdPatch

> UpnRewriteSet UpnrewritesetsIdPatch(ctx, id, operation)

Change a specific UpnRewriteSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**UpnRewriteSet**](UpnRewriteSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpnrewritesetsIdPut

> UpnRewriteSet UpnrewritesetsIdPut(ctx, id, upnRewriteSet)

Replace/change a specific UpnRewriteSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**upnRewriteSet** | [**UpnRewriteSet**](UpnRewriteSet.md)|  | 

### Return type

[**UpnRewriteSet**](UpnRewriteSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpnrewritesetsPost

> []UpnRewriteSet UpnrewritesetsPost(ctx, upnRewriteSet)

Create a new UpnRewriteSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**upnRewriteSet** | [**UpnRewriteSet**](UpnRewriteSet.md)|  | 

### Return type

[**[]UpnRewriteSet**](UpnRewriteSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

