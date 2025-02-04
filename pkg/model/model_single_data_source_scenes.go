/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 数据源分发详情
type SingleDataSourceScenes struct {
	DataSourceId *int64                `json:"data_source_id,omitempty"`
	Scenes       *[]SingleUpdateScenes `json:"scenes,omitempty"`
	SwitchType   FileDispatchSwitch    `json:"switch_type,omitempty"`
}
