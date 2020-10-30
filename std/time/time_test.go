package time__test

import (
	"fmt"
	"math/rand"
	"strconv"
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
	tt, err := time.Parse("2006年01月02日", "2020年10月24日")
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

func TestCalcYearsold(t *testing.T) {
	yearsold := func() string {
		birthdate, err := time.Parse("2006-01-02", "1979-06-20")
		if err != nil {
			return ""
		}

		now := time.Now()
		yearsold := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Sub(birthdate).Hours() / (24 * 365)

		return strconv.Itoa(int(yearsold))
	}()
	fmt.Println(yearsold)
}

// 定时器
func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
