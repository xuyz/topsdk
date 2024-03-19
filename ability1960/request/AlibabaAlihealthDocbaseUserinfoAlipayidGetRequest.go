package request


type AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest struct {
    /*
        阿里健康ID     */
    AlihealthUserId  *string `json:"alihealth_user_id" required:"true" `
    /*
        渠道alipay taobao uc gaode     */
    AppChannel  *string `json:"app_channel" required:"true" `
}

func (s *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAlihealthUserId(v string) *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest {
    s.AlihealthUserId = &v
    return s
}
func (s *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) SetAppChannel(v string) *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest {
    s.AppChannel = &v
    return s
}

func (req *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AlihealthUserId != nil) {
        paramMap["alihealth_user_id"] = *req.AlihealthUserId
    }
    if(req.AppChannel != nil) {
        paramMap["app_channel"] = *req.AppChannel
    }
    return paramMap
}

func (req *AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}