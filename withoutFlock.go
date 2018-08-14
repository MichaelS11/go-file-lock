// +build !darwin,!freebsd,!linux

package filelock

import (
	"fmt"
	"os"
	"strconv"
)

func (lockHandle *LockHandle) lock() error {
	err := os.Remove(lockHandle.filename)
	if err != nil && len(err.Error()) > 79 &&
		err.Error()[len(err.Error())-79:] == "The process cannot access the file because it is being used by another process." {
		return ErrFileIsBeingUsed
	}
	if err != nil && len(err.Error()) > 42 &&
		err.Error()[len(err.Error())-42:] != "The system cannot find the file specified." {
		return fmt.Errorf("Remove error: %v", err)
	}

	lockHandle.file, err = os.OpenFile(lockHandle.filename, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0600)
	if err != nil && len(err.Error()) > 79 &&
		err.Error()[len(err.Error())-79:] == "The process cannot access the file because it is being used by another process." {
		return ErrFileIsBeingUsed
	}
	if err != nil {
		return fmt.Errorf("OpenFile error: %v", err)
	}

	_, err = lockHandle.file.WriteString(strconv.FormatInt(int64(os.Getpid()), 10))
	if err != nil {
		return fmt.Errorf("WriteString error: %v", err)
	}

	return nil
}

func (lockHandle *LockHandle) unlock() error {
	err := lockHandle.file.Close()

	lockHandle.file = nil

	if err != nil {
		return fmt.Errorf("Close error: %v", err)
	}

	err = os.Remove(lockHandle.filename)
	if err != nil && len(err.Error()) > 79 &&
		err.Error()[len(err.Error())-79:] == "The process cannot access the file because it is being used by another process." {
		return nil
	}
	if err != nil && len(err.Error()) > 42 &&
		err.Error()[len(err.Error())-42:] != "The system cannot find the file specified." {
		return fmt.Errorf("Remove error: %v", err)
	}

	return nil
}
