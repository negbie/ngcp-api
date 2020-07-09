# \PeeringServerApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PeeringserversGet**](PeeringServerApi.md#PeeringserversGet) | **Get** /peeringservers/ | Get PeeringServer items
[**PeeringserversIdDelete**](PeeringServerApi.md#PeeringserversIdDelete) | **Delete** /peeringservers/{id} | Delete a specific PeeringServer
[**PeeringserversIdGet**](PeeringServerApi.md#PeeringserversIdGet) | **Get** /peeringservers/{id} | Get a specific PeeringServer
[**PeeringserversIdPatch**](PeeringServerApi.md#PeeringserversIdPatch) | **Patch** /peeringservers/{id} | Change a specific PeeringServer
[**PeeringserversIdPut**](PeeringServerApi.md#PeeringserversIdPut) | **Put** /peeringservers/{id} | Replace/change a specific PeeringServer
[**PeeringserversPost**](PeeringServerApi.md#PeeringserversPost) | **Post** /peeringservers/ | Create a new PeeringServer



## PeeringserversGet

> []PeeringServer PeeringserversGet(ctx, optional)

Get PeeringServer items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PeeringserversGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PeeringserversGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **optional.String**| Filter for peering server group | 
 **name** | **optional.String**| Filter for peering server name | 
 **host** | **optional.String**| Filter for peering server host | 
 **ip** | **optional.String**| Filter for peering server ip | 
 **enabled** | **optional.String**| Filter for peering server enabled flag | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]PeeringServer**](PeeringServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringserversIdDelete

> PeeringserversIdDelete(ctx, id)

Delete a specific PeeringServer

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


## PeeringserversIdGet

> PeeringServer PeeringserversIdGet(ctx, id)

Get a specific PeeringServer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**PeeringServer**](PeeringServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringserversIdPatch

> PeeringServer PeeringserversIdPatch(ctx, id, operation)

Change a specific PeeringServer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**PeeringServer**](PeeringServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringserversIdPut

> PeeringServer PeeringserversIdPut(ctx, id, peeringServer)

Replace/change a specific PeeringServer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**peeringServer** | [**PeeringServer**](PeeringServer.md)|  | 

### Return type

[**PeeringServer**](PeeringServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PeeringserversPost

> []PeeringServer PeeringserversPost(ctx, peeringServer)

Create a new PeeringServer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peeringServer** | [**PeeringServer**](PeeringServer.md)|  | 

### Return type

[**[]PeeringServer**](PeeringServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

