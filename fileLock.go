package fileLock

import (
	"errors"
	"os"
)

type lockHandleStruct struct {
	filename string
	file     *os.File
}

var FileIsBeingUsed = errors.New("file is being used by another process")

func New(filename string) (*lockHandleStruct, error) {
	lockHandle := &lockHandleStruct{filename: filename}

	err := lockFile(lockHandle)

	if err != nil {
		return nil, err
	}

	return lockHandle, nil
}

func (lockHandle *lockHandleStruct) Unlock() error {
	return unlockFile(lockHandle)
}
