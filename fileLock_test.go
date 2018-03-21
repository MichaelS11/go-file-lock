package fileLock

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	TestBaseDir  string
	TestLockFile string
)

func init() {
	TestBaseDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	TestLockFile = TestBaseDir + "/testLockFile.lock"
}

func TestWithoutNew(t *testing.T) {
	lockHandle := LockHandle{}

	err := lockHandle.Lock()
	if err == nil || err.Error() != "filename empty" {
		t.Fatal("Lock:", err)
	}

	err = lockHandle.Unlock()
	if err != ErrFileAlreadyUnlocked {
		t.Fatal("Unlock:", err)
	}
}

func TestNewFilenameEmpty(t *testing.T) {
	lockHandle, err := New("")

	if lockHandle != nil {
		t.Fatal("lockHandle not empty")
	}
	if err == nil || err.Error() != "filename empty" {
		t.Fatal("New:", err)
	}
}

func TestNewThenUnlock(t *testing.T) {
	lockHandle, err := New(TestLockFile)
	if err != nil {
		t.Fatal("New:", err)
	}

	fileInfo, err := lockHandle.file.Stat()
	if fileInfo.Size() < 1 {
		t.Fatal("Size:", fileInfo.Size())
	}

	err = lockHandle.Unlock()
	if err != nil {
		t.Fatal("Unlock:", err)
	}
}

func TestTwoNew(t *testing.T) {
	chanInt := make(chan int, 1)

	go newWaitUnlock(t, chanInt)

	<-chanInt

	if t.Failed() {
		return
	}

	lockHandle, err := New(TestLockFile)
	if err != ErrFileIsBeingUsed {
		t.Error("New:", err)
		chanInt <- 21
		return
	}
	if lockHandle == nil {
		t.Error("lockHandle nil")
		chanInt <- 22
		return
	}

	chanInt <- 23

	<-chanInt
}

func newWaitUnlock(t *testing.T, chanInt chan int) {
	lockHandle, err := New(TestLockFile)
	if err != nil {
		t.Error("New:", err)
		chanInt <- 11
		return
	}

	chanInt <- 12

	<-chanInt

	if t.Failed() {
		return
	}

	err = lockHandle.Unlock()
	if err != nil {
		t.Error("Unlock:", err)
	}

	chanInt <- 13
}

func TestNewTwiceThenUnlock(t *testing.T) {
	lockHandle, err := New(TestLockFile)
	if err != nil {
		t.Fatal("New:", err)
	}

	fileInfo, err := lockHandle.file.Stat()
	if fileInfo.Size() < 1 {
		t.Fatal("Size:", fileInfo.Size())
	}

	lockHandle2, err := New(TestLockFile)
	if err != ErrFileIsBeingUsed {
		t.Fatal("New 2:", err)
	}
	if lockHandle2 == nil {
		t.Fatal("lockHandle2 nil")
	}

	err = lockHandle.Unlock()
	if err != nil {
		t.Fatal("Unlock:", err)
	}
}

func TestNewThenUnlockTwice(t *testing.T) {
	lockHandle, err := New(TestLockFile)
	if err != nil {
		t.Fatal("New:", err)
	}

	fileInfo, err := lockHandle.file.Stat()
	if fileInfo.Size() < 1 {
		t.Fatal("Size:", fileInfo.Size())
	}

	err = lockHandle.Unlock()
	if err != nil {
		t.Fatal("Unlock:", err)
	}

	err = lockHandle.Unlock()
	if err != ErrFileAlreadyUnlocked {
		t.Fatal("Unlock:", err)
	}
}

func TestNewUnlockLockUnlock(t *testing.T) {
	lockHandle, err := New(TestLockFile)
	if err != nil {
		t.Fatal("New:", err)
	}

	fileInfo, err := lockHandle.file.Stat()
	if fileInfo.Size() < 1 {
		t.Fatal("Size:", fileInfo.Size())
	}

	err = lockHandle.Unlock()
	if err != nil {
		t.Fatal("Unlock:", err)
	}

	err = lockHandle.Lock()
	if err != nil {
		t.Fatal("Lock:", err)
	}

	err = lockHandle.Unlock()
	if err != nil {
		t.Fatal("Unlock:", err)
	}
}

func TestNewInvalidFile(t *testing.T) {
	null := 0
	lockHandle, err := New(TestLockFile + string(null))
	if err == nil {
		t.Fatal("New no error")
	}

	err = lockHandle.Unlock()
	if err != ErrFileAlreadyUnlocked {
		t.Fatal("Unlock:", err)
	}
}
