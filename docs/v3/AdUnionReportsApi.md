# TencentAds\AdUnionReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdUnionReportsGet**](AdUnionReportsApi.md#AdUnionReportsGet) | **Get** /ad_union_reports/get | 联盟广告位报表接口


# **AdUnionReportsGet**
> AdUnionReportsGetResponse AdUnionReportsGet(ctx, accountId, dateRange, fields, optional)
联盟广告位报表接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dateRange** | [**ReportDateRange**](ReportDateRange.md)|  | 
  **fields** | [**[]string**](string.md)|  | 
 **optional** | ***AdUnionReportsApiAdUnionReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdUnionReportsApiAdUnionReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filtering** | [**optional.Interface of UnionReportFiltering**](UnionReportFiltering.md)|  | 
 **groupBy** | [**optional.Interface of []string**](string.md)|  | 
 **orderBy** | [**optional.Interface of []OrderByStruct**](OrderByStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 

### Return type

[**AdUnionReportsGetResponse**](AdUnionReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

