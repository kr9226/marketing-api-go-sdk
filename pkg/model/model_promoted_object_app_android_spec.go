/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// android应用信息
type PromotedObjectAppAndroidSpec struct {
	Packname                  *string                     `json:"packname,omitempty"`
	Version                   *string                     `json:"version,omitempty"`
	Icon                      *string                     `json:"icon,omitempty"`
	AverageRating             *string                     `json:"average_rating,omitempty"`
	PackageSize               *string                     `json:"package_size,omitempty"`
	Genres                    *[]string                   `json:"genres,omitempty"`
	PackageDownloadUrl        *string                     `json:"package_download_url,omitempty"`
	ChannelPackageSpec        *[]ChannelPackageSpecStruct `json:"channel_package_spec,omitempty"`
	ChildrenPrivacyProtection ChildrenPrivacyProtection   `json:"children_privacy_protection,omitempty"`
}
