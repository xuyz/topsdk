package domain


type AlibabaAlihealthMedicalDoctorSyncServiceResult struct {
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

func (s *AlibabaAlihealthMedicalDoctorSyncServiceResult) SetErrMessage(v string) *AlibabaAlihealthMedicalDoctorSyncServiceResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthMedicalDoctorSyncServiceResult) SetData(v int64) *AlibabaAlihealthMedicalDoctorSyncServiceResult {
    s.Data = &v
    return s
}
func (s *AlibabaAlihealthMedicalDoctorSyncServiceResult) SetErrCode(v string) *AlibabaAlihealthMedicalDoctorSyncServiceResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalDoctorSyncServiceResult) SetSuccess(v bool) *AlibabaAlihealthMedicalDoctorSyncServiceResult {
    s.Success = &v
    return s
}
