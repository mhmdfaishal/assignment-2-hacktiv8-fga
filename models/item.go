package models

type Item struct {
	LineItemID  int    `json:"lineItemID" gorm:"primary_key;auto_increment"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
