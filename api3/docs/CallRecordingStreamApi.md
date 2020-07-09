# \CallRecordingStreamApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallrecordingstreamsGet**](CallRecordingStreamApi.md#CallrecordingstreamsGet) | **Get** /callrecordingstreams/ | Get CallRecordingStream items
[**CallrecordingstreamsIdDelete**](CallRecordingStreamApi.md#CallrecordingstreamsIdDelete) | **Delete** /callrecordingstreams/{id} | Delete a specific CallRecordingStream
[**CallrecordingstreamsIdGet**](CallRecordingStreamApi.md#CallrecordingstreamsIdGet) | **Get** /callrecordingstreams/{id} | Get a specific CallRecordingStream



## CallrecordingstreamsGet

> []CallRecordingStream CallrecordingstreamsGet(ctx, optional)

Get CallRecordingStream items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CallrecordingstreamsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CallrecordingstreamsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordingId** | **optional.String**| Filter for callrecording streams belonging to a specific recording session. | 
 **type_** | **optional.String**| Filter for callrecording streams with a specific type (\&quot;single\&quot; or \&quot;mixed\&quot;) | 
 **tz** | **optional.String**| Format start_time according to the optional time zone provided here, e.g. Europe/Berlin. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CallRecordingStream**](CallRecordingStream.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallrecordingstreamsIdDelete

> CallrecordingstreamsIdDelete(ctx, id)

Delete a specific CallRecordingStream

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


## CallrecordingstreamsIdGet

> CallRecordingStream CallrecordingstreamsIdGet(ctx, id)

Get a specific CallRecordingStream

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CallRecordingStream**](CallRecordingStream.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

