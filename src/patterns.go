package src

import (
	"bytes"
	"encoding/hex"
	"regexp"
)

var separatePattern = regexp.MustCompile("\\s+")

func stringToByteArray(pattern string) ([]byte, []bool, error) {
	bytes := separatePattern.Split(pattern, -1)
	result := make([]byte, len(bytes))
	wildcards := make([]bool, len(bytes))

	for i, b := range bytes {
		if b == "?" || b == "??" {
			result[i] = 0x0
			wildcards[i] = true
		} else {
			parsed, err := hex.DecodeString(b)
			if err != nil {
				return nil, nil, err
			}
			result[i] = byte(parsed[0])
		}
	}

	return result, wildcards, nil
}

// hexToUintptr converts b into a uintptr.
// It's optimized to assume the input will not be invalid.
// (I.e., that /proc/$$/maps won't produce a garbage value.)
func hexToUintptr(b []byte) (n uintptr) {
	for _, d := range b {
		n *= 16
		switch {
		case '0' <= d && d <= '9':
			n += uintptr(d - '0')
		case 'a' <= d && d <= 'z':
			n += uintptr(d - 'a' + 10)
		case 'A' <= d && d <= 'Z':
			n += uintptr(d - 'A' + 10)
		default:
			return 0
		}
	}
	return n
}

// parseUint parses b into a uint64. See hexToUintptr for more.
func parseUint(b []byte) (n uint64) {
	for _, d := range b {
		n *= 10
		switch {
		case '0' <= d && d <= '9':
			n += uint64(d - '0')
		case 'a' <= d && d <= 'z':
			n += uint64(d - 'a' + 10)
		case 'A' <= d && d <= 'Z':
			n += uint64(d - 'A' + 10)
		default:
			return 0
		}
	}
	return n
}

// splitOn splits b in half on the first occurance of c.
func splitOn(b []byte, c byte) (p1, p2 []byte) {
	i := bytes.IndexByte(b, c)
	if i < 0 {
		return nil, nil
	}
	return b[:i], b[i+1:]
}
