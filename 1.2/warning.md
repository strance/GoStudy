- 正确关闭文件。

```go
f, err := os.Open("/tmp/file.md")
if err != nil {
	return err
}

defer func() {
	closeErr := f.Close()
	if closeErr != nil {
		if err == nil {
			err = closeErr
		} else {
			log.Println("Error occured while closing the file :", closeErr)
		}
	}
}()
return err

```
