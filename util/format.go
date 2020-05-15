package util

import "fmt"

func FormatBusinessCode(id uint64) (code string) {
	code = fmt.Sprintf("%04d", id)
	return
}
