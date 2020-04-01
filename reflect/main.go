package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type fullName struct {
	FName string `json:"fname"`
	MName string `json:"mName"`
	LName string `json:"lname"`
}

type people struct {
	Name   fullName
	Sex    string `json:"sex"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	cat    cat
}

func (p people) String() string {
	return "my name is " + p.Name.FName + " " + p.Name.MName + " " + p.Name.LName + "," +
		"sex is " + p.Sex + ",Height is " + strconv.Itoa(p.Height) + ",Weight is " + strconv.Itoa(p.Weight)
}

type cat struct {
	Name  *fullName
	Color string `json:"color"`
}

type dog struct {
	name  fullName
	Color string `json:"color"`
	Breed string `json:"breed"`
}

type Foo struct {
	A int `tag1:"Tag1" tag2:"Second Tag"`
	B string
}

func main() {
	// 反射的使用
	// s := "String字符串"
	// fo := Foo{A: 10, B: "字段String字符串"}
	//
	// sVal := reflect.ValueOf(s)
	// // 在没有获取指针的前提下，我们只能读取变量的值。
	// fmt.Println(sVal.Interface())
	//
	// sPtr := reflect.ValueOf(&s)
	// sPtr.Elem().Set(reflect.ValueOf("修改值1"))
	// sPtr.Elem().SetString("修改值2")
	// // 修改指针指向的值，原变量改变
	// fmt.Println(s)
	// fmt.Println(sPtr) // 要注意这是一个指针变量，其值是一个指针地址
	//
	// foType := reflect.TypeOf(fo)
	// fmt.Println(foType)
	// foVal := reflect.New(foType)
	// // foVal.Elem().Field(0).SetString("A") // 引发panic
	// foVal.Elem().Field(0).SetInt(1)
	// foVal.Elem().Field(1).SetString("B")
	// f2 := foVal.Elem().Interface().(Foo)
	// fmt.Printf("%+v, %d, %s\n", f2, f2.A, f2.B)
	// var d = dog{fullName{"dahuang", "dahuang", "dahuang"}, "zzz", "abc"}
	// bdata, err := json.Marshal(d)
	// fmt.Println(string(bdata), err)

	// num := reflect.ValueOf(d).NumField()
	// for i := 0; i < num; i++ {
	// 	fmt.Println(reflect.TypeOf(d).Field(i).Name)
	// }

	// var jsonData = `{"hotelCode":"4393","mebid":258919958,"overseaOrder":0,"totalFee":152400,"appID":"A00006","appVersion":"3.1.2","clientIp":"1.1.1.1","extensions":"{\u0027business\u0027:null,\u0027channelSourceType\u0027:1}","merchantNo":"00001","merchantOrderNo":"102256423338","nonce":"gFKNeEs77FHQlpdOW1Ge4ow6ns2","os":"android","phoneModel":"MI 3W","sid":"sid","signData":"40288C81B6B16B8AE224770EB4CE4E53","timestamp":1553499141,"token":"4840edd883179b9072acccffaa35ca1c"}`
	// var m map[string]interface{}
	// json.Unmarshal([]byte(jsonData), &m)
	//
	// jsonBytes, _ := json.Marshal(m)
	// fmt.Println(string(jsonBytes))

	// var keys []string
	// for k := range m {
	// 	fmt.Print(k, " ")
	// 	keys = append(keys, k)
	// }
	// fmt.Println()
	//
	// sort.Strings(keys)
	// fmt.Println(keys)
	//
	// var m2 = map[string]interface{}{}
	// for _, v := range keys {
	// 	m2[v] = m[v]
	// }
	//
	// for k := range m2 {
	// 	fmt.Print(k, " ")
	// }
	// fmt.Println()
	type siteAccountInfo struct {
		Pulling int32
	}

	var siteAccountInfos = map[int]siteAccountInfo{}
	siteAccountInfos[1] = siteAccountInfo{100}

	if siteAccountInfo, ok := siteAccountInfos[1]; ok {
		fmt.Printf("%p\n", &siteAccountInfo)
		siteAccountInfos[1].Pulling = 2
	}

}

func setName(param interface{}, resp interface{}) {

	// param 不同的参数输入 具有同一个Name结构体字段
	full := reflect.ValueOf(param).Elem().FieldByName("Name") // Elem 获取接口包含的原始值

	// 修改值
	full.Set(reflect.ValueOf(fullName{FName: "FFdog", MName: "SSdog", LName: "LLdog"}))

	b, _ := json.Marshal(full.Interface()) // 编码结构体的一部分
	fmt.Println(string(b))

	// 解码到结构体的指定字段
	r := reflect.ValueOf(resp).Elem().FieldByName("Name")
	err := json.Unmarshal(b, r.Addr().Interface()) // r是结构体Name对应的reflect.Value类型值
	fmt.Println(err)

	fmt.Println(r)
}
