package model

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"math/rand"
	"time"
	"workspace/ginweb/define"
	"workspace/ginweb/pkg/mongodb"
)

func Order(req *define.OrderReq) (err error) {
	// 1.	chief_id cannot be null;
	if req.ChefId == "" {
		err = errors.New("Invalid chef_id")
		return
	}

	// 2.	menu_items cannot be less than or equal to 0;
	if len(req.MenuItems) <= 0 {
		err = errors.New("Invalid menus")
		return
	}

	// 3.	price cannot be less than or equal to 0;
	if req.Price <= 0 {
		err = errors.New("Invalid price")
		return
	}

	// 4.	deliver_time cannot be less than the current time;
	timeNow := int(time.Now().Unix())
	if req.DeliverTime < timeNow {
		err = errors.New("Your delivery time cannot be less than the current time")
		return
	}

	// 5.	address cannot be null;
	if req.Address == "" {
		err = errors.New("Invalid address")
		return
	}

	// 6.	zip_code should be correct (here is just 10001 and 22030, you can add more in collection hungry_list);
	var chefMenu define.MenuList
	_ = mongodb.FindOne(define.TABLE_MENU, bson.M{"_id": bson.ObjectIdHex(req.ChefId)}, nil, &chefMenu)
	//verify
	if chefMenu.Zipcode != req.Zipcode {
		err = errors.New("Zip code does not match")
		return
	}
	var menus = make(map[string]define.MenuItem)
	for _, v := range chefMenu.MenuItems {
		menus[v.Id.Hex()] = v
	}

	var price int
	for k, v := range req.MenuItems {
		if val, ok := menus[v.MenuID]; ok {
			price += val.Price * v.Count
			req.MenuItems[k].Id = bson.ObjectId(val.Id)
		}
	}

	if price != req.Price {
		err = errors.New("Invalid price")
		return
	}

	err = mongodb.Insert(define.TABLE_ORDER, bson.M{
		"order_no":     fmt.Sprintf("%v%v", time.Now().Nanosecond(), rand.Int31n(9999)),
		"username":     req.Username,
		"phone":        req.Phone,
		"remark":       req.Remark,
		"address":      req.Address,
		"zipcode":      req.Zipcode,
		"deliver_time": req.DeliverTime,
		"price":        price,
		"tip":          req.Tip,
		"tax":          req.Tax,
		"total":        req.Total,
		"menus":        req.MenuItems,
	})
	if err != nil {
		err = errors.New("order fail")
		return
	}

	return
}
