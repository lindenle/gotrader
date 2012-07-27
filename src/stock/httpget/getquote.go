package httpget

import (
	"fmt"
//	"time"
)

type Quote struct {
	Date int
	Open, High, Low, Close, Volume  float32 
	Adj int
}

func Getquote (ticker string) Quote {
	// here we need to get the csv url 
	baseurl := "http://ichart.finance.yahoo.com/table.csv?s="
	opturl := "&d=0&e=28&f=2012&g=m&a=3&b=12&c=2000&ignore=.csv"
	fmt.Printf("%s %s%s%s\n",ticker,baseurl,ticker,opturl)
	p := new(Quote)
	return *p
}


