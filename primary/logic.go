package primary

import (
	"math/rand"
	"os"
	"time"
)

// GenerateBytesMask - Generate random string
func GenerateBytesMask(n int, withString bool) string {
	var src = rand.NewSource(time.Now().UnixNano())

	var letterBytes string

	if withString == false {
		letterBytes = "123456789"
	} else {
		letterBytes = "1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	}

	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// DeleteAFile -- Delete a file
func DeleteAFile(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		return err
	}

	return nil
}
