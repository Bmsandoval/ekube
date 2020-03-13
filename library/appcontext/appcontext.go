package appcontext

import (
	"context"
	"github.com/bmsandoval/ekube/configs"
)

type Context struct {
	Config configs.Configuration
	GoContext context.Context
}