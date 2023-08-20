package drop_root_rights

import (
	"errors"
	"os"
	"syscall"
)

type DropRootRights struct{}

func (drr *DropRootRights) Drop() error {
	euid := os.Geteuid()

	if euid != 0 {
		// not root
		return nil
	}

	uid := os.Getuid()

	if uid == 0 {
		// root without suid bit, cannot drop privilege
		return errors.New("program was run from root shell, not normal user (without SUID bit)")
	}

	return syscall.Seteuid(uid)
}

func (drr *DropRootRights) IsEffectiveRoot() bool {
	return os.Geteuid() == 0
}
