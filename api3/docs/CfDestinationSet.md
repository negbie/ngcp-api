# CfDestinationSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the destination set. | [optional] 
**Destinations** | [**[]CfDestinationSetDestinations**](CFDestinationSet_destinations.md) | An array of destinations, each containing the keys \&quot;destination\&quot;, \&quot;timeout\&quot;, \&quot;priority\&quot;. \&quot;destination\&quot; is an address to forward a call. \&quot;timeout\&quot; is the durationin seconds, how long to try this destination before the call is forwarded to the next destination or given up. \&quot;priority\&quot; is an order of this destiation among others destinations for this forward type. \&quot;simple_destination\&quot; is not a control field and is used only for representation purposes as a simple destination format, e.g. \&quot;4312345\&quot; if it is a number, or \&quot;user@domain\&quot; if it is a URI | 
**SubscriberId** | **float32** | The subscriber id this destination set belongs to. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


