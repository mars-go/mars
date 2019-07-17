package db

import (
	"reflect"
	"time"
)

var modelType = reflect.TypeOf(Model{})

type Modeler interface {
	PK() uint64
	SetPK(id uint64)
}

type Model struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Model) PK() uint64 {
	return m.ID
}

func (m *Model) SetPK(id uint64) {
	m.ID = id
}

func checkModel(v interface{}) (*Model, bool) {
	var model *Model
	var ok bool
	vf := reflect.ValueOf(v)
	if vf.Kind() == reflect.Ptr {
		tModel := vf.Elem().FieldByName("Model")
		model, ok = tModel.Addr().Interface().(*Model)
	}
	return model, ok
}
