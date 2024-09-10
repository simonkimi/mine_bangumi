package model

import "gorm.io/gorm"

type Torrent struct {
	gorm.Model
	BangumiId    uint64   ``
	Bangumi      *Bangumi `gorm:"foreignKey:BangumiId"`
	RssItemId    uint64   ``
	RssItem      *RssItem `gorm:"foreignKey:RssItemId"`
	Name         string   `gorm:"default:''"`    // 种子名称
	Link         string   `gorm:"default:''"`    // 种子链接
	HomePage     string   `gorm:"default:''"`    // 种子主页
	IsDownloaded bool     `gorm:"default:false"` // 是否已下载
}
