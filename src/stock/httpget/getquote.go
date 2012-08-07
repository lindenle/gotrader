package httpget

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	url := fmt.Sprintf("%s%s%s",baseurl,ticker,opturl)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatal(err)
		
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range strings.Split(string(body),"\n") {
		if line != "" {
			new Quote(line.split(","))
		}
	}

	p := new(Quote)
	return *p
}


