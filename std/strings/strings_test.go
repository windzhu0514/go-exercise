package strings

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrim(t *testing.T) {
	fmt.Println(strings.TrimRight("normalCheckInMap_psgrName12", "432"))
	fmt.Println(strings.TrimRight("normalCheckInMap_psgrName12", "12"))
}

func TestField(t *testing.T) {
	isWide := true
	isVipSeat := true
	rowInfo := "######"
	if !isWide && !isVipSeat {
		rowInfo = rowInfo[:3] + "=" + rowInfo[3:]
	} else if !isWide && isVipSeat {
		rowInfo = rowInfo[:2] + "=" + rowInfo[2:]
	} else if isWide && !isVipSeat {
		rowInfo = rowInfo[:2] + "=" + rowInfo[2:]
		rowInfo = rowInfo[:7] + "=" + rowInfo[7:]
	} else {
		rowInfo = rowInfo[:2] + "=" + rowInfo[2:]
		rowInfo = rowInfo[:5] + "=" + rowInfo[5:]
	}
	fmt.Println(rowInfo)
}
