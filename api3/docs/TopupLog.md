# TopupLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileAfterId** | **float32** | The contract&#39;s actual billing profile after the top-up. | 
**Type** | **string** | The top-up request type. | [optional] 
**LockLevelBefore** | **string** | The contract&#39;s subscribers&#39; lock levels before the top-up. | 
**CashBalanceBefore** | **float32** | The contract&#39;s cash balance before the top-up in Euro/USD/etc. | 
**PackageAfterId** | **float32** | The contract&#39;s profile package after the top-up. | 
**Username** | **string** | The user that attempted the topup. | [optional] 
**CashBalanceAfter** | **float32** | The contract&#39;s cash balance after the top-up in Euro/USD/etc. | 
**Amount** | **float32** | The top-up amount in Euro/USD/etc. | 
**RequestToken** | **string** | The external ID to identify top-up request. | 
**VoucherId** | **float32** | The voucher in case of a voucher top-up. | 
**Outcome** | **string** | The top-up operation outcome. | [optional] 
**ContractBalanceAfterId** | **float32** | The contract&#39;s balance interval after the top-up. | 
**LockLevelAfter** | **string** | The contract&#39;s subscribers&#39; lock levels after the top-up. | 
**Message** | **string** | The top-up request response message (error reason). | 
**ContractBalanceBeforeId** | **float32** | The contract&#39;s balance interval before the top-up. | 
**ContractId** | **float32** | The subscriber&#39;s customer contract. | 
**PackageBeforeId** | **float32** | The contract&#39;s profile package before the top-up. | 
**SubscriberId** | **float32** | The subscriber for which to topup the balance. | 
**ProfileBeforeId** | **float32** | The contract&#39;s actual billing profile before the top-up. | 
**Timestamp** | [**time.Time**](time.Time.md) | The timestamp of the topup attempt. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


