// utils/cursor.go
package utils

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func EncodeCursor(cursorMap map[string]interface{}) string {
	parts := make([]string, 0, len(cursorMap))
	for key, value := range cursorMap {
		parts = append(parts, fmt.Sprintf("%s=%v", key, value))
	}
	return base64.StdEncoding.EncodeToString([]byte(strings.Join(parts, "|")))
}

func DecodeCursor(cursor string) (map[string]string, error) {
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(decoded), "|")
	cursorMap := make(map[string]string, len(parts))
	for _, part := range parts {
		keyValue := strings.Split(part, "=")
		if len(keyValue) == 2 {
			cursorMap[keyValue[0]] = keyValue[1]
		}
	}
	return cursorMap, nil
}
