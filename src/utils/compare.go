/*
 * compare.go
 * Copyright (c) ti-bone 2022.
 */

package utils

import "strings"

// Compare - Check if string have words from string array, returns true if it has.
func Compare(str string, arr []string) bool {
	for _, filter := range arr {
		if strings.Contains(strings.ToLower(str), strings.ToLower(filter)) {
			return true
		}
	}
	return false
}
