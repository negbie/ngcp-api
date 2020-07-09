# \BillingFeeApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingfeesGet**](BillingFeeApi.md#BillingfeesGet) | **Get** /billingfees/ | Get BillingFee items
[**BillingfeesIdDelete**](BillingFeeApi.md#BillingfeesIdDelete) | **Delete** /billingfees/{id} | Delete a specific BillingFee
[**BillingfeesIdGet**](BillingFeeApi.md#BillingfeesIdGet) | **Get** /billingfees/{id} | Get a specific BillingFee
[**BillingfeesIdPatch**](BillingFeeApi.md#BillingfeesIdPatch) | **Patch** /billingfees/{id} | Change a specific BillingFee
[**BillingfeesIdPut**](BillingFeeApi.md#BillingfeesIdPut) | **Put** /billingfees/{id} | Replace/change a specific BillingFee
[**BillingfeesPost**](BillingFeeApi.md#BillingfeesPost) | **Post** /billingfees/ | Create a new BillingFee



## BillingfeesGet

> []BillingFee BillingfeesGet(ctx, optional)

Get BillingFee items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingfeesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingfeesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingProfileId** | **optional.String**| Filter for fees belonging to a specific billing profile | 
 **billingZoneId** | **optional.String**| Filter for fees of a specific billing zone | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]BillingFee**](BillingFee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingfeesIdDelete

> BillingfeesIdDelete(ctx, id)

Delete a specific BillingFee

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


## BillingfeesIdGet

> BillingFee BillingfeesIdGet(ctx, id)

Get a specific BillingFee

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**BillingFee**](BillingFee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingfeesIdPatch

> BillingFee BillingfeesIdPatch(ctx, id, operation)

Change a specific BillingFee

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**BillingFee**](BillingFee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingfeesIdPut

> BillingFee BillingfeesIdPut(ctx, id, billingFee)

Replace/change a specific BillingFee

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**billingFee** | [**BillingFee**](BillingFee.md)|  | 

### Return type

[**BillingFee**](BillingFee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingfeesPost

> []BillingFee BillingfeesPost(ctx, billingFee)

Create a new BillingFee

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingFee** | [**BillingFee**](BillingFee.md)|  | 

### Return type

[**[]BillingFee**](BillingFee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

