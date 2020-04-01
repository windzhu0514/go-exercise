package main

import (
	"log"

	"go-exercise/cgo/cgotest3/meituanclib"
)

func main() {
	var input = `POST /uuid/v2/register __reqTraceID=a371cff3-c592-4dc9-a596-cf749b70de1c&__skck=6a375bce8c66a0dc293860dfa83833ef&__skno=50c1ff27-7acf-4ec3-82bd-c553e1af7856&__skts=1528448412542&__skua=01eea360c6bbf0024c0390f2d1a38a45&__skvs=1.1&__vhost=api.mobile.meituan.com&android_id=8c353e0422a71951&app_name=tower&app_version=1.6.1&brand=Cubietech&ci=&ci=&ctype=android&devid=26E61D9DF653EC9&model=CC-A80&msid=&msid=&serial_number=38908c0030d01ea6e68f&userid=-1&utm_campaign=AtowerBgroupC0E0&utm_content=26E61D9DF653EC9&utm_medium=android&utm_source=yyb-lx&utm_term=570&version_name=8.7`

	if output, err := meituanclib.Sign(input); err != nil {
		log.Println(err.Error())
	} else {
		log.Println("output is: ", string(output))
	}
}
