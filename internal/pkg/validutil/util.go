package validutil

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"web-api-scaffold/internal/pkg/constant"
)

func ValidateTokenSignature(token, sign string) bool {
	var (
		chars = []byte(token)
		hash  = md5.New()
		err   error
	)

	sort.Slice(chars, func(i int, j int) bool {
		return chars[i] < chars[j]
	})

	_, err = hash.Write(chars)
	_, err = hash.Write([]byte(constant.DeviceSecret))

	//logger.Instance().Print(hex.EncodeToString(hash.Sum(nil)))

	return err == nil && hex.EncodeToString(hash.Sum(nil)) == sign
}
