package pkg

import "context"

type Application struct {
	Envs               map[string]string
	ApplicationContext *context.Context
}
