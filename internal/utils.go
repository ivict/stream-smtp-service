package internal

import (
	"os"
	"strconv"
)

func ParseUint16(number string) uint16 {
	result, err := strconv.ParseUint(number, 0, 16)
	if err != nil {
		return 0
	}
	return uint16(result)
}

func FormatAddr(host string, port uint16) string {
	return host + ":" + strconv.FormatUint(uint64(port), 10)
}

func Last[E any](slice []E) E {
	if len(slice) == 0 {
		var zero E
		return zero
	}
	return slice[len(slice)-1]
}

func GetenvInt(key string, fallback int) int {
	keyValue, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}
	return keyValue
}
