# \SoundSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoundsetsGet**](SoundSetApi.md#SoundsetsGet) | **Get** /soundsets/ | Get SoundSet items
[**SoundsetsIdDelete**](SoundSetApi.md#SoundsetsIdDelete) | **Delete** /soundsets/{id} | Delete a specific SoundSet
[**SoundsetsIdGet**](SoundSetApi.md#SoundsetsIdGet) | **Get** /soundsets/{id} | Get a specific SoundSet
[**SoundsetsIdPatch**](SoundSetApi.md#SoundsetsIdPatch) | **Patch** /soundsets/{id} | Change a specific SoundSet
[**SoundsetsIdPut**](SoundSetApi.md#SoundsetsIdPut) | **Put** /soundsets/{id} | Replace/change a specific SoundSet
[**SoundsetsPost**](SoundSetApi.md#SoundsetsPost) | **Post** /soundsets/ | Create a new SoundSet



## SoundsetsGet

> []SoundSet SoundsetsGet(ctx, optional)

Get SoundSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoundsetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SoundsetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **optional.String**| Filter for sound sets of a specific customer | 
 **resellerId** | **optional.String**| Filter for sound sets of a specific reseller | 
 **name** | **optional.String**| Filter for sound sets with a specific name (wildcard pattern allowed) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SoundSet**](SoundSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundsetsIdDelete

> SoundsetsIdDelete(ctx, id)

Delete a specific SoundSet

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


## SoundsetsIdGet

> SoundSet SoundsetsIdGet(ctx, id)

Get a specific SoundSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SoundSet**](SoundSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundsetsIdPatch

> SoundSet SoundsetsIdPatch(ctx, id, operation)

Change a specific SoundSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**SoundSet**](SoundSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundsetsIdPut

> SoundSet SoundsetsIdPut(ctx, id, soundSet)

Replace/change a specific SoundSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**soundSet** | [**SoundSet**](SoundSet.md)|  | 

### Return type

[**SoundSet**](SoundSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundsetsPost

> []SoundSet SoundsetsPost(ctx, soundSet)

Create a new SoundSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soundSet** | [**SoundSet**](SoundSet.md)|  | 

### Return type

[**[]SoundSet**](SoundSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

