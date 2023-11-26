package main

import (
	"os"
	"path/filepath"

	"github.com/emersion/go-autostart"
)

type AutostartInstaller struct {
	autostartApp *autostart.App
}

func (ai *AutostartInstaller) Init() {
	ai.autostartApp = &autostart.App{
		Name:        ABOUT_TITLE,
		DisplayName: ABOUT_TITLE,
	}
}

func (ai *AutostartInstaller) IsAutostartEnabled() bool {
	return ai.autostartApp.IsEnabled()
}

func (ai *AutostartInstaller) EnableAutostart(enable bool) error {
	if enable {
		exeAbsPath, _ := filepath.Abs(os.Args[0])

		ai.autostartApp.Exec = []string{
			exeAbsPath,
			RUN_COMMAND}

		return ai.autostartApp.Enable()
	} else {
		return ai.autostartApp.Disable()
	}
}
