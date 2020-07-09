# \ProfilePackageApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfilepackagesGet**](ProfilePackageApi.md#ProfilepackagesGet) | **Get** /profilepackages/ | Get ProfilePackage items
[**ProfilepackagesIdDelete**](ProfilePackageApi.md#ProfilepackagesIdDelete) | **Delete** /profilepackages/{id} | Delete a specific ProfilePackage
[**ProfilepackagesIdGet**](ProfilePackageApi.md#ProfilepackagesIdGet) | **Get** /profilepackages/{id} | Get a specific ProfilePackage
[**ProfilepackagesIdPatch**](ProfilePackageApi.md#ProfilepackagesIdPatch) | **Patch** /profilepackages/{id} | Change a specific ProfilePackage
[**ProfilepackagesIdPut**](ProfilePackageApi.md#ProfilepackagesIdPut) | **Put** /profilepackages/{id} | Replace/change a specific ProfilePackage
[**ProfilepackagesPost**](ProfilePackageApi.md#ProfilepackagesPost) | **Post** /profilepackages/ | Create a new ProfilePackage



## ProfilepackagesGet

> []ProfilePackage ProfilepackagesGet(ctx, optional)

Get ProfilePackage items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProfilepackagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ProfilepackagesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerId** | **optional.String**| Filter for profile packages belonging to a specific reseller | 
 **name** | **optional.String**| Filter for profile packages with a specific name | 
 **profileName** | **optional.String**| Filter for profile packages containing a billing profile with specific name | 
 **networkName** | **optional.String**| Filter for profile packages containing a billing network with specific name | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]ProfilePackage**](ProfilePackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilepackagesIdDelete

> ProfilepackagesIdDelete(ctx, id)

Delete a specific ProfilePackage

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


## ProfilepackagesIdGet

> ProfilePackage ProfilepackagesIdGet(ctx, id)

Get a specific ProfilePackage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**ProfilePackage**](ProfilePackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilepackagesIdPatch

> ProfilePackage ProfilepackagesIdPatch(ctx, id, operation)

Change a specific ProfilePackage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**operation** | [**[]Operation**](operation.md)| A JSON patch document specifying modifications | 

### Return type

[**ProfilePackage**](ProfilePackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilepackagesIdPut

> ProfilePackage ProfilepackagesIdPut(ctx, id, profilePackage)

Replace/change a specific ProfilePackage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 
**profilePackage** | [**ProfilePackage**](ProfilePackage.md)|  | 

### Return type

[**ProfilePackage**](ProfilePackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilepackagesPost

> []ProfilePackage ProfilepackagesPost(ctx, profilePackage)

Create a new ProfilePackage

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profilePackage** | [**ProfilePackage**](ProfilePackage.md)|  | 

### Return type

[**[]ProfilePackage**](ProfilePackage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

