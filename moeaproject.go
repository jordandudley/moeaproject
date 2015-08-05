package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	moeaURL := "http://www.occeweb.com/MOEAsearch/"
	t1 := time.Now()
	response, httpError := http.Get(moeaURL)
	t2 := time.Now()
	fmt.Printf("Http call took %v",t2.Sub(t1))
	defer response.Body.Close()
	if httpError == nil {
		bytes, ioError := ioutil.ReadAll(response.Body)
		fmt.Println(string(bytes))
		fmt.Println(ioError)
	} else {
		fmt.Println("Error:", httpError)
	}

}
