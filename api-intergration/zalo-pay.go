package api_intergration

import (
	"encoding/json"
	"fmt"
	"github.com/VuiPhan/web-service-gin/models"
	"github.com/zpmep/hmacutil"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type object map[string]interface{}

var (
	//app_id = "200002"
	//key1   = "c53PpfuEujVA7IQPjSyMzgbrpfJdtEgS"
	//key2   = "VZkD5d3sD4nnovNeUd5s9gShkle9sox9"

	app_id = "2554"
	key1   = "sdngKKJmqEMzvh5QQcdD2A9XBSKUNaYn"
	key2   = "trMrHtvjo6myautxDUiAcYsVtaeQ8nhf"
)

func CallApiZaloPayToCreateOrder(infoOrder models.Order) models.ResponseCreateOrderFromZaloPay {
	rand.Seed(time.Now().UnixNano())
	embedData, _ := json.Marshal(object{})
	items, _ := json.Marshal(infoOrder.OrderDetail)
	// request data
	params := make(url.Values)
	params.Add("app_id", app_id)
	params.Add("amount", strconv.Itoa(int(infoOrder.TotalAmount)))
	params.Add("app_user", "user123")
	params.Add("embed_data", string(embedData))
	params.Add("item", string(items))
	params.Add("description", "Thanh toán đơn hàng #"+strconv.Itoa(infoOrder.OrderID)+" tại Shop VuiPhan")
	params.Add("bank_code", "zalopayapp")

	now := time.Now()
	params.Add("app_time", strconv.FormatInt(now.UnixNano()/int64(time.Millisecond), 10)) // miliseconds

	params.Add("app_trans_id", fmt.Sprintf("%02d%02d%02d_%v", now.Year()%100, int(now.Month()), now.Day(), infoOrder.OrderID)) // translation missing: vi.docs.shared.sample_code.comments.app_trans_id

	// appid|app_trans_id|appuser|amount|apptime|embeddata|item
	data := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v", params.Get("app_id"), params.Get("app_trans_id"), params.Get("app_user"),
		params.Get("amount"), params.Get("app_time"), params.Get("embed_data"), params.Get("item"))
	params.Add("mac", hmacutil.HexStringEncode(hmacutil.SHA256, key1, data))

	// Content-Type: application/x-www-form-urlencoded
	res, err := http.PostForm("https://sb-openapi.zalopay.vn/v2/create", params)

	// parse response
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var result models.ResponseCreateOrderFromZaloPay

	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	return result
}
