# TencentAds\AsyncReportFilesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncReportFilesGet**](AsyncReportFilesApi.md#AsyncReportFilesGet) | **Get** /async_report_files/get | 获取文件接口


# **AsyncReportFilesGet**
> string AsyncReportFilesGet(ctx, taskId, fileId, optional)
获取文件接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **int64**|  | 
  **fileId** | **int64**|  | 
 **optional** | ***AsyncReportFilesApiAsyncReportFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsyncReportFilesApiAsyncReportFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.Int64**|  | 
 **organizationId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

**string**

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

