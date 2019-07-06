package debug

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
type Record struct {
	Value interface{} `json:"value"`
	Time time.Time `json:"time"`
}
var recordList []Record

func Check (err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func htmlController(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("debug.html");Check(err)
	html := string(htmlByte)
	w.Header().Set("Content-Type", "text/html")
	_, err = fmt.Fprintf(w, html);Check(err)
}
func debugController(w http.ResponseWriter, r *http.Request) {
	var dataVO = map[string]interface{}{
		"type": "pass",
		"data": recordList,
	}
	//log.Print(recordList)
	dataJO, _ := json.Marshal(dataVO)
	_, _ = fmt.Fprintf(w, string(dataJO))
}
func startServer () {
	http.HandleFunc("/", htmlController)
	http.HandleFunc("/debug", debugController)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Print("listen success")
	}
}
func Debug (v ...interface{}) {
	for _, value := range v {
		recordList = append(recordList, Record{Time: time.Now(), Value: value})
		output := map[string]interface{}{
			"value": value,
		}
		jsonbyte, _ := json.Marshal(output)
		jsonstring := string(jsonbyte)
		log.Printf("%+v", jsonstring)
	}
	startServer()
}
