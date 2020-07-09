# ProfilePackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopupLockLevel** | **string** | The lock level to reset all customer&#39;s subscribers to after a successful top-up (usually 0). | 
**Description** | **string** | Arbitrary text. | [optional] 
**UnderrunLockThreshold** | **float32** | The balance threshold (in cents) for the underrun lock level to come into effect. | 
**InitialProfiles** | [**[]ProfilePackageInitialProfiles**](ProfilePackage_initial_profiles.md) | An array of objects with keys \&quot;profile_id\&quot; and \&quot;network_id\&quot; to create profile mappings from when applying this profile package to a customer. | 
**BalanceIntervalUnit** | **string** | The temporal unit for the balance interval. | 
**CarryOverMode** | **string** | Options to carry over the customer&#39;s balance to the next balance interval. | 
**BalanceIntervalValue** | **float32** | The balance interval in temporal units. | 
**InitialBalance** | **float32** | The initial balance (in cents) that will be set for the very first balance interval. | 
**ResellerId** | **float32** | The reseller id this profile package belongs to. | 
**UnderrunProfiles** | [**[]ProfilePackageInitialProfiles**](ProfilePackage_initial_profiles.md) | An array of objects with keys \&quot;profile_id\&quot; and \&quot;network_id\&quot; to create profile mappings from when the balance underruns the \&quot;underrun_profile_threshold\&quot; value. | 
**UnderrunLockLevel** | **string** | The lock level to set all customer&#39;s subscribers to in case the balance underruns \&quot;underrun_lock_threshold\&quot;. | 
**NotopupDiscardIntervals** | **float32** | The balance will be discarded if no top-up happened for the given number of balance interval units. | 
**UnderrunProfileThreshold** | **float32** | The balance threshold (in cents) for underrun profiles to come into effect. | 
**ServiceCharge** | **float32** | The service charge amount (in cents) will be subtracted from the voucher amount. | 
**TimelyDurationUnit** | **string** | The temporal unit for the \&quot;timely\&quot; interval. | 
**BalanceIntervalStartMode** | **string** | This mode determines when balance intervals start. | 
**Name** | **string** | The unique name of the profile package. | [optional] 
**TopupProfiles** | [**[]ProfilePackageTopupProfiles**](ProfilePackage_topup_profiles.md) | An array of objects with keys \&quot;profile_id\&quot; and \&quot;network_id\&quot; to create profile mappings from when a customer top-ups with a voucher associated with this profile package. | 
**TimelyDurationValue** | **float32** | The \&quot;timely\&quot; interval in temporal units. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


