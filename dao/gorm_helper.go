package dao

import (
	"go-template/framework"
	"gorm.io/gorm"
)

type GormHelper[T any] struct {
	*framework.Gorm
}

func NewGormHelper[T any]() *GormHelper[T] {
	return &GormHelper[T]{
		Gorm: framework.GormInstance(),
	}
}

func (receiver *GormHelper[T]) Model() *T {
	var entity T
	return &entity
}

func (receiver *GormHelper[T]) Models() []*T {
	var entities []*T
	return entities
}

func (receiver *GormHelper[T]) FindById(id int64) *T {
	return receiver.Builder().Eq("id", id).One()
}

func (receiver *GormHelper[T]) All() []*T {
	return receiver.Builder().List()
}

func (receiver *GormHelper[T]) UpdateFieldById(id int64, field string, value any) int64 {
	return receiver.Builder().Eq("id", id).UpdateField(field, value)
}

func (receiver *GormHelper[T]) UpdateModelById(id int64, value T) int64 {
	return receiver.Builder().Eq("id", id).UpdateModel(value)
}

func (receiver *GormHelper[T]) UpdateFieldsById(id int64, value map[string]any) int64 {
	return receiver.Builder().Eq("id", id).UpdateFields(value)
}

func (receiver *GormHelper[T]) DeleteById(id int64) int64 {
	return receiver.Builder().Eq("id", id).Delete()
}

func (receiver *GormHelper[T]) Save(value interface{}) int64 {
	tx := receiver.Session.Save(value)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return tx.RowsAffected
}

func (receiver *GormHelper[T]) Delete(value interface{}) int64 {
	tx := receiver.Session.Delete(value)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return tx.RowsAffected
}

func (receiver *GormHelper[T]) Builder() *SqlBuilder[T, any] {
	s := &SqlBuilder[T, any]{columns: []string{}, conditions: []string{}, sessionEnable: true}
	s.tx = receiver.Session.Session(&gorm.Session{}).Where("deleted = 0")
	return s
}
