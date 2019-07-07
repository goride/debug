package debug

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)
type Record struct {
	Value interface{} `json:"value"`
	Time time.Time `json:"time"`
}
var recordList []Record

func check (err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func Debug (v ...interface{}) {
	recordList = append(recordList, Record{Time: time.Now(), Value: v})
	for _, value := range v {
		output := map[string]interface{}{
			"value": value,
		}
		jsonbyte, err := json.Marshal(output)
		check(err)
		jsonstring := string(jsonbyte)
		log.Printf("%+v", jsonstring)
	}
	postValueJSONbyte, err := json.Marshal(recordList)
	check(err)
	postValueJSON := string(postValueJSONbyte)
	resp, err := http.Post("http://127.0.0.1:9999/debug/add",
		"application/x-www-form-urlencoded",
		strings.NewReader("json=" + postValueJSON))
	check(err)
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//check(err)
	//log.Println(string(body))

}
