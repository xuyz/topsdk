package domain


type AlibabaAlihealthMedicalHospitalSyncCommonRequest4Top struct {
    /*
        hosNo:渠道医院ID,hosName:医院名称,status:医院状态     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalHospitalSyncCommonRequest4Top) SetBizContent(v string) *AlibabaAlihealthMedicalHospitalSyncCommonRequest4Top {
    s.BizContent = &v
    return s
}
