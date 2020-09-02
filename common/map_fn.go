package common

import (
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type ParseValueFunc func(m map[string]interface{}) interface{}

func GetMapInt64Value(m map[string]interface{}, key string, value *int64) {
	if v, ok := m[key].(int64); ok {
		*value = v
	}
	if v, ok := m[key].(float64); ok {
		*value = int64(v)
	}

	if v, ok := m[key].(string); ok {
		s, _ := strconv.ParseInt(v, 10, 64)
		*value = s
	}

}

func GetMapFloat64Value(m map[string]interface{}, key string, value *float64) {
	if v, ok := m[key].(float64); ok {
		*value = v
	}

	if v, ok := m[key].(decimal.Decimal); ok {
		f, _ := v.Float64()
		*value = f
	}
	if v, ok := m[key].(decimal.NullDecimal); ok {
		f, _ := v.Decimal.Float64()
		*value = f
	}
}

func GetMapStringValue(m map[string]interface{}, key string, value *string) {
	if v, ok := m[key].(string); ok {
		*value = v
	}
	if v, ok := m[key].(float64); ok {
		*value = strconv.FormatFloat(v, 'E', -1, 64)
	}

	if v, ok := m[key].(int64); ok {
		*value = strconv.FormatInt(v, 64)
	}
	if v, ok := m[key].(time.Time); ok {
		*value = StringfyDateToRFC3339(v)
	}

}

func GetMapBoolValue(m map[string]interface{}, key string, value *bool) {
	if v, ok := m[key].(bool); ok {
		*value = v
	}

}
