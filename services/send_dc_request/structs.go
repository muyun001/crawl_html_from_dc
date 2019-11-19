package send_dc_request

type DcApiResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	RData  string `json:"rdata"`
}
