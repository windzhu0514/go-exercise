package meituanclib

import (
	"log"
	"testing"
)

func TestSign(t *testing.T) {
	var input = `POST /uuid/v2/register __reqTraceID=a371cff3-c592-4dc9-a596-cf749b70de1c&__skck=6a375bce8c66a0dc293860dfa83833ef&__skno=50c1ff27-7acf-4ec3-82bd-c553e1af7856&__skts=1528448412542&__skua=01eea360c6bbf0024c0390f2d1a38a45&__skvs=1.1&__vhost=api.mobile.meituan.com&android_id=8c353e0422a71951&app_name=tower&app_version=1.6.1&brand=Cubietech&ci=&ci=&ctype=android&devid=26E61D9DF653EC9&model=CC-A80&msid=&msid=&serial_number=38908c0030d01ea6e68f&userid=-1&utm_campaign=AtowerBgroupC0E0&utm_content=26E61D9DF653EC9&utm_medium=android&utm_source=yyb-lx&utm_term=570&version_name=8.7`

	if output, err := Sign(input); err != nil {
		log.Println(err.Error())
	} else {
		log.Println("output is: ", string(output))
	}
}

func TestSiua(t *testing.T) {
	var input = `1.0}}nanopi3|Android|aosp_nanopi3|nanopi3|LMY48G|en|US|FriendlyARM (www.arm9.net)|AOSP on NanoPi 3|5.1.1|22|dev-keys|Android/aosp_nanopi3/nanopi3:5.1.1/LMY48G/root04051713:userdebug/dev-keys|nanopi3|jensen|userdebug|nanopi3|armeabi-v7a|armeabi|aosp_nanopi3-userdebug 5.1.1 LMY48G eng.root.20170405.171324 dev-keys|0|1|}}mtp,adb|mtp,adb|mtp,adb|||wlan0|NOT_READY||0|0|1|1|1|0|1|1|0|1|1|1}}8|-|-|-|-|}}-|-|-|672*1280|5GB|5GB|94:a1:a2:bc:fe:f6||wifi}}-|0|1.0|0|0||46e1dd97-678c-4b45-82b1-502b044a0d54}}0|0|0|}}Android|com.ly.meituan|1.0|22|-|2018-05-15 06:33:18:018}}0.0|0.0|wu|50:64:2b:94:1f:3b|1|-67|-|-|-|}}`

	if output, err := Siua(input); err != nil {
		log.Println(err.Error())
	} else {
		log.Println("output is: ", string(output))
	}
}
