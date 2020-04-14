package field

func noError(val interface{}, err error) interface{} {
	if err != nil {
		return err.Error()
	}
	return val
}
