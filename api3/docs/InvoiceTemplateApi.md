# \InvoiceTemplateApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoicetemplatesGet**](InvoiceTemplateApi.md#InvoicetemplatesGet) | **Get** /invoicetemplates/ | Get InvoiceTemplate items
[**InvoicetemplatesIdGet**](InvoiceTemplateApi.md#InvoicetemplatesIdGet) | **Get** /invoicetemplates/{id} | Get a specific InvoiceTemplate



## InvoicetemplatesGet

> []InvoiceTemplate InvoicetemplatesGet(ctx, optional)

Get InvoiceTemplate items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InvoicetemplatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvoicetemplatesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for invoice templates belonging to a specific reseller | 
 **name** | **optional.String**| Filter for invoice templates with a specific name | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]InvoiceTemplate**](InvoiceTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicetemplatesIdGet

> InvoiceTemplate InvoicetemplatesIdGet(ctx, id)

Get a specific InvoiceTemplate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**InvoiceTemplate**](InvoiceTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

