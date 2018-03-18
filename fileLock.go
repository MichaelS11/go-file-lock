package fileLock

import (
	"errors"
	"os"
)

type lockHandleStruct struct {
	filename string
	file     *os.File
}

var (
	// error returned when file is in use by another program
	ErrFileIsBeingUsed = errors.New("file is being used by another process")
)

// New creates a new file lock
func New(filename string) (*lockHandleStruct, error) {
	lockHandle := &lockHandleStruct{filename: filename}

	err := lockFile(lockHandle)

	if err != nil {
		return nil, err
	}

	return lockHandle, nil
}

// Unlock unlocks the file lock
func (lockHandle *lockHandleStruct) Unlock() error {
	return unlockFile(lockHandle)
}
