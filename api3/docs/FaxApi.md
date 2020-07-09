# \FaxApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FaxesGet**](FaxApi.md#FaxesGet) | **Get** /faxes/ | Get Fax items
[**FaxesIdGet**](FaxApi.md#FaxesIdGet) | **Get** /faxes/{id} | Get a specific Fax
[**FaxesPost**](FaxApi.md#FaxesPost) | **Post** /faxes/ | Create a new Fax



## FaxesGet

> []Fax FaxesGet(ctx, optional)

Get Fax items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FaxesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FaxesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tz** | **optional.String**| Format timestamp according to the optional time zone provided here, e.g. Europe/Berlin. | 
 **useOwnerTz** | **optional.String**| Format timestamp according to the filtered customer&#39;s/subscribers&#39;s inherited time zone. | 
 **subscriberId** | **optional.String**| Filter for faxes belonging to a specific subscriber | 
 **timeFrom** | **optional.String**| Filter for faxes performed after or at the given timestamp. | 
 **timeTo** | **optional.String**| Filter for faxes performed before or at the given timestamp. | 
 **sid** | **optional.String**| Filter for a fax with the specific session id | 
 **numberRewriteMode** | **optional.String**| Force the number normalization logic (available: &#39;default&#39;, &#39;extended&#39;). | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]Fax**](Fax.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FaxesIdGet

> Fax FaxesIdGet(ctx, id)

Get a specific Fax

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**Fax**](Fax.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FaxesPost

> []Fax FaxesPost(ctx, fax)

Create a new Fax

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fax** | [**Fax**](Fax.md)|  | 

### Return type

[**[]Fax**](Fax.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

