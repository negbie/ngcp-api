# \RtcNetworkApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RtcnetworksGet**](RtcNetworkApi.md#RtcnetworksGet) | **Get** /rtcnetworks/ | Get RtcNetwork items
[**RtcnetworksIdGet**](RtcNetworkApi.md#RtcnetworksIdGet) | **Get** /rtcnetworks/{id} | Get a specific RtcNetwork
[**RtcnetworksIdPatch**](RtcNetworkApi.md#RtcnetworksIdPatch) | **Patch** /rtcnetworks/{id} | Change a specific RtcNetwork
[**RtcnetworksIdPut**](RtcNetworkApi.md#RtcnetworksIdPut) | **Put** /rtcnetworks/{id} | Replace/change a specific RtcNetwork



## RtcnetworksGet

> []RtcNetwork RtcnetworksGet(ctx, optional)

Get RtcNetwork items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RtcnetworksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RtcnetworksGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]RtcNetwork**](RtcNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RtcnetworksIdGet

> RtcNetwork RtcnetworksIdGet(ctx, id)

Get a specific RtcNetwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**RtcNetwork**](RtcNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RtcnetworksIdPatch

> RtcNetwork RtcnetworksIdPatch(ctx, id, operation)

Change a specific RtcNetwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**RtcNetwork**](RtcNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RtcnetworksIdPut

> RtcNetwork RtcnetworksIdPut(ctx, id, rtcNetwork)

Replace/change a specific RtcNetwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**rtcNetwork** | [**RtcNetwork**](RtcNetwork.md)|  | 

### Return type

[**RtcNetwork**](RtcNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

