# \CFTimeSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CftimesetsGet**](CFTimeSetApi.md#CftimesetsGet) | **Get** /cftimesets/ | Get CFTimeSet items
[**CftimesetsIdDelete**](CFTimeSetApi.md#CftimesetsIdDelete) | **Delete** /cftimesets/{id} | Delete a specific CFTimeSet
[**CftimesetsIdGet**](CFTimeSetApi.md#CftimesetsIdGet) | **Get** /cftimesets/{id} | Get a specific CFTimeSet
[**CftimesetsIdPatch**](CFTimeSetApi.md#CftimesetsIdPatch) | **Patch** /cftimesets/{id} | Change a specific CFTimeSet
[**CftimesetsIdPut**](CFTimeSetApi.md#CftimesetsIdPut) | **Put** /cftimesets/{id} | Replace/change a specific CFTimeSet
[**CftimesetsPost**](CFTimeSetApi.md#CftimesetsPost) | **Post** /cftimesets/ | Create a new CFTimeSet



## CftimesetsGet

> []CfTimeSet CftimesetsGet(ctx, optional)

Get CFTimeSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CftimesetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CftimesetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for timesets belonging to a specific subscriber | 
 **name** | **optional.String**| Filter for contacts matching a timeset name pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CfTimeSet**](CFTimeSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CftimesetsIdDelete

> CftimesetsIdDelete(ctx, id)

Delete a specific CFTimeSet

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


## CftimesetsIdGet

> CfTimeSet CftimesetsIdGet(ctx, id)

Get a specific CFTimeSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CfTimeSet**](CFTimeSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CftimesetsIdPatch

> CfTimeSet CftimesetsIdPatch(ctx, id, operation)

Change a specific CFTimeSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CfTimeSet**](CFTimeSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CftimesetsIdPut

> CfTimeSet CftimesetsIdPut(ctx, id, cfTimeSet)

Replace/change a specific CFTimeSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**cfTimeSet** | [**CfTimeSet**](CfTimeSet.md)|  | 

### Return type

[**CfTimeSet**](CFTimeSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CftimesetsPost

> []CfTimeSet CftimesetsPost(ctx, cfTimeSet)

Create a new CFTimeSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfTimeSet** | [**CfTimeSet**](CfTimeSet.md)|  | 

### Return type

[**[]CfTimeSet**](CFTimeSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

