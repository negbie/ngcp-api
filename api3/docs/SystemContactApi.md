# \SystemContactApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemcontactsGet**](SystemContactApi.md#SystemcontactsGet) | **Get** /systemcontacts/ | Get SystemContact items
[**SystemcontactsIdDelete**](SystemContactApi.md#SystemcontactsIdDelete) | **Delete** /systemcontacts/{id} | Delete a specific SystemContact
[**SystemcontactsIdGet**](SystemContactApi.md#SystemcontactsIdGet) | **Get** /systemcontacts/{id} | Get a specific SystemContact
[**SystemcontactsIdPatch**](SystemContactApi.md#SystemcontactsIdPatch) | **Patch** /systemcontacts/{id} | Change a specific SystemContact
[**SystemcontactsIdPut**](SystemContactApi.md#SystemcontactsIdPut) | **Put** /systemcontacts/{id} | Replace/change a specific SystemContact
[**SystemcontactsPost**](SystemContactApi.md#SystemcontactsPost) | **Post** /systemcontacts/ | Create a new SystemContact



## SystemcontactsGet

> []SystemContact SystemcontactsGet(ctx, optional)

Get SystemContact items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemcontactsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SystemcontactsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Filter for contacts matching an email pattern | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SystemContact**](SystemContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemcontactsIdDelete

> SystemcontactsIdDelete(ctx, id)

Delete a specific SystemContact

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


## SystemcontactsIdGet

> SystemContact SystemcontactsIdGet(ctx, id)

Get a specific SystemContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SystemContact**](SystemContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemcontactsIdPatch

> SystemContact SystemcontactsIdPatch(ctx, id, operation)

Change a specific SystemContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**SystemContact**](SystemContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemcontactsIdPut

> SystemContact SystemcontactsIdPut(ctx, id, systemContact)

Replace/change a specific SystemContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**systemContact** | [**SystemContact**](SystemContact.md)|  | 

### Return type

[**SystemContact**](SystemContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemcontactsPost

> []SystemContact SystemcontactsPost(ctx, systemContact)

Create a new SystemContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemContact** | [**SystemContact**](SystemContact.md)|  | 

### Return type

[**[]SystemContact**](SystemContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

