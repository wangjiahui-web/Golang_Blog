package api

import (
	"blogProject/internal/pkg/resp"
	"blogProject/internal/service"
	"blogProject/model/vo"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

var BlogApi blogApi
type blogApi struct {}



//
// GetPagedPictureList 分页获取用户信息
// @Summary     分页获取用户信息
// @Description 分页获取用户信息
// @Tags        blogApi
// @Accept      json
// @Produce     json
// @Param       pageNum  query    int             true  "页码(从 1 开始)"
// @Param       pageSize query    int             true  "每页的记录数"
// @Success     200      {object}  vo.BackPicture     "OK"
// @Failure     400      {object} swagger.Http400 "Bad Request"
// @Failure     404      {object} swagger.Http404 "Page Not Found"
// @Failure     500      {object} swagger.Http500 "InternalError"
// @Router      /blog/GetPagedPictureList [get]
func (blogApi) GetPagedPictureList(c *gin.Context) {
	var num = c.Query("pageNum")
	var size = c.Query("pageSize")
	pageNum, err := strconv.Atoi(num)
	if err != nil {
		resp.InternalServerError(c, "pageNum不是数字"+err.Error())
		return
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		resp.InternalServerError(c, "pageSize不是数字"+err.Error())
		return
	}
	res ,err := service.BlogService.GetPagedPictureList(pageNum, pageSize)
	if err != nil {
		resp.InternalServerError(c, err.Error())
		zap.S().Errorf("获取图片信息失败: %s\n", err)
		return
	}
	picturesList := make([]vo.BackPicture, 0, len(res))
	for _, picture := range res {
		var information = vo.BackPicture{
			Id:        int(picture.ID),
			ImagePath: picture.ImagePath,
		}
		picturesList = append(picturesList, information)
	}
	resp.Ok(c, "获取图片信息成功", picturesList)
}


//
// GetBlogSummaryList 分页获取博客信息
// @Summary     分页获取博客信息
// @Description 分页获取博客信息
// @Tags        blogApi
// @Accept      json
// @Produce     json
// @Param       pageNum  query    int             true  "页码(从 1 开始)"
// @Param       pageSize query    int             true  "每页的记录数"
// @Success     200      {object}  vo.BackBlog     "OK"
// @Failure     400      {object} swagger.Http400 "Bad Request"
// @Failure     404      {object} swagger.Http404 "Page Not Found"
// @Failure     500      {object} swagger.Http500 "InternalError"
// @Router      /blog/GetBlogSummaryList [get]
func (blogApi) GetBlogSummaryList(c *gin.Context) {
	var num = c.Query("pageNum")
	var size = c.Query("pageSize")
	pageNum, err := strconv.Atoi(num)
	if err != nil {
		resp.InternalServerError(c, "pageNum不是数字"+err.Error())
		return
	}
	pageSize, err := strconv.Atoi(size)
	if err != nil {
		resp.InternalServerError(c, "pageSize不是数字"+err.Error())
		return
	}
	res ,err := service.BlogService.GetBlogSummaryList(pageNum, pageSize)
	if err != nil {
		resp.InternalServerError(c, err.Error())
		zap.S().Errorf("获取图片信息失败: %s\n", err)
		return
	}
	blogsList := make([]vo.BackBlog, 0, len(res))
	for _ , blog := range res {
		var information = vo.BackBlog{
		 Id :blog.ID,
		 Title: blog.Title,
		 Summary: blog.Summary,
		 PostTime: blog.PostTime,
		 FirstImage: blog.FirstImage,
}
		blogsList = append(blogsList, information)
	}
	resp.Ok(c, "获取博客信息成功", blogsList)
}

// GetBlogDetail 根据博客 ID 查询出博客详情
//
//	@Summary		根据博客 ID 查询出博客详情
//	@Description	根据博客 ID 查询出博客详情
//	@Tags			blogApi
//	@Accept			json
//	@Produce		json
//	@Param			id	   path		    int			false	"博客 ID"
//	@Success		200		{object}	vo.BlogDetailWithComment	"OK"
//	@Failure		400		{object}	swagger.Http400	"Bad Request"
//	@Failure		404		{object}	swagger.Http404	"Page Not Found"
//	@Failure		500		{object}	swagger.Http500	"InternalError"
//	@Router			/blog/GetBlogDetail/id/{id} [get]
func (blogApi) GetBlogDetail(c *gin.Context) {
	// /blog/GetBlogDetail/id/:id
	bid := c.Param("id")
	id, err := strconv.Atoi(bid)
	if err != nil {
		resp.InternalServerError(c, "id不是数字"+err.Error())
		return
	}
	blogDetail, err := service.BlogService.GetBlogDetailWithComment(id) // 调用服务层方法获取博客详情和评论信息
	if err != nil {
		resp.InternalServerError(c, err.Error())
		zap.S().Errorf("获取博客详情和评论信息失败: %s\n", err)
		return
	}
	resp.Ok(c, "获取博客详情和评论信息成功", blogDetail)
}