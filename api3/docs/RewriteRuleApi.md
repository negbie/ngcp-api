# \RewriteRuleApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RewriterulesGet**](RewriteRuleApi.md#RewriterulesGet) | **Get** /rewriterules/ | Get RewriteRule items
[**RewriterulesIdDelete**](RewriteRuleApi.md#RewriterulesIdDelete) | **Delete** /rewriterules/{id} | Delete a specific RewriteRule
[**RewriterulesIdGet**](RewriteRuleApi.md#RewriterulesIdGet) | **Get** /rewriterules/{id} | Get a specific RewriteRule
[**RewriterulesIdPatch**](RewriteRuleApi.md#RewriterulesIdPatch) | **Patch** /rewriterules/{id} | Change a specific RewriteRule
[**RewriterulesIdPut**](RewriteRuleApi.md#RewriterulesIdPut) | **Put** /rewriterules/{id} | Replace/change a specific RewriteRule
[**RewriterulesPost**](RewriteRuleApi.md#RewriterulesPost) | **Post** /rewriterules/ | Create a new RewriteRule



## RewriterulesGet

> []RewriteRule RewriterulesGet(ctx, optional)

Get RewriteRule items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RewriterulesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RewriterulesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **optional.String**| Filter rules for a certain description (wildcards possible). | 
 **setId** | **optional.String**| Filter for rules belonging to a specific rewriteruleset. | 
 **resellerId** | **optional.String**| Filter for rules belonging to a specific reseller. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]RewriteRule**](RewriteRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesIdDelete

> RewriterulesIdDelete(ctx, id)

Delete a specific RewriteRule

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


## RewriterulesIdGet

> RewriteRule RewriterulesIdGet(ctx, id)

Get a specific RewriteRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**RewriteRule**](RewriteRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesIdPatch

> RewriteRule RewriterulesIdPatch(ctx, id, operation)

Change a specific RewriteRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**RewriteRule**](RewriteRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesIdPut

> RewriteRule RewriterulesIdPut(ctx, id, rewriteRule)

Replace/change a specific RewriteRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**rewriteRule** | [**RewriteRule**](RewriteRule.md)|  | 

### Return type

[**RewriteRule**](RewriteRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RewriterulesPost

> []RewriteRule RewriterulesPost(ctx, rewriteRule)

Create a new RewriteRule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rewriteRule** | [**RewriteRule**](RewriteRule.md)|  | 

### Return type

[**[]RewriteRule**](RewriteRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

