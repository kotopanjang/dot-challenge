package helper

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

const (
	RoundingAuto = "RoundAuto"
	RoundingUp   = "RoundUp"
	RoundingDown = "RoundDown"
)

// ToString is to convert any type to sring
func ToString(s any) string {
	switch s := s.(type) {
	case float64:
		return decimal.NewFromFloat(s).String()
	default:
		return fmt.Sprintf("%v", s)
	}
}

// ToInt is to convert any type to Int64
// if the convertion error, will return 0 instead
func ToInt(i any) int64 {
	switch i := i.(type) {
	case string:
		res, err := strconv.Atoi(i)
		if err != nil {
			return 0
		}
		return int64(res)
	case float64:
		return int64(i)
	}

	return 0
}

// ToFloat64 is to convert any type to float
func ToFloat64(o interface{}, decimalPoint int, rounding string) float64 {
	//fmt.Printf("\ndec: %v\n", decimalPoint)
	if IsPointer(o) {
		return float64(0)
	}

	var f float64
	var e error

	t := strings.ToLower(typeName(o))
	v := value(o)

	if t != "interface{}" && strings.HasPrefix(t, "int") {
		f = float64(v.Int())
	} else if strings.HasPrefix(t, "uint") {
		f = float64(v.Uint())
	} else if strings.HasPrefix(t, "float") {
		f = float64(v.Float())
	} else {
		f, e = strconv.ParseFloat(v.String(), 64)
		if e != nil {
			return 0
		}
	}

	//fmt.Printf("\ndec: %v\n", decimalPoint)
	switch rounding {
	case RoundingAuto:
		return RoundingAuto64(f, decimalPoint)
	case RoundingDown:
		return RoundingDown64(f, decimalPoint)
	case RoundingUp:
		return RoundingUp64(f, decimalPoint)
	}

	if math.IsNaN(f) || math.IsInf(f, 0) {
		f = 0
	}

	return f
}

func IsPointer(o interface{}) bool {
	v := reflect.ValueOf(o)
	return v.Kind() == reflect.Ptr
}

func RoundingAuto64(f float64, decimalPoint int) (retValue float64) {

	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow

	if f < 0 {
		f = math.Ceil(f - 0.5)
	} else {
		f = math.Floor(f + 0.5)
	}

	retValue = f / tempPow
	return
}

func RoundingDown64(f float64, decimalPoint int) (retValue float64) {
	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow
	f = math.Floor(f)
	retValue = f / tempPow
	return
}

func RoundingUp64(f float64, decimalPoint int) (retValue float64) {
	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow
	f = math.Ceil(f)
	retValue = f / tempPow
	return
}

func value(o interface{}) reflect.Value {
	return reflect.ValueOf(o)
}

func typeName(o interface{}) string {
	typeName := fmt.Sprintf("%T", o)

	switch o.(type) {
	case string:
		typeName = "string"
	case int:
		typeName = "int"
	case int32:
		typeName = "int32"
	case float64:
		typeName = "float64"
	case bool:
		typeName = "bool"
	}

	return typeName
}
