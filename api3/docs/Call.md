# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeerAuthUser** | **string** | The authentication username used for outbound authentication. | 
**DestinationCarrierBillingFeeId** | **float32** | The billing fee id used to calculate the destination carrier cost. | 
**DestinationGpp1** | **string** | General Purpose Parameter 1. | 
**SourceCustomerBillingZoneId** | **float32** | The source billing zone id (from billing.billing_zones_history) attached to the customer billing cost. | 
**SourceCarrierFreeTime** | **float32** | The number of free seconds of the carrier contract for this call. | 
**RatedAt** | **string** | The timestamp when the rating occurred. | 
**SourceGpp8** | **string** | General Purpose Parameter 8. | 
**SourceResellerBillingZoneId** | **float32** | The source billing zone id (from billing.billing_zones_history) attached to the reseller billing cost. | 
**ExportedAt** | **string** | The timestamp when the exporting occurred. | 
**SourceCarrierCost** | **float32** | The cost for the operator towards the peering carrier. | 
**DestinationResellerCost** | **float32** | The cost for the reseller of the called party towards the system operator. | 
**DestinationResellerBillingFeeId** | **float32** | The billing fee id used to calculate the destination reseller cost. | 
**DestinationProviderId** | **string** | Reseller contract id of called subscriber, or contract id of peer if to external. | 
**SourceIp** | **string** | IP address of calling party. | 
**SourceUserId** | **string** | UUID of calling subscriber, or 0 if from external. | 
**DestinationDomainIn** | **string** | Destination domain as received by the system from calling party after internal rewriting. | [optional] 
**SourceGpp3** | **string** | General Purpose Parameter 3. | 
**StartTime** | **string** | The timestamp of the call connection. | [optional] 
**DestinationUserDialed** | **string** | Destination username or number as received by the system from calling party before any internal rewriting. | [optional] 
**DestinationCustomerFreeTime** | **float32** | The number of free seconds of the customer used for this call. | 
**DestinationGpp0** | **string** | General Purpose Parameter 0. | 
**PeerAuthRealm** | **string** | The authentication realm (domain) used for outbound authentication. | 
**DestinationCarrierFreeTime** | **float32** | The number of free seconds of the carrier contract for this call. | 
**DestinationGpp4** | **string** | General Purpose Parameter 4. | 
**SourceLnpPrefix** | **string** | LNP prefix for the calling peer from looking up the calling party. | [optional] 
**SourceGpp9** | **string** | General Purpose Parameter 9. | 
**DestinationExternalSubscriberId** | **string** | External id of called subscriber, if local. | 
**SourceCarrierBillingFeeId** | **float32** | The billing fee id used to calculate the source carrier cost. | 
**DestinationGpp9** | **string** | General Purpose Parameter 9. | 
**CallId** | **string** | The value of the SIP Call-ID header for this call. | [optional] 
**SourceCarrierBillingZoneId** | **float32** | The source billing zone id (from billing.billing_zones_history) attached to the carrier billing cost. | 
**SourceCli** | **string** | Normalized CLI (usually E164) of calling party. | [optional] 
**SourceGpp1** | **string** | General Purpose Parameter 1. | 
**SourceCustomerId** | **string** | Customer id of calling subscriber, if local. | 
**DestinationUserId** | **string** | UUID of called subscriber, or 0 if to external. | 
**SourceCustomerCost** | **float32** | The cost for the calling party customer towards the reseller. | 
**CallType** | **string** | The type of call, one of call, cfu, cfb, cft, cfna, cfs, cfr. | [optional] 
**Duration** | **float32** | The duration of the call. | [optional] 
**SourceDomain** | **string** | Domain of calling party. | [optional] 
**CallCode** | **string** | The final SIP response code of the call. | [optional] 
**SourceUser** | **string** | Username of calling party. | [optional] 
**ExportStatus** | **string** | The status of the exporting process. | 
**RatingStatus** | **string** | The status of the rating process. | 
**DestinationGpp6** | **string** | General Purpose Parameter 6. | 
**SourceProviderId** | **string** | Reseller contract id of calling subscriber, or contract id of peer if from external. | 
**SourceGpp4** | **string** | General Purpose Parameter 4. | 
**DestinationCustomerBillingFeeId** | **float32** | The billing fee id used to calculate the destination customer cost. | 
**DestinationCustomerCost** | **float32** | The cost for the called party customer towards the reseller. | 
**SourceGpp6** | **string** | General Purpose Parameter 6. | 
**DestinationGpp2** | **string** | General Purpose Parameter 2. | 
**SourceExternalSubscriberId** | **string** | External ID of calling subscriber, if local. | 
**DestinationUserIn** | **string** | Destination username or number as received by the system from calling party after internal rewriting. | [optional] 
**DestinationGpp8** | **string** | General Purpose Parameter 8. | 
**DestinationUserOut** | **string** | Called party username/number as it leaves the system, after applying called party&#39;s out rewrite rules. | [optional] 
**SourceGpp0** | **string** | General Purpose Parameter 0. | 
**DestinationLnpPrefix** | **string** | LNP prefix for the called party or it&#39;s destination peer. | [optional] 
**DestinationUser** | **string** | Username or number of called party. | [optional] 
**SourceCustomerBillingFeeId** | **float32** | The billing fee id used to calculate the source customer cost. | 
**SourceResellerCost** | **float32** | The cost for the reseller of the calling party towards the system operator. | 
**DestinationCarrierCost** | **float32** | The cost for the carrier towards the operator (e.g. for 800-numbers). | 
**SourceGpp7** | **string** | General Purpose Parameter 7. | 
**SourceUserOut** | **string** | Calling party username/number as it leaves the system, after applying called party&#39;s out rewrite rules. | [optional] 
**DestinationGpp5** | **string** | General Purpose Parameter 5. | 
**SourceExternalContractId** | **string** | External ID of the calling subscriber&#39;s customer, if local. | 
**DestinationCustomerBillingZoneId** | **float32** | The destination billing zone id (from billing.billing_zones_history) attached to the customer billing cost. | 
**DestinationResellerBillingZoneId** | **float32** | The destination billing zone id (from billing.billing_zones_history) attached to the reseller billing cost. | 
**DestinationGpp7** | **string** | General Purpose Parameter 7. | 
**DestinationResellerFreeTime** | **float32** | The number of free seconds of the reseller used for this call. | 
**DestinationCarrierBillingZoneId** | **float32** | The destination billing zone id (from billing.billing_zones_history) attached to the carrier billing cost. | 
**SourceGpp2** | **string** | General Purpose Parameter 2. | 
**CallStatus** | **string** | The status of the call, one of ok, busy, noanswer, cancel, offline, timeout, other. | [optional] 
**DestinationGpp3** | **string** | General Purpose Parameter 3. | 
**SourceResellerFreeTime** | **float32** | The number of free seconds of the reseller used for this call. | 
**DestinationDomain** | **string** | Domain of called party. | [optional] 
**SourceResellerBillingFeeId** | **float32** | The billing fee id used to calculate the source reseller cost. | 
**InitTime** | **string** | The timestamp of the call initiation. | [optional] 
**SourceGpp5** | **string** | General Purpose Parameter 5. | 
**DestinationExternalContractId** | **string** | External id of called subscriber&#39;s customer, if local. | 
**DestinationCustomerId** | **string** | Customer id of called subscriber, if local. | 
**SourceClir** | **bool** | Whether calling party number was suppressed (CLIR). | [optional] 
**SourceCustomerFreeTime** | **float32** | The number of free seconds of the customer used for this call. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


