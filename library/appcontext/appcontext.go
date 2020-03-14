package appcontext

import (
	"context"
	"github.com/bmsandoval/ekube/configs"
	"github.com/bmsandoval/ekube/db"
)

type Context struct {
	DB db.Connection
	Config configs.Configuration
	GoContext context.Context
}