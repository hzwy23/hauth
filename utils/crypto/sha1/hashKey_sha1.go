package sha1

import (
	"crypto/sha1"
	"fmt"
)

func GenSha1Key(gpname string, keys ...string) string {
	sh := sha1.New()
	sh.Write([]byte(gpname))
	sh.Write([]byte("_join_"))
	for _, val := range keys {
		sh.Write([]byte(val))
		sh.Write([]byte("_join_"))
	}
	return fmt.Sprintf("%x", sh.Sum(nil))
}
