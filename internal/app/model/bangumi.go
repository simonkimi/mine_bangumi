package model

import (
	"gorm.io/gorm"
)

type Bangumi struct {
	gorm.Model
	OfficialTitle   string `gorm:"default:'official_title'"` // 番剧名称
	Year            string `gorm:"default:''"`               // 番剧年份
	TitleRaw        string `gorm:"default:'title_raw'"`      // 番剧原名
	Season          int    `gorm:"default:1"`                // 番剧季度
	SeasonRaw       string `gorm:"default:''"`               // 番剧季度原名
	GroupName       string `gorm:"default:''"`               // 字幕组
	Dpi             string `gorm:"default:''"`               // 分辨率
	Source          string `gorm:"default:''"`               // 来源
	Subtitle        string `gorm:"default:''"`               // 字幕
	EpsCollect      bool   `gorm:"default:false"`            // 是否已收集
	Offset          int    `gorm:"default:0"`                // 番剧偏移量
	BlackListFilter string `gorm:"default:'720,\\d+-\\d+'"`  // 黑名单过滤器
	WhiteListFilter string `gorm:"default:''"`               // 白名单过滤器
	RssLink         string `gorm:"default:''"`               // 番剧RSS链接
	PosterLink      string `gorm:"default:''"`               // 番剧海报链接
	Added           bool   `gorm:"default:false"`            // 是否已添加
	RuleName        string `gorm:"default:''"`               // 番剧规则名
	SavePath        string `gorm:"default:''"`               // 番剧保存路径
	IsDeleted       bool   `gorm:"default:false"`            // 是否已删除
}
