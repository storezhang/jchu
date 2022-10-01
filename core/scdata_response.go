package core

// SCDataResponse 省大数据中心响应
type SCDataResponse struct {
	// 授权码
	Key string `json:"key"`
	// 签名数据
	SignatureData string `json:"signatureData"`
	// 数据
	Data string `json:"data"`
}

func (r *SCDataResponse) Check() bool {
	return `` != r.Key && `` != r.SignatureData && `` != r.Data
}
