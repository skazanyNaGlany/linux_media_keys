package main_window

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainWindow struct {
	RunTestButtonCallback         func()
	KillButtonCallback            func()
	ToggleAutostartButtonCallback func()
	ExitButtonCallback            func()
	window                        *tview.Flex
	menu                          *tview.List
	aboutModal                    *tview.Modal
	windowContents                *tview.Flex
	windowPages                   *tview.Pages
	runTestButtonEnabled          bool
	killButtonEnabled             bool
	runButtonIndex                int
	killButtonIndex               int
	toggleAutostartButtonIndex    int
}

func (mw *MainWindow) Init() {
	mw.buildWindow()
	mw.buildWindowContents()
	mw.buildAboutModal()
	mw.buildWindowPages()

	mw.window.AddItem(mw.windowPages, 0, 1, true)

	mw.windowPages.SwitchToPage(PAGE_MAIN)
}

func (mw *MainWindow) buildWindow() {
	mw.window = tview.NewFlex()
	mw.window.Box = tview.NewBox()
	mw.window.SetBackgroundColor(tcell.ColorBlack)
	mw.window.SetDirection(tview.FlexRow)
}

func (mw *MainWindow) buildWindowPages() {
	mw.windowPages = tview.NewPages()

	mw.windowPages.AddPage("main", mw.windowContents, true, false)
	mw.windowPages.AddPage("about", mw.aboutModal, true, false)
}

func (mw *MainWindow) buildAboutModal() {
	mw.aboutModal = tview.NewModal()
	mw.aboutModal.SetText(fmt.Sprintf(ABOUT_MESSAGE, APP_VERSION))
	mw.aboutModal.AddButtons([]string{"OK"})
	mw.aboutModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		mw.windowPages.SwitchToPage(PAGE_MAIN)
	})
}

func (mw *MainWindow) buildWindowContents() {
	mw.menu = tview.NewList()
	mw.menu.SetBorderPadding(10, 10, 10, 0)
	mw.menu.AddItem(RUN_TEST_BUTTON_ENABLE, "", 'r', mw.runTestButtonCallback)
	mw.menu.AddItem(KILL_BUTTON_ENABLE, "", 'k', mw.killButtonCallback)
	mw.menu.AddItem(TOGGLE_AUTOSTART_BUTTON_ENABLE, "", 'e', mw.toggleAutostartButtonCallback)
	mw.menu.AddItem(ABOUT_BUTTON, "", 'a', mw.aboutButtonCallback)
	mw.menu.AddItem(EXIT_BUTTON, "", 'q', mw.exitButtonCallback)
	mw.menu.SetBorder(true)
	mw.menu.SetTitle(fmt.Sprintf(TITLE, APP_VERSION))

	mw.runButtonIndex = 0
	mw.killButtonIndex = 1
	mw.toggleAutostartButtonIndex = 2

	mw.windowContents = tview.NewFlex()
	mw.windowContents.AddItem(nil, 0, 1, false)
	mw.windowContents.AddItem(mw.menu, 0, 2, true)
	mw.windowContents.AddItem(nil, 0, 1, false)

}

func (mw *MainWindow) runTestButtonCallback() {
	if !mw.runTestButtonEnabled {
		return
	}

	if mw.RunTestButtonCallback != nil {
		mw.RunTestButtonCallback()
	}
}

func (mw *MainWindow) killButtonCallback() {
	if !mw.killButtonEnabled {
		return
	}

	if mw.KillButtonCallback != nil {
		mw.KillButtonCallback()
	}
}

func (mw *MainWindow) toggleAutostartButtonCallback() {
	if mw.ToggleAutostartButtonCallback != nil {
		mw.ToggleAutostartButtonCallback()
	}
}

func (mw *MainWindow) aboutButtonCallback() {
	mw.windowPages.SwitchToPage(PAGE_ABOUT)
}

func (mw *MainWindow) exitButtonCallback() {
	if mw.ExitButtonCallback != nil {
		mw.ExitButtonCallback()
	}
}

func (mw *MainWindow) GetWindow() tview.Primitive {
	return mw.window
}

func (mw *MainWindow) SetToggleAutostartButtonState(enable bool) {
	if enable {
		mw.menu.SetItemText(mw.toggleAutostartButtonIndex, TOGGLE_AUTOSTART_BUTTON_ENABLE, "")
	} else {
		mw.menu.SetItemText(mw.toggleAutostartButtonIndex, TOGGLE_AUTOSTART_BUTTON_DISABLE, "")
	}
}

func (mw *MainWindow) RunTestButtonEnable(enable bool) {
	if enable {
		mw.menu.SetItemText(mw.runButtonIndex, RUN_TEST_BUTTON_ENABLE, "")
	} else {
		mw.menu.SetItemText(mw.runButtonIndex, RUN_TEST_BUTTON_DISABLE, "")
	}

	mw.runTestButtonEnabled = enable
}

func (mw *MainWindow) KillButtonEnable(enable bool) {
	if enable {
		mw.menu.SetItemText(mw.killButtonIndex, KILL_BUTTON_ENABLE, "")
	} else {
		mw.menu.SetItemText(mw.killButtonIndex, KILL_BUTTON_DISABLE, "")
	}

	mw.killButtonEnabled = enable
}
