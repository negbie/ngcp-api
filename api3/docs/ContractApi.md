# \ContractApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractsGet**](ContractApi.md#ContractsGet) | **Get** /contracts/ | Get Contract items
[**ContractsIdGet**](ContractApi.md#ContractsIdGet) | **Get** /contracts/{id} | Get a specific Contract
[**ContractsIdPatch**](ContractApi.md#ContractsIdPatch) | **Patch** /contracts/{id} | Change a specific Contract
[**ContractsIdPut**](ContractApi.md#ContractsIdPut) | **Put** /contracts/{id} | Replace/change a specific Contract
[**ContractsPost**](ContractApi.md#ContractsPost) | **Post** /contracts/ | Create a new Contract



## ContractsGet

> []Contract ContractsGet(ctx, optional)

Get Contract items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContractsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ContractsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactId** | **optional.String**| Filter for contracts with a specific contact id | 
 **status** | **optional.String**| Filter for contracts with a specific status (except \&quot;terminated\&quot;) | 
 **externalId** | **optional.String**| Filter for contracts with a specific external id | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsIdGet

> Contract ContractsIdGet(ctx, id)

Get a specific Contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsIdPatch

> Contract ContractsIdPatch(ctx, id, operation)

Change a specific Contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsIdPut

> Contract ContractsIdPut(ctx, id, contract)

Replace/change a specific Contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**contract** | [**Contract**](Contract.md)|  | 

### Return type

[**Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContractsPost

> []Contract ContractsPost(ctx, contract)

Create a new Contract

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contract** | [**Contract**](Contract.md)|  | 

### Return type

[**[]Contract**](Contract.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

