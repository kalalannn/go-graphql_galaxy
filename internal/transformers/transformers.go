package transformers

import (
	"math"
	"time"
)

const TimeFormat = time.RFC3339

type TimeOrPointer interface {
	time.Time | *time.Time
}

func TransformTimeToString[T TimeOrPointer](t T) *string {
	var result string
	switch v := any(t).(type) {
	case time.Time:
		result = v.Format(TimeFormat)
	case *time.Time:
		if v != nil {
			result = v.Format(TimeFormat)
		} else {
			return nil
		}
	}
	return &result
}

func RoundFloat(num float64, decimals int) float64 {
	mul := math.Pow(10, float64(decimals))
	return math.Round(num*mul) / mul
}
