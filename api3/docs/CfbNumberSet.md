# CfbNumberSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the bnumber set | [optional] 
**Mode** | **string** | The bnumber set mode. A blacklist matches everything except numbers in the list, a whitelist only matches numbers (or expressions) in this list. | [optional] 
**Bnumbers** | [**[]CfbNumberSetBnumbers**](CFBNumberSet_bnumbers.md) | An array of bnumbers, each containing the key \&quot;bnumber\&quot; which will be matched against the called party number (callee) to determine whether to apply the callforward or not. \&quot;bnumber\&quot; is the callee&#39;s number in E164 format to match. They are formatted either as regular expressions or shell patterns depending on the value of the is_regex flag. | 
**IsRegex** | **bool** | A flag indicating, whether the numbers in this set are regular expressions. If true, all bnumbers will be interepreted as perl compatible regular expressions and matched against the B-Number of the calls. If false, shell pattern the whole numbers are matched while shell patterns like 431* or 49123[1-5]67 are possible. If true, capturing groups can be formed using parentheses and referenced in the destinations via \\\\1, \\\\2,... . | 
**SubscriberId** | **float32** | The subscriber id this b-number set belongs to | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


