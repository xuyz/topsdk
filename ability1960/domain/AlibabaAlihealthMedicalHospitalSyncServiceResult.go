package domain


type AlibabaAlihealthMedicalHospitalSyncServiceResult struct {
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

func (s *AlibabaAlihealthMedicalHospitalSyncServiceResult) SetErrMessage(v string) *AlibabaAlihealthMedicalHospitalSyncServiceResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthMedicalHospitalSyncServiceResult) SetData(v int64) *AlibabaAlihealthMedicalHospitalSyncServiceResult {
    s.Data = &v
    return s
}
func (s *AlibabaAlihealthMedicalHospitalSyncServiceResult) SetErrCode(v string) *AlibabaAlihealthMedicalHospitalSyncServiceResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalHospitalSyncServiceResult) SetSuccess(v bool) *AlibabaAlihealthMedicalHospitalSyncServiceResult {
    s.Success = &v
    return s
}
