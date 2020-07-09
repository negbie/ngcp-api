# \RewriteRuleSetApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RewriterulesetsGet**](RewriteRuleSetApi.md#RewriterulesetsGet) | **Get** /rewriterulesets/ | Get RewriteRuleSet items
[**RewriterulesetsIdDelete**](RewriteRuleSetApi.md#RewriterulesetsIdDelete) | **Delete** /rewriterulesets/{id} | Delete a specific RewriteRuleSet
[**RewriterulesetsIdGet**](RewriteRuleSetApi.md#RewriterulesetsIdGet) | **Get** /rewriterulesets/{id} | Get a specific RewriteRuleSet
[**RewriterulesetsIdPatch**](RewriteRuleSetApi.md#RewriterulesetsIdPatch) | **Patch** /rewriterulesets/{id} | Change a specific RewriteRuleSet
[**RewriterulesetsIdPut**](RewriteRuleSetApi.md#RewriterulesetsIdPut) | **Put** /rewriterulesets/{id} | Replace/change a specific RewriteRuleSet
[**RewriterulesetsPost**](RewriteRuleSetApi.md#RewriterulesetsPost) | **Post** /rewriterulesets/ | Create a new RewriteRuleSet



## RewriterulesetsGet

> []RewriteRuleSet RewriterulesetsGet(ctx, optional)

Get RewriteRuleSet items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RewriterulesetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RewriterulesetsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for rewriterulesets belonging to a specific reseller | 
 **description** | **optional.String**| Filter rulesets for a certain description (wildcards possible). | 
 **name** | **optional.String**| Filter rulesets for a certain name (wildcards possible). | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]RewriteRuleSet**](RewriteRuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesetsIdDelete

> RewriterulesetsIdDelete(ctx, id)

Delete a specific RewriteRuleSet

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


## RewriterulesetsIdGet

> RewriteRuleSet RewriterulesetsIdGet(ctx, id)

Get a specific RewriteRuleSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**RewriteRuleSet**](RewriteRuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesetsIdPatch

> RewriteRuleSet RewriterulesetsIdPatch(ctx, id, operation)

Change a specific RewriteRuleSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**RewriteRuleSet**](RewriteRuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesetsIdPut

> RewriteRuleSet RewriterulesetsIdPut(ctx, id, rewriteRuleSet)

Replace/change a specific RewriteRuleSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**rewriteRuleSet** | [**RewriteRuleSet**](RewriteRuleSet.md)|  | 

### Return type

[**RewriteRuleSet**](RewriteRuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesetsPost

> []RewriteRuleSet RewriterulesetsPost(ctx, rewriteRuleSet)

Create a new RewriteRuleSet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rewriteRuleSet** | [**RewriteRuleSet**](RewriteRuleSet.md)|  | 

### Return type

[**[]RewriteRuleSet**](RewriteRuleSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

