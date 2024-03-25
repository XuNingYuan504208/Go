package slice

/*
slice 遍历
遍历过程中可更改slice内容，引用类型
*/
func SliceErgodic(slice1 []int) {
	// for _,value := range slice1
	for index, value := range slice1 {
		println("slice1 :: index : ", index, ", value : ", value)
	}
}

/*
slice append
返回一个新的 slice
*/
func SliceAppend(slice1 []int) {
	slice1 = append(slice1, 4)
}
