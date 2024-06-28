package web

import "mall/models"

// 营销数据映射模型
type Market struct {
	Id          uint64 `gorm:"primaryKey"`   // 编号
	Name        string `gorm:"name"`         // 名称
	Type        int    `gorm:"type"`         // 类型
	BannerImage string `json:"banner_image"` // 活动图片
	BeginTime   string `gorm:"begin_time"`   // 开始时间
	OverTime    string `gorm:"over_time"`    // 结束时间
	GoodsIds    string `gorm:"goods_ids"`    // 关联商品
	Status      int    `gorm:"status"`       // 状态，1-开启，2-关闭
	Created     string `gorm:"created"`      // 创建时间
	Updated     string `gorm:"updated"`      // 更新时间
	Sid         uint64 `gorm:"sid"`          // 店铺编号
}

// 营销创建参数模型
type MarketCreateParam struct {
	Name        string `json:"name" binding:"required"`
	Type        int    `json:"type" binding:"required,gt=0"`
	BannerImage string `json:"bannerImage" binding:"required"`
	BeginTime   string `json:"beginTime" binding:"required"`
	OverTime    string `json:"overTime" binding:"required"`
	GoodsIds    string `json:"goodsIds" binding:"required"`
	Sid         uint64 `json:"sid" binding:"required,gt=0"`
}

// 营销删除参数模型
type MarketDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// 营销更新参数模型
type MarketUpdateParam struct {
	Id          uint64 `json:"id" binding:"required,gt=0"`
	Name        string `json:"name" binding:"required"`
	Type        int    `json:"type" binding:"required"`
	BannerImage string `json:"bannerImage" binding:"required"`
	BeginTime   string `json:"beginTime" binding:"required"`
	OverTime    string `json:"overTime" binding:"required"`
	GoodsIds    string `json:"goodsIds" binding:"required"`
	Status      int    `json:"status" binding:"required,gt=0"`
}

// 营销状态更新参数模型
type MarketStatusUpdateParam struct {
	Id     uint64 `json:"id" binding:"required,gt=0"`
	Status int    `json:"status" binding:"required,gt=0"`
}

// 营销查询参数模型
type MarketQueryParam struct {
	Page   models.Page
	Id     uint64 `form:"id"`
	Type   int    `form:"type"`
	Status int    `form:"status"`
	Sid    uint64 `form:"sid" binding:"required,gt=0"`
}

// 营销商品传输模型
type MarketGoodsQueryParma struct {
	GoodsIds string `form:"goodsIds"`
}

// 营销列表传输模型
type MarketList struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	BannerImage string `json:"bannerImage"`
	BeginTime   string `json:"beginTime"`
	OverTime    string `json:"overTime"`
	GoodsIds    string `json:"goodsIds"`
	Status      int    `json:"status"`
}

// 营销商品传输模型
type MarketGoods struct {
	Id       uint64  `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ImageUrl string  `json:"imageUrl"`
}
