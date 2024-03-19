package domain


type AlibabaAlihealthMedicalRegistrationSyncnewCommonRequest4Top struct {
    /*
        业务数据json     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalRegistrationSyncnewCommonRequest4Top) SetBizContent(v string) *AlibabaAlihealthMedicalRegistrationSyncnewCommonRequest4Top {
    s.BizContent = &v
    return s
}
