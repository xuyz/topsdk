package domain


type AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel struct {
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

func (s *AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel) SetMsgInfo(v string) *AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel) SetMsgCode(v string) *AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel) SetModel(v string) *AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel {
    s.Model = &v
    return s
}
