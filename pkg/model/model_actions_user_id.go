/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 用户标识，app数据上报时必填，web数据上报时可以不填user_id，但建议填写，方便后续优化
type ActionsUserId struct {
	HashImei      *string `json:"hash_imei,omitempty"`
	Md5Sha256Imei *string `json:"md5_sha256_imei,omitempty"`
	HashIdfa      *string `json:"hash_idfa,omitempty"`
	Md5Sha256Idfa *string `json:"md5_sha256_idfa,omitempty"`
	GdtOpenid     *string `json:"gdt_openid,omitempty"`
	HashPhone     *string `json:"hash_phone,omitempty"`
	Sha256Phone   *string `json:"sha256_phone,omitempty"`
	HashAndroidId *string `json:"hash_android_id,omitempty"`
	HashOaid      *string `json:"hash_oaid,omitempty"`
	Md5Sha256Oaid *string `json:"md5_sha256_oaid,omitempty"`
	WechatOpenid  *string `json:"wechat_openid,omitempty"`
	WechatUnionid *string `json:"wechat_unionid,omitempty"`
	WechatAppId   *string `json:"wechat_app_id,omitempty"`
	Caid          *string `json:"caid,omitempty"`
	CaidVersion   *int64  `json:"caid_version,omitempty"`
}
