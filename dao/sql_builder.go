package dao

import (
	"fmt"
	"gorm.io/gorm"
)

type SqlBuilder[M any, F any] struct {
	tx            *gorm.DB
	columns       []string
	conditions    []string
	sort          string
	offset        int
	limit         int
	rowsAffected  int64
	sessionEnable bool
}

func (s *SqlBuilder[M, F]) Select(fields ...string) *SqlBuilder[M, F] {
	s.tx.Select(fields)
	return s
}

func (s *SqlBuilder[M, F]) Page(pageNo int, pageSize int) *SqlBuilder[M, F] {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (pageNo - 1) * pageSize
	s.tx.Offset(offset).Limit(pageSize)

	s.offset = offset
	s.limit = pageSize
	return s
}

func (s *SqlBuilder[M, F]) Eq(field string, value any) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s = ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s = %v", field, value))
	return s
}

func (s *SqlBuilder[M, F]) Ne(field string, value any) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s != ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s != %s", field, value))
	return s
}

func (s *SqlBuilder[M, F]) Gt(field string, value any) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s > ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s > %s", field, value))
	return s
}

func (s *SqlBuilder[M, F]) Lt(field string, value any) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s < ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s < %s", field, value))
	return s
}

func (s *SqlBuilder[M, F]) Ge(field string, value any) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s >= ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s >= %s", field, value))
	return s
}

func (s *SqlBuilder[M, F]) Le(field string, value any) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s <= ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s <= %s", field, value))
	return s
}

func (s *SqlBuilder[M, F]) In(field string, value any) *SqlBuilder[M, F] {
	s.checkListParam(value)
	s.tx.Where(fmt.Sprintf("%s in ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s in (%s)", field, value))
	return s
}

func (s *SqlBuilder[M, F]) NotIn(field string, value any) *SqlBuilder[M, F] {
	s.checkListParam(value)

	s.tx.Where(fmt.Sprintf("%s not in ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s not in (%s)", field, value))
	return s
}

func (s *SqlBuilder[M, F]) IsNull(field string) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s is null", field))
	s.conditions = append(s.conditions, fmt.Sprintf("%s is null", field))
	return s
}

func (s *SqlBuilder[M, F]) NotNull(field string) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s is not null", field))
	s.conditions = append(s.conditions, fmt.Sprintf("%s is not null", field))
	return s
}

func (s *SqlBuilder[M, F]) Like(field string, value any) *SqlBuilder[M, F] {
	s.tx.Where(fmt.Sprintf("%s like ?", field), value)
	s.conditions = append(s.conditions, fmt.Sprintf("%s like '%v'", field, value))
	return s
}

func (s *SqlBuilder[M, F]) Order(order string) *SqlBuilder[M, F] {
	s.tx.Order(order)
	s.sort = order
	return s
}

func (s *SqlBuilder[M, F]) One() *M {
	var t []M
	tx := s.tx.Find(&t)
	if tx.Error != nil {
		panic(tx.Error)
	}
	if len(t) == 0 {
		return nil
	}

	return &(t[0])
}

func (s *SqlBuilder[M, F]) List() []*M {
	var t []*M
	tx := s.tx.Find(&t)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return t
}

func (s *SqlBuilder[M, F]) Delete() int64 {
	if len(s.conditions) == 0 {
		panic("需要删除条件")
	}

	var t M
	tx := s.tx.Delete(&t)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return tx.RowsAffected
}

func (s *SqlBuilder[M, F]) Count() int64 {
	var count int64
	var model M
	tx := s.tx.Model(model).Count(&count)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return count
}

func (s *SqlBuilder[M, F]) UpdateField(field string, value any) int64 {
	if len(s.conditions) == 0 {
		panic("需要更新条件")
	}

	var model M
	tx := s.tx.Model(model).UpdateColumn(field, value)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return tx.RowsAffected
}

func (s *SqlBuilder[M, F]) UpdateModel(m M) int64 {
	if len(s.conditions) == 0 {
		panic("需要更新条件")
	}

	var model M
	tx := s.tx.Model(model).UpdateColumns(m)
	if tx.Error != nil {
		panic(tx.Error)
	}

	return tx.RowsAffected
}

func (s *SqlBuilder[M, F]) UpdateFields(m map[string]any) int64 {
	if len(s.conditions) == 0 {
		panic("需要更新条件")
	}

	var model M
	tx := s.tx.Model(model).UpdateColumns(m)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return tx.RowsAffected
}

func (s *SqlBuilder[M, F]) checkListParam(value any) {
	invalid := true
	values, ok := value.([]interface{})
	if ok {
		if len(values) == 0 {
			panic("查询条件不能为空")
		}
		invalid = false
	}

	int64Values, ok2 := value.([]int64)
	if ok2 {
		if len(int64Values) == 0 {
			panic("查询条件不能为空")
		}
		invalid = false
	}

	stringValues, ok3 := value.([]string)
	if ok3 {
		if len(stringValues) == 0 {
			panic("查询条件不能为空")
		}
		invalid = false
	}

	if invalid {
		panic(fmt.Sprintf("字段类型异常: %T", value))
	}
}
