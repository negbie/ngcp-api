# \CustomerContactApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomercontactsGet**](CustomerContactApi.md#CustomercontactsGet) | **Get** /customercontacts/ | Get CustomerContact items
[**CustomercontactsIdDelete**](CustomerContactApi.md#CustomercontactsIdDelete) | **Delete** /customercontacts/{id} | Delete a specific CustomerContact
[**CustomercontactsIdGet**](CustomerContactApi.md#CustomercontactsIdGet) | **Get** /customercontacts/{id} | Get a specific CustomerContact
[**CustomercontactsIdPatch**](CustomerContactApi.md#CustomercontactsIdPatch) | **Patch** /customercontacts/{id} | Change a specific CustomerContact
[**CustomercontactsIdPut**](CustomerContactApi.md#CustomercontactsIdPut) | **Put** /customercontacts/{id} | Replace/change a specific CustomerContact
[**CustomercontactsPost**](CustomerContactApi.md#CustomercontactsPost) | **Post** /customercontacts/ | Create a new CustomerContact



## CustomercontactsGet

> []CustomerContact CustomercontactsGet(ctx, optional)

Get CustomerContact items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomercontactsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CustomercontactsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Filter for contacts matching an email pattern | 
 **resellerId** | **optional.String**| Filter for contacts belonging to a specific reseller | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CustomerContact**](CustomerContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomercontactsIdDelete

> CustomercontactsIdDelete(ctx, id)

Delete a specific CustomerContact

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


## CustomercontactsIdGet

> CustomerContact CustomercontactsIdGet(ctx, id)

Get a specific CustomerContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CustomerContact**](CustomerContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomercontactsIdPatch

> CustomerContact CustomercontactsIdPatch(ctx, id, operation)

Change a specific CustomerContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**CustomerContact**](CustomerContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomercontactsIdPut

> CustomerContact CustomercontactsIdPut(ctx, id, customerContact)

Replace/change a specific CustomerContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**customerContact** | [**CustomerContact**](CustomerContact.md)|  | 

### Return type

[**CustomerContact**](CustomerContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomercontactsPost

> []CustomerContact CustomercontactsPost(ctx, customerContact)

Create a new CustomerContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerContact** | [**CustomerContact**](CustomerContact.md)|  | 

### Return type

[**[]CustomerContact**](CustomerContact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

