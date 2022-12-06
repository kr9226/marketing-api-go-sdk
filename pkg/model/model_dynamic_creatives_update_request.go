/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DynamicCreativesUpdateRequest struct {
	DynamicCreativeId                *int64                     `json:"dynamic_creative_id,omitempty"`
	DynamicCreativeTemplateId        *int64                     `json:"dynamic_creative_template_id,omitempty"`
	DynamicCreativeElements          *DynamicCreativeElements   `json:"dynamic_creative_elements,omitempty"`
	DeepLinkUrl                      *string                    `json:"deep_link_url,omitempty"`
	ImpressionTrackingUrl            *string                    `json:"impression_tracking_url,omitempty"`
	ClickTrackingUrl                 *string                    `json:"click_tracking_url,omitempty"`
	FeedsVideoCommentSwitch          *bool                      `json:"feeds_video_comment_switch,omitempty"`
	UnionMarketSwitch                *bool                      `json:"union_market_switch,omitempty"`
	PageType                         DestinationType            `json:"page_type,omitempty"`
	PageSpec                         *DynamicCreativePageSpec   `json:"page_spec,omitempty"`
	LinkPageType                     LinkPageType               `json:"link_page_type,omitempty"`
	LinkNameType                     LinkUrlLinkNameType        `json:"link_name_type,omitempty"`
	LinkPageSpec                     *LinkPageSpec              `json:"link_page_spec,omitempty"`
	ConversionDataType               ConversionDataType         `json:"conversion_data_type,omitempty"`
	ConversionTargetType             ConversionTargetType       `json:"conversion_target_type,omitempty"`
	QqMiniGameTrackingQueryString    *string                    `json:"qq_mini_game_tracking_query_string,omitempty"`
	AndroidDeepLinkAppId             *string                    `json:"android_deep_link_app_id,omitempty"`
	IosDeepLinkAppId                 *string                    `json:"ios_deep_link_app_id,omitempty"`
	UniversalLinkUrl                 *string                    `json:"universal_link_url,omitempty"`
	ShareContentSpec                 *ShareContentSpec          `json:"share_content_spec,omitempty"`
	ProfileId                        *int64                     `json:"profile_id,omitempty"`
	ComponentId                      *int64                     `json:"component_id,omitempty"`
	OnlineEnabled                    *bool                      `json:"online_enabled,omitempty"`
	RevisedAdcreativeSpec            *RevisedAdcreativeSpec     `json:"revised_adcreative_spec,omitempty"`
	VideoEndPage                     *VideoEndPageSpec          `json:"video_end_page,omitempty"`
	WebviewUrl                       *string                    `json:"webview_url,omitempty"`
	SimpleCanvasSubType              SimpleCanvasSubType        `json:"simple_canvas_sub_type,omitempty"`
	FloatingZone                     *FloatingZone              `json:"floating_zone,omitempty"`
	MarketingPendantImageId          *string                    `json:"marketing_pendant_image_id,omitempty"`
	BarrageList                      *[]BarrageListCreateStruct `json:"barrage_list,omitempty"`
	CountdownSwitch                  *bool                      `json:"countdown_switch,omitempty"`
	AppGiftPackCode                  *AppGiftPackCode           `json:"app_gift_pack_code,omitempty"`
	UnionMarketSpec                  *UnionMarketSpec           `json:"union_market_spec,omitempty"`
	AutoDerivedProgramCreativeSwitch *bool                      `json:"auto_derived_program_creative_switch,omitempty"`
	HeadClickType                    HeadClickType              `json:"head_click_type,omitempty"`
	HeadClickSpec                    *HeadClickSpec             `json:"head_click_spec,omitempty"`
	CampaignId                       *int64                     `json:"campaign_id,omitempty"`
	IndustryLabel                    *string                    `json:"industry_label,omitempty"`
	AccountId                        *int64                     `json:"account_id,omitempty"`
}
