# Go file lock

[![GoDoc Reference](https://godoc.org/github.com/MichaelS11/go-file-lock?status.svg)](http://godoc.org/github.com/MichaelS11/go-file-lock)
[![Build Status](https://travis-ci.org/MichaelS11/go-file-lock.svg)](https://travis-ci.org/MichaelS11/go-file-lock)
[![Coverage](https://gocover.io/_badge/github.com/MichaelS11/go-file-lock)](https://gocover.io/github.com/MichaelS11/go-file-lock#)
[![Go Report Card](https://goreportcard.com/badge/github.com/MichaelS11/go-file-lock)](https://goreportcard.com/report/github.com/MichaelS11/go-file-lock)

Golang file lock to run only one instance of an application at a time

## Get

go get github.com/MichaelS11/go-file-lock

## Usage

```Go
import (
	"github.com/MichaelS11/go-file-lock"
)

func Run() error {

  lockHandle, err := fileLock.New("myLockFile.lock")

  if err != nil && err == fileLock.ErrFileIsBeingUsed {
    return nil
  }

  if err != nil {
    return err
  }
  
  # do main program

  return lockHandle.Unlock()
  
}
```

If you want the current directory of your application, you can do:

```Go
baseDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
```
