package vo

import (
	"blogProject/model"
	"blogProject/pkg/ctime"
)

type BackPicture struct {
	Id        int     `json:"id" gorm:"column:id"`
	ImagePath *string `json:"image_path" gorm:"column:image_path"`
}

type BackBlog struct {
	Id        int64        `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键" json:"id"`
	Title      string       `gorm:"column:title;type:varchar(255);not null;comment:标题" json:"title"`
	Summary    *string      `gorm:"column:summary;type:varchar(512);comment:博客摘要" json:"summary"`
	PostTime   *ctime.CTime `gorm:"column:post_time;type:datetime;default:CURRENT_TIMESTAMP;comment:发表时间" json:"post_time"`
	FirstImage *string      `gorm:"column:first_image;type:varchar(255);comment:博客首图" json:"first_image"`
}


// Other functions and structures, like BlogService, CommentService, BlogDao, etc.
// ...

type CommentTree struct {
	Comment        *model.Comment `json:"comment"`
	ChildComment   []*CommentTree `json:"child_comment"`
	MappedComments map[int64]*CommentTree `json:"-"`
}


type BlogDetailWithComment struct {
	Blog    *model.Blog      `json:"blog"`
	Comments []*CommentTree `json:"comments"`
}