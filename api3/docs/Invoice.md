# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodEnd** | [**time.Time**](time.Time.md) | The end of the invoice period. | 
**Period** | **string** | Invoice period. | 
**TemplateId** | **float32** | The template is used to generate invoice. | [optional] 
**AmountVat** | **float32** | The vat amount of the invoice in USD, EUR etc. | 
**SentDate** | [**time.Time**](time.Time.md) | The date the invoice has been sent by email or null if not sent. | 
**CustomerId** | **float32** | The customer this invoice belongs to. | [optional] 
**PeriodStart** | [**time.Time**](time.Time.md) | The start of the invoice period. | 
**Serial** | **string** | The invoice serial number. | 
**AmountTotal** | **float32** | The vat amount of the invoice in USD, EUR etc. | 
**AmountNet** | **float32** | The net amount of the invoice in USD, EUR etc. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


