package services

import (
	"github.com/bmsandoval/ekube/library/appcontext"
	"github.com/bmsandoval/ekube/services/helloworld"
	"reflect"
)

type Bundle struct {
	HelloworldSvc helloworld.Service
}

var bundlables = []bundlable{
	helloworld.Helpable{},
}

type bundlable interface {
	NewHelper(appCtx appcontext.Context) (interface{}, error)
	ServiceName() string
}

func NewBundle(appCtx appcontext.Context) (*Bundle, error) {
	bundle := &Bundle{}

	for _, bundlable := range bundlables {
		helper, err := bundlable.NewHelper(appCtx)
		if err != nil {
			return nil, err
		}
		SetField(bundle, bundlable.ServiceName(), helper)
	}

	return bundle, nil
}

func SetField(bundler *Bundle, field string, value interface{}) {
	v := reflect.ValueOf(bundler).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
}