package goMap

/*
判断map的key是否存在
*/
func existKey(m map[string]string) bool {
	value, ok := m["key"]
	if ok {
		// 键存在
		println(value)
	} else {
		// 键不存在
	}
	return ok
}
