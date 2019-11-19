package main

import (
	"crawl_html_from_dc/main/send_receive"
	"crawl_html_from_dc/services/api"
	"crawl_html_from_dc/services/get_response_html"
	"fmt"
	"strings"
)

var basicUrl = "http://www.baidu.com/s?wd=%s"
var keywords = []string{
	"白银橡胶地板", "定西橡胶地板", "临夏亚麻地板", "罗拉橡胶地板", "商洛橡胶地板革", "临夏亚麻地板厚度", "临夏亚麻地板清洁", "固原罗拉橡胶地板", "临夏亚麻地板焊线", "临夏亚麻地板规格", "重庆管家婆财务软件价格", "重庆管家婆财务软件手机版", "重庆管家婆财务软件电话", "重庆管家婆财务软件售后服务电话", "重庆管家婆手机版端口", "重庆管家婆手机版t9", "财贸双全凭证打印模板", "重庆管家婆手机版赠品", "管家婆工贸反记账步骤", "重庆管家婆财务软件a8", "婚车租赁店名", "奔驰gla婚庆租车", "亳州婚车租赁店名", "合肥婚车租赁店名", "宣城悍马婚车租赁", "芜湖婚车租赁合同", "宿州婚车租赁名字", "宿州个人婚庆租车", "芜湖婚车租赁利润", "蚌埠本田婚车租赁", "重庆纸箱加工设备", "重庆气泡膜机", "重庆万州区气泡膜打包", "重庆合川区瓦楞纸箱", "重庆合川区纸箱打包机", "重庆万州区打包气泡膜", "重庆万州区PE气泡膜", "重庆万州区定做气泡膜", "重庆黔江区纸箱打包机", "重庆合川区纸箱瓦楞纸", "道路护栏工厂", "高端道路护栏", "知名道路护栏", "便宜桥梁护栏厂", "云南道路护栏", "四川桥梁护栏厂", "云南道路护栏厂", "知名道路护栏厂", "知名桥梁护栏定制", "知名桥梁护栏公司",
}

func main() {
	sendRequests()

	asynReceiveHtml()
	//syncReceiveHtml()
}

func sendRequests() {
	for i, keyword := range keywords {
		sendRequest := &api.SendRequest{Url: fmt.Sprintf(basicUrl, keyword)}
		sendRequest.Headers.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36"
		sendRequest.Headers.Cookie = "BAIDUID=5028A9DDCD923E023F0FEEC0B22370EB:FG=1; BIDUPSID=5028A9DDCD923E023F0FEEC0B22370EB; PSTM=1568778486; BD_UPN=12314753; MSA_WH=1920_937; H_WISE_SIDS=135669_137150_137735_133103_136909_136651_136293_134725_113879_128065_136294_134982_136436_120195_137456_136659_137716_136366_132911_136455_135847_131247_137750_132378_131517_118881_118864_118849_118832_118788_136687_107319_132782_136799_136429_136091_133351_137222_136862_129649_136196_133847_132551_134047_131423_135232_136164_136753_110085_127969_131951_136612_137253_127416_136636_137097_137207_134349_132467_137619_137449_136987_100457; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; MCITY=-224%3A; delPer=0; BD_CK_SAM=1; BD_HOME=0; BDRCVFR[PowvNBqg-GC]=mbxnW11j9Dfmh7GuZR8mvqV; H_PS_PSSID=1461_21104_30073_29567_29221_26350_22159; PSINO=5; shifen[135882633018_59048]=1573543069; BCLID=11134997986530099843; BDSFRCVID=B7_OJeC627-0XRTw1RSOuQfJBejzquvTH6ao_9S7Ug_hlypMK8FoEG0PHx8g0Kub2p1TogKKL2OTHmuF_2uxOjjg8UtVJeC6EG0Ptf8g0M5; H_BDCLCKID_SF=tJAHoKLaJC83H43TqRrEKtFD-frQ5C62aKDsQpQ7BhcqEIL406Ah2xku5NLJQJotLCcZ5tjafR7kMxbSj4QohtAJ5Goi2t4OW6kJ2hoX3p5nhMJS257JDMP0-xQEXqQy523i2IovQpnVfqQ3DRoWXPIqbN7P-p5Z5mAqKl0MLPbtbb0xXj_0D6J3eaLHJ58s56bL3RTsH4jaKROvhDTjh6PYjnn9BtQmJJufsCJ9LfbbhfobXnoGbxIYbf6EbRQqQg-q3R77fx8bSJ33M-vBKMuUe-jy0x-jLgbOVn0MW-5Dh4tl3-nJyUPTD4nnBPrt3H8HL4nv2JcJbM5m3x6qLTKkQN3T-PKO5bRu_CcJ-J8XhDL4D5JP; H_PS_645EC=26d96KiTpOumgRlNnVQgkrrdhp69Dbb8Zrv8gqgnAMoQmglfLdlxs8xPebw; BDSVRTM=79; COOKIE_SESSION=3694_4_8_9_7_17_0_2_8_5_1_4_3461_0_0_0_1573540281_1573543069_1573549657%7C9%236730_3_1573543069%7C2; WWW_ST=1573551101383"

		err := send_receive.Send(sendRequest)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i, "/", len(keywords), keyword, "请求发送成功")
		}
	}
}

func asynReceiveHtml() {
	//var urls = []string{}
	//var unCompleteUrls = []string{}

	for _, keyword := range keywords {
		url := fmt.Sprintf(basicUrl, keyword)
		dcResponse, err := send_receive.AsynReceive(url) // 异步获取结果
		if err != nil {
			fmt.Println(err)
			continue
		}

		if dcResponse.Status != get_response_html.DcCrawlSucceed {
			fmt.Println(fmt.Sprintf(keyword, "下载中心抓取数据失败，错误信息：%s", dcResponse.Msg))
			continue
		}

		if !strings.Contains(dcResponse.RData, "result") {
			fmt.Println(keyword, "此返回结果已被取走，或者下载中心未抓到有效数据")
			continue
		}

		html, err := getHtmlFromResponse(dcResponse)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(html)
	}

	//unCompleteUrls = urls
	//for _, url := range unCompleteUrls {
	//	dcResponse, err := crawl_html.AsynReceive(url) // 异步获取结果
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	if dcResponse.Status != get_response_html.DcCrawlSucceed {
	//		fmt.Println(fmt.Sprintf("下载中心抓取数据失败，错误信息：%s", dcResponse.Msg))
	//		continue
	//	}
	//
	//	html, err := getHtmlFromResponse(dcResponse)
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	fmt.Println(html)
	//
	//}
}

func getHtmlFromResponse(dcResponse *get_response_html.DcApiResponse) (string, error) {
	rDataMap, err := get_response_html.ResponseRDataMap(dcResponse.RData)
	if err != nil {
		return "", err
	}

	html, err := get_response_html.DecodeHtml(rDataMap)
	if err != nil {
		return "", err
	}

	return html, nil
}

func syncReceiveHtml() {
	for _, keyword := range keywords {
		url := fmt.Sprintf(basicUrl, keyword)
		html, err := send_receive.SyncReceive(url) // 同步获取结果
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(html)
	}
}
