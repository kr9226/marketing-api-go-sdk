/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type FileSchemaGetResponseData struct {
	Scenes *[]string          `json:"scenes,omitempty"`
	Schema *[]MktSchemaDefine `json:"schema,omitempty"`
}
