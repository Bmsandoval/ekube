package hello

import (
	"github.com/bmsandoval/ekube/db/models"
	"github.com/bmsandoval/ekube/library/appcontext"
)

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
	return "HelloSvc"
}

type Service interface {
	Create(greetingModel models.Greetings) (*models.Greetings, error)
	Get() ([]models.Greetings, error)
}
