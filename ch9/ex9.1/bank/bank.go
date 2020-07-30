// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

type wTask struct {
	amount int
	res chan bool
}
var withdraw = make(chan wTask)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) {
	res := make(chan bool, 1)
	withdraw <- wTask{amount:amount, res: res}
	if <-res {
		fmt.Println("withdraw succ")
	} else {
		fmt.Println("withdraw fail")
	}
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdraw:
			if balance>=w.amount {
				balance -= w.amount
				w.res <- true
			} else {
				w.res <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
