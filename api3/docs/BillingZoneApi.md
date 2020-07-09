# \BillingZoneApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingzonesGet**](BillingZoneApi.md#BillingzonesGet) | **Get** /billingzones/ | Get BillingZone items
[**BillingzonesIdDelete**](BillingZoneApi.md#BillingzonesIdDelete) | **Delete** /billingzones/{id} | Delete a specific BillingZone
[**BillingzonesIdGet**](BillingZoneApi.md#BillingzonesIdGet) | **Get** /billingzones/{id} | Get a specific BillingZone
[**BillingzonesIdPatch**](BillingZoneApi.md#BillingzonesIdPatch) | **Patch** /billingzones/{id} | Change a specific BillingZone
[**BillingzonesIdPut**](BillingZoneApi.md#BillingzonesIdPut) | **Put** /billingzones/{id} | Replace/change a specific BillingZone
[**BillingzonesPost**](BillingZoneApi.md#BillingzonesPost) | **Post** /billingzones/ | Create a new BillingZone



## BillingzonesGet

> []BillingZone BillingzonesGet(ctx, optional)

Get BillingZone items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingzonesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingzonesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingProfileId** | **optional.String**| Filter for zones belonging to a specific billing profile | 
 **zone** | **optional.String**| Filter for zone name | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]BillingZone**](BillingZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingzonesIdDelete

> BillingzonesIdDelete(ctx, id)

Delete a specific BillingZone

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


## BillingzonesIdGet

> BillingZone BillingzonesIdGet(ctx, id)

Get a specific BillingZone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**BillingZone**](BillingZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingzonesIdPatch

> BillingZone BillingzonesIdPatch(ctx, id, operation)

Change a specific BillingZone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**BillingZone**](BillingZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingzonesIdPut

> BillingZone BillingzonesIdPut(ctx, id, billingZone)

Replace/change a specific BillingZone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**billingZone** | [**BillingZone**](BillingZone.md)|  | 

### Return type

[**BillingZone**](BillingZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingzonesPost

> []BillingZone BillingzonesPost(ctx, billingZone)

Create a new BillingZone

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingZone** | [**BillingZone**](BillingZone.md)|  | 

### Return type

[**[]BillingZone**](BillingZone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

