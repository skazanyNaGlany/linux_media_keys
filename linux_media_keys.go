package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/rivo/tview"
	"github.com/skazanyNaGlany/linux_media_keys/windows/main_window"
)

func main() {
	checkPlatform()
	changeCwd()

	if isRunCommand() {
		run()
	} else {
		userInterface()
	}
}

func checkPlatform() {
	if runtime.GOOS != "linux" {
		log.Panicln("This program can be run only on Linux.")
	}
}

func changeCwd() {
	os.Chdir(
		filepath.Dir(os.Args[0]))
}

func isRunCommand() bool {
	return len(os.Args) == 2 && os.Args[1] == RUN_COMMAND
}

func run() {
	emulator := MediaKeysEmulator{}
	emulator.Init()

	emulator.EmulateInLoop()
}

func userInterface() {
	app := tview.NewApplication()

	installer := AutostartInstaller{}
	installer.Init()

	instance := Instance{}
	instance.Init()

	mainWindow := main_window.MainWindow{}
	mainWindow.Init()

	if err := instance.Refresh(); err != nil {
		log.Println(err)
	}

	running := instance.IsRunnnig()

	mainWindow.RunTestButtonEnable(!running)
	mainWindow.KillButtonEnable(running)

	mainWindow.SetToggleAutostartButtonState(
		!installer.IsAutostartEnabled())

	// if installer.IsAutostartEnabled() {
	// }

	mainWindow.RunTestButtonCallback = func() {
		if err := instance.Refresh(); err != nil {
			log.Println(err)
		}

		if instance.IsRunnnig() {
			log.Println("Already running.")
		}

		exec.Command(
			instance.GetExePathname(),
			RUN_COMMAND).Start()

		mainWindow.RunTestButtonEnable(false)
		mainWindow.KillButtonEnable(true)
	}

	mainWindow.KillButtonCallback = func() {
		if err := instance.Refresh(); err != nil {
			log.Println(err)
		}

		instance.IsRunnnig()
		instance.Kill()

		mainWindow.RunTestButtonEnable(true)
		mainWindow.KillButtonEnable(false)
	}

	mainWindow.ExitButtonCallback = func() {
		app.Stop()
	}

	mainWindow.ToggleAutostartButtonCallback = func() {
		if err := installer.EnableAutostart(
			!installer.IsAutostartEnabled()); err != nil {
			log.Println(err)
		}

		mainWindow.SetToggleAutostartButtonState(!installer.IsAutostartEnabled())

		// if installer.IsAutostartEnabled() {
		// } else {
		// }
	}

	if err := app.SetRoot(mainWindow.GetWindow(), true).EnableMouse(true).Run(); err != nil {
		log.Panicln(err)
	}
}
