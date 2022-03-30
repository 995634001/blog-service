package model

// 标签管理
type BlogTag struct {
	*Model
	ID    uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name  string `gorm:"column:name"`            // 标签名称
	State uint   `gorm:"column:state;default:1"` // 状态 0 为禁用、1 为启用
}

func (m *BlogTag) TableName() string {
	return "blog_tag"
}
