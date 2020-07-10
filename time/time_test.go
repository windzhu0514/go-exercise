package time__test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println("1589881203731")
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Nanosecond() / 1e6)
}

func TestMilliSecond2Time(t *testing.T) {
	// 毫秒转时间
	tt := time.Unix(1593655422587/1000, 1593655422587%1000000)
	fmt.Println(tt.Format("2006-01-02 15:04:05"))
}

func TestSecond2Time(t *testing.T) {
	// 秒转时间
	tt := time.Unix(1494668822, 0)
	fmt.Println(tt.Format("2006-01-02 15:04:05"))
}

func TestParse(t *testing.T) {
	tt, err := time.Parse("1月02日15点", "5月28日18点")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tt)
	fmt.Println(tt.Month(), tt.Day(), tt.Hour())
}

func TestAdd(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
	now = now.Add(-time.Second * 10)
	fmt.Println(now.Unix())
}

func TestRandTime(tt *testing.T) {
	for i := 0; i < 10; i++ {
		y := rand.Intn(3) + 1  // 提前1到3年
		m := rand.Intn(12) - 6 // -6 > +5
		d := rand.Intn(28)
		t := time.Now().AddDate(-y, m, d)
		t = t.Add(time.Hour * time.Duration((rand.Intn(12) - 6)))
		fmt.Println(t.Format("2006-01-02 15:04:05"))
	}
}

func TestDuration(t *testing.T) {
	d := time.Hour*3 + time.Minute*5
	fmt.Println(d)
	fmt.Println(d.Round(time.Hour))
}
