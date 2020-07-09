# \VoicemailGreetingApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VoicemailgreetingsGet**](VoicemailGreetingApi.md#VoicemailgreetingsGet) | **Get** /voicemailgreetings/ | Get VoicemailGreeting items
[**VoicemailgreetingsIdDelete**](VoicemailGreetingApi.md#VoicemailgreetingsIdDelete) | **Delete** /voicemailgreetings/{id} | Delete a specific VoicemailGreeting
[**VoicemailgreetingsIdGet**](VoicemailGreetingApi.md#VoicemailgreetingsIdGet) | **Get** /voicemailgreetings/{id} | Get a specific VoicemailGreeting
[**VoicemailgreetingsIdPut**](VoicemailGreetingApi.md#VoicemailgreetingsIdPut) | **Put** /voicemailgreetings/{id} | Replace/change a specific VoicemailGreeting
[**VoicemailgreetingsPost**](VoicemailGreetingApi.md#VoicemailgreetingsPost) | **Post** /voicemailgreetings/ | Create a new VoicemailGreeting



## VoicemailgreetingsGet

> []VoicemailGreeting VoicemailgreetingsGet(ctx, optional)

Get VoicemailGreeting items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VoicemailgreetingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VoicemailgreetingsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for registrations of a specific subscriber | 
 **type_** | **optional.String**| Filter for the greeting type | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]VoicemailGreeting**](VoicemailGreeting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailgreetingsIdDelete

> VoicemailgreetingsIdDelete(ctx, id)

Delete a specific VoicemailGreeting

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


## VoicemailgreetingsIdGet

> VoicemailGreeting VoicemailgreetingsIdGet(ctx, id)

Get a specific VoicemailGreeting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**VoicemailGreeting**](VoicemailGreeting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailgreetingsIdPut

> VoicemailGreeting VoicemailgreetingsIdPut(ctx, id, voicemailGreeting)

Replace/change a specific VoicemailGreeting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**voicemailGreeting** | [**VoicemailGreeting**](VoicemailGreeting.md)|  | 

### Return type

[**VoicemailGreeting**](VoicemailGreeting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoicemailgreetingsPost

> []VoicemailGreeting VoicemailgreetingsPost(ctx, voicemailGreeting)

Create a new VoicemailGreeting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**voicemailGreeting** | [**VoicemailGreeting**](VoicemailGreeting.md)|  | 

### Return type

[**[]VoicemailGreeting**](VoicemailGreeting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

