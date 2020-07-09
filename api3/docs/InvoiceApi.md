# \InvoiceApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoicesGet**](InvoiceApi.md#InvoicesGet) | **Get** /invoices/ | Get Invoice items
[**InvoicesIdDelete**](InvoiceApi.md#InvoicesIdDelete) | **Delete** /invoices/{id} | Delete a specific Invoice
[**InvoicesIdGet**](InvoiceApi.md#InvoicesIdGet) | **Get** /invoices/{id} | Get a specific Invoice
[**InvoicesPost**](InvoiceApi.md#InvoicesPost) | **Post** /invoices/ | Create a new Invoice



## InvoicesGet

> []Invoice InvoicesGet(ctx, optional)

Get Invoice items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoicesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvoicesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **periodStartGe** | **optional.String**| Filter for invoices newer or equal to the given date (YYYY-MM-DDThh:mm:ss) | 
 **periodEndLe** | **optional.String**| Filter for invoices older or equal to the given date (YYYY-MM-DDThh:mm:ss) | 
 **resellerId** | **optional.String**| Filter for invoices of customers belonging to a certain reseller | 
 **customerId** | **optional.String**| Filter for invoices belonging to a specific customer | 
 **serial** | **optional.String**| Filter for invoices matching a serial (patterns allowed) | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesIdDelete

> InvoicesIdDelete(ctx, id)

Delete a specific Invoice

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


## InvoicesIdGet

> Invoice InvoicesIdGet(ctx, id)

Get a specific Invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesPost

> []Invoice InvoicesPost(ctx, invoice)

Create a new Invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoice** | [**Invoice**](Invoice.md)|  | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

