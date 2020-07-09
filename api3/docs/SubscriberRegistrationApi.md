# \SubscriberRegistrationApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriberregistrationsGet**](SubscriberRegistrationApi.md#SubscriberregistrationsGet) | **Get** /subscriberregistrations/ | Get SubscriberRegistration items
[**SubscriberregistrationsIdDelete**](SubscriberRegistrationApi.md#SubscriberregistrationsIdDelete) | **Delete** /subscriberregistrations/{id} | Delete a specific SubscriberRegistration
[**SubscriberregistrationsIdGet**](SubscriberRegistrationApi.md#SubscriberregistrationsIdGet) | **Get** /subscriberregistrations/{id} | Get a specific SubscriberRegistration
[**SubscriberregistrationsIdPatch**](SubscriberRegistrationApi.md#SubscriberregistrationsIdPatch) | **Patch** /subscriberregistrations/{id} | Change a specific SubscriberRegistration
[**SubscriberregistrationsIdPut**](SubscriberRegistrationApi.md#SubscriberregistrationsIdPut) | **Put** /subscriberregistrations/{id} | Replace/change a specific SubscriberRegistration
[**SubscriberregistrationsPost**](SubscriberRegistrationApi.md#SubscriberregistrationsPost) | **Post** /subscriberregistrations/ | Create a new SubscriberRegistration



## SubscriberregistrationsGet

> []SubscriberRegistration SubscriberregistrationsGet(ctx, optional)

Get SubscriberRegistration items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubscriberregistrationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SubscriberregistrationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriberId** | **optional.String**| Filter for registrations of a specific subscriber | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]SubscriberRegistration**](SubscriberRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberregistrationsIdDelete

> SubscriberregistrationsIdDelete(ctx, id)

Delete a specific SubscriberRegistration

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


## SubscriberregistrationsIdGet

> SubscriberRegistration SubscriberregistrationsIdGet(ctx, id)

Get a specific SubscriberRegistration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**SubscriberRegistration**](SubscriberRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberregistrationsIdPatch

> SubscriberRegistration SubscriberregistrationsIdPatch(ctx, id, operation)

Change a specific SubscriberRegistration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**SubscriberRegistration**](SubscriberRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberregistrationsIdPut

> SubscriberRegistration SubscriberregistrationsIdPut(ctx, id, subscriberRegistration)

Replace/change a specific SubscriberRegistration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**subscriberRegistration** | [**SubscriberRegistration**](SubscriberRegistration.md)|  | 

### Return type

[**SubscriberRegistration**](SubscriberRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberregistrationsPost

> []SubscriberRegistration SubscriberregistrationsPost(ctx, subscriberRegistration)

Create a new SubscriberRegistration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberRegistration** | [**SubscriberRegistration**](SubscriberRegistration.md)|  | 

### Return type

[**[]SubscriberRegistration**](SubscriberRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

