package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// GenerateSign generates SHA-512 signature based on Rivo's rules
func GenerateSign(params map[string]interface{}, key string) string {
	if params == nil {
		return ""
	}

	// 1. Collect and filter parameters
	var keys []string
	for k, v := range params {
		if k == "sign" || v == nil {
			continue
		}
		val := strings.TrimSpace(fmt.Sprint(v))
		if val == "" {
			continue
		}
		keys = append(keys, k)
	}

	// 2. Sort keys in ASCII ascending order
	sort.Strings(keys)

	// 3. Concatenate key=value
	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, strings.TrimSpace(fmt.Sprint(params[k]))))
	}
	signStr := strings.Join(parts, "&")

	// 4. Append merchant secret key
	signStr = fmt.Sprintf("%s&key=%s", signStr, key)

	// 5. SHA-512 hash
	hash := sha512.New()
	hash.Write([]byte(signStr))
	return hex.EncodeToString(hash.Sum(nil))
}

// VerifySign verifies the signature of the data object
func VerifySign(params map[string]interface{}, key string, signToVerify string) bool {
	generatedSign := GenerateSign(params, key)
	return strings.EqualFold(generatedSign, strings.TrimSpace(signToVerify))
}
