package dao

import (
	"blogProject/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var CommentDao commentDao
type commentDao struct {}

func (commentDao) GetBlogComments(db *gorm.DB, id int) ([]model.Comment, error) {
	var comments []model.Comment
	// 根据博客 id 查询所有的评论,同时预加载 Blog 信息
	find := db.Table("comment").
		Preload("Blog").
		Where("blog_id = ?", id).Find(&comments)
	if find.Error != nil {
		zap.S().Errorf("查询评论列表失败: %v", find.Error)
		return nil, find.Error
	}
	return comments, nil
}