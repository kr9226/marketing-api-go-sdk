/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意元素
type AdcreativeCreativeElementsWithOptions struct {
	Image                      *string                            `json:"image,omitempty"`
	Image2                     *string                            `json:"image2,omitempty"`
	Image3                     *string                            `json:"image3,omitempty"`
	Title                      *string                            `json:"title,omitempty"`
	Description                *string                            `json:"description,omitempty"`
	Corporate                  *AdcreativeCorporate               `json:"corporate,omitempty"`
	Video                      *string                            `json:"video,omitempty"`
	ImageOptions               *[]string                          `json:"image_options,omitempty"`
	TitleOptions               *[]string                          `json:"title_options,omitempty"`
	DescriptionOptions         *[]string                          `json:"description_options,omitempty"`
	VideoOptions               *[]string                          `json:"video_options,omitempty"`
	DeepLinkType               *string                            `json:"deep_link_type,omitempty"`
	LinkNameType               LinkNameType                       `json:"link_name_type,omitempty"`
	ImageList                  *[]string                          `json:"image_list,omitempty"`
	ElementStory               *[]AdcreativeElementStoryArrayItem `json:"element_story,omitempty"`
	Url                        *string                            `json:"url,omitempty"`
	ButtonText                 *string                            `json:"button_text,omitempty"`
	BottomText                 *string                            `json:"bottom_text,omitempty"`
	ExcitationText             *string                            `json:"excitation_text,omitempty"`
	CountdownBegin             *int64                             `json:"countdown_begin,omitempty"`
	CountdownExpiringTimestamp *int64                             `json:"countdown_expiring_timestamp,omitempty"`
	CountdownPrice             *string                            `json:"countdown_price,omitempty"`
	CountdownTimeType          AdCreativeCountdownTimeType        `json:"countdown_time_type,omitempty"`
	MiniProgramId              *string                            `json:"mini_program_id,omitempty"`
	MiniProgramPath            *string                            `json:"mini_program_path,omitempty"`
	Label                      *[]CreativeLabel                   `json:"label,omitempty"`
	ProductTags                *[]string                          `json:"product_tags,omitempty"`
	LogoDescription            *string                            `json:"logo_description,omitempty"`
	Logo                       *string                            `json:"logo,omitempty"`
	LeftBottomTxt              *string                            `json:"left_bottom_txt,omitempty"`
	AnimationEffect            *string                            `json:"animation_effect,omitempty"`
	Phone                      *string                            `json:"phone,omitempty"`
	ShortVideoStruct           *ShortVideoStruct                  `json:"short_video_struct,omitempty"`
	LongVideoStruct            *LongVideoStruct                   `json:"long_video_struct,omitempty"`
	BannerContent              *AdcreativeBannerContent           `json:"banner_content,omitempty"`
	CardContent                *AdcreativeCardContent             `json:"card_content,omitempty"`
	VideoPopupUrl              *string                            `json:"video_popup_url,omitempty"`
	VideoPopupButton           *AdcreativeVideoPopupButton        `json:"video_popup_button,omitempty"`
	ButtonUrl                  *string                            `json:"button_url,omitempty"`
	Brand                      *AdCreativeBrand                   `json:"brand,omitempty"`
	Caption                    *string                            `json:"caption,omitempty"`
	LabelledImg                *AdcreativeLabelledImg             `json:"labelled_img,omitempty"`
	FullScreenImage            *string                            `json:"full_screen_image,omitempty"`
	ZipUrl                     *string                            `json:"zip_url,omitempty"`
	EndPage                    *AdCreativeEndPage                 `json:"end_page,omitempty"`
	HeadLine                   *string                            `json:"head_line,omitempty"`
	ShopImageStruct            *AdCreativeShopImageStruct         `json:"shop_image_struct,omitempty"`
	ChosenButton               *ChosenButton                      `json:"chosen_button,omitempty"`
	LivingDescStruct           *AdCreativeLivingDescStruct        `json:"living_desc_struct,omitempty"`
	FloatingZoneStruct         *FloatingZone                      `json:"floating_zone_struct,omitempty"`
	LeftButton                 *string                            `json:"left_button,omitempty"`
	RightButton                *string                            `json:"right_button,omitempty"`
	LeftCanvas                 *string                            `json:"left_canvas,omitempty"`
	RightCanvas                *string                            `json:"right_canvas,omitempty"`
	WegameInfoSpec             *WegameInfoSpec                    `json:"wegame_info_spec,omitempty"`
	WechatChannelsSpec         *AdCreativeWechatChannelsSpec      `json:"wechat_channels_spec,omitempty"`
	FinderObjectVisibility     *bool                              `json:"finder_object_visibility,omitempty"`
	ImageListJumpInfo          *[]LandingPageStructure            `json:"image_list_jump_info,omitempty"`
	OriginVideo                *string                            `json:"origin_video,omitempty"`
	RedEnvelopeStruct          *RedEnvelopeStruct                 `json:"red_envelope_struct,omitempty"`
	ButtonTextJumpInfo         *LandingPageStructure              `json:"button_text_jump_info,omitempty"`
	WxgamePlayablePageSpec     *WxgamePlayablePageSpec            `json:"wxgame_playable_page_spec,omitempty"`
	MainJumpInfo               *[]LandingPageStructure            `json:"main_jump_info,omitempty"`
}
