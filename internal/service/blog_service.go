package service

import (
	"blogProject/internal/repository/dao"
	"blogProject/model"
	"blogProject/model/vo"
	"fmt"
)

var BlogService blogService
type blogService struct {}

func (blogService) GetPagedPictureList(pageNum int, pageSize int) (pictureList []model.Picture ,err error){
	res ,err := dao.BlogDao.GetPagedPictureList(dao.DB, pageNum, pageSize)
	if err != nil {
		return
	}
	fmt.Println("service res",res)
	return res ,nil
}


func (blogService) GetBlogSummaryList(pageNum int, pageSize int) (blogList []model.Blog ,err error){
	res ,err := dao.BlogDao.GetBlogSummaryList(dao.DB, pageNum, pageSize)
	if err != nil {
		return
	}
	return res ,nil
}

/*func buildCommentTree(comments []model.Comment) ([]*vo.CommentTree, error) {
	rootNodes := make([]*vo.CommentTree, 0)

	// Map comments by ID
	mappedComments := make(map[int64]*vo.CommentTree)
	for i := range comments {
		mappedComments[comments[i].ID] = &vo.CommentTree{Comment: &comments[i]}
	}

	// Add comments to their parent nodes
	for _, comment := range comments {
		if comment.ParentID == nil || *comment.ParentID == 0 {
			// Comment has no parent, add to root nodes
			rootNode := mappedComments[comment.ID]
			rootNodes = append(rootNodes, rootNode)
		} else {
			// Comment has parent, add to parent's child comments
			parentNode, ok := mappedComments[*comment.ParentID]
			if !ok {
				continue // Skip this iteration if parentNode is nil
			}
			childNode, ok := mappedComments[comment.ID]
			if !ok {
				continue // Skip this iteration if childNode is nil
			}
			parentNode.ChildComment = append(parentNode.ChildComment, childNode)
		}
		fmt.Println("comment",comment)

	}
	return rootNodes, nil
}
*/
func buildCommentTree(mappedComments map[int64]*vo.CommentTree, comments []*model.Comment, parentID int64) []*vo.CommentTree {
	var result []*vo.CommentTree
	fmt.Println(parentID)
	for _, comment := range comments {
		if comment.ParentID != nil && *comment.ParentID == parentID {
			commentTree := &vo.CommentTree{
				Comment:        comment,
				MappedComments: mappedComments,
			}
			commentTree.ChildComment = buildCommentTree(mappedComments, comments, comment.ID)
			result = append(result, commentTree)
			mappedComments[comment.ID] = commentTree
		}
	}
	return result
}

func (blogService) GetBlogDetailWithComment(id int) (vo.BlogDetailWithComment, error) {
	blog, err := dao.BlogDao.GetBlogByID(dao.DB, id)
	if err != nil {
		return vo.BlogDetailWithComment{}, err
	}

	comments, err := dao.CommentDao.GetBlogComments(dao.DB, id)
	if err != nil {
		return vo.BlogDetailWithComment{}, err
	}

	var c []*model.Comment
	for i := range comments {
		c = append(c, &comments[i])
	}

	// Build comment tree
	mappedComments := make(map[int64]*vo.CommentTree)

	commentTree := buildCommentTree(mappedComments, c, 0)

	var blogDetailWithComment vo.BlogDetailWithComment
	if blog != nil {
		blogDetailWithComment = vo.BlogDetailWithComment{
			Blog:     blog,
			Comments: commentTree,
		}
	} else {
		blogDetailWithComment = vo.BlogDetailWithComment{
			Blog:     nil,
			Comments: commentTree,
		}
	}

	return blogDetailWithComment, nil
}