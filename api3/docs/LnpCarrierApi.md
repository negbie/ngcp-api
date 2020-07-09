# \LnpCarrierApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LnpcarriersGet**](LnpCarrierApi.md#LnpcarriersGet) | **Get** /lnpcarriers/ | Get LnpCarrier items
[**LnpcarriersIdDelete**](LnpCarrierApi.md#LnpcarriersIdDelete) | **Delete** /lnpcarriers/{id} | Delete a specific LnpCarrier
[**LnpcarriersIdGet**](LnpCarrierApi.md#LnpcarriersIdGet) | **Get** /lnpcarriers/{id} | Get a specific LnpCarrier
[**LnpcarriersIdPatch**](LnpCarrierApi.md#LnpcarriersIdPatch) | **Patch** /lnpcarriers/{id} | Change a specific LnpCarrier
[**LnpcarriersIdPut**](LnpCarrierApi.md#LnpcarriersIdPut) | **Put** /lnpcarriers/{id} | Replace/change a specific LnpCarrier
[**LnpcarriersPost**](LnpCarrierApi.md#LnpcarriersPost) | **Post** /lnpcarriers/ | Create a new LnpCarrier



## LnpcarriersGet

> []LnpCarrier LnpcarriersGet(ctx, optional)

Get LnpCarrier items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LnpcarriersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LnpcarriersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| Filter for LNP carriers with a specific prefix (wildcards possible) | 
 **name** | **optional.String**| Filter for LNP carriers with a specific name (wildcards possible) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]LnpCarrier**](LnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpcarriersIdDelete

> LnpcarriersIdDelete(ctx, id)

Delete a specific LnpCarrier

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


## LnpcarriersIdGet

> LnpCarrier LnpcarriersIdGet(ctx, id)

Get a specific LnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**LnpCarrier**](LnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpcarriersIdPatch

> LnpCarrier LnpcarriersIdPatch(ctx, id, operation)

Change a specific LnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**LnpCarrier**](LnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpcarriersIdPut

> LnpCarrier LnpcarriersIdPut(ctx, id, lnpCarrier)

Replace/change a specific LnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**lnpCarrier** | [**LnpCarrier**](LnpCarrier.md)|  | 

### Return type

[**LnpCarrier**](LnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LnpcarriersPost

> []LnpCarrier LnpcarriersPost(ctx, lnpCarrier)

Create a new LnpCarrier

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lnpCarrier** | [**LnpCarrier**](LnpCarrier.md)|  | 

### Return type

[**[]LnpCarrier**](LnpCarrier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

