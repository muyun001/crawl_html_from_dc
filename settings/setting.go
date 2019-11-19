package settings

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var CrawlHtmlFromDc bool

var DcApi string
var DcApiInner string
var DcApiOuter string
var DcUserId string
var DcIsInner bool

func init() {
	checkEnv()
	LoadSetting()
}

func checkEnv() {
	_ = godotenv.Load()
	needChecks := []string{
		"CRAWL_HTML_FROM_DC",
		"DC_API_OUTER", "DC_API_INNER", "DC_IS_INNER", "DC_USER_ID",
	}

	for _, envKey := range needChecks {
		if os.Getenv(envKey) == "" {
			log.Fatalf("env %s missed", envKey)
		}
	}
}

func LoadSetting() {
	var err error
	CrawlHtmlFromDc, err = strconv.ParseBool(os.Getenv("CRAWL_HTML_FROM_DC"))
	if err != nil {
		log.Fatalf("CRAWL_HTML_FROM_DC parse bool error")
	}
	DcUserId = os.Getenv("DC_USER_ID")

	DcApiOuter = os.Getenv("DC_API_OUTER")
	DcApiInner = os.Getenv("DC_API_INNER")

	//DcIsInner, err = strconv.ParseBool(os.Getenv("DC_IS_INNER"))
	//if err != nil {
	//	log.Fatalf("DcIsInner parse bool error")
	//}

	// 读取exe参数
	DcIsInner := flag.Bool("inner", true, "dc work in inner")
	flag.Parse()

	if *DcIsInner {
		DcApi = DcApiInner
	} else {
		DcApi = DcApiOuter
	}
}
