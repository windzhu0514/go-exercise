package main

import (
	"fmt"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestOne(t *testing.T) {
	js := `var stCitySelectArr = [
{id:'619395',name:'安康市',pinyinUrl:'ankangshi',pinyin:'ankang',jianpin:'A',provinceId:'24',provinceName:'陕西',cityFlag:'ticket',cityNameDesc:'安康',repeatFlag:'N',upName:'陕西',showName:'安康',orderFlag :'Y'},
{id:'84656',name:'阿克塞县',pinyinUrl:'akesaixian',pinyin:'akesai',jianpin:'A',provinceId:'25',provinceName:'甘肃',cityFlag:'ticket',cityNameDesc:'阿克塞',repeatFlag:'N',upName:'酒泉市',showName:'阿克塞',orderFlag :'Y'}
];
var stlowerCityArr = [];
var stupperCityArr = [];`

	vm := otto.New()
	fmt.Println(vm.Run(js))
	v, err := vm.Run("stCitySelectArr[0].id")
	fmt.Println(err)
	fmt.Println(v.Export())
}
