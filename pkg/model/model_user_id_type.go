/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// UserIdType : 号码包用户ID类型
type UserIdType string

// List of UserIdType
const (
	UserIdType_GDT_OPENID          UserIdType = "GDT_OPENID"
	UserIdType_HASH_IDFA           UserIdType = "HASH_IDFA"
	UserIdType_HASH_IMEI           UserIdType = "HASH_IMEI"
	UserIdType_HASH_MAC            UserIdType = "HASH_MAC"
	UserIdType_HASH_MOBILE_PHONE   UserIdType = "HASH_MOBILE_PHONE"
	UserIdType_HASH_QQ             UserIdType = "HASH_QQ"
	UserIdType_IDFA                UserIdType = "IDFA"
	UserIdType_IMEI                UserIdType = "IMEI"
	UserIdType_MAC                 UserIdType = "MAC"
	UserIdType_MOBILE_QQ_OPENID    UserIdType = "MOBILE_QQ_OPENID"
	UserIdType_QQ                  UserIdType = "QQ"
	UserIdType_WX_OPENID           UserIdType = "WX_OPENID"
	UserIdType_WX_UNIONID          UserIdType = "WX_UNIONID"
	UserIdType_WECHAT_OPENID       UserIdType = "WECHAT_OPENID"
	UserIdType_SALTED_HASH_IMEI    UserIdType = "SALTED_HASH_IMEI"
	UserIdType_SALTED_HASH_IDFA    UserIdType = "SALTED_HASH_IDFA"
	UserIdType_OAID                UserIdType = "OAID"
	UserIdType_HASH_OAID           UserIdType = "HASH_OAID"
	UserIdType_SHA256_MOBILE_PHONE UserIdType = "SHA256_MOBILE_PHONE"
	UserIdType_MD5_SHA256_IMEI     UserIdType = "MD5_SHA256_IMEI"
	UserIdType_MD5_SHA256_IDFA     UserIdType = "MD5_SHA256_IDFA"
	UserIdType_MD5_SHA256_OAID     UserIdType = "MD5_SHA256_OAID"
)
