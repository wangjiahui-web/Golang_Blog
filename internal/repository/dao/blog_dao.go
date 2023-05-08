package dao

import (
	"blogProject/model"
	"blogProject/pkg/orm"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var BlogDao blogDao

type blogDao struct {}

//
// GetPagedPictureList 分页查询图片信息
// @parameter pageNum 页码(从 1 开始)
// @parameter pageSize 每页的记录数
// @return    *orm.Page 分页对象
//

func (blogDao) GetPagedPictureList(db *gorm.DB, pageNum int, pageSize int)(pictureList []model.Picture, err error) {
	var totalRows int64
	count := db.Table("picture").Count(&totalRows)
	if count.Error != nil {
		zap.S().Errorf("查询个数失败: %v", count.Error)
		return nil, count.Error
	}
	pictureList = make([]model.Picture, 0)

	find := db.Table("picture").
		Scopes(orm.Paginate(pageNum, pageSize)).
		Find(&pictureList)
	if find.Error != nil {
		zap.S().Errorf("查询图片列表失败: %v", find.Error)
		return nil, find.Error
	}
	return pictureList , nil
}



func (blogDao) GetBlogSummaryList(db *gorm.DB, pageNum int, pageSize int)(blogList []model.Blog, err error) {
	var totalRows int64
	count := db.Table("blog").Count(&totalRows)
	if count.Error != nil {
		zap.S().Errorf("查询个数失败: %v", count.Error)
		return nil, count.Error
	}
	blogList = make([]model.Blog, 0)

	find := db.Table("blog").
		Scopes(orm.Paginate(pageNum, pageSize)).
		Find(&blogList)
	if find.Error != nil {
		zap.S().Errorf("查询博客列表失败: %v", find.Error)
		return nil, find.Error
	}
	return blogList, nil
}


/*func (blogDao) GetBlogDetail(db *gorm.DB, id int)(AllList []model.Comment, err error) {
	var comments []model.Comment
	find := db.Table("comment").
		Preload("Blog").
		Where("blog_id = ?", id).Find(&comments)
	if find.Error != nil {
		zap.S().Errorf("查询评论列表失败: %v", find.Error)
		return nil, find.Error
	}
	return comments, nil
}
*/

func (blogDao) GetBlogByID(db *gorm.DB, id int) (*model.Blog, error) {
	var blog model.Blog
	if err := db.Where("id = ?", id).First(&blog).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("blog with ID %d not found", id)
		}
		return nil, err
	}
	return &blog, nil
}
func (blogDao) GetBlogComments(db *gorm.DB, id int) ([]*model.Comment, error) {
	var comments []*model.Comment
	find := db.Table("comment").
		Preload("Blog").
		Where("blog_id = ?", id).Find(&comments)
	if find.Error != nil {
		zap.S().Errorf("查询评论列表失败: %v", find.Error)
		return nil, find.Error
	}
	return comments, nil
}