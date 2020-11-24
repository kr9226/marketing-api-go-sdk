/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 任务所需条件
type BatchAsyncRequestTaskSpec struct {
	UpdateUnionPositionPackageSpec              *[]UpdateUnionPositionPackageItem              `json:"update_union_position_package_spec,omitempty"`
	UpdateExcludeUnionPositionPackageSpec       *[]UpdateExcludeUnionPositionPackageItem       `json:"update_exclude_union_position_package_spec,omitempty"`
	UpdateTargetingIdSpec                       *[]UpdateTargetingIdItem                       `json:"update_targeting_id_spec,omitempty"`
	UpdateBidStrategySpec                       *[]UpdateBidStrategyItem                       `json:"update_bid_strategy_spec,omitempty"`
	UpdateDeepConversionBehaviorBidSpec         *[]UpdateDeepConversionBehaviorBidItem         `json:"update_deep_conversion_behavior_bid_spec,omitempty"`
	UpdateAdgroupAppAndroidChannelPackageIdSpec *[]UpdateAdgroupAppAndroidChannelPackageIdItem `json:"update_adgroup_app_android_channel_package_id_spec,omitempty"`
	UpdateCampaignSpeedModeSpec                 *[]UpdateCampaignSpeedModeItem                 `json:"update_campaign_speed_mode_spec,omitempty"`
	DeleteCampaignSpec                          *[]DeleteCampaignItem                          `json:"delete_campaign_spec,omitempty"`
	DeleteAdgroupSpec                           *[]DeleteAdgroupItem                           `json:"delete_adgroup_spec,omitempty"`
	DeleteAdSpec                                *[]DeleteAdItem                                `json:"delete_ad_spec,omitempty"`
}
