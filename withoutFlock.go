// +build !darwin,!freebsd,!linux

package fileLock

import (
	"fmt"
	"os"
	"strconv"
)

func lockFile(lockHandle *lockHandleStruct) error {
	err := os.Remove(lockHandle.filename)

	if err != nil && len(err.Error()) > 79 &&
		err.Error()[len(err.Error())-79:] == "The process cannot access the file because it is being used by another process." {
		return ErrFileIsBeingUsed
	}

	if err != nil && len(err.Error()) > 42 &&
		err.Error()[len(err.Error())-42:] != "The system cannot find the file specified." {
		return err
	}

	lockHandle.file, err = os.OpenFile(lockHandle.filename, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0600)

	if err != nil && len(err.Error()) > 79 &&
		err.Error()[len(err.Error())-79:] == "The process cannot access the file because it is being used by another process." {
		return ErrFileIsBeingUsed
	}

	if err != nil {
		return err
	}

	_, err = lockHandle.file.WriteString(strconv.FormatInt(int64(os.Getpid()), 10))

	if err != nil {
		return err
	}

	return nil
}

func unlockFile(lockHandle *lockHandleStruct) error {
	if lockHandle.file == nil {
		return fmt.Errorf("nil file pointer")
	}

	err := lockHandle.file.Close()

	lockHandle.file = nil

	if err != nil {
		return err
	}

	err = os.Remove(lockHandle.filename)

	if err != nil && len(err.Error()) > 79 &&
		err.Error()[len(err.Error())-79:] == "The process cannot access the file because it is being used by another process." {
		return nil
	}

	if err != nil && len(err.Error()) > 42 &&
		err.Error()[len(err.Error())-42:] != "The system cannot find the file specified." {
		return err
	}

	return nil
}
