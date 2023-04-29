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

func newWorkExperience(db *gorm.DB, opts ...gen.DOOption) workExperience {
	_workExperience := workExperience{}

	_workExperience.workExperienceDo.UseDB(db, opts...)
	_workExperience.workExperienceDo.UseModel(&model.WorkExperience{})

	tableName := _workExperience.workExperienceDo.TableName()
	_workExperience.ALL = field.NewAsterisk(tableName)
	_workExperience.ID = field.NewInt64(tableName, "id")
	_workExperience.AdminID = field.NewInt64(tableName, "admin_id")
	_workExperience.ComanyName = field.NewString(tableName, "comany_name")
	_workExperience.Post = field.NewString(tableName, "post")
	_workExperience.WorkContent = field.NewString(tableName, "work_content")
	_workExperience.StartDate = field.NewTime(tableName, "start_date")
	_workExperience.EndDate = field.NewString(tableName, "end_date")

	_workExperience.fillFieldMap()

	return _workExperience
}

type workExperience struct {
	workExperienceDo

	ALL         field.Asterisk
	ID          field.Int64
	AdminID     field.Int64  // 管理员 ID
	ComanyName  field.String // 所在公司名称
	Post        field.String // 岗位
	WorkContent field.String // 工作内容
	StartDate   field.Time   // 在该公司的入职时间
	EndDate     field.String // 在该公司的离职时间，便于使用“至今”

	fieldMap map[string]field.Expr
}

func (w workExperience) Table(newTableName string) *workExperience {
	w.workExperienceDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w workExperience) As(alias string) *workExperience {
	w.workExperienceDo.DO = *(w.workExperienceDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *workExperience) updateTableName(table string) *workExperience {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewInt64(table, "id")
	w.AdminID = field.NewInt64(table, "admin_id")
	w.ComanyName = field.NewString(table, "comany_name")
	w.Post = field.NewString(table, "post")
	w.WorkContent = field.NewString(table, "work_content")
	w.StartDate = field.NewTime(table, "start_date")
	w.EndDate = field.NewString(table, "end_date")

	w.fillFieldMap()

	return w
}

func (w *workExperience) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *workExperience) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 7)
	w.fieldMap["id"] = w.ID
	w.fieldMap["admin_id"] = w.AdminID
	w.fieldMap["comany_name"] = w.ComanyName
	w.fieldMap["post"] = w.Post
	w.fieldMap["work_content"] = w.WorkContent
	w.fieldMap["start_date"] = w.StartDate
	w.fieldMap["end_date"] = w.EndDate
}

func (w workExperience) clone(db *gorm.DB) workExperience {
	w.workExperienceDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w workExperience) replaceDB(db *gorm.DB) workExperience {
	w.workExperienceDo.ReplaceDB(db)
	return w
}

type workExperienceDo struct{ gen.DO }

func (w workExperienceDo) Debug() *workExperienceDo {
	return w.withDO(w.DO.Debug())
}

func (w workExperienceDo) WithContext(ctx context.Context) *workExperienceDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w workExperienceDo) ReadDB() *workExperienceDo {
	return w.Clauses(dbresolver.Read)
}

func (w workExperienceDo) WriteDB() *workExperienceDo {
	return w.Clauses(dbresolver.Write)
}

func (w workExperienceDo) Session(config *gorm.Session) *workExperienceDo {
	return w.withDO(w.DO.Session(config))
}

func (w workExperienceDo) Clauses(conds ...clause.Expression) *workExperienceDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w workExperienceDo) Returning(value interface{}, columns ...string) *workExperienceDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w workExperienceDo) Not(conds ...gen.Condition) *workExperienceDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w workExperienceDo) Or(conds ...gen.Condition) *workExperienceDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w workExperienceDo) Select(conds ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w workExperienceDo) Where(conds ...gen.Condition) *workExperienceDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w workExperienceDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *workExperienceDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w workExperienceDo) Order(conds ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w workExperienceDo) Distinct(cols ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w workExperienceDo) Omit(cols ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w workExperienceDo) Join(table schema.Tabler, on ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w workExperienceDo) LeftJoin(table schema.Tabler, on ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w workExperienceDo) RightJoin(table schema.Tabler, on ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w workExperienceDo) Group(cols ...field.Expr) *workExperienceDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w workExperienceDo) Having(conds ...gen.Condition) *workExperienceDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w workExperienceDo) Limit(limit int) *workExperienceDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w workExperienceDo) Offset(offset int) *workExperienceDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w workExperienceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *workExperienceDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w workExperienceDo) Unscoped() *workExperienceDo {
	return w.withDO(w.DO.Unscoped())
}

func (w workExperienceDo) Create(values ...*model.WorkExperience) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w workExperienceDo) CreateInBatches(values []*model.WorkExperience, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w workExperienceDo) Save(values ...*model.WorkExperience) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w workExperienceDo) First() (*model.WorkExperience, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkExperience), nil
	}
}

func (w workExperienceDo) Take() (*model.WorkExperience, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkExperience), nil
	}
}

func (w workExperienceDo) Last() (*model.WorkExperience, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkExperience), nil
	}
}

func (w workExperienceDo) Find() ([]*model.WorkExperience, error) {
	result, err := w.DO.Find()
	return result.([]*model.WorkExperience), err
}

func (w workExperienceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WorkExperience, err error) {
	buf := make([]*model.WorkExperience, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w workExperienceDo) FindInBatches(result *[]*model.WorkExperience, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w workExperienceDo) Attrs(attrs ...field.AssignExpr) *workExperienceDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w workExperienceDo) Assign(attrs ...field.AssignExpr) *workExperienceDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w workExperienceDo) Joins(fields ...field.RelationField) *workExperienceDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w workExperienceDo) Preload(fields ...field.RelationField) *workExperienceDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w workExperienceDo) FirstOrInit() (*model.WorkExperience, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkExperience), nil
	}
}

func (w workExperienceDo) FirstOrCreate() (*model.WorkExperience, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WorkExperience), nil
	}
}

func (w workExperienceDo) FindByPage(offset int, limit int) (result []*model.WorkExperience, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w workExperienceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w workExperienceDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w workExperienceDo) Delete(models ...*model.WorkExperience) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *workExperienceDo) withDO(do gen.Dao) *workExperienceDo {
	w.DO = *do.(*gen.DO)
	return w
}