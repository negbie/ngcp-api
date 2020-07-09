# \BillingNetworkApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingnetworksGet**](BillingNetworkApi.md#BillingnetworksGet) | **Get** /billingnetworks/ | Get BillingNetwork items
[**BillingnetworksIdGet**](BillingNetworkApi.md#BillingnetworksIdGet) | **Get** /billingnetworks/{id} | Get a specific BillingNetwork
[**BillingnetworksIdPatch**](BillingNetworkApi.md#BillingnetworksIdPatch) | **Patch** /billingnetworks/{id} | Change a specific BillingNetwork
[**BillingnetworksIdPut**](BillingNetworkApi.md#BillingnetworksIdPut) | **Put** /billingnetworks/{id} | Replace/change a specific BillingNetwork
[**BillingnetworksPost**](BillingNetworkApi.md#BillingnetworksPost) | **Post** /billingnetworks/ | Create a new BillingNetwork



## BillingnetworksGet

> []BillingNetwork BillingnetworksGet(ctx, optional)

Get BillingNetwork items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingnetworksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingnetworksGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for billing networks belonging to a specific reseller | 
 **ip** | **optional.String**| Filter for billing networks containing a specific IP address | 
 **name** | **optional.String**| Filter for billing networks matching a name pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]BillingNetwork**](BillingNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingnetworksIdGet

> BillingNetwork BillingnetworksIdGet(ctx, id)

Get a specific BillingNetwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**BillingNetwork**](BillingNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingnetworksIdPatch

> BillingNetwork BillingnetworksIdPatch(ctx, id, operation)

Change a specific BillingNetwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**BillingNetwork**](BillingNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingnetworksIdPut

> BillingNetwork BillingnetworksIdPut(ctx, id, billingNetwork)

Replace/change a specific BillingNetwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**billingNetwork** | [**BillingNetwork**](BillingNetwork.md)|  | 

### Return type

[**BillingNetwork**](BillingNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingnetworksPost

> []BillingNetwork BillingnetworksPost(ctx, billingNetwork)

Create a new BillingNetwork

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingNetwork** | [**BillingNetwork**](BillingNetwork.md)|  | 

### Return type

[**[]BillingNetwork**](BillingNetwork.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

