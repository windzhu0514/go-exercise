package main

func main() {
	// p1 := &china{word: "我是中国人"}
	// p2 := &african{word: "我是非洲人"}
	// p3 := &european{word: "我是欧洲人"}
	//
	// // speakers := []speaker{p1, p2}
	// // for n := range speakers {
	// // 	fmt.Println("people: ", speakers[n])
	// // 	speakers[n].speak()
	// // }
	//
	// // peopleSay(p1)
	// // peopleSay(p2)
	// // peopleSay(p3)
	//
	// manyPeopleSay(p1, p2, p3)
	var logginer Loginer
	logginer.Login("ci", "logPrefix")
}

type Loginer interface {
	Login(ci, logPrefix string) (int, error)
}
