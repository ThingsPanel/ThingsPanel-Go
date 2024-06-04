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

	"project/model"
)

func newSysVersion(db *gorm.DB, opts ...gen.DOOption) sysVersion {
	_sysVersion := sysVersion{}

	_sysVersion.sysVersionDo.UseDB(db, opts...)
	_sysVersion.sysVersionDo.UseModel(&model.SysVersion{})

	tableName := _sysVersion.sysVersionDo.TableName()
	_sysVersion.ALL = field.NewAsterisk(tableName)
	_sysVersion.VersionNumber = field.NewInt32(tableName, "version_number")
	_sysVersion.Version = field.NewString(tableName, "version")

	_sysVersion.fillFieldMap()

	return _sysVersion
}

type sysVersion struct {
	sysVersionDo

	ALL           field.Asterisk
	VersionNumber field.Int32
	Version       field.String

	fieldMap map[string]field.Expr
}

func (s sysVersion) Table(newTableName string) *sysVersion {
	s.sysVersionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysVersion) As(alias string) *sysVersion {
	s.sysVersionDo.DO = *(s.sysVersionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysVersion) updateTableName(table string) *sysVersion {
	s.ALL = field.NewAsterisk(table)
	s.VersionNumber = field.NewInt32(table, "version_number")
	s.Version = field.NewString(table, "version")

	s.fillFieldMap()

	return s
}

func (s *sysVersion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysVersion) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 2)
	s.fieldMap["version_number"] = s.VersionNumber
	s.fieldMap["version"] = s.Version
}

func (s sysVersion) clone(db *gorm.DB) sysVersion {
	s.sysVersionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysVersion) replaceDB(db *gorm.DB) sysVersion {
	s.sysVersionDo.ReplaceDB(db)
	return s
}

type sysVersionDo struct{ gen.DO }

type ISysVersionDo interface {
	gen.SubQuery
	Debug() ISysVersionDo
	WithContext(ctx context.Context) ISysVersionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysVersionDo
	WriteDB() ISysVersionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysVersionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysVersionDo
	Not(conds ...gen.Condition) ISysVersionDo
	Or(conds ...gen.Condition) ISysVersionDo
	Select(conds ...field.Expr) ISysVersionDo
	Where(conds ...gen.Condition) ISysVersionDo
	Order(conds ...field.Expr) ISysVersionDo
	Distinct(cols ...field.Expr) ISysVersionDo
	Omit(cols ...field.Expr) ISysVersionDo
	Join(table schema.Tabler, on ...field.Expr) ISysVersionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysVersionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysVersionDo
	Group(cols ...field.Expr) ISysVersionDo
	Having(conds ...gen.Condition) ISysVersionDo
	Limit(limit int) ISysVersionDo
	Offset(offset int) ISysVersionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysVersionDo
	Unscoped() ISysVersionDo
	Create(values ...*model.SysVersion) error
	CreateInBatches(values []*model.SysVersion, batchSize int) error
	Save(values ...*model.SysVersion) error
	First() (*model.SysVersion, error)
	Take() (*model.SysVersion, error)
	Last() (*model.SysVersion, error)
	Find() ([]*model.SysVersion, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysVersion, err error)
	FindInBatches(result *[]*model.SysVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysVersion) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysVersionDo
	Assign(attrs ...field.AssignExpr) ISysVersionDo
	Joins(fields ...field.RelationField) ISysVersionDo
	Preload(fields ...field.RelationField) ISysVersionDo
	FirstOrInit() (*model.SysVersion, error)
	FirstOrCreate() (*model.SysVersion, error)
	FindByPage(offset int, limit int) (result []*model.SysVersion, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysVersionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysVersionDo) Debug() ISysVersionDo {
	return s.withDO(s.DO.Debug())
}

func (s sysVersionDo) WithContext(ctx context.Context) ISysVersionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysVersionDo) ReadDB() ISysVersionDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysVersionDo) WriteDB() ISysVersionDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysVersionDo) Session(config *gorm.Session) ISysVersionDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysVersionDo) Clauses(conds ...clause.Expression) ISysVersionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysVersionDo) Returning(value interface{}, columns ...string) ISysVersionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysVersionDo) Not(conds ...gen.Condition) ISysVersionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysVersionDo) Or(conds ...gen.Condition) ISysVersionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysVersionDo) Select(conds ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysVersionDo) Where(conds ...gen.Condition) ISysVersionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysVersionDo) Order(conds ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysVersionDo) Distinct(cols ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysVersionDo) Omit(cols ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysVersionDo) Join(table schema.Tabler, on ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysVersionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysVersionDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysVersionDo) Group(cols ...field.Expr) ISysVersionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysVersionDo) Having(conds ...gen.Condition) ISysVersionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysVersionDo) Limit(limit int) ISysVersionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysVersionDo) Offset(offset int) ISysVersionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysVersionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysVersionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysVersionDo) Unscoped() ISysVersionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysVersionDo) Create(values ...*model.SysVersion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysVersionDo) CreateInBatches(values []*model.SysVersion, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysVersionDo) Save(values ...*model.SysVersion) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysVersionDo) First() (*model.SysVersion, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysVersion), nil
	}
}

func (s sysVersionDo) Take() (*model.SysVersion, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysVersion), nil
	}
}

func (s sysVersionDo) Last() (*model.SysVersion, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysVersion), nil
	}
}

func (s sysVersionDo) Find() ([]*model.SysVersion, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysVersion), err
}

func (s sysVersionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysVersion, err error) {
	buf := make([]*model.SysVersion, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysVersionDo) FindInBatches(result *[]*model.SysVersion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysVersionDo) Attrs(attrs ...field.AssignExpr) ISysVersionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysVersionDo) Assign(attrs ...field.AssignExpr) ISysVersionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysVersionDo) Joins(fields ...field.RelationField) ISysVersionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysVersionDo) Preload(fields ...field.RelationField) ISysVersionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysVersionDo) FirstOrInit() (*model.SysVersion, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysVersion), nil
	}
}

func (s sysVersionDo) FirstOrCreate() (*model.SysVersion, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysVersion), nil
	}
}

func (s sysVersionDo) FindByPage(offset int, limit int) (result []*model.SysVersion, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysVersionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysVersionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysVersionDo) Delete(models ...*model.SysVersion) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysVersionDo) withDO(do gen.Dao) *sysVersionDo {
	s.DO = *do.(*gen.DO)
	return s
}