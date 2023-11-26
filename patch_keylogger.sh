KEYLOGGER_GO_PATH=`find /home/$USER/go/pkg/mod/github.com -name keylogger.go`
KEYLOGGER_GO_PATCH_PATH=`realpath keylogger.go.patch`

if [ -n "$KEYLOGGER_GO_PATH" ]; then
    pushd .

    cd `dirname "$KEYLOGGER_GO_PATH"`

    chmod a+w .
    chmod a+w "$KEYLOGGER_GO_PATH"

    patch -p1 < "$KEYLOGGER_GO_PATCH_PATH"

    chmod a-w .
    chmod a-w "$KEYLOGGER_GO_PATH"

    popd
fi
