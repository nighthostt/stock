package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Record struct {
	Date     string
	Open     string
	High     string
	Close    string
	Volume   string
	Chg      string
	Pchg     string
	Ma5      string
	Ma10     string
	Ma20     string
	Vma5     string
	Vma10    string
	Vma20    string
	Turnover string
}
type Stocks struct {
	Stock []string
}
type Records struct {
	Recs string `json:"record"`
}

func Run() {
	resp, err := http.Get("http://api.finance.ifeng.com/akdaily/?code=sh689009&type=last")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//var re Records
	//_ = json.Unmarshal([]byte(body), &re)
	//fmt.Printf("re = %v\n", re)
	data := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&data)
	fmt.Printf("data = %v", data["record"])
	//fmt.Println(string(body))
}
