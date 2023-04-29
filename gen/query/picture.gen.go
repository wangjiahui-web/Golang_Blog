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

func newPicture(db *gorm.DB, opts ...gen.DOOption) picture {
	_picture := picture{}

	_picture.pictureDo.UseDB(db, opts...)
	_picture.pictureDo.UseModel(&model.Picture{})

	tableName := _picture.pictureDo.TableName()
	_picture.ALL = field.NewAsterisk(tableName)
	_picture.ID = field.NewInt64(tableName, "id")
	_picture.Title = field.NewString(tableName, "title")
	_picture.ImagePath = field.NewString(tableName, "image_path")
	_picture.Address = field.NewString(tableName, "address")
	_picture.CreateTime = field.NewTime(tableName, "create_time")

	_picture.fillFieldMap()

	return _picture
}

type picture struct {
	pictureDo

	ALL        field.Asterisk
	ID         field.Int64  // 照片 ID
	Title      field.String // 照片标题
	ImagePath  field.String // 照片路径
	Address    field.String // 照片拍摄地点
	CreateTime field.Time   // 照片上传日期

	fieldMap map[string]field.Expr
}

func (p picture) Table(newTableName string) *picture {
	p.pictureDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p picture) As(alias string) *picture {
	p.pictureDo.DO = *(p.pictureDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *picture) updateTableName(table string) *picture {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.Title = field.NewString(table, "title")
	p.ImagePath = field.NewString(table, "image_path")
	p.Address = field.NewString(table, "address")
	p.CreateTime = field.NewTime(table, "create_time")

	p.fillFieldMap()

	return p
}

func (p *picture) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *picture) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["id"] = p.ID
	p.fieldMap["title"] = p.Title
	p.fieldMap["image_path"] = p.ImagePath
	p.fieldMap["address"] = p.Address
	p.fieldMap["create_time"] = p.CreateTime
}

func (p picture) clone(db *gorm.DB) picture {
	p.pictureDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p picture) replaceDB(db *gorm.DB) picture {
	p.pictureDo.ReplaceDB(db)
	return p
}

type pictureDo struct{ gen.DO }

func (p pictureDo) Debug() *pictureDo {
	return p.withDO(p.DO.Debug())
}

func (p pictureDo) WithContext(ctx context.Context) *pictureDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pictureDo) ReadDB() *pictureDo {
	return p.Clauses(dbresolver.Read)
}

func (p pictureDo) WriteDB() *pictureDo {
	return p.Clauses(dbresolver.Write)
}

func (p pictureDo) Session(config *gorm.Session) *pictureDo {
	return p.withDO(p.DO.Session(config))
}

func (p pictureDo) Clauses(conds ...clause.Expression) *pictureDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pictureDo) Returning(value interface{}, columns ...string) *pictureDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pictureDo) Not(conds ...gen.Condition) *pictureDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pictureDo) Or(conds ...gen.Condition) *pictureDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pictureDo) Select(conds ...field.Expr) *pictureDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pictureDo) Where(conds ...gen.Condition) *pictureDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pictureDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *pictureDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pictureDo) Order(conds ...field.Expr) *pictureDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pictureDo) Distinct(cols ...field.Expr) *pictureDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pictureDo) Omit(cols ...field.Expr) *pictureDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pictureDo) Join(table schema.Tabler, on ...field.Expr) *pictureDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pictureDo) LeftJoin(table schema.Tabler, on ...field.Expr) *pictureDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pictureDo) RightJoin(table schema.Tabler, on ...field.Expr) *pictureDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pictureDo) Group(cols ...field.Expr) *pictureDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pictureDo) Having(conds ...gen.Condition) *pictureDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pictureDo) Limit(limit int) *pictureDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pictureDo) Offset(offset int) *pictureDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pictureDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *pictureDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pictureDo) Unscoped() *pictureDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pictureDo) Create(values ...*model.Picture) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pictureDo) CreateInBatches(values []*model.Picture, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pictureDo) Save(values ...*model.Picture) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pictureDo) First() (*model.Picture, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Picture), nil
	}
}

func (p pictureDo) Take() (*model.Picture, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Picture), nil
	}
}

func (p pictureDo) Last() (*model.Picture, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Picture), nil
	}
}

func (p pictureDo) Find() ([]*model.Picture, error) {
	result, err := p.DO.Find()
	return result.([]*model.Picture), err
}

func (p pictureDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Picture, err error) {
	buf := make([]*model.Picture, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pictureDo) FindInBatches(result *[]*model.Picture, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pictureDo) Attrs(attrs ...field.AssignExpr) *pictureDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pictureDo) Assign(attrs ...field.AssignExpr) *pictureDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pictureDo) Joins(fields ...field.RelationField) *pictureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pictureDo) Preload(fields ...field.RelationField) *pictureDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pictureDo) FirstOrInit() (*model.Picture, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Picture), nil
	}
}

func (p pictureDo) FirstOrCreate() (*model.Picture, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Picture), nil
	}
}

func (p pictureDo) FindByPage(offset int, limit int) (result []*model.Picture, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pictureDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pictureDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pictureDo) Delete(models ...*model.Picture) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pictureDo) withDO(do gen.Dao) *pictureDo {
	p.DO = *do.(*gen.DO)
	return p
}