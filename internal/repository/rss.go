package repository

import "gorm.io/gorm"

type RssItem struct {
	gorm.Model
	Name        string `gorm:"default:''"`   // RSS名称
	Link        string `gorm:"default:''"`   // RSS链接
	IsAggregate string `gorm:"default:''"`   // 聚合器
	Parser      string `gorm:"default:''"`   // 解析器
	IsEnabled   bool   `gorm:"default:true"` // 是否启用
}
