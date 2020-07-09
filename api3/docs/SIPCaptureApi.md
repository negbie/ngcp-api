# \SIPCaptureApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SipcapturesGet**](SIPCaptureApi.md#SipcapturesGet) | **Get** /sipcaptures/ | Get SIPCapture items
[**SipcapturesIdGet**](SIPCaptureApi.md#SipcapturesIdGet) | **Get** /sipcaptures/{id} | Get a specific SIPCapture



## SipcapturesGet

> []SipCapture SipcapturesGet(ctx, optional)

Get SIPCapture items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SipcapturesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SipcapturesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callId** | **optional.String**| Filter for a particular call_id | 
 **startGe** | **optional.String**| Filter for data starting greater or equal the specified time stamp. | 
 **startLe** | **optional.String**| Filter for data starting lower or equal the specified time stamp. | 
 **method** | **optional.String**| Filter for a particular SIP method | 
 **subscriberId** | **optional.String**| End time of the captured SIP data | 
 **tz** | **optional.String**| Format start_time according to the optional time zone provided here, e.g. Europe/Berlin. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SipCapture**](SIPCapture.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SipcapturesIdGet

> SipCapture SipcapturesIdGet(ctx, id)

Get a specific SIPCapture

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SipCapture**](SIPCapture.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

