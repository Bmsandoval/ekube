package helloworld

import "github.com/bmsandoval/ekube/library/appcontext"

type Helper struct {
	AppCtx appcontext.Context
}
type Helpable struct{}

func(h Helpable) NewHelper(appCtx appcontext.Context) (interface{}, error) {
	return &Helper{
		AppCtx: appCtx,
	}, nil
}

func (h Helpable) ServiceName() string {
	return "HelloworldSvc"
}

type Service interface {
	Create()
}
