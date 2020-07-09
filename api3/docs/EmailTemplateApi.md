# \EmailTemplateApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EmailtemplatesGet**](EmailTemplateApi.md#EmailtemplatesGet) | **Get** /emailtemplates/ | Get EmailTemplate items
[**EmailtemplatesIdDelete**](EmailTemplateApi.md#EmailtemplatesIdDelete) | **Delete** /emailtemplates/{id} | Delete a specific EmailTemplate
[**EmailtemplatesIdGet**](EmailTemplateApi.md#EmailtemplatesIdGet) | **Get** /emailtemplates/{id} | Get a specific EmailTemplate
[**EmailtemplatesIdPatch**](EmailTemplateApi.md#EmailtemplatesIdPatch) | **Patch** /emailtemplates/{id} | Change a specific EmailTemplate
[**EmailtemplatesIdPut**](EmailTemplateApi.md#EmailtemplatesIdPut) | **Put** /emailtemplates/{id} | Replace/change a specific EmailTemplate
[**EmailtemplatesPost**](EmailTemplateApi.md#EmailtemplatesPost) | **Post** /emailtemplates/ | Create a new EmailTemplate



## EmailtemplatesGet

> []EmailTemplate EmailtemplatesGet(ctx, optional)

Get EmailTemplate items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailtemplatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EmailtemplatesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for email templates belonging to a specific reseller | 
 **name** | **optional.String**| Filter for email templates with a specific name | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]EmailTemplate**](EmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailtemplatesIdDelete

> EmailtemplatesIdDelete(ctx, id)

Delete a specific EmailTemplate

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


## EmailtemplatesIdGet

> EmailTemplate EmailtemplatesIdGet(ctx, id)

Get a specific EmailTemplate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**EmailTemplate**](EmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailtemplatesIdPatch

> EmailTemplate EmailtemplatesIdPatch(ctx, id, operation)

Change a specific EmailTemplate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**EmailTemplate**](EmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailtemplatesIdPut

> EmailTemplate EmailtemplatesIdPut(ctx, id, emailTemplate)

Replace/change a specific EmailTemplate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**emailTemplate** | [**EmailTemplate**](EmailTemplate.md)|  | 

### Return type

[**EmailTemplate**](EmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailtemplatesPost

> []EmailTemplate EmailtemplatesPost(ctx, emailTemplate)

Create a new EmailTemplate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailTemplate** | [**EmailTemplate**](EmailTemplate.md)|  | 

### Return type

[**[]EmailTemplate**](EmailTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

