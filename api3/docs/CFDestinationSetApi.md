# \CFDestinationSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CfdestinationsetsGet**](CFDestinationSetApi.md#CfdestinationsetsGet) | **Get** /cfdestinationsets/ | Get CFDestinationSet items
[**CfdestinationsetsIdDelete**](CFDestinationSetApi.md#CfdestinationsetsIdDelete) | **Delete** /cfdestinationsets/{id} | Delete a specific CFDestinationSet
[**CfdestinationsetsIdGet**](CFDestinationSetApi.md#CfdestinationsetsIdGet) | **Get** /cfdestinationsets/{id} | Get a specific CFDestinationSet
[**CfdestinationsetsIdPatch**](CFDestinationSetApi.md#CfdestinationsetsIdPatch) | **Patch** /cfdestinationsets/{id} | Change a specific CFDestinationSet
[**CfdestinationsetsIdPut**](CFDestinationSetApi.md#CfdestinationsetsIdPut) | **Put** /cfdestinationsets/{id} | Replace/change a specific CFDestinationSet
[**CfdestinationsetsPost**](CFDestinationSetApi.md#CfdestinationsetsPost) | **Post** /cfdestinationsets/ | Create a new CFDestinationSet



## CfdestinationsetsGet

> []CfDestinationSet CfdestinationsetsGet(ctx, optional)

Get CFDestinationSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CfdestinationsetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CfdestinationsetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for destination sets belonging to a specific subscriber | 
 **name** | **optional.String**| Filter for contacts matching a destination set name pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CfDestinationSet**](CFDestinationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfdestinationsetsIdDelete

> CfdestinationsetsIdDelete(ctx, id)

Delete a specific CFDestinationSet

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


## CfdestinationsetsIdGet

> CfDestinationSet CfdestinationsetsIdGet(ctx, id)

Get a specific CFDestinationSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CfDestinationSet**](CFDestinationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfdestinationsetsIdPatch

> CfDestinationSet CfdestinationsetsIdPatch(ctx, id, operation)

Change a specific CFDestinationSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CfDestinationSet**](CFDestinationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfdestinationsetsIdPut

> CfDestinationSet CfdestinationsetsIdPut(ctx, id, cfDestinationSet)

Replace/change a specific CFDestinationSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**cfDestinationSet** | [**CfDestinationSet**](CfDestinationSet.md)|  | 

### Return type

[**CfDestinationSet**](CFDestinationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CfdestinationsetsPost

> []CfDestinationSet CfdestinationsetsPost(ctx, cfDestinationSet)

Create a new CFDestinationSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cfDestinationSet** | [**CfDestinationSet**](CfDestinationSet.md)|  | 

### Return type

[**[]CfDestinationSet**](CFDestinationSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

