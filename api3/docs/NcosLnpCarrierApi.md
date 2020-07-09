# \NcosLnpCarrierApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NcoslnpcarriersGet**](NcosLnpCarrierApi.md#NcoslnpcarriersGet) | **Get** /ncoslnpcarriers/ | Get NcosLnpCarrier items
[**NcoslnpcarriersIdDelete**](NcosLnpCarrierApi.md#NcoslnpcarriersIdDelete) | **Delete** /ncoslnpcarriers/{id} | Delete a specific NcosLnpCarrier
[**NcoslnpcarriersIdGet**](NcosLnpCarrierApi.md#NcoslnpcarriersIdGet) | **Get** /ncoslnpcarriers/{id} | Get a specific NcosLnpCarrier
[**NcoslnpcarriersIdPatch**](NcosLnpCarrierApi.md#NcoslnpcarriersIdPatch) | **Patch** /ncoslnpcarriers/{id} | Change a specific NcosLnpCarrier
[**NcoslnpcarriersIdPut**](NcosLnpCarrierApi.md#NcoslnpcarriersIdPut) | **Put** /ncoslnpcarriers/{id} | Replace/change a specific NcosLnpCarrier
[**NcoslnpcarriersPost**](NcosLnpCarrierApi.md#NcoslnpcarriersPost) | **Post** /ncoslnpcarriers/ | Create a new NcosLnpCarrier



## NcoslnpcarriersGet

> []NcosLnpCarrier NcoslnpcarriersGet(ctx, optional)

Get NcosLnpCarrier items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NcoslnpcarriersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NcoslnpcarriersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ncosLevelId** | **optional.String**| Filter for NCOS LNP entries belonging to a specific NCOS level. | 
 **carrierId** | **optional.String**| Filter for NCOS LNP entries belonging to a specific LNP carrier. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]NcosLnpCarrier**](NcosLnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslnpcarriersIdDelete

> NcoslnpcarriersIdDelete(ctx, id)

Delete a specific NcosLnpCarrier

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


## NcoslnpcarriersIdGet

> NcosLnpCarrier NcoslnpcarriersIdGet(ctx, id)

Get a specific NcosLnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**NcosLnpCarrier**](NcosLnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslnpcarriersIdPatch

> NcosLnpCarrier NcoslnpcarriersIdPatch(ctx, id, operation)

Change a specific NcosLnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**NcosLnpCarrier**](NcosLnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslnpcarriersIdPut

> NcosLnpCarrier NcoslnpcarriersIdPut(ctx, id, ncosLnpCarrier)

Replace/change a specific NcosLnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**ncosLnpCarrier** | [**NcosLnpCarrier**](NcosLnpCarrier.md)|  | 

### Return type

[**NcosLnpCarrier**](NcosLnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NcoslnpcarriersPost

> []NcosLnpCarrier NcoslnpcarriersPost(ctx, ncosLnpCarrier)

Create a new NcosLnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ncosLnpCarrier** | [**NcosLnpCarrier**](NcosLnpCarrier.md)|  | 

### Return type

[**[]NcosLnpCarrier**](NcosLnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

