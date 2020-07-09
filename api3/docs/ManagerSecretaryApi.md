# \ManagerSecretaryApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManagersecretaryGet**](ManagerSecretaryApi.md#ManagersecretaryGet) | **Get** /managersecretary/ | Get ManagerSecretary items
[**ManagersecretaryIdDelete**](ManagerSecretaryApi.md#ManagersecretaryIdDelete) | **Delete** /managersecretary/{id} | Delete a specific ManagerSecretary
[**ManagersecretaryIdGet**](ManagerSecretaryApi.md#ManagersecretaryIdGet) | **Get** /managersecretary/{id} | Get a specific ManagerSecretary
[**ManagersecretaryIdPut**](ManagerSecretaryApi.md#ManagersecretaryIdPut) | **Put** /managersecretary/{id} | Replace/change a specific ManagerSecretary



## ManagersecretaryGet

> []ManagerSecretary ManagersecretaryGet(ctx, optional)

Get ManagerSecretary items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagersecretaryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ManagersecretaryGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]ManagerSecretary**](ManagerSecretary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagersecretaryIdDelete

> ManagersecretaryIdDelete(ctx, id)

Delete a specific ManagerSecretary

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


## ManagersecretaryIdGet

> ManagerSecretary ManagersecretaryIdGet(ctx, id)

Get a specific ManagerSecretary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**ManagerSecretary**](ManagerSecretary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagersecretaryIdPut

> ManagerSecretary ManagersecretaryIdPut(ctx, id, managerSecretary)

Replace/change a specific ManagerSecretary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**managerSecretary** | [**ManagerSecretary**](ManagerSecretary.md)|  | 

### Return type

[**ManagerSecretary**](ManagerSecretary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

