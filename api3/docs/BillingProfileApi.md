# \BillingProfileApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingprofilesGet**](BillingProfileApi.md#BillingprofilesGet) | **Get** /billingprofiles/ | Get BillingProfile items
[**BillingprofilesIdGet**](BillingProfileApi.md#BillingprofilesIdGet) | **Get** /billingprofiles/{id} | Get a specific BillingProfile
[**BillingprofilesIdPatch**](BillingProfileApi.md#BillingprofilesIdPatch) | **Patch** /billingprofiles/{id} | Change a specific BillingProfile
[**BillingprofilesIdPut**](BillingProfileApi.md#BillingprofilesIdPut) | **Put** /billingprofiles/{id} | Replace/change a specific BillingProfile
[**BillingprofilesPost**](BillingProfileApi.md#BillingprofilesPost) | **Post** /billingprofiles/ | Create a new BillingProfile



## BillingprofilesGet

> []BillingProfile BillingprofilesGet(ctx, optional)

Get BillingProfile items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BillingprofilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a BillingprofilesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for billing profiles belonging to a specific reseller | 
 **handle** | **optional.String**| Filter for billing profiles with a specific handle | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingprofilesIdGet

> BillingProfile BillingprofilesIdGet(ctx, id)

Get a specific BillingProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingprofilesIdPatch

> BillingProfile BillingprofilesIdPatch(ctx, id, operation)

Change a specific BillingProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingprofilesIdPut

> BillingProfile BillingprofilesIdPut(ctx, id, billingProfile)

Replace/change a specific BillingProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**billingProfile** | [**BillingProfile**](BillingProfile.md)|  | 

### Return type

[**BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingprofilesPost

> []BillingProfile BillingprofilesPost(ctx, billingProfile)

Create a new BillingProfile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billingProfile** | [**BillingProfile**](BillingProfile.md)|  | 

### Return type

[**[]BillingProfile**](BillingProfile.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

