package models

type StockRating struct{
	ID string `json:"id`
    Ticker string `json:"ticker`
	Target_from string `json:"target_from`
	Target_to string `json:"target_to`
	Company string `json:"company`
	Action string `json:"action`
	Brokerage string `json:"brokerage`
	Rating_from string `json:"rating_from`
	Rating_to string `json:"rating_to`
	time string `json:"time`
}