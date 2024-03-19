package domain


type AlibabaAlihealthMedicalbaseDoctorSyncnewTopChannelDoctorSyncDTO struct {
    /*
        医生下架参数     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseDoctorSyncnewTopChannelDoctorSyncDTO) SetBizContent(v string) *AlibabaAlihealthMedicalbaseDoctorSyncnewTopChannelDoctorSyncDTO {
    s.BizContent = &v
    return s
}
