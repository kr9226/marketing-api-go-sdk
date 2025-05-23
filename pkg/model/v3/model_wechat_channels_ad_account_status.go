/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// WechatChannelsAdAccountStatus : 视频号开户状态
type WechatChannelsAdAccountStatus string

// List of WechatChannelsAdAccountStatus
const (
	WechatChannelsAdAccountStatus_PENDING           WechatChannelsAdAccountStatus = "PENDING"
	WechatChannelsAdAccountStatus_CREATED           WechatChannelsAdAccountStatus = "CREATED"
	WechatChannelsAdAccountStatus_EXPIRED           WechatChannelsAdAccountStatus = "EXPIRED"
	WechatChannelsAdAccountStatus_AUDIT_PENDING     WechatChannelsAdAccountStatus = "AUDIT_PENDING"
	WechatChannelsAdAccountStatus_AUDIT_REFUSED     WechatChannelsAdAccountStatus = "AUDIT_REFUSED"
	WechatChannelsAdAccountStatus_FROZEN            WechatChannelsAdAccountStatus = "FROZEN"
	WechatChannelsAdAccountStatus_BANNED            WechatChannelsAdAccountStatus = "BANNED"
	WechatChannelsAdAccountStatus_DEACTIVATED       WechatChannelsAdAccountStatus = "DEACTIVATED"
	WechatChannelsAdAccountStatus_CHECK_FAIL        WechatChannelsAdAccountStatus = "CHECK_FAIL"
	WechatChannelsAdAccountStatus_PROCESSING_LOGOUT WechatChannelsAdAccountStatus = "PROCESSING_LOGOUT"
)
