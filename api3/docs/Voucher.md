# Voucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | **float32** | The profile package the customer will switch with the top-up. | 
**ValidUntil** | [**time.Time**](time.Time.md) | The date until this voucher is valid (YYYY-MM-DD hh:mm:ss). | [optional] 
**ResellerId** | **float32** | The reseller id this voucher belongs to. | [optional] 
**Code** | **string** | The voucher code. | [optional] 
**Amount** | **float32** | The amount of the voucher in cents of Euro/USD/etc. | [optional] 
**CustomerId** | **float32** | The customer contract this voucher can be used by (optional). | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


