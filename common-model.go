package util

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// CommonModel model必须包含字段
type CommonModel struct {
	// ID 主键ID uuidv1
	ID string `gorm:"primarykey;column:id;type:uuid" json:"id"`
	// 创建时间
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	// 最后更新时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
	// DeletedAt 删除时间
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"-"`
}

// BeforeCreate BeforeCreate
func (u *CommonModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV1().String()
	return
}
