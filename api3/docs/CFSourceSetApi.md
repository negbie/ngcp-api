# \CFSourceSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfsourcesetsGet**](CFSourceSetApi.md#CfsourcesetsGet) | **Get** /cfsourcesets/ | Get CFSourceSet items
[**CfsourcesetsIdDelete**](CFSourceSetApi.md#CfsourcesetsIdDelete) | **Delete** /cfsourcesets/{id} | Delete a specific CFSourceSet
[**CfsourcesetsIdGet**](CFSourceSetApi.md#CfsourcesetsIdGet) | **Get** /cfsourcesets/{id} | Get a specific CFSourceSet
[**CfsourcesetsIdPatch**](CFSourceSetApi.md#CfsourcesetsIdPatch) | **Patch** /cfsourcesets/{id} | Change a specific CFSourceSet
[**CfsourcesetsIdPut**](CFSourceSetApi.md#CfsourcesetsIdPut) | **Put** /cfsourcesets/{id} | Replace/change a specific CFSourceSet
[**CfsourcesetsPost**](CFSourceSetApi.md#CfsourcesetsPost) | **Post** /cfsourcesets/ | Create a new CFSourceSet



## CfsourcesetsGet

> []CfSourceSet CfsourcesetsGet(ctx, optional)

Get CFSourceSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CfsourcesetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CfsourcesetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for source sets belonging to a specific subscriber | 
 **name** | **optional.String**| Filter for items matching a source set name pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CfSourceSet**](CFSourceSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfsourcesetsIdDelete

> CfsourcesetsIdDelete(ctx, id)

Delete a specific CFSourceSet

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


## CfsourcesetsIdGet

> CfSourceSet CfsourcesetsIdGet(ctx, id)

Get a specific CFSourceSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CfSourceSet**](CFSourceSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfsourcesetsIdPatch

> CfSourceSet CfsourcesetsIdPatch(ctx, id, operation)

Change a specific CFSourceSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CfSourceSet**](CFSourceSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfsourcesetsIdPut

> CfSourceSet CfsourcesetsIdPut(ctx, id, cfSourceSet)

Replace/change a specific CFSourceSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**cfSourceSet** | [**CfSourceSet**](CfSourceSet.md)|  | 

### Return type

[**CfSourceSet**](CFSourceSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfsourcesetsPost

> []CfSourceSet CfsourcesetsPost(ctx, cfSourceSet)

Create a new CFSourceSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfSourceSet** | [**CfSourceSet**](CfSourceSet.md)|  | 

### Return type

[**[]CfSourceSet**](CFSourceSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

