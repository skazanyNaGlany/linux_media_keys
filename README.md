# Linux Media Keys
Media keys on Linux, without actual media keys.


### Keymap for version 0.2
Esc-Up - VOLUME UP

Esc-Down - VOLUME DOWN

Esc-Left - PREVIOUS MEDIA

Esc-Right - NEXT MEDIA

Esc-F4 - SUSPEND COMPUTER

Esc-F7 - RESTART COMPUTER

Esc-F8 - SHUTDOWN COMPUTER

Esc-F11 - TOGGLE MUTE

Esc-F12 - TOGGLE PLAY

### Compiling
```
$ go mod tidy
$ ./patch_keylogger.sh
$ ./build.sh
```
