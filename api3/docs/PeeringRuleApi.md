# \PeeringRuleApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PeeringrulesGet**](PeeringRuleApi.md#PeeringrulesGet) | **Get** /peeringrules/ | Get PeeringRule items
[**PeeringrulesIdDelete**](PeeringRuleApi.md#PeeringrulesIdDelete) | **Delete** /peeringrules/{id} | Delete a specific PeeringRule
[**PeeringrulesIdGet**](PeeringRuleApi.md#PeeringrulesIdGet) | **Get** /peeringrules/{id} | Get a specific PeeringRule
[**PeeringrulesIdPatch**](PeeringRuleApi.md#PeeringrulesIdPatch) | **Patch** /peeringrules/{id} | Change a specific PeeringRule
[**PeeringrulesIdPut**](PeeringRuleApi.md#PeeringrulesIdPut) | **Put** /peeringrules/{id} | Replace/change a specific PeeringRule
[**PeeringrulesPost**](PeeringRuleApi.md#PeeringrulesPost) | **Post** /peeringrules/ | Create a new PeeringRule



## PeeringrulesGet

> []PeeringRule PeeringrulesGet(ctx, optional)

Get PeeringRule items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeeringrulesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PeeringrulesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.String**| Filter for peering group | 
 **description** | **optional.String**| Filter for peering rules description | 
 **enabled** | **optional.String**| Filter for peering rules enabled flag | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PeeringRule**](PeeringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringrulesIdDelete

> PeeringrulesIdDelete(ctx, id)

Delete a specific PeeringRule

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


## PeeringrulesIdGet

> PeeringRule PeeringrulesIdGet(ctx, id)

Get a specific PeeringRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PeeringRule**](PeeringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringrulesIdPatch

> PeeringRule PeeringrulesIdPatch(ctx, id, operation)

Change a specific PeeringRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PeeringRule**](PeeringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringrulesIdPut

> PeeringRule PeeringrulesIdPut(ctx, id, peeringRule)

Replace/change a specific PeeringRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**peeringRule** | [**PeeringRule**](PeeringRule.md)|  | 

### Return type

[**PeeringRule**](PeeringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringrulesPost

> []PeeringRule PeeringrulesPost(ctx, peeringRule)

Create a new PeeringRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peeringRule** | [**PeeringRule**](PeeringRule.md)|  | 

### Return type

[**[]PeeringRule**](PeeringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

