# \CustomerApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersGet**](CustomerApi.md#CustomersGet) | **Get** /customers/ | Get Customer items
[**CustomersIdGet**](CustomerApi.md#CustomersIdGet) | **Get** /customers/{id} | Get a specific Customer
[**CustomersIdPatch**](CustomerApi.md#CustomersIdPatch) | **Patch** /customers/{id} | Change a specific Customer
[**CustomersIdPut**](CustomerApi.md#CustomersIdPut) | **Put** /customers/{id} | Replace/change a specific Customer
[**CustomersPost**](CustomerApi.md#CustomersPost) | **Post** /customers/ | Create a new Customer



## CustomersGet

> []Customer CustomersGet(ctx, optional)

Get Customer items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Filter for customers with a specific status (comma-separated list of statuses to include possible) | 
 **notStatus** | **optional.String**| Filter for customers not having a specific status (comma-separated list of statuses to exclude possible) | 
 **resellerId** | **optional.String**| Filter for customers belonging to a specific reseller | 
 **externalId** | **optional.String**| Filter for customer with specific external_id | 
 **contactId** | **optional.String**| Filter for customers belonging to a specific contact | 
 **packageId** | **optional.String**| Filter for customers with specific profile package id | 
 **createTimestampGt** | **optional.String**| Filter for customers with create_timestamp greater then specified value | 
 **createTimestampLt** | **optional.String**| Filter for customers with create_timestamp less then specified value | 
 **activateTimestampGt** | **optional.String**| Filter for customers with activate_timestamp greater then specified value | 
 **activateTimestampLt** | **optional.String**| Filter for customers with activate_timestamp less then specified value | 
 **modifyTimestampGt** | **optional.String**| Filter for customers with modify_timestamp greater then specified value | 
 **modifyTimestampLt** | **optional.String**| Filter for customers with modify_timestamp less then specified value | 
 **terminateTimestampGt** | **optional.String**| Filter for customers with terminate_timestamp greater then specified value | 
 **terminateTimestampLt** | **optional.String**| Filter for customers with terminate_timestamp less then specified value | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersIdGet

> Customer CustomersIdGet(ctx, id)

Get a specific Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersIdPatch

> Customer CustomersIdPatch(ctx, id, operation)

Change a specific Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersIdPut

> Customer CustomersIdPut(ctx, id, customer)

Replace/change a specific Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**customer** | [**Customer**](Customer.md)|  | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersPost

> []Customer CustomersPost(ctx, customer)

Create a new Customer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customer** | [**Customer**](Customer.md)|  | 

### Return type

[**[]Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

