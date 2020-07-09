# \PeeringGroupApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PeeringgroupsGet**](PeeringGroupApi.md#PeeringgroupsGet) | **Get** /peeringgroups/ | Get PeeringGroup items
[**PeeringgroupsIdDelete**](PeeringGroupApi.md#PeeringgroupsIdDelete) | **Delete** /peeringgroups/{id} | Delete a specific PeeringGroup
[**PeeringgroupsIdGet**](PeeringGroupApi.md#PeeringgroupsIdGet) | **Get** /peeringgroups/{id} | Get a specific PeeringGroup
[**PeeringgroupsIdPatch**](PeeringGroupApi.md#PeeringgroupsIdPatch) | **Patch** /peeringgroups/{id} | Change a specific PeeringGroup
[**PeeringgroupsIdPut**](PeeringGroupApi.md#PeeringgroupsIdPut) | **Put** /peeringgroups/{id} | Replace/change a specific PeeringGroup
[**PeeringgroupsPost**](PeeringGroupApi.md#PeeringgroupsPost) | **Post** /peeringgroups/ | Create a new PeeringGroup



## PeeringgroupsGet

> []PeeringGroup PeeringgroupsGet(ctx, optional)

Get PeeringGroup items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeeringgroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PeeringgroupsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter for peering group name | 
 **description** | **optional.String**| Filter for peering group description | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PeeringGroup**](PeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringgroupsIdDelete

> PeeringgroupsIdDelete(ctx, id)

Delete a specific PeeringGroup

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


## PeeringgroupsIdGet

> PeeringGroup PeeringgroupsIdGet(ctx, id)

Get a specific PeeringGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PeeringGroup**](PeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringgroupsIdPatch

> PeeringGroup PeeringgroupsIdPatch(ctx, id, operation)

Change a specific PeeringGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PeeringGroup**](PeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringgroupsIdPut

> PeeringGroup PeeringgroupsIdPut(ctx, id, peeringGroup)

Replace/change a specific PeeringGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**peeringGroup** | [**PeeringGroup**](PeeringGroup.md)|  | 

### Return type

[**PeeringGroup**](PeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringgroupsPost

> []PeeringGroup PeeringgroupsPost(ctx, peeringGroup)

Create a new PeeringGroup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peeringGroup** | [**PeeringGroup**](PeeringGroup.md)|  | 

### Return type

[**[]PeeringGroup**](PeeringGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

