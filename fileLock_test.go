package fileLock

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewThenUnlock(t *testing.T) {
	baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	lockHandle, err := New(baseDir + "/lockFileLock.lock")

	if err != nil {
		t.Error("New:", err)
		return
	}

	fileInfo, err := lockHandle.file.Stat()

	if fileInfo.Size() < 1 {
		t.Error("Size:", fileInfo.Size())
		return
	}

	err = lockHandle.Unlock()

	if err != nil {
		t.Error("Unlock:", err)
		return
	}
}

func TestTwoNew(t *testing.T) {

	baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	chanInt := make(chan int, 1)

	go newWaitUnlock(t, chanInt)

	<-chanInt

	if t.Failed() {
		return
	}

	lockHandle, err := New(baseDir + "/lockFileLock.lock")

	if err != nil && lockHandle != nil {
		chanInt <- 11
		t.Error("New - lockHandle not nil:", err)
		return
	}

	if err == nil {
		chanInt <- 12
		t.Error("expected error")
		return
	}

	if err != nil && err.Error() != "file is being used by another process" {
		chanInt <- 12
		t.Error("New:", err)
		return
	}

	chanInt <- 13

	<-chanInt
}

func newWaitUnlock(t *testing.T, chanInt chan int) {
	baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	lockHandle, err := New(baseDir + "/lockFileLock.lock")

	if err != nil {
		t.Error("New:", err)
		chanInt <- 1
		return
	}

	chanInt <- 2

	<-chanInt

	if t.Failed() {
		return
	}

	err = lockHandle.Unlock()

	if err != nil {
		t.Error("Unlock:", err)
	}

	chanInt <- 3
}

func TestNewTwiceThenUnlock(t *testing.T) {
	baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	lockHandle, err := New(baseDir + "/lockFileLock.lock")

	if err != nil {
		t.Error("New:", err)
		return
	}

	fileInfo, err := lockHandle.file.Stat()

	if fileInfo.Size() < 1 {
		t.Error("Size:", fileInfo.Size())
		return
	}

	lockHandle2, err := New(baseDir + "/lockFileLock.lock")

	if err != nil && lockHandle2 != nil {
		t.Error("New - lockHandle2 not nil:", err)
		return
	}

	if err != nil && err.Error() != "file is being used by another process" {
		t.Error("New:", err)
		return
	}

	if err == nil {
		t.Error("expected error")
		return
	}

	err = lockHandle.Unlock()

	if err != nil {
		t.Error("Unlock:", err)
		return
	}
}

func TestNewThenUnlockTwice(t *testing.T) {
	baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	lockHandle, err := New(baseDir + "/lockFileLock.lock")

	if err != nil {
		t.Error("New:", err)
		return
	}

	fileInfo, err := lockHandle.file.Stat()

	if fileInfo.Size() < 1 {
		t.Error("Size:", fileInfo.Size())
		return
	}

	err = lockHandle.Unlock()

	if err != nil {
		t.Error("Unlock:", err)
		return
	}

	err = lockHandle.Unlock()

	if err != nil && err.Error() != "nil file pointer" {
		t.Error("Unlock:", err)
		return
	}

	if err == nil {
		t.Error("expected error")
		return
	}
}
