package domain


type AlibabaAlihealthMedicalbaseDeptSyncnewTopChannelDeptSyncDto struct {
    /*
        医院ID+科室ID+状态     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseDeptSyncnewTopChannelDeptSyncDto) SetBizContent(v string) *AlibabaAlihealthMedicalbaseDeptSyncnewTopChannelDeptSyncDto {
    s.BizContent = &v
    return s
}
