# BillingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FraudIntervalLimit** | **float32** | The fraud detection threshold per month (in cents, e.g. 10000). | 
**FraudDailyNotify** | **string** | Comma-Separated list of Email addresses to send notifications when tresholds are exceeded. | 
**Currency** | **string** | The currency symbol or ISO code, used on invoices and webinterfaces. | 
**PeaktimeSpecial** | [**[]BillingProfilePeaktimeSpecial**](BillingProfile_peaktime_special.md) | The &#39;special&#39; peak-time schedule for this billing profile. It is represented by an array of objects, each containing the keys \&quot;start\&quot; (YYYY-MM-DD HH:mm:ss) and \&quot;stop\&quot; (YYYY-MM-DD HH:mm:ss). Each time range provided determines when to use a fee&#39;s offpeak rates. | 
**FraudIntervalLock** | **string** | Options to lock customer if the monthly limit is exceeded. | 
**IntervalCharge** | **float32** | The base fee charged per billing interval (a monthly fixed fee, e.g. 100) in cents. This fee can be used on the invoice. | 
**FraudIntervalNotify** | **string** | Comma-Separated list of Email addresses to send notifications when tresholds are exceeded. | 
**AdviceOfCharge** | **bool** | Enable Advice of Charge support for the billing profile. | 
**FraudUseResellerRates** | **bool** | Use rates of the reseller&#39;s billing profile for fraud control. | 
**Name** | **string** | A human readable profile name. | [optional] 
**IntervalFreeCash** | **float32** | The included free money per billing interval (in cents, e.g. 10000). | 
**Prepaid** | **bool** | Whether customers using this profile are handled prepaid. | 
**PrepaidLibrary** | **string** | Which prepaid rating library to use. | 
**ResellerId** | **float32** | The reseller id this profile belongs to. | [optional] 
**IntervalFreeTime** | **float32** | The included free minutes per billing interval (in seconds, e.g. 60000 for 1000 free minutes). | 
**FraudDailyLock** | **string** | Options to lock customer if the daily limit is exceeded. | 
**Handle** | **string** | A unique identifier string (only alphanumeric chars and _). | [optional] 
**FraudDailyLimit** | **float32** | The fraud detection threshold per day (in cents, e.g. 1000). | 
**PeaktimeWeekdays** | [**[]BillingProfilePeaktimeWeekdays**](BillingProfile_peaktime_weekdays.md) | The &#39;weekday&#39; peak-time schedule for this billing profile. It is represented by an array of objects, each containing the keys \&quot;weekday\&quot; (0 .. Monday, 6 .. Sunday), \&quot;start\&quot; (HH:mm:ss) and \&quot;stop\&quot; (HH:mm:ss). Each time range provided determines when to use a fee&#39;s offpeak rates. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


