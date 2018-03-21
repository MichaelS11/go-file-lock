package fileLock_test

import (
	"fmt"

	"github.com/MichaelS11/go-file-lock"
)

func Example_fileLockNewAndUnlock() {
	lockHandle, err := fileLock.New("myLockFile.lock")

	if err != nil && err == fileLock.ErrFileIsBeingUsed {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	// do main program
	fmt.Println("running")

	err = lockHandle.Unlock()

	if err != nil {
		fmt.Println(err)
	}

	// output: running
}
