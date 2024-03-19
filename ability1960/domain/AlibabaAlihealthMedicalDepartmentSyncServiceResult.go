package domain


type AlibabaAlihealthMedicalDepartmentSyncServiceResult struct {
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

func (s *AlibabaAlihealthMedicalDepartmentSyncServiceResult) SetErrMessage(v string) *AlibabaAlihealthMedicalDepartmentSyncServiceResult {
    s.ErrMessage = &v
    return s
}
func (s *AlibabaAlihealthMedicalDepartmentSyncServiceResult) SetData(v int64) *AlibabaAlihealthMedicalDepartmentSyncServiceResult {
    s.Data = &v
    return s
}
func (s *AlibabaAlihealthMedicalDepartmentSyncServiceResult) SetErrCode(v string) *AlibabaAlihealthMedicalDepartmentSyncServiceResult {
    s.ErrCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalDepartmentSyncServiceResult) SetSuccess(v bool) *AlibabaAlihealthMedicalDepartmentSyncServiceResult {
    s.Success = &v
    return s
}
