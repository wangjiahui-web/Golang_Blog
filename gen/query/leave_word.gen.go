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

func newLeaveWord(db *gorm.DB, opts ...gen.DOOption) leaveWord {
	_leaveWord := leaveWord{}

	_leaveWord.leaveWordDo.UseDB(db, opts...)
	_leaveWord.leaveWordDo.UseModel(&model.LeaveWord{})

	tableName := _leaveWord.leaveWordDo.TableName()
	_leaveWord.ALL = field.NewAsterisk(tableName)
	_leaveWord.ID = field.NewInt64(tableName, "id")
	_leaveWord.Name = field.NewString(tableName, "name")
	_leaveWord.Email = field.NewString(tableName, "email")
	_leaveWord.Subject = field.NewString(tableName, "subject")
	_leaveWord.Content = field.NewString(tableName, "content")
	_leaveWord.CreateTime = field.NewTime(tableName, "create_time")

	_leaveWord.fillFieldMap()

	return _leaveWord
}

type leaveWord struct {
	leaveWordDo

	ALL        field.Asterisk
	ID         field.Int64
	Name       field.String // 姓名
	Email      field.String // 邮箱
	Subject    field.String // 主题
	Content    field.String // 留言内容
	CreateTime field.Time   // 留言时间

	fieldMap map[string]field.Expr
}

func (l leaveWord) Table(newTableName string) *leaveWord {
	l.leaveWordDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l leaveWord) As(alias string) *leaveWord {
	l.leaveWordDo.DO = *(l.leaveWordDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *leaveWord) updateTableName(table string) *leaveWord {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.Name = field.NewString(table, "name")
	l.Email = field.NewString(table, "email")
	l.Subject = field.NewString(table, "subject")
	l.Content = field.NewString(table, "content")
	l.CreateTime = field.NewTime(table, "create_time")

	l.fillFieldMap()

	return l
}

func (l *leaveWord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *leaveWord) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 6)
	l.fieldMap["id"] = l.ID
	l.fieldMap["name"] = l.Name
	l.fieldMap["email"] = l.Email
	l.fieldMap["subject"] = l.Subject
	l.fieldMap["content"] = l.Content
	l.fieldMap["create_time"] = l.CreateTime
}

func (l leaveWord) clone(db *gorm.DB) leaveWord {
	l.leaveWordDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l leaveWord) replaceDB(db *gorm.DB) leaveWord {
	l.leaveWordDo.ReplaceDB(db)
	return l
}

type leaveWordDo struct{ gen.DO }

func (l leaveWordDo) Debug() *leaveWordDo {
	return l.withDO(l.DO.Debug())
}

func (l leaveWordDo) WithContext(ctx context.Context) *leaveWordDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l leaveWordDo) ReadDB() *leaveWordDo {
	return l.Clauses(dbresolver.Read)
}

func (l leaveWordDo) WriteDB() *leaveWordDo {
	return l.Clauses(dbresolver.Write)
}

func (l leaveWordDo) Session(config *gorm.Session) *leaveWordDo {
	return l.withDO(l.DO.Session(config))
}

func (l leaveWordDo) Clauses(conds ...clause.Expression) *leaveWordDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l leaveWordDo) Returning(value interface{}, columns ...string) *leaveWordDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l leaveWordDo) Not(conds ...gen.Condition) *leaveWordDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l leaveWordDo) Or(conds ...gen.Condition) *leaveWordDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l leaveWordDo) Select(conds ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l leaveWordDo) Where(conds ...gen.Condition) *leaveWordDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l leaveWordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *leaveWordDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l leaveWordDo) Order(conds ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l leaveWordDo) Distinct(cols ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l leaveWordDo) Omit(cols ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l leaveWordDo) Join(table schema.Tabler, on ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l leaveWordDo) LeftJoin(table schema.Tabler, on ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l leaveWordDo) RightJoin(table schema.Tabler, on ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l leaveWordDo) Group(cols ...field.Expr) *leaveWordDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l leaveWordDo) Having(conds ...gen.Condition) *leaveWordDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l leaveWordDo) Limit(limit int) *leaveWordDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l leaveWordDo) Offset(offset int) *leaveWordDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l leaveWordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *leaveWordDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l leaveWordDo) Unscoped() *leaveWordDo {
	return l.withDO(l.DO.Unscoped())
}

func (l leaveWordDo) Create(values ...*model.LeaveWord) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l leaveWordDo) CreateInBatches(values []*model.LeaveWord, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l leaveWordDo) Save(values ...*model.LeaveWord) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l leaveWordDo) First() (*model.LeaveWord, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaveWord), nil
	}
}

func (l leaveWordDo) Take() (*model.LeaveWord, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaveWord), nil
	}
}

func (l leaveWordDo) Last() (*model.LeaveWord, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaveWord), nil
	}
}

func (l leaveWordDo) Find() ([]*model.LeaveWord, error) {
	result, err := l.DO.Find()
	return result.([]*model.LeaveWord), err
}

func (l leaveWordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LeaveWord, err error) {
	buf := make([]*model.LeaveWord, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l leaveWordDo) FindInBatches(result *[]*model.LeaveWord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l leaveWordDo) Attrs(attrs ...field.AssignExpr) *leaveWordDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l leaveWordDo) Assign(attrs ...field.AssignExpr) *leaveWordDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l leaveWordDo) Joins(fields ...field.RelationField) *leaveWordDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l leaveWordDo) Preload(fields ...field.RelationField) *leaveWordDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l leaveWordDo) FirstOrInit() (*model.LeaveWord, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaveWord), nil
	}
}

func (l leaveWordDo) FirstOrCreate() (*model.LeaveWord, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LeaveWord), nil
	}
}

func (l leaveWordDo) FindByPage(offset int, limit int) (result []*model.LeaveWord, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l leaveWordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l leaveWordDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l leaveWordDo) Delete(models ...*model.LeaveWord) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *leaveWordDo) withDO(do gen.Dao) *leaveWordDo {
	l.DO = *do.(*gen.DO)
	return l
}