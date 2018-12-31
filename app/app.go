package app

// MakeJSON make formatted data
func MakeJSON(code int, msg string, v interface{}) []byte {
	bs, _ := json.Marshal(map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": v,
	})

	return bs
}

// JSON make formatted data
func JSON(v interface{}) []byte {
	return MakeJSON(0, "successful", v)
}

// ErrJSON make formatted data
func ErrJSON(code int, msg string) []byte {
	return MakeJSON(code, msg, []int{})
}
