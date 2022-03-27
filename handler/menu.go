package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workspace/ginweb/define"
	"workspace/ginweb/model"
	"workspace/ginweb/pkg/logs"
)

func SearchMenus(c *gin.Context) {
	var err error
	rsp := define.ResultWithPage{
		Code: 0,
		Msg:  "success",
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	req := &define.SearchMenusReq{}
	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}
	logs.Debugf("req:%+v", req)

	if req.DeliverTime == 0 {
		rsp.Code = 500
		rsp.Msg = "please enter your deliver time"
		return
	}
	if req.Zipcode == ""{
		rsp.Code = 500
		rsp.Msg = "please enter your zipcode"
		return
	}

	if req.Page == 0 {
		rsp.Page.Page = 1
	}
	if req.PageSize == 0 {
		rsp.Page.PageSize = 10
	}
	resp, total, err := model.GetMenus(req)
	if err != nil {
		rsp.Code = 500
	} else {
		rsp.Data = resp
		rsp.Page.Total = total
	}
}

