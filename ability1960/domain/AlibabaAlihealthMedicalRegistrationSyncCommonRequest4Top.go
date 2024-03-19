package domain


type AlibabaAlihealthMedicalRegistrationSyncCommonRequest4Top struct {
    /*
        业务数据json     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalRegistrationSyncCommonRequest4Top) SetBizContent(v string) *AlibabaAlihealthMedicalRegistrationSyncCommonRequest4Top {
    s.BizContent = &v
    return s
}
