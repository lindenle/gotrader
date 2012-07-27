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

func get_quote (ticker string) Quote {
	// here we need to get the csv url 
	fmt.Printf("%s\n",ticker)
	p := new(Quote)
	return *p
}


