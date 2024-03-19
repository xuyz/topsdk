package domain


type AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult struct {
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
        响应数据      */
    Data  *string `json:"data,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult) SetSuccess(v bool) *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult {
    s.Success = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult) SetErrCode(v string) *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult) SetErrMessage(v string) *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult) SetData(v string) *AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult {
    s.Data = &v
    return s
}
