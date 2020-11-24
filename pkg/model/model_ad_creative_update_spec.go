/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告创意结构
type AdCreativeUpdateSpec struct {
	AdcreativeName     string                          `json:"adcreative_name,omitempty"`
	AdcreativeElements *DpAdcreativeCreativeElementsMp `json:"adcreative_elements,omitempty"`
	PageSpec           *DpPageSpec                     `json:"page_spec,omitempty"`
	ButtonTips         string                          `json:"button_tips,omitempty"`
	CouponTitle        string                          `json:"coupon_title,omitempty"`
}
