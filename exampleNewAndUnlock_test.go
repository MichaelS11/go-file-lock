package filelock_test

import (
	"fmt"

	"github.com/MichaelS11/go-file-lock"
)

func Example_fileLockNewAndUnlock() {
	lockHandle, err := filelock.New("myLockFile.lock")

	if err != nil && err == filelock.ErrFileIsBeingUsed {
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
