package internal

import "strconv"

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
