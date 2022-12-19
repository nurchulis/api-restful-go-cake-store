package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"bytes"
)

type ListTest struct {
	ListTest []TestList `json:"list_test"`
}
 
type TestList struct {
	Method string `json: "method"`
	Module string `json: "module"`
	Url string `json: "url"`
	Result string `json: "result"`
	Payload string `json: "payload"`
}

func TestGetEntries(t *testing.T) {

	file, _ := ioutil.ReadFile("test.json")
 
	data := ListTest{}
 
	_ = json.Unmarshal([]byte(file), &data)
 
	for i := 0; i < len(data.ListTest); i++ {
		fmt.Println("Method : ", data.ListTest[i].Method)
		fmt.Println("Module : ", data.ListTest[i].Module)
		fmt.Println("Url : ", data.ListTest[i].Url)
		//fmt.Println("Expect Result : ",  strings.ReplaceAll( data.ListTest[i].Result, `'`, `"`))
		fmt.Println("Payload : ",  strings.ReplaceAll( data.ListTest[i].Payload, `'`, `"`))

		var jsonStr = []byte(`{"title":"Blueberry cheesecake","description":"cheesecake made of Blueberry","rating":4,"image":"https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11sunny-lemon-cheesecake-102220-1.jpeg"}`)
		req, err := http.NewRequest(data.ListTest[i].Method, data.ListTest[i].Url, bytes.NewBuffer(jsonStr))

		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CakeService)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		expected := strings.ReplaceAll( data.ListTest[i].Result, `'`, `"`)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}else{
			fmt.Println("Testing Passed : ", data.ListTest[i].Module )
		}

	}


}
