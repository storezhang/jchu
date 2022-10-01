package core

type (
	// FtLicenseUploadReq 52号文授权协议上传请求
	FtLicenseUploadReq struct {
		// 企业名
		Name string `json:"enterpriseName"`
		// 统一代码
		Code string `json:"uniscId"`
		// 授权信息
		AuthorizedInfos string `json:"authorizedInfos"`
		// 授权码
		AuthorizedCode string `json:"authorizedOrgsCode"`
		// 平台
		PlatformId string `json:"platformId"`
		// 数量
		Count string `json:"count"`
		// 融资阶段
		LoanStage string `json:"loanStage"`
		// 文件类型
		FileType string `json:"licenseFileType"`
		// 电 子签章供应商
		CaSupplier string `json:"caSupplier"`
		// 验证地址
		ValidateUrl string `json:"validateUrl"`
		// 开始时间
		AuthorizedStartTime string `json:"authorizedStartTime"`
		// 结束时间
		AuthorizedEndTime string `json:"authorizedEndTime"`
		// 授权文件摘要
		HashCode string `json:"hashCode"`
	}

	// FtLicenseUploadRsp 52号文授权协议上传响应
	FtLicenseUploadRsp struct {
		// 授权码
		LicenseId string `json:"licenseId"`
	}
)
