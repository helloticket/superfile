package helper

func Less(a, b interface{}) bool {
	switch a.(type) {
	case int:
		if ai, ok := a.(int); ok {
			if bi, ok := b.(int); ok {
				return ai < bi
			}
		}
	case int64:
		if ai, ok := a.(int64); ok {
			if bi, ok := b.(int64); ok {
				return ai < bi
			}
		}
	case string:
		if ai, ok := a.(string); ok {
			if bi, ok := b.(string); ok {
				return ai < bi
			}
		}
	default:
	}
	return false
}
