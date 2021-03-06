# CustomerLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The unique name of the location. | [optional] 
**Description** | **string** | Arbitrary text. | [optional] 
**ContractId** | **float32** | The contract id this location belongs to. | [optional] 
**Blocks** | [**[]CustomerLocationBlocks**](CustomerLocation_blocks.md) | An array of location blocks, each containing the keys (base) \&quot;ip\&quot; address and an optional \&quot;mask\&quot; to specify the network portion (subnet prefix length). The specified blocks must not overlap and can uniformly contain either IPv6 addresses or IPv4 addresses. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


