package fileLock

import (
	"errors"
	"os"
)

type lockHandleStruct struct {
	filename string
	file     *os.File
}

var ErrFileIsBeingUsed = errors.New("file is being used by another process") // error returned when file is in use by another program

// Create a new file lock
func New(filename string) (*lockHandleStruct, error) {
	lockHandle := &lockHandleStruct{filename: filename}

	err := lockFile(lockHandle)

	if err != nil {
		return nil, err
	}

	return lockHandle, nil
}

// Unlocks file lock
func (lockHandle *lockHandleStruct) Unlock() error {
	return unlockFile(lockHandle)
}
