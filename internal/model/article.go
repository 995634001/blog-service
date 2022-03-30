package model

// 文章管理
type BlogArticle struct {
	*Model
	ID            uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title         string `gorm:"column:title"`           // 文章标题
	Desc          string `gorm:"column:desc"`            // 文章简述
	CoverImageUrl string `gorm:"column:cover_image_url"` // 封面图片地址
	Content       string `gorm:"column:content"`         // 文章内容
	State         uint   `gorm:"column:state;default:1"` // 状态 0 为禁用、1 为启用
}

func (m *BlogArticle) TableName() string {
	return "blog_article"
}
