package main

import (
	"context"

	"github.com/Peripli/example-plugin/myplugin"
	"github.com/Peripli/service-manager/pkg/sm"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	env := sm.DefaultEnv()
	serviceManager := sm.New(ctx, cancel, env)

	serviceManager.RegisterPlugins(&myplugin.MyPlugin{})

	sm := serviceManager.Build()
	sm.Run()
}
