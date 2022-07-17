package models

type InfoCallbackFromZaloPay struct {
	Data string `json:"data"`
	Mac  string `json:"mac"`
	Type int    `json:"type"`
}
type InfoOrderCallBack struct {
	AppId      int    `json:"app_id"`
	AppTransId string `json:"app_trans_id"`
	AppTime    int64  `json:"app_time"`
	AppUser    string `json:"app_user"`
	Amount     int    `json:"amount"`
	EmbedData  struct {
		Merchantinfo  string `json:"merchantinfo"`
		Promotioninfo string `json:"promotioninfo"`
	} `json:"embed_data"`
	Item []struct {
		Itemid       string `json:"itemid"`
		Itemname     string `json:"itemname"`
		Itemprice    int    `json:"itemprice"`
		Itemquantity int    `json:"itemquantity"`
	} `json:"item"`
	ZpTransId      int64  `json:"zp_trans_id"`
	ServerTime     int64  `json:"server_time"`
	Channel        int    `json:"channel"`
	MerchantUserId string `json:"merchant_user_id"`
	UserFeeAmount  int    `json:"user_fee_amount"`
	DiscountAmount int    `json:"discount_amount"`
}
