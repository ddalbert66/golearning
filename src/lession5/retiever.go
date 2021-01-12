package main

import (
	"fmt"
	"goLearning20200930/src/lession5/mock"
	"goLearning20200930/src/lession5/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.ltn.com.tw")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is www.con.tw"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents", v.Content)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}
