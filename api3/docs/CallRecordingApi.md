# \CallRecordingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallrecordingsGet**](CallRecordingApi.md#CallrecordingsGet) | **Get** /callrecordings/ | Get CallRecording items
[**CallrecordingsIdDelete**](CallRecordingApi.md#CallrecordingsIdDelete) | **Delete** /callrecordings/{id} | Delete a specific CallRecording
[**CallrecordingsIdGet**](CallRecordingApi.md#CallrecordingsIdGet) | **Get** /callrecordings/{id} | Get a specific CallRecording



## CallrecordingsGet

> []CallRecording CallrecordingsGet(ctx, optional)

Get CallRecording items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CallrecordingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CallrecordingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for callrecordings belonging to a specific reseller | 
 **status** | **optional.String**| Filter for callrecordings with a specific status | 
 **subscriberId** | **optional.String**| Filter for callrecordings where the subscriber with the given id is involved. | 
 **tz** | **optional.String**| Format start_time according to the optional time zone provided here, e.g. Europe/Berlin. | 
 **forceDelete** | **optional.String**| Force callrecording info deletion from database despite callrecording files deletion errors. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CallRecording**](CallRecording.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CallrecordingsIdDelete

> CallrecordingsIdDelete(ctx, id)

Delete a specific CallRecording

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


## CallrecordingsIdGet

> CallRecording CallrecordingsIdGet(ctx, id)

Get a specific CallRecording

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CallRecording**](CallRecording.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

