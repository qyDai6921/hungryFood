package define

import (
	"github.com/globalsign/mgo/bson"
)

type SearchMenusReq struct {
	Zipcode     string `form:"zip_code"`
	DeliverTime int    `form:"deliver_time"`
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
}

type MenuList struct {
	Id                bson.ObjectId `bson:"_id" json:"id"`                                    //mongoid
	Name              string        `bson:"name" json:"name"`                                 // menu name
	Rating            int           `bson:"rating" json:"rating"`                             // rating of chef
	AvgPricePerPerson float32       `bson:"avg_price_per_person" json:"avg_price_per_person"` // average price per person
	Tags              []string      `bson:"tags" json:"tags"`                                 // tags
	Description       string        `bson:"description" json:"description"`                   // descriptions
	MenuItems         []MenuItem    `bson:"menu_items" json:"menu_items"`                     // menu items
	Zipcode           string        `bson:"zipcode" json:"zip_code"`                          // zip code
}

type MenuItem struct {
	Id     bson.ObjectId `bson:"id" json:"id"`
	MenuID string        `bson:"menu_id" json:"menu_id"` // menu id
	Name   string        `bson:"name" json:"name"`       // item name
	Tags   []string      `bson:"tags" json:"tags"`       // tags
	Price  int           `bson:"price" json:"price"`     // price of each
	Count  int           `bson:"count" json:"count"`     // count
}
