# \NumberApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NumbersGet**](NumberApi.md#NumbersGet) | **Get** /numbers/ | Get Number items
[**NumbersIdGet**](NumberApi.md#NumbersIdGet) | **Get** /numbers/{id} | Get a specific Number
[**NumbersIdPatch**](NumberApi.md#NumbersIdPatch) | **Patch** /numbers/{id} | Change a specific Number
[**NumbersIdPut**](NumberApi.md#NumbersIdPut) | **Put** /numbers/{id} | Replace/change a specific Number



## NumbersGet

> []Number NumbersGet(ctx, optional)

Get Number items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NumbersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a NumbersGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for numbers assigned to subscribers belonging to a specific reseller | 
 **customerId** | **optional.String**| Filter for numbers assigned to subscribers of a specific customer. | 
 **subscriberId** | **optional.String**| Filter for numbers assigned to a specific subscriber. | 
 **ac** | **optional.String**| Filter for ac field of the number. | 
 **cc** | **optional.String**| Filter for cc field c the number. | 
 **sn** | **optional.String**| Filter for sn field c the number. | 
 **type_** | **optional.String**| Filter for number type, either \&quot;primary\&quot; or \&quot;alias\&quot;. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Number**](Number.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NumbersIdGet

> Number NumbersIdGet(ctx, id)

Get a specific Number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Number**](Number.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NumbersIdPatch

> Number NumbersIdPatch(ctx, id, operation)

Change a specific Number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**Number**](Number.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NumbersIdPut

> Number NumbersIdPut(ctx, id, number)

Replace/change a specific Number

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**number** | [**Number**](Number.md)|  | 

### Return type

[**Number**](Number.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

