package domain


type AlibabaAlihealthMedicalbaseDoctorSyncnewTopResultModel struct {
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

func (s *AlibabaAlihealthMedicalbaseDoctorSyncnewTopResultModel) SetMsgInfo(v string) *AlibabaAlihealthMedicalbaseDoctorSyncnewTopResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseDoctorSyncnewTopResultModel) SetMsgCode(v string) *AlibabaAlihealthMedicalbaseDoctorSyncnewTopResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseDoctorSyncnewTopResultModel) SetModel(v string) *AlibabaAlihealthMedicalbaseDoctorSyncnewTopResultModel {
    s.Model = &v
    return s
}
