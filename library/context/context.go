package context

import (
	"context"
	"github.com/spf13/viper"
)

type Context struct {
	Configuration viper.Viper
	GoContext context.Context
}