package model

import (
	"github.com/globalsign/mgo/bson"
	"workspace/ginweb/define"
	"workspace/ginweb/pkg/mongodb"
)

func GetMenus(req *define.SearchMenusReq) (resp []define.MenuList,total int, err error) {
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	where := bson.M{}
	if req.DeliverTime > 0 {
		where["opentime"] = bson.M{"$lt":req.DeliverTime}
		where["closetime"] = bson.M{"$gt":req.DeliverTime}
	}
	if req.Zipcode != "" {
		where["zipcode"] = req.Zipcode
	}

	total, _ = mongodb.FindAllWithPage(define.TABLE_MENU, where, nil, req.PageSize, req.Page, &resp)

	return
}
