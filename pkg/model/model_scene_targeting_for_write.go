/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 场景定向
type SceneTargetingForWrite struct {
	UnionPositionPackage        *[]int64     `json:"union_position_package,omitempty"`
	ExcludeUnionPositionPackage *[]int64     `json:"exclude_union_position_package,omitempty"`
	DisplayScene                *[]string    `json:"display_scene,omitempty"`
	TencentNews                 *[]string    `json:"tencent_news,omitempty"`
	WechatScene                 *WechatScene `json:"wechat_scene,omitempty"`
	WechatPosition              *[]int64     `json:"wechat_position,omitempty"`
	QbsearchScene               *[]string    `json:"qbsearch_scene,omitempty"`
	WechatChannelsScene         *[]int64     `json:"wechat_channels_scene,omitempty"`
	WechatSearchScene           *[]string    `json:"wechat_search_scene,omitempty"`
	PcScene                     *[]string    `json:"pc_scene,omitempty"`
}
