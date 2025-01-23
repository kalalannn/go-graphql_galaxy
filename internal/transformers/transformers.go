package transformers

import "time"

const timeFormat = time.RFC3339

type TimeOrPointer interface {
	time.Time | *time.Time
}

func TransformTimeToString[T TimeOrPointer](t T) *string {
	var result string
	switch v := any(t).(type) {
	case time.Time:
		result = v.Format(timeFormat)
	case *time.Time:
		if v != nil {
			result = v.Format(timeFormat)
		} else {
			return nil
		}
	}
	return &result
}
