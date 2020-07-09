# \PeeringInboundRuleApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PeeringinboundrulesGet**](PeeringInboundRuleApi.md#PeeringinboundrulesGet) | **Get** /peeringinboundrules/ | Get PeeringInboundRule items
[**PeeringinboundrulesIdDelete**](PeeringInboundRuleApi.md#PeeringinboundrulesIdDelete) | **Delete** /peeringinboundrules/{id} | Delete a specific PeeringInboundRule
[**PeeringinboundrulesIdGet**](PeeringInboundRuleApi.md#PeeringinboundrulesIdGet) | **Get** /peeringinboundrules/{id} | Get a specific PeeringInboundRule
[**PeeringinboundrulesIdPatch**](PeeringInboundRuleApi.md#PeeringinboundrulesIdPatch) | **Patch** /peeringinboundrules/{id} | Change a specific PeeringInboundRule
[**PeeringinboundrulesIdPut**](PeeringInboundRuleApi.md#PeeringinboundrulesIdPut) | **Put** /peeringinboundrules/{id} | Replace/change a specific PeeringInboundRule
[**PeeringinboundrulesPost**](PeeringInboundRuleApi.md#PeeringinboundrulesPost) | **Post** /peeringinboundrules/ | Create a new PeeringInboundRule



## PeeringinboundrulesGet

> []PeeringInboundRule PeeringinboundrulesGet(ctx, optional)

Get PeeringInboundRule items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeeringinboundrulesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PeeringinboundrulesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.String**| Filter for peering group | 
 **field** | **optional.String**| Filter for peering rules field (wildcards possible) | 
 **enabled** | **optional.String**| Filter for peering rules enabled flag | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PeeringInboundRule**](PeeringInboundRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringinboundrulesIdDelete

> PeeringinboundrulesIdDelete(ctx, id)

Delete a specific PeeringInboundRule

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


## PeeringinboundrulesIdGet

> PeeringInboundRule PeeringinboundrulesIdGet(ctx, id)

Get a specific PeeringInboundRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PeeringInboundRule**](PeeringInboundRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringinboundrulesIdPatch

> PeeringInboundRule PeeringinboundrulesIdPatch(ctx, id, operation)

Change a specific PeeringInboundRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PeeringInboundRule**](PeeringInboundRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringinboundrulesIdPut

> PeeringInboundRule PeeringinboundrulesIdPut(ctx, id, peeringInboundRule)

Replace/change a specific PeeringInboundRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**peeringInboundRule** | [**PeeringInboundRule**](PeeringInboundRule.md)|  | 

### Return type

[**PeeringInboundRule**](PeeringInboundRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringinboundrulesPost

> []PeeringInboundRule PeeringinboundrulesPost(ctx, peeringInboundRule)

Create a new PeeringInboundRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peeringInboundRule** | [**PeeringInboundRule**](PeeringInboundRule.md)|  | 

### Return type

[**[]PeeringInboundRule**](PeeringInboundRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

