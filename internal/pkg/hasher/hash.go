package hasher

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"io"
)

func GenerateUniqueUUID() string {
	return uuid.NewV4().String()
}

func GenerateUniqueHexUUID() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}

func CalculateMD5Hash(s string) string {
	hash := md5.New()
	if _, err := hash.Write([]byte(s)); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func CalculateSHA1Hash(s string) string {
	hash := sha1.New()
	if _, err := hash.Write([]byte(s)); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func CalculateSmallHashByReader(r io.ReadSeeker) string {
	hash := md5.New()
	if _, err := io.CopyN(hash, r, 1<<16); err != nil {
		return ""
	}
	if _, err := r.Seek(0, 0); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))[8:24]
}

func CalculateSHA256HashByReader(r io.ReadSeeker) string {
	hash := sha256.New()
	if _, err := io.Copy(hash, r); err != nil {
		return ""
	}
	if _, err := r.Seek(0, 0); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}
