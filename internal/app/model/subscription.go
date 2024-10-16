package model

import "gorm.io/gorm"
import "gorm.io/datatypes"

// Subscription 订阅
type Subscription struct {
	gorm.Model
	Link            string                      `gorm:"column:link"`              // RSS链接
	IsAggregate     bool                        `gorm:"column:is_aggregate"`      // 聚合器
	Source          string                      `gorm:"column:source"`            // 来源
	BlackListFilter datatypes.JSONSlice[string] `gorm:"column:black_list_filter"` // 黑名单过滤器
	WhiteListFilter datatypes.JSONSlice[string] `gorm:"column:white_list_filter"` // 白名单过滤器
	IsEnabled       bool                        `gorm:"column:is_enabled"`        // 是否启用
	IsActive        bool                        `gorm:"column:is_active"`         // 是否激活(判断完结, 是否有更新)
}

// SubscriptionFile 订阅文件项
type SubscriptionFile struct {
	gorm.Model
	SubscriptionID   uint          `gorm:"column:subscription_id"`
	Subscription     *Subscription `gorm:"foreignKey:SubscriptionID"` // 订阅
	Name             string        `gorm:"column:name"`               // 文件名
	Link             string        `gorm:"column:link"`               // 文件链接
	IsDownload       bool          `gorm:"column:is_download"`        // 是否已经发送给下载器
	DownloaderTaskID string        `gorm:"column:downloader_task_id"` // 下载器任务ID
}
