package rivo

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func signPayload(payload interface{}, secretKey string) (string, error) {
	data, err := structToSignData(payload)
	if err != nil {
		return "", err
	}
	return signWithData(data, secretKey), nil
}

func verifyPayloadSign(payload interface{}, secretKey, sign string) (bool, error) {
	expected, err := signPayload(payload, secretKey)
	if err != nil {
		return false, err
	}
	return strings.EqualFold(expected, strings.TrimSpace(sign)), nil
}

func signWithData(data map[string]string, secretKey string) string {
	keys := make([]string, 0, len(data))
	for k, v := range data {
		if strings.EqualFold(k, "sign") || strings.TrimSpace(v) == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", k, data[k]))
	}
	signStr := strings.Join(parts, "&")
	signStr = fmt.Sprintf("%s&key=%s", signStr, secretKey)

	hash := sha512.New()
	hash.Write([]byte(signStr))
	return hex.EncodeToString(hash.Sum(nil))
}

func structToSignData(payload interface{}) (map[string]string, error) {
	val := reflect.ValueOf(payload)
	for val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return map[string]string{}, nil
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("payload must be struct or pointer to struct")
	}

	typ := val.Type()
	out := make(map[string]string, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		fv := val.Field(i)
		if !sf.IsExported() {
			continue
		}

		tag := sf.Tag.Get("json")
		name, omitempty, skip := parseJSONTag(tag)
		if skip {
			continue
		}
		if name == "" {
			name = sf.Name
		}

		if omitempty && isZeroValue(fv) {
			continue
		}
		s, ok := valueToString(fv)
		if !ok {
			continue
		}
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		out[name] = s
	}

	return out, nil
}

func parseJSONTag(tag string) (name string, omitempty bool, skip bool) {
	if tag == "-" {
		return "", false, true
	}
	if tag == "" {
		return "", false, false
	}
	parts := strings.Split(tag, ",")
	name = parts[0]
	for i := 1; i < len(parts); i++ {
		if parts[i] == "omitempty" {
			omitempty = true
		}
	}
	return name, omitempty, false
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		return v.IsNil()
	}
	return v.IsZero()
}

func valueToString(v reflect.Value) (string, bool) {
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return "", false
		}
		return valueToString(v.Elem())
	}
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return "", false
		}
		return valueToString(v.Elem())
	}

	switch v.Kind() {
	case reflect.String:
		return v.String(), true
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), true
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32), true
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), true
	default:
		return fmt.Sprint(v.Interface()), true
	}
}
