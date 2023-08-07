package env_printer_plugin_test

import (
	"github.com/sinlov/drone-env-printer/env_printer_plugin"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPlugin(t *testing.T) {
	// mock Plugin
	t.Logf("~> mock Plugin")
	p := env_printer_plugin.Plugin{
		Name:    mockName,
		Version: mockVersion,
	}
	// do Plugin
	t.Logf("~> do Plugin")
	if envCheck(t) {
		return
	}

	// use env:ENV_DEBUG
	p.Config.Debug = envDebug

	//err := p.Exec()
	//if nil == err {
	//	t.Fatal("args [ webhook ] empty error should be catch!")
	//}

	p.Config.PaddingLeftMax = 36
	p.Config.EnvPrintKeys = []string{
		"GOPATH",
		"GOBIN",
	}

	p.Drone = *drone_info.MockDroneInfo("success")
	err := p.Exec()
	// verify Plugin

	err = p.CleanResultEnv()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "", os.Getenv(env_printer_plugin.EnvPluginResultShareHost))
}
