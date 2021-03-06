# \SoundHandleApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoundhandlesGet**](SoundHandleApi.md#SoundhandlesGet) | **Get** /soundhandles/ | Get SoundHandle items
[**SoundhandlesIdGet**](SoundHandleApi.md#SoundhandlesIdGet) | **Get** /soundhandles/{id} | Get a specific SoundHandle



## SoundhandlesGet

> []SoundHandle SoundhandlesGet(ctx, optional)

Get SoundHandle items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoundhandlesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SoundhandlesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**| Filter for sound handles of a specific group | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SoundHandle**](SoundHandle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundhandlesIdGet

> SoundHandle SoundhandlesIdGet(ctx, id)

Get a specific SoundHandle

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SoundHandle**](SoundHandle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

