package domain


type AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult struct {
    /*
        errMessage     */
    ErrMessage  *string `json:"err_message,omitempty" `

    /*
        返回数据对象     */
    Data  *string `json:"data,omitempty" `

    /*
        errCode     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult) SetErrMessage(v string) *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult) SetData(v string) *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult {
    s.Data = &v
    return s
}
func (s *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult) SetErrCode(v string) *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult) SetSuccess(v bool) *AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult {
    s.Success = &v
    return s
}
