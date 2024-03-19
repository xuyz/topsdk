package domain


type AlibabaAlihealthMedicalbaseThirdOrderSyncResult struct {
    /*
        执行是否成功 true成功 false 失败     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误编码     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        错误信息     */
    ErrMessage  *string `json:"err_message,omitempty" `

    /*
        响应数据     */
    Data  *string `json:"data,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseThirdOrderSyncResult) SetSuccess(v bool) *AlibabaAlihealthMedicalbaseThirdOrderSyncResult {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseThirdOrderSyncResult) SetErrCode(v string) *AlibabaAlihealthMedicalbaseThirdOrderSyncResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseThirdOrderSyncResult) SetErrMessage(v string) *AlibabaAlihealthMedicalbaseThirdOrderSyncResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseThirdOrderSyncResult) SetData(v string) *AlibabaAlihealthMedicalbaseThirdOrderSyncResult {
    s.Data = &v
    return s
}
