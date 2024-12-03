package service

import (
	"strings"
)

func ParseFFprobeOutput(output []byte) map[string]interface{} {
	info := make(map[string]interface{})

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := parts[0]
			value := parts[1]
			info[key] = value
		}
	}

	return info
}
