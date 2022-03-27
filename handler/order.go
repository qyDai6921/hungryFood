package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workspace/ginweb/define"
	"workspace/ginweb/model"
	"workspace/ginweb/pkg/logs"
)

func Order(c *gin.Context) {
	var err error
	rsp := define.Result{
		Code: 0,
		Msg:  "success",
	}
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()

	req := &define.OrderReq{}
	if err = c.ShouldBindJSON(&req); err != nil {
		logs.Debugf("err:%+v", err)
		return
	}
	logs.Debugf("req:%+v", req)

	err = model.Order(req)
	if err != nil {
		rsp.Code = 500
		rsp.Msg = err.Error()
	}
}
