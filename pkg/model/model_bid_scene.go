/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// BidScene : 出价场景，游戏行业投放端又称\"投放场景\"，该能力与出价类型(smart_bid_type)/出价策略(bid_strategy)互斥使用用
type BidScene string

// List of BidScene
const (
	BidScene_UNKNOWN        BidScene = "BID_SCENE_UNKNOWN"
	BidScene_NORMAL_AVERAGE BidScene = "BID_SCENE_NORMAL_AVERAGE"
	BidScene_NORMAL_TARGET  BidScene = "BID_SCENE_NORMAL_TARGET"
	BidScene_NORMAL_MAX     BidScene = "BID_SCENE_NORMAL_MAX"
)
