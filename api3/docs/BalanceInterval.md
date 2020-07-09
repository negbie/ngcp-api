# BalanceInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnderrunProfiles** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) when the cash balance underran the profile package&#39;s underrun_profile_threshold and underrun profiles were applied. | 
**TimelyTopupStart** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) pointing the begin of the time range when top-ups will be considered &#39;timely&#39;. | 
**Start** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) pointing the first second belonging to the balance interval. | 
**IsActual** | **bool** | Is this balance interval the actual one? | 
**Stop** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) pointing the last second belonging to the balance interval. | 
**CashDebit** | **float32** | The amount spent during this interval in EUR/USD/etc. | 
**UnderrunLock** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) when the cash balance underran the profile package&#39;s underrun_lock_threshold and subscribers&#39; lock levels were set. | 
**TopupCount** | **float32** | The number of top-ups performed in this interval. | 
**BillingProfileId** | **float32** | The id of the billing profile at the first second of the balance interval. | 
**NotopupDiscardExpiry** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) pointing the deadline, when the cash balance will be discarded if no top-up was performed. | 
**FreeTimeSpent** | **float32** | The free-time spent during this interval in seconds. | 
**FreeTimeBalance** | **float32** | The interval&#39;s free-time balance of the contract in seconds. | 
**InitialFreeTimeBalance** | **float32** | The interval&#39;s initial free-time balance of the contract in seconds. | 
**TimelyTopupStop** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) pointing the end of the time range until top-ups will be considered &#39;timely&#39;. | 
**TimelyTopupCount** | **float32** | The number of top-ups performed in the &#39;timely&#39; span of this interval. | 
**InitialCashBalance** | **float32** | The interval&#39;s initial cash balance of the contract in EUR/USD/etc. | 
**CashBalance** | **float32** | The interval&#39;s cash balance of the contract in EUR/USD/etc. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


