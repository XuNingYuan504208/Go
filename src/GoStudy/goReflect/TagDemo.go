package goReflect

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" db:"user_name"`
	Age  int    `json:"age" db:"user_age"`
}

func main() {
	// 创建一个 User 结构体实例
	user := User{Name: "Alice", Age: 30}

	// 获取 User 结构体类型
	userType := reflect.TypeOf(user)

	//
	//。
	/*
		Elem() 方法用于获取指针、数组、切片、字典、通道或接口的元素类型
		因此，如果 user 是一个指针类型，那么这行代码获取的是该指针指向的类型信息，而不是指针本身的类型
		如果 user 是一个指向结构体的指针，那么这行代码获取的就是该结构体的类型信息
	*/
	//userType := reflect.TypeOf(user).Elem()

	// 遍历 User 结构体的字段
	for i := 0; i < userType.NumField(); i++ {
		// 获取字段信息
		field := userType.Field(i)
		// 获取字段的名称
		fieldName := field.Name
		// 获取字段的类型
		fieldType := field.Type
		// 获取字段的 Tag
		fieldTag := field.Tag
		// 字段Tag的json key(在已知Key的前提)
		jsonInfo := field.Tag.Get("json")

		// 打印字段信息和 Tag
		fmt.Printf("Field Name: %s, Type: %s, Tag: %s\n , jsonTag : %s\n", fieldName, fieldType, fieldTag, jsonInfo)
	}

}

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year" `
	Price  int      `json:"rmb" `
	Actors []string `json:"actors"`
}

func mainTag() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	// json.Marshal() 将接口类型的数据结构（通常是结构体、切片、数组、map 等）序列化为 JSON 格式的字节切片-> 就是把对象转成json
	// 返回一个表示 JSON 格式的字节切片和一个错误值。
	jsonBytes, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}

	// 打印 JSON 格式的字节切片
	/*
		json.Marshal() 方法会根据结构体字段的标签（Tag）来确定 JSON 字段的名称。例如，Price 字段的标签为 json:"rmb"，因此在生成的 JSON 中，该字段会被命名为 rmb
	*/
	fmt.Println("jsonBytes : ", string(jsonBytes))

	// 解码 jsonStr -> struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonBytes, &myMovie)
	if err != nil {
		println("json unmarshal error : ", err)
		return
	}
	println(myMovie)
}
