package actions

import (
	"crawl_html_from_dc/main/send_receive"
	"crawl_html_from_dc/services/api"
	"crawl_html_from_dc/services/get_response_html"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Receive struct {
	Url string `json:"url"`
}

const (
	CodeReceiveGetResultSuccess    = 200 // 获取结果成功
	CodeReceiveTaskInquiring       = 201 // 任务正在查询
	CodeReceiveDcNoThisTask        = 300 // 下载中心暂时没有此任务
	CodeReceiveGetRDataMapError    = 500 // 从结果提取rData结构的map出错
	CodeReceiveReceiveFromApiError = 501 // 从下载中心接口取结果出错
	CodeReceiveStatusError         = 502 // 返回状态码有问题(不是1或2)
	CodeReceiveRequestFormError    = 400 // 访问此接口时参数格式不正确
)

func DcReceive(c *gin.Context) {
	receiveRequest := &api.SendRequest{}
	if err := c.BindJSON(receiveRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": CodeReceiveRequestFormError,
			"msg":  "请求格式不正确",
			"err":  err.Error(),
		})
		return
	}

	dcResponse, err := send_receive.AsynReceive(receiveRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": CodeReceiveReceiveFromApiError,
			"msg":  "获取结果失败",
			"err":  err.Error(),
		})
		return
	}

	if dcResponse.RData == "" || dcResponse.RData == "{}" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": CodeReceiveDcNoThisTask,
			"msg":  "获取结果失败",
			"err":  "下载中心的下载队列中暂时没有该任务，请1分钟后重试。若重试之后还是返回该结果，请重新发送该任务到下载中心",
		})
		return
	}

	rDataMap, err := get_response_html.ResponseRDataMap(dcResponse.RData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": CodeReceiveGetRDataMapError,
			"msg":  "获取结果失败",
			"err":  err.Error(),
		})
		return
	}

	for uniqueMd5 := range rDataMap {
		switch rDataMap[uniqueMd5].Status {
		case 0, 1:
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": CodeReceiveTaskInquiring,
				"msg":  "获取结果失败",
				"err":  "任务还在查询中",
			})
			return

		case 2:
			c.JSON(http.StatusOK, gin.H{
				"code": CodeReceiveGetResultSuccess,
				"msg":  "获取结果成功",
				"data": dcResponse,
			})
			return

		default:
			c.JSON(http.StatusOK, gin.H{
				"code": CodeReceiveStatusError,
				"msg":  "获取结果失败",
				"err":  "返回状态码未知",
				"data": dcResponse,
			})
		}
	}
}
