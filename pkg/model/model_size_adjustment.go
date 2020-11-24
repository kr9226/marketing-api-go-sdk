/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 调整尺寸
type SizeAdjustment struct {
	TargetSize           SizeType                 `json:"target_size,omitempty"`
	SizeAdjustmentMethod SizeAdjustmentMethodEnum `json:"size_adjustment_method,omitempty"`
	Color                string                   `json:"color,omitempty"`
}
