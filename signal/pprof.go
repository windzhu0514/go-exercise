// +build !windows

package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

var openpprof bool

func pprofHandler() {
	go func() {
		chSig := make(chan os.Signal, 1)
		signal.Notify(chSig, syscall.SIGUSR1, syscall.SIGUSR2)
		for {
			switch <-chSig {
			case syscall.SIGUSR1:
				openpprof = true
				fmt.Println("syscall.SIGUSR1")
			case syscall.SIGUSR2:
				openpprof = false
				fmt.Println("syscall.SIGUSR2")
			}
		}
	}()
}

func index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if openpprof {
			fmt.Println("openpprof")
			pprof.Index(w, r)
			return
		}

		fmt.Println("not openpprof")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func cmdline() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if openpprof {
			fmt.Println("openpprof")
			pprof.Cmdline(w, r)
			return
		}

		fmt.Println("not openpprof")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func profile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if openpprof {
			fmt.Println("openpprof")
			pprof.Profile(w, r)
			return
		}

		fmt.Println("not openpprof")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
func symbol() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if openpprof {
			fmt.Println("openpprof")
			pprof.Symbol(w, r)
			return
		}

		fmt.Println("not openpprof")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
func trace() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if openpprof {
			fmt.Println("openpprof")
			pprof.Trace(w, r)
			return
		}

		fmt.Println("not openpprof")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
