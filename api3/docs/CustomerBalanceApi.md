# \CustomerBalanceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerbalancesGet**](CustomerBalanceApi.md#CustomerbalancesGet) | **Get** /customerbalances/ | Get CustomerBalance items
[**CustomerbalancesIdGet**](CustomerBalanceApi.md#CustomerbalancesIdGet) | **Get** /customerbalances/{id} | Get a specific CustomerBalance
[**CustomerbalancesIdPatch**](CustomerBalanceApi.md#CustomerbalancesIdPatch) | **Patch** /customerbalances/{id} | Change a specific CustomerBalance
[**CustomerbalancesIdPut**](CustomerBalanceApi.md#CustomerbalancesIdPut) | **Put** /customerbalances/{id} | Replace/change a specific CustomerBalance



## CustomerbalancesGet

> []CustomerBalance CustomerbalancesGet(ctx, optional)

Get CustomerBalance items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerbalancesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomerbalancesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for customer balances belonging to a specific reseller | 
 **contactId** | **optional.String**| Filter for contracts with a specific contact id | 
 **status** | **optional.String**| Filter for contracts with a specific status (except \&quot;terminated\&quot;) | 
 **externalId** | **optional.String**| Filter for contracts with a specific external id | 
 **domain** | **optional.String**| Filter for contracts with subscribers belonging to a specific domain | 
 **prepaid** | **optional.String**| Filter for contracts with a prepaid billing profile | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CustomerBalance**](CustomerBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerbalancesIdGet

> CustomerBalance CustomerbalancesIdGet(ctx, id)

Get a specific CustomerBalance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CustomerBalance**](CustomerBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerbalancesIdPatch

> CustomerBalance CustomerbalancesIdPatch(ctx, id, operation)

Change a specific CustomerBalance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CustomerBalance**](CustomerBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerbalancesIdPut

> CustomerBalance CustomerbalancesIdPut(ctx, id, customerBalance)

Replace/change a specific CustomerBalance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**customerBalance** | [**CustomerBalance**](CustomerBalance.md)|  | 

### Return type

[**CustomerBalance**](CustomerBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

