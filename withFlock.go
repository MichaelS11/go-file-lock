// +build darwin freebsd linux

package fileLock

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func lockFile(lockHandle *lockHandleStruct) error {
	var err error

	lockHandle.file, err = os.OpenFile(lockHandle.filename, os.O_CREATE|os.O_RDWR, 0600)

	if err != nil {
		return err
	}

	_, err = lockHandle.file.WriteString(strconv.FormatInt(int64(os.Getpid()), 10))

	if err != nil {
		return err
	}

	err = syscall.Flock(int(lockHandle.file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)

	if err != nil && err.Error() == "resource temporarily unavailable" {
		return ErrFileIsBeingUsed
	}

	if err != nil {
		return err
	}

	return nil
}

func unlockFile(lockHandle *lockHandleStruct) error {
	if lockHandle.file == nil {
		return fmt.Errorf("nil file pointer")
	}

	err := syscall.Flock(int(lockHandle.file.Fd()), syscall.LOCK_UN|syscall.LOCK_NB)

	if err != nil {
		return err
	}

	err = lockHandle.file.Close()

	lockHandle.file = nil

	if err != nil {
		return err
	}

	return nil
}
