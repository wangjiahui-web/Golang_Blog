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

func newAdmin(db *gorm.DB, opts ...gen.DOOption) admin {
	_admin := admin{}

	_admin.adminDo.UseDB(db, opts...)
	_admin.adminDo.UseModel(&model.Admin{})

	tableName := _admin.adminDo.TableName()
	_admin.ALL = field.NewAsterisk(tableName)
	_admin.ID = field.NewInt64(tableName, "id")
	_admin.Username = field.NewString(tableName, "username")
	_admin.Nickname = field.NewString(tableName, "nickname")
	_admin.Password = field.NewString(tableName, "password")
	_admin.Website = field.NewString(tableName, "website")
	_admin.Birth = field.NewTime(tableName, "birth")
	_admin.Tel = field.NewString(tableName, "tel")
	_admin.Email = field.NewString(tableName, "email")
	_admin.Avatar = field.NewString(tableName, "avatar")
	_admin.City = field.NewString(tableName, "city")
	_admin.School = field.NewString(tableName, "school")

	_admin.fillFieldMap()

	return _admin
}

type admin struct {
	adminDo

	ALL      field.Asterisk
	ID       field.Int64
	Username field.String // 用户名
	Nickname field.String // 昵称
	Password field.String // 密码
	Website  field.String // 博主个人首页
	Birth    field.Time   // 博主生日
	Tel      field.String // 电话号码
	Email    field.String // 电子邮箱
	Avatar   field.String // 头像地址
	City     field.String // 所在城市
	School   field.String // 毕业学校

	fieldMap map[string]field.Expr
}

func (a admin) Table(newTableName string) *admin {
	a.adminDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a admin) As(alias string) *admin {
	a.adminDo.DO = *(a.adminDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *admin) updateTableName(table string) *admin {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Username = field.NewString(table, "username")
	a.Nickname = field.NewString(table, "nickname")
	a.Password = field.NewString(table, "password")
	a.Website = field.NewString(table, "website")
	a.Birth = field.NewTime(table, "birth")
	a.Tel = field.NewString(table, "tel")
	a.Email = field.NewString(table, "email")
	a.Avatar = field.NewString(table, "avatar")
	a.City = field.NewString(table, "city")
	a.School = field.NewString(table, "school")

	a.fillFieldMap()

	return a
}

func (a *admin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *admin) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["id"] = a.ID
	a.fieldMap["username"] = a.Username
	a.fieldMap["nickname"] = a.Nickname
	a.fieldMap["password"] = a.Password
	a.fieldMap["website"] = a.Website
	a.fieldMap["birth"] = a.Birth
	a.fieldMap["tel"] = a.Tel
	a.fieldMap["email"] = a.Email
	a.fieldMap["avatar"] = a.Avatar
	a.fieldMap["city"] = a.City
	a.fieldMap["school"] = a.School
}

func (a admin) clone(db *gorm.DB) admin {
	a.adminDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a admin) replaceDB(db *gorm.DB) admin {
	a.adminDo.ReplaceDB(db)
	return a
}

type adminDo struct{ gen.DO }

func (a adminDo) Debug() *adminDo {
	return a.withDO(a.DO.Debug())
}

func (a adminDo) WithContext(ctx context.Context) *adminDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminDo) ReadDB() *adminDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminDo) WriteDB() *adminDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminDo) Session(config *gorm.Session) *adminDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminDo) Clauses(conds ...clause.Expression) *adminDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminDo) Returning(value interface{}, columns ...string) *adminDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminDo) Not(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminDo) Or(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminDo) Select(conds ...field.Expr) *adminDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminDo) Where(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *adminDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a adminDo) Order(conds ...field.Expr) *adminDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminDo) Distinct(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminDo) Omit(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminDo) Join(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminDo) LeftJoin(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminDo) RightJoin(table schema.Tabler, on ...field.Expr) *adminDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminDo) Group(cols ...field.Expr) *adminDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminDo) Having(conds ...gen.Condition) *adminDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminDo) Limit(limit int) *adminDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminDo) Offset(offset int) *adminDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *adminDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminDo) Unscoped() *adminDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminDo) Create(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminDo) CreateInBatches(values []*model.Admin, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminDo) Save(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminDo) First() (*model.Admin, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Take() (*model.Admin, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Last() (*model.Admin, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Find() ([]*model.Admin, error) {
	result, err := a.DO.Find()
	return result.([]*model.Admin), err
}

func (a adminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Admin, err error) {
	buf := make([]*model.Admin, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminDo) FindInBatches(result *[]*model.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminDo) Attrs(attrs ...field.AssignExpr) *adminDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminDo) Assign(attrs ...field.AssignExpr) *adminDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminDo) Joins(fields ...field.RelationField) *adminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminDo) Preload(fields ...field.RelationField) *adminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminDo) FirstOrInit() (*model.Admin, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FirstOrCreate() (*model.Admin, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FindByPage(offset int, limit int) (result []*model.Admin, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminDo) Delete(models ...*model.Admin) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminDo) withDO(do gen.Dao) *adminDo {
	a.DO = *do.(*gen.DO)
	return a
}
