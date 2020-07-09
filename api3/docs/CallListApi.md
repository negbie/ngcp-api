# \CallListApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalllistsGet**](CallListApi.md#CalllistsGet) | **Get** /calllists/ | Get CallList items
[**CalllistsIdGet**](CallListApi.md#CalllistsIdGet) | **Get** /calllists/{id} | Get a specific CallList



## CalllistsGet

> []CallList CalllistsGet(ctx, optional)

Get CallList items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CalllistsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CalllistsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tz** | **optional.String**| Format start_time according to the optional time zone provided here, e.g. Europe/Berlin. | 
 **useOwnerTz** | **optional.String**| Format start_time according to the filtered customer&#39;s/subscribers&#39;s inherited time zone. | 
 **subscriberId** | **optional.String**| Filter for calls for a specific subscriber. Either this or customer_id is mandatory if called by admin, reseller or subscriberadmin to filter list down to a specific subscriber in order to properly determine the direction of calls. | 
 **customerId** | **optional.String**| Filter for calls for a specific customer. Either this or subscriber_id is mandatory if called by admin, reseller or subscriberadmin to filter list down to a specific customer. For calls within the same customer_id, the direction will always be \&quot;out\&quot;. | 
 **aliasField** | **optional.String**| Set this parameter for example to \&quot;gpp0\&quot; if you store alias numbers in the gpp0 preference and want to have that value shown as other CLI for calls from or to such a local subscriber. | 
 **status** | **optional.String**| Filter for calls with a specific status. One of \&quot;ok\&quot;, \&quot;busy\&quot;, \&quot;noanswer\&quot;, \&quot;cancel\&quot;, \&quot;offline\&quot;, \&quot;timeout\&quot;, \&quot;other\&quot;. | 
 **statusNe** | **optional.String**| Filter for calls not having a specific status. One of \&quot;ok\&quot;, \&quot;busy\&quot;, \&quot;noanswer\&quot;, \&quot;cancel\&quot;, \&quot;offline\&quot;, \&quot;timeout\&quot;, \&quot;other\&quot;. | 
 **ratingStatus** | **optional.String**| Filter for calls having a specific rating status. Comma separated list of \&quot;ok\&quot;, \&quot;unrated\&quot;, \&quot;failed\&quot;. | 
 **ratingStatusNe** | **optional.String**| Filter for calls not having a specific rating status. Comma separated list of \&quot;ok\&quot;, \&quot;unrated\&quot;, \&quot;failed\&quot;. | 
 **type_** | **optional.String**| Filter for calls with a specific type. One of \&quot;call\&quot;, \&quot;cfu\&quot;, \&quot;cfb\&quot;, \&quot;cft\&quot;, \&quot;cfna\&quot;, \&quot;cfs\&quot;, \&quot;cfr\&quot;. | 
 **typeNe** | **optional.String**| Filter for calls not having a specific type. One of \&quot;call\&quot;, \&quot;cfu\&quot;, \&quot;cfb\&quot;, \&quot;cft\&quot;, \&quot;cfna\&quot;, \&quot;cfs\&quot;, \&quot;cfr\&quot;. | 
 **direction** | **optional.String**| Filter for calls with a specific direction. One of \&quot;in\&quot;, \&quot;out\&quot;. | 
 **startGe** | **optional.String**| Filter for calls starting greater or equal the specified time stamp. | 
 **startLe** | **optional.String**| Filter for calls starting lower or equal the specified time stamp. | 
 **initGe** | **optional.String**| Filter for calls initiated greater or equal the specified time stamp. | 
 **initLe** | **optional.String**| Filter for calls initiated lower or equal the specified time stamp. | 
 **callId** | **optional.String**| Filter for a particular call_id prefix and sort by call leg depth. | 
 **ownCli** | **optional.String**| Filter calls by a specific number that is a part of in or out calls. | 
 **orderBy** | **optional.String**| Order collection by a specific attribute. | 
 **orderByDirection** | **optional.String**| Direction which the collection should be ordered by. Possible values are: asc (default), desc. | 
 **page** | **optional.Int32**| Pagination page which should be displayed (default: 1) | 
 **rows** | **optional.Int32**| Number of rows in one pagination page (default: 10) | 

### Return type

[**[]CallList**](CallList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CalllistsIdGet

> CallList CalllistsIdGet(ctx, id)

Get a specific CallList

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**|  | 

### Return type

[**CallList**](CallList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

