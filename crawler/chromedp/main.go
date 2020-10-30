// Command simple is a chromedp example demonstrating how to do a simple google
// search.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-exercise/crawler/chromedp/driver"

	"github.com/chromedp/chromedp"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	defer driver.Begin().End()

	// run task list
	var site, res string
	err := login()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("登录成功")

	err = querycoach()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("查询成功")

	err = createorder()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("创单成功")

	fmt.Printf("saved screenshot from search result listing `%s` (%s)\n", res, site)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGKILL, syscall.SIGTERM)
	<-termChan
}

// 登录
func login() error {
	// 打开登录页面
	var n int
	if err := driver.RunWithTimeout3(time.Second*10, chromedp.Tasks{
		driver.Navigate(`https://www.changtu.com/`),
		chromedp.WaitReady(`#headerPageCode`, chromedp.ByID),
		chromedp.ActionFunc(func(context.Context) error {
			fmt.Println("打开主页成功", n)
			return errors.New("打开主页成功")
		}),
		chromedp.Click(`//ul[@id='loginNav_unlogin']/li[3]/a`),
		chromedp.WaitReady(`#comLoginSubmit`, chromedp.ByID),
		chromedp.ActionFunc(func(context.Context) error {
			fmt.Println("打开登录页成功")
			return nil
		}),
		chromedp.Sleep(time.Second * 3),
	}); err != nil {
		log.Println(err)
		return err
	}

	if err := driver.Run(chromedp.Tasks{
		chromedp.SendKeys(`#comUserName`, "15862430163", chromedp.ByID),
		chromedp.Sleep(time.Second * 3),
		chromedp.SendKeys(`#password`, "windzhu0514", chromedp.ByID),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click("#comLoginSubmit", chromedp.ByID),
		chromedp.WaitVisible(`#qaPageNickname`, chromedp.ByID),
		chromedp.ActionFunc(func(context.Context) error {
			fmt.Println("登录成功")
			return nil
		}),
	}); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func querycoach() error {

	chooseCity := `//*[@id='trip8080chooseCity0'][@title='` + `苏州` + `']`

	// 出发城市
	if err := driver.RunWithTimeout3(time.Second*3, chromedp.Tasks{
		chromedp.Clear(`#bookStartCityName`, chromedp.ByID),
		chromedp.Sleep(time.Second),
		chromedp.SendKeys(`#bookStartCityName`, "苏州", chromedp.ByID),
		chromedp.Clear(`#bookStartCityName`, chromedp.ByID),
		chromedp.Sleep(time.Second),
		chromedp.WaitVisible(`#trip8080citySelectChoose`, chromedp.ByID),
		// chromedp.Text(`#trip8080chooseCity0[@title]`, &cityName, chromedp.BySearch),
		// chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		// 	fmt.Println("cityName", cityName)
		// 	if cityName != "苏州" {
		// 		return errors.New("未找到出发城市")
		// 	}

		// 	return nil
		// }),
		chromedp.Click(chooseCity, chromedp.BySearch),
	}); err != nil {
		return err
	}

	chooseCity = `//*[@id="trip8080chooseCity0"][@title='` + `上海` + `']`

	// 到达城市
	if err := driver.RunWithTimeout3(time.Second*3, chromedp.Tasks{
		chromedp.Clear(`#bookEndCityName`, chromedp.ByID),
		chromedp.Sleep(time.Second),
		chromedp.SendKeys(`#bookEndCityName`, "上海", chromedp.ByID),
		chromedp.Sleep(time.Second),
		chromedp.WaitVisible(`#trip8080citySelectChoose`, chromedp.NodeVisible, chromedp.ByID),
		// chromedp.Text(`#trip8080chooseCity0[@title]`, &cityName, chromedp.ByID),
		// chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		// 	fmt.Println("cityName", cityName)
		// 	if cityName != "上海" {
		// 		return errors.New("未找到到达城市")
		// 	}

		// 	return nil
		// }),
		chromedp.Click(chooseCity, chromedp.BySearch),
	}); err != nil {
		return err
	}

	date := time.Now().AddDate(0, 0, 1).Format("2006-01-02") // 第二天
	dataCell := `//div[@id='trip8080datePicker']//td[a[@id='` + date + `']]`

	// 选择日期
	if err := driver.RunWithTimeout3(time.Second*3, chromedp.Tasks{
		chromedp.Click(`#bookStartDate`, chromedp.ByID),
		chromedp.WaitVisible(`#trip8080datePicker`, chromedp.ByID),
		chromedp.Sleep(time.Second),
		chromedp.Click(dataCell, chromedp.BySearch),
		chromedp.Sleep(time.Second),
	}); err != nil {
		return err
	}

	// date := time.Now().AddDate(0, 0, 1).Format("2006-01-02") // 第二天
	// if err := driver.Run(chromedp.Tasks{
	// 	chromedp.Clear(`#bookStartCityName`,  chromedp.ByID),
	// 	chromedp.SendKeys(`#bookStartCityName`, "苏州", chromedp.ByID),
	// 	chromedp.Sleep(time.Second),
	//
	// 	chromedp.Clear(`#bookEndCityName`,  chromedp.ByID),
	// 	chromedp.SendKeys(`#bookEndCityName`, "上海",  chromedp.ByID),
	// 	chromedp.Sleep(time.Second),
	//
	// 	chromedp.SendKeys(`#bookStartDate`, date,  chromedp.ByID),
	// 	chromedp.Sleep(time.Second),
	// }); err != nil {
	// 	return err
	// }

	// 查询车次
	if err := driver.Run(chromedp.Tasks{
		chromedp.Click("#bookBtn", chromedp.ByID),
		chromedp.WaitReady(`#go_list`, chromedp.ByID),
		chromedp.Sleep(time.Second),
	}); err != nil {
		return err
	}

	//dptTime, dptStationName, arrStationName, price := "05:40", "苏州北广场站", "上海虹桥机场", "53.00"
	//[text()[contains(., '05:40')]][text()[contains(., '苏州北广场站')]][text()[contains(., '上海虹桥机场')]][text()[contains(., '53.00')]]
	// [contains(text(), '05:40')]]
	// //li[0][text()[contains(., '05:40')]]
	// //li[text()[contains(., '06:00')]]
	if err := driver.Run(chromedp.Tasks{
		chromedp.Click(`//div[@class='gocomme_cj_list schListDiv']//li[7]/input[2][contains(@onclick,"06:15")][contains(@onclick,"同里镇汽车站")][contains(@onclick,"上海南站")]`, chromedp.BySearch),
		chromedp.Sleep(time.Second * 1),
		chromedp.WaitVisible(`#ticketTime`, chromedp.ByID),
		chromedp.Click(`//div[@id="ticketTime"]//input[@onclick='ticketTimeContinue();']`, chromedp.BySearch),
		chromedp.WaitReady(`#header`, chromedp.ByID),
		chromedp.Sleep(time.Second * 3),
	}); err != nil {
		fmt.Println("************", err)
		return err
	}

	return nil
}

func createorder() error {
	if err := driver.RunWithTimeout(time.Second*20, chromedp.Tasks{
		// 添加乘客
		chromedp.Click(`#addPassengerButton`, chromedp.NodeVisible, chromedp.ByID),
		chromedp.SendKeys(`//*[@id="addPassengers_body"]/div[1]/span/span[1]/input`, "章茂林"),
		chromedp.SendKeys(`//*[@id="addPassengers_body"]/div[3]/span/span[1]/input`, "430381198207197113", chromedp.BySearch),
		chromedp.SendKeys(`//*[@id="addPassengers_body"]/div[5]/span/span[1]/input`, "15862430163", chromedp.BySearch),
		chromedp.Sleep(time.Second * 3),
		chromedp.Click(`//*[@id="addPassengers_body"]/div[6]/input`),

		chromedp.Sleep(time.Second * 2),
		chromedp.Click(`//*[@id="goSafe"]/div[2]/div[1]/ul/li[2]/div/div[6]/div[2]/input`),
		chromedp.Sleep(time.Second * 1),
		chromedp.Click(`//*[@id="goSafe"]/div[2]/div[1]/ul/li[2]/div/div[6]/div[1]/p[1]/a`, chromedp.NodeVisible),
		chromedp.Sleep(time.Second * 1),

		chromedp.Click(`//*[@id="contact_nameqita"]`), // 选择其他联系人
		chromedp.Sleep(time.Second * 1),
		chromedp.Clear(`//*[@id="contact_mes"]/ul/li[1]/span[1]/input`),
		chromedp.SendKeys(`//*[@id="contact_mes"]/ul/li[1]/span[1]/input`, "章茂林"),
		chromedp.Sleep(time.Second * 1),
		chromedp.Clear(`//*[@id="contactIdCode"]`),
		chromedp.SendKeys(`//*[@id="contactIdCode"]`, "430381198207197113"),
		chromedp.Sleep(time.Second * 1),

		//chromedp.Clear(`//*[@id="contact_phone"]`, chromedp.BySearch),
		//chromedp.SendKeys(`//*[@id="contact_phone"]`, "", chromedp.BySearch),
		//chromedp.SendKeys(`//*[@id="contact_phone"]`, "15862430163", chromedp.BySearch),

		chromedp.Sleep(time.Second * 2),
		chromedp.Click(`//input[@id="submitBtn"]`), // 提交订单

		chromedp.WaitVisible(`#submitOnlinePay`, chromedp.ByID),
		chromedp.Sleep(time.Second * 5),
	}); err != nil {
		fmt.Println("************", err)
		return err
	}

	return nil
}
