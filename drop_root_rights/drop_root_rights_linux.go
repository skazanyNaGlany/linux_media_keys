package drop_root_rights

import (
	"errors"
	"os"
	"syscall"
)

const LINUX_ROOT_UID = 0

type DropRootRights struct{}

func (drr *DropRootRights) Drop() error {
	if !drr.IsEffectiveRoot() {
		// not root
		return nil
	}

	uid := os.Getuid()

	if uid == LINUX_ROOT_UID {
		// root without suid bit, cannot drop privilege
		return errors.New("program was run from root shell, not normal user (without SUID bit)")
	}

	return syscall.Seteuid(uid)
}

func (drr *DropRootRights) IsEffectiveRoot() bool {
	return os.Geteuid() == LINUX_ROOT_UID
}
