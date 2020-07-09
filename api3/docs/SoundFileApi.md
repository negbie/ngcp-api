# \SoundFileApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoundfilesGet**](SoundFileApi.md#SoundfilesGet) | **Get** /soundfiles/ | Get SoundFile items
[**SoundfilesIdDelete**](SoundFileApi.md#SoundfilesIdDelete) | **Delete** /soundfiles/{id} | Delete a specific SoundFile
[**SoundfilesIdGet**](SoundFileApi.md#SoundfilesIdGet) | **Get** /soundfiles/{id} | Get a specific SoundFile
[**SoundfilesIdPut**](SoundFileApi.md#SoundfilesIdPut) | **Put** /soundfiles/{id} | Replace/change a specific SoundFile
[**SoundfilesPost**](SoundFileApi.md#SoundfilesPost) | **Post** /soundfiles/ | Create a new SoundFile



## SoundfilesGet

> []SoundFile SoundfilesGet(ctx, optional)

Get SoundFile items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoundfilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SoundfilesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setId** | **optional.String**| Filter for sound files of a specific sound set | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SoundFile**](SoundFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundfilesIdDelete

> SoundfilesIdDelete(ctx, id)

Delete a specific SoundFile

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


## SoundfilesIdGet

> SoundFile SoundfilesIdGet(ctx, id)

Get a specific SoundFile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SoundFile**](SoundFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundfilesIdPut

> SoundFile SoundfilesIdPut(ctx, id, soundFile)

Replace/change a specific SoundFile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**soundFile** | [**SoundFile**](SoundFile.md)|  | 

### Return type

[**SoundFile**](SoundFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoundfilesPost

> []SoundFile SoundfilesPost(ctx, soundFile)

Create a new SoundFile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soundFile** | [**SoundFile**](SoundFile.md)|  | 

### Return type

[**[]SoundFile**](SoundFile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

