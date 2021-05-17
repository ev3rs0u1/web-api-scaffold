package fileutil

import (
	"fmt"
	"os"
	"strings"
	"web-api-scaffold/internal/pkg/devio"
)

func Ext(fn string) string {
	for i := len(fn) - 1; i >= 0 && !os.IsPathSeparator(fn[i]); i-- {
		if fn[i] == '.' {
			return strings.ToLower(fn[i+1:])
		}
	}
	return ""
}

func WithoutExt(fn string) string {
	return strings.TrimSuffix(fn, "."+Ext(fn))
}

func NewCopyName(fn string) string {
	suffix := Ext(fn)
	if suffix == "" {
		return fmt.Sprintf("%s_(%d)", fn, devio.GetUnixTimestamp())
	}

	prefix := WithoutExt(fn)
	return fmt.Sprintf("%s_(%d).%s", prefix, devio.GetUnixTimestamp(), suffix)
}
