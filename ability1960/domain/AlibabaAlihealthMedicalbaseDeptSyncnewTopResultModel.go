package domain


type AlibabaAlihealthMedicalbaseDeptSyncnewTopResultModel struct {
    /*
        操作说明     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        操作码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        isvId     */
    Model  *bool `json:"model,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseDeptSyncnewTopResultModel) SetMsgInfo(v string) *AlibabaAlihealthMedicalbaseDeptSyncnewTopResultModel {
    s.MsgInfo = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseDeptSyncnewTopResultModel) SetMsgCode(v string) *AlibabaAlihealthMedicalbaseDeptSyncnewTopResultModel {
    s.MsgCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseDeptSyncnewTopResultModel) SetModel(v bool) *AlibabaAlihealthMedicalbaseDeptSyncnewTopResultModel {
    s.Model = &v
    return s
}
