package actions

import (
	"crawl_html_from_dc/main/send_receive"
	"crawl_html_from_dc/services/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSendRequestSuccess   = 200 // 发送请求成功
	CodeSendRequestRepeat    = 201 // 重复发送请求
	CodeSendRequestToDcError = 500 // 发送请求到下载中心失败
	CodeSendRequestFormError = 400 // 访问此接口时参数格式不正确
)

func DcSend(c *gin.Context) {
	sendRequest := &api.SendRequest{}
	if err := c.BindJSON(sendRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeSendRequestFormError,
			"msg":  "请求格式不正确",
			"err":  err.Error(),
		})
		return
	}

	if err := send_receive.Send(sendRequest); err != nil {
		if err.Error() == "重复插入数据" {
			c.JSON(http.StatusOK, gin.H{
				"code": CodeSendRequestRepeat,
				"msg":  "重复插入数据",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code": CodeSendRequestToDcError,
			"msg":  "请求发送失败",
			"err":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSendRequestSuccess,
		"msg":  "请求发送成功",
	})
}
