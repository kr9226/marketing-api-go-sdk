/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意形式详情
type CreativePermissions struct {
	CreativeComponents           *[]CreativeComponentsPermit `json:"creative_components,omitempty"`
	SupportImpressionTrackingUrl *bool                       `json:"support_impression_tracking_url,omitempty"`
	SupportClickTrackingUrl      *bool                       `json:"support_click_tracking_url,omitempty"`
	SupportPageTrackUrl          *bool                       `json:"support_page_track_url,omitempty"`
}
