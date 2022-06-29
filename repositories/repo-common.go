package repositories

import "fmt"

var (
	ErrNotFound     = fmt.Errorf("record Not Found")
	ErrInsertFailed = fmt.Errorf("db insert error")
	ErrUpdateFailed = fmt.Errorf("db update error")
	ErrDeleteFailed = fmt.Errorf("db delete error")
)
