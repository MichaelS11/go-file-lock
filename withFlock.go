// +build darwin freebsd linux

package filelock

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func (lockHandle *LockHandle) lock() error {
	var err error

	lockHandle.file, err = os.OpenFile(lockHandle.filename, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return fmt.Errorf("OpenFile error: %v", err)
	}

	_, err = lockHandle.file.WriteString(strconv.FormatInt(int64(os.Getpid()), 10))
	if err != nil {
		return fmt.Errorf("WriteString error: %v", err)
	}

	err = syscall.Flock(int(lockHandle.file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil && err.Error() == "resource temporarily unavailable" {
		return ErrFileIsBeingUsed
	}
	if err != nil {
		return fmt.Errorf("Flock error: %v", err)
	}

	return nil
}

func (lockHandle *LockHandle) unlock() error {
	err := syscall.Flock(int(lockHandle.file.Fd()), syscall.LOCK_UN|syscall.LOCK_NB)
	if err != nil {
		return fmt.Errorf("Flock error: %v", err)
	}

	err = lockHandle.file.Close()

	lockHandle.file = nil

	if err != nil {
		return fmt.Errorf("Close error: %v", err)
	}

	return nil
}
