# \CFBNumberSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfbnumbersetsGet**](CFBNumberSetApi.md#CfbnumbersetsGet) | **Get** /cfbnumbersets/ | Get CFBNumberSet items
[**CfbnumbersetsIdDelete**](CFBNumberSetApi.md#CfbnumbersetsIdDelete) | **Delete** /cfbnumbersets/{id} | Delete a specific CFBNumberSet
[**CfbnumbersetsIdGet**](CFBNumberSetApi.md#CfbnumbersetsIdGet) | **Get** /cfbnumbersets/{id} | Get a specific CFBNumberSet
[**CfbnumbersetsIdPatch**](CFBNumberSetApi.md#CfbnumbersetsIdPatch) | **Patch** /cfbnumbersets/{id} | Change a specific CFBNumberSet
[**CfbnumbersetsIdPut**](CFBNumberSetApi.md#CfbnumbersetsIdPut) | **Put** /cfbnumbersets/{id} | Replace/change a specific CFBNumberSet
[**CfbnumbersetsPost**](CFBNumberSetApi.md#CfbnumbersetsPost) | **Post** /cfbnumbersets/ | Create a new CFBNumberSet



## CfbnumbersetsGet

> []CfbNumberSet CfbnumbersetsGet(ctx, optional)

Get CFBNumberSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CfbnumbersetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CfbnumbersetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for B-Number sets belonging to a specific subscriber | 
 **name** | **optional.String**| Filter for items matching a B-Number Set name pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CfbNumberSet**](CFBNumberSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfbnumbersetsIdDelete

> CfbnumbersetsIdDelete(ctx, id)

Delete a specific CFBNumberSet

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


## CfbnumbersetsIdGet

> CfbNumberSet CfbnumbersetsIdGet(ctx, id)

Get a specific CFBNumberSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CfbNumberSet**](CFBNumberSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfbnumbersetsIdPatch

> CfbNumberSet CfbnumbersetsIdPatch(ctx, id, operation)

Change a specific CFBNumberSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CfbNumberSet**](CFBNumberSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfbnumbersetsIdPut

> CfbNumberSet CfbnumbersetsIdPut(ctx, id, cfbNumberSet)

Replace/change a specific CFBNumberSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**cfbNumberSet** | [**CfbNumberSet**](CfbNumberSet.md)|  | 

### Return type

[**CfbNumberSet**](CFBNumberSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfbnumbersetsPost

> []CfbNumberSet CfbnumbersetsPost(ctx, cfbNumberSet)

Create a new CFBNumberSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfbNumberSet** | [**CfbNumberSet**](CfbNumberSet.md)|  | 

### Return type

[**[]CfbNumberSet**](CFBNumberSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

