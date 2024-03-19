package domain


type AlibabaAlihealthMedicalRegistrationSyncnewServiceResult struct {
    /*
        errMessage     */
    ErrMessage  *string `json:"err_message,omitempty" `

    /*
        返回数据对象     */
    Data  *int64 `json:"data,omitempty" `

    /*
        errCode     */
    ErrCode  *string `json:"err_code,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult) SetErrMessage(v string) *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult) SetData(v int64) *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult {
    s.Data = &v
    return s
}
func (s *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult) SetErrCode(v string) *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult) SetSuccess(v bool) *AlibabaAlihealthMedicalRegistrationSyncnewServiceResult {
    s.Success = &v
    return s
}
