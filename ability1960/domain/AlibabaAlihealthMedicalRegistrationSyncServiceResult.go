package domain


type AlibabaAlihealthMedicalRegistrationSyncServiceResult struct {
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

func (s *AlibabaAlihealthMedicalRegistrationSyncServiceResult) SetErrMessage(v string) *AlibabaAlihealthMedicalRegistrationSyncServiceResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthMedicalRegistrationSyncServiceResult) SetData(v int64) *AlibabaAlihealthMedicalRegistrationSyncServiceResult {
    s.Data = &v
    return s
}
func (s *AlibabaAlihealthMedicalRegistrationSyncServiceResult) SetErrCode(v string) *AlibabaAlihealthMedicalRegistrationSyncServiceResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalRegistrationSyncServiceResult) SetSuccess(v bool) *AlibabaAlihealthMedicalRegistrationSyncServiceResult {
    s.Success = &v
    return s
}
