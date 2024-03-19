package domain


type AlibabaAlihealthMedicalbaseHospitalSyncTopResultModel struct {
    /*
        操作说明     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        操作码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        isvId     */
    Model  *string `json:"model,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseHospitalSyncTopResultModel) SetMsgInfo(v string) *AlibabaAlihealthMedicalbaseHospitalSyncTopResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncTopResultModel) SetMsgCode(v string) *AlibabaAlihealthMedicalbaseHospitalSyncTopResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncTopResultModel) SetModel(v string) *AlibabaAlihealthMedicalbaseHospitalSyncTopResultModel {
    s.Model = &v
    return s
}
