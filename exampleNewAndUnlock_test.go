package fileLock_test

import (
	"fmt"

	"github.com/MichaelS11/go-file-lock"
)

func Example_fileLockNewAndUnlock() {
	lockHandle, err := fileLock.New("myLockFile.lock")

	if err != nil && err == fileLock.ErrFileIsBeingUsed {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	// do main program
	fmt.Println("running")

	err = lockHandle.Unlock()

	if err != nil {
		panic(err)
	}

	// output: running
}
