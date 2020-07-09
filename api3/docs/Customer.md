# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProfileDefinition** | **string** | Explicitly declare the way how you want to set billing profiles for this API call. | 
**InvoiceEmailTemplateId** | **float32** | The email template used to notify users about invoice. | 
**BillingProfileId** | **float32** | The billing profile id used to charge this contract, which will become active immediately. This field is required if the profile definition mode is not defined or the &#39;id&#39; mode is used. | 
**Type** | **string** | Either \&quot;sipaccount\&quot; or \&quot;pbxaccount\&quot;. | [optional] 
**Status** | **string** | The status of the contract. | [optional] 
**PassresetEmailTemplateId** | **float32** | The email template used to notify users about password reset. | 
**CreateTimestamp** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) of the creation. | 
**BillingProfiles** | [**[]CustomerBillingProfiles**](Customer_billing_profiles.md) | The billing profile / billing network interval schedule used to charge this contract can be specified. It is represented by an array of objects, each containing the keys \&quot;start\&quot;, \&quot;stop\&quot;, \&quot;profile_id\&quot; and \&quot;network_id\&quot; (/api/customers/ only). When POSTing, it has to contain a single interval with empty \&quot;start\&quot; and \&quot;stop\&quot; fields. Only intervals beginning in the future can be updated afterwards. This field is required if the &#39;profiles&#39; profile definition mode is used. | 
**ActivateTimestamp** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) of the activation. | 
**TerminateTimestamp** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) of the termination. | 
**SubscriberEmailTemplateId** | **float32** | The email template used to notify users about subscriber creation. | 
**InvoiceTemplateId** | **float32** | The invoice template for invoice generation. If none is assigned, no invoice will be generated for this customer. | 
**ModifyTimestamp** | [**time.Time**](time.Time.md) | The datetime (YYYY-MM-DD HH:mm:ss) of the modification. | 
**ContactId** | **float32** | The contact id this contract belongs to. | [optional] 
**VatRate** | **float32** | The VAT rate in percentage (e.g. 20). | 
**ProfilePackageId** | **float32** | The profile package&#39;s id, whose initial profile/networks are to be used to charge this contract. This field is required if the &#39;package&#39; profile definition mode is used. | 
**MaxSubscribers** | **float32** | Optionally set the maximum number of subscribers for this contract. Leave empty for unlimited. | 
**ExternalId** | **string** | A non-unique external ID e.g., provided by a 3rd party provisioning | 
**AddVat** | **bool** | Whether to charge VAT in invoices. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


