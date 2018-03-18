package fileLock

import (
	"errors"
	"os"
)

// LockHandle to be used for locking and unlocking the file
type LockHandle struct {
	filename string
	file     *os.File
}

var (
	// ErrFileIsBeingUsed is returned when file is in use by another program
	ErrFileIsBeingUsed = errors.New("file is being used by another process")
	// ErrFileAlreadyUnlocked is returned when file has already been unlocked or was not locked
	ErrFileAlreadyUnlocked = errors.New("file already unlocked")
)

// New creates a new file lock
func New(filename string) (*LockHandle, error) {
	if len(filename) < 1 {
		return nil, errors.New("filename empty")
	}

	lockHandle := &LockHandle{filename: filename}
	return lockHandle, lockHandle.lock()
}

// Lock locks the file lock
func (lockHandle *LockHandle) Lock() error {
	if len(lockHandle.filename) < 1 {
		return errors.New("filename empty")
	}

	return lockHandle.lock()
}

// Unlock unlocks the file lock
func (lockHandle *LockHandle) Unlock() error {
	if lockHandle.file == nil {
		return ErrFileAlreadyUnlocked
	}

	return lockHandle.unlock()
}
