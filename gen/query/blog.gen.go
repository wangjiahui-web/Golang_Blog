// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"blogProject/gen/model"
)

func newBlog(db *gorm.DB, opts ...gen.DOOption) blog {
	_blog := blog{}

	_blog.blogDo.UseDB(db, opts...)
	_blog.blogDo.UseModel(&model.Blog{})

	tableName := _blog.blogDo.TableName()
	_blog.ALL = field.NewAsterisk(tableName)
	_blog.ID = field.NewInt64(tableName, "id")
	_blog.Title = field.NewString(tableName, "title")
	_blog.Summary = field.NewString(tableName, "summary")
	_blog.Content = field.NewString(tableName, "content")
	_blog.PostTime = field.NewTime(tableName, "post_time")
	_blog.ReadCount = field.NewInt32(tableName, "read_count")
	_blog.AdminID = field.NewInt64(tableName, "admin_id")
	_blog.IsDelete = field.NewInt32(tableName, "is_delete")
	_blog.CategoryID = field.NewInt64(tableName, "category_id")
	_blog.FirstImage = field.NewString(tableName, "first_image")

	_blog.fillFieldMap()

	return _blog
}

type blog struct {
	blogDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	Title      field.String // 标题
	Summary    field.String // 博客摘要
	Content    field.String // 内容
	PostTime   field.Time   // 发表时间
	ReadCount  field.Int32  // 阅读量
	AdminID    field.Int64  // 博客发布人 ID
	IsDelete   field.Int32  // 标识是否删除记录
	CategoryID field.Int64  // 分类外键
	FirstImage field.String // 博客首图

	fieldMap map[string]field.Expr
}

func (b blog) Table(newTableName string) *blog {
	b.blogDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b blog) As(alias string) *blog {
	b.blogDo.DO = *(b.blogDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *blog) updateTableName(table string) *blog {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewInt64(table, "id")
	b.Title = field.NewString(table, "title")
	b.Summary = field.NewString(table, "summary")
	b.Content = field.NewString(table, "content")
	b.PostTime = field.NewTime(table, "post_time")
	b.ReadCount = field.NewInt32(table, "read_count")
	b.AdminID = field.NewInt64(table, "admin_id")
	b.IsDelete = field.NewInt32(table, "is_delete")
	b.CategoryID = field.NewInt64(table, "category_id")
	b.FirstImage = field.NewString(table, "first_image")

	b.fillFieldMap()

	return b
}

func (b *blog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *blog) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 10)
	b.fieldMap["id"] = b.ID
	b.fieldMap["title"] = b.Title
	b.fieldMap["summary"] = b.Summary
	b.fieldMap["content"] = b.Content
	b.fieldMap["post_time"] = b.PostTime
	b.fieldMap["read_count"] = b.ReadCount
	b.fieldMap["admin_id"] = b.AdminID
	b.fieldMap["is_delete"] = b.IsDelete
	b.fieldMap["category_id"] = b.CategoryID
	b.fieldMap["first_image"] = b.FirstImage
}

func (b blog) clone(db *gorm.DB) blog {
	b.blogDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b blog) replaceDB(db *gorm.DB) blog {
	b.blogDo.ReplaceDB(db)
	return b
}

type blogDo struct{ gen.DO }

func (b blogDo) Debug() *blogDo {
	return b.withDO(b.DO.Debug())
}

func (b blogDo) WithContext(ctx context.Context) *blogDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b blogDo) ReadDB() *blogDo {
	return b.Clauses(dbresolver.Read)
}

func (b blogDo) WriteDB() *blogDo {
	return b.Clauses(dbresolver.Write)
}

func (b blogDo) Session(config *gorm.Session) *blogDo {
	return b.withDO(b.DO.Session(config))
}

func (b blogDo) Clauses(conds ...clause.Expression) *blogDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b blogDo) Returning(value interface{}, columns ...string) *blogDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b blogDo) Not(conds ...gen.Condition) *blogDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b blogDo) Or(conds ...gen.Condition) *blogDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b blogDo) Select(conds ...field.Expr) *blogDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b blogDo) Where(conds ...gen.Condition) *blogDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b blogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *blogDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b blogDo) Order(conds ...field.Expr) *blogDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b blogDo) Distinct(cols ...field.Expr) *blogDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b blogDo) Omit(cols ...field.Expr) *blogDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b blogDo) Join(table schema.Tabler, on ...field.Expr) *blogDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b blogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *blogDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b blogDo) RightJoin(table schema.Tabler, on ...field.Expr) *blogDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b blogDo) Group(cols ...field.Expr) *blogDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b blogDo) Having(conds ...gen.Condition) *blogDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b blogDo) Limit(limit int) *blogDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b blogDo) Offset(offset int) *blogDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b blogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *blogDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b blogDo) Unscoped() *blogDo {
	return b.withDO(b.DO.Unscoped())
}

func (b blogDo) Create(values ...*model.Blog) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b blogDo) CreateInBatches(values []*model.Blog, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b blogDo) Save(values ...*model.Blog) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b blogDo) First() (*model.Blog, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blog), nil
	}
}

func (b blogDo) Take() (*model.Blog, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blog), nil
	}
}

func (b blogDo) Last() (*model.Blog, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blog), nil
	}
}

func (b blogDo) Find() ([]*model.Blog, error) {
	result, err := b.DO.Find()
	return result.([]*model.Blog), err
}

func (b blogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Blog, err error) {
	buf := make([]*model.Blog, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b blogDo) FindInBatches(result *[]*model.Blog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b blogDo) Attrs(attrs ...field.AssignExpr) *blogDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b blogDo) Assign(attrs ...field.AssignExpr) *blogDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b blogDo) Joins(fields ...field.RelationField) *blogDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b blogDo) Preload(fields ...field.RelationField) *blogDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b blogDo) FirstOrInit() (*model.Blog, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blog), nil
	}
}

func (b blogDo) FirstOrCreate() (*model.Blog, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Blog), nil
	}
}

func (b blogDo) FindByPage(offset int, limit int) (result []*model.Blog, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b blogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b blogDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b blogDo) Delete(models ...*model.Blog) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *blogDo) withDO(do gen.Dao) *blogDo {
	b.DO = *do.(*gen.DO)
	return b
}
