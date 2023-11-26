package main

import (
	"errors"
	"os/exec"
	"time"

	"github.com/MarinX/keylogger"
	"github.com/micmonay/keybd_event"
)

type MediaKeysEmulator struct {
	keyLogger *keylogger.KeyLogger
	kbEvent   keybd_event.KeyBonding
}

func NewMediaKeysEmulator() (*MediaKeysEmulator, error) {
	var err error
	var mke MediaKeysEmulator

	keyboard := keylogger.FindKeyboardDevice()

	if len(keyboard) <= 0 {
		return nil, errors.New("no keyboard found")
	}

	mke.keyLogger, err = keylogger.New(keyboard)

	if err != nil {
		return nil, err
	}

	mke.kbEvent, err = keybd_event.NewKeyBonding()

	if err != nil {
		mke.keyLogger.Close()

		return nil, err
	}

	// keybd_event - for linux, it is very important to wait 2 seconds
	time.Sleep(2 * time.Second)

	return &mke, nil
}

func (mke *MediaKeysEmulator) Close() error {
	return mke.keyLogger.Close()
}

func (mke *MediaKeysEmulator) EmulateInLoop() {
	events := mke.keyLogger.Read()
	pressed := make(map[string]bool)

	for iEvent := range events {
		switch iEvent.Type {
		case keylogger.EvKey:
			keyString := iEvent.KeyString()

			if iEvent.KeyPress() {
				pressed[keyString] = true
			}

			if iEvent.KeyRelease() {
				delete(pressed, keyString)
			}

			mke.executeKeyboardAction(pressed)
		}
	}
}

func (mke *MediaKeysEmulator) executeKeyboardAction(pressed map[string]bool) {
	_, escapePressed := pressed[VK_ESCAPE]

	if !escapePressed {
		return
	}

	_, downPressed := pressed[VK_DOWN]
	_, upPressed := pressed[VK_UP]
	_, rightPressed := pressed[VK_RIGHT]
	_, leftPressed := pressed[VK_LEFT]
	_, f4Pressed := pressed[VK_F4]
	_, f7Pressed := pressed[VK_F7]
	_, f8Pressed := pressed[VK_F8]
	_, f11Pressed := pressed[VK_F11]
	_, f12Pressed := pressed[VK_F12]

	if downPressed {
		// volume down
		mke.kbEvent.SetKeys(keybd_event.VK_VOLUMEDOWN)
		mke.kbEvent.Launching()
	} else if upPressed {
		// volume up
		mke.kbEvent.SetKeys(keybd_event.VK_VOLUMEUP)
		mke.kbEvent.Launching()
	} else if rightPressed {
		// media next
		mke.kbEvent.SetKeys(keybd_event.VK_NEXTSONG)
		mke.kbEvent.Launching()
	} else if leftPressed {
		// media prev
		mke.kbEvent.SetKeys(keybd_event.VK_PREVIOUSSONG)
		mke.kbEvent.Launching()
	} else if f4Pressed {
		// sleep
		exec.Command("pm-suspend").Start()
	} else if f7Pressed {
		// restart
		exec.Command("reboot").Start()
	} else if f8Pressed {
		// shutdown
		exec.Command("shutdown", "-h", "now").Start()
	} else if f11Pressed {
		// mute/unmute
		mke.kbEvent.SetKeys(keybd_event.VK_MUTE)
		mke.kbEvent.Launching()
	} else if f12Pressed {
		// play/pause
		mke.kbEvent.SetKeys(keybd_event.VK_PLAYPAUSE)
		mke.kbEvent.Launching()
	}
}
