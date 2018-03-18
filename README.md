# Go file lock

Golang file lock to run only one instance of an application at a time

## Get

go get github.com/MichaelS11/go-file-lock

## Usage

```Go
func Run() error {

  lockHandle, err := fileLock.New("myLockFile.lock")

  if err != nil && err == fileLock.FileIsBeingUsed {
    return nil
  }

  if err != nil {
    return err
  }
  
  # do main program

  return lockHandle.Unlock()
  
}
```