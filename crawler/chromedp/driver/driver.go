package driver

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

var helper *driver

type driver struct {
	cancelFunc context.CancelFunc
	ctx        context.Context
	cdp        *chromedp.CDP
}

func Begin() interface{ End() } {
	helper = &driver{}
	helper.ctx, helper.cancelFunc = context.WithCancel(context.Background())

	// create chrome instance
	var err error
	helper.cdp, err = chromedp.New(helper.ctx, chromedp.WithLog(log.Printf))
	// helper.cdp, err = chromedp.New(helper.ctx, chromedp.WithTargets(client.New().WatchPageTargets(helper.ctx)), chromedp.WithLog(log.Printf))
	// helper.cdp, err = chromedp.New(helper.ctx, chromedp.WithLog(log.Printf),		chromedp.WithRunnerOptions(runner.Flag("headless", true), runner.Flag("disable-gpu", true)))

	if err != nil {
		log.Fatal(err)
	}

	return helper
}

func (helper *driver) End() {
	err := helper.cdp.Shutdown(helper.ctx)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = helper.cdp.Wait()
	if err != nil {
		log.Fatal(err)
	}

	helper.cancelFunc()
}

// 返回访问错误
func Navigate(urlstr string) chromedp.Action {
	return chromedp.ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		th, ok := h.(*chromedp.TargetHandler)
		if !ok {
			return chromedp.ErrInvalidHandler
		}

		frameID, _, errText, err := page.Navigate(urlstr).Do(ctxt, th)
		if err != nil {
			return err
		}

		err = th.SetActive(ctxt, frameID)
		if err != nil {
			return errors.New(err.Error() + "-" + errText)
		}

		if errText != "" {
			return errors.New(errText)
		}

		return nil
	})
}

func Run(a chromedp.Action) error {
	if helper == nil {
		log.Fatal("must call Begin first")
	}

	return helper.cdp.Run(helper.ctx, a)
}

func RunWithTimeout(d time.Duration, a chromedp.Action) error {
	if helper == nil {
		log.Fatal("must call Begin first")
	}

	ctx, cancel := context.WithTimeout(helper.ctx, d)
	defer cancel()

	return helper.cdp.Run(ctx, a)
}

func RunWithTimeout3(d time.Duration, a chromedp.Action) (err error) {
	if helper == nil {
		log.Fatal("must call Begin first")
	}

	for index := 0; index < 3; index++ {
		ctx, cancel := context.WithTimeout(helper.ctx, d)
		err = helper.cdp.Run(ctx, a)
		if err == nil {
			cancel()

			break
		}
		cancel()
	}

	return
}

// NodesCount 查询节点的一级子节点个数 xpath
func NodesCount(sel interface{}, count *int) chromedp.Action {
	return chromedp.QueryAfter(sel, func(ctxt context.Context, h *chromedp.TargetHandler, n ...*cdp.Node) error {
		*count = len(n)
		return nil
	})
}
