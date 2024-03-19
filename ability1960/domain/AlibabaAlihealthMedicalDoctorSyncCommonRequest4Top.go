package domain


type AlibabaAlihealthMedicalDoctorSyncCommonRequest4Top struct {
    /*
        hosNo:渠道医院ID,deptNo:渠道科室ID,doctorNo:渠道医生ID,doctorName:医生姓名,status:状态     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalDoctorSyncCommonRequest4Top) SetBizContent(v string) *AlibabaAlihealthMedicalDoctorSyncCommonRequest4Top {
    s.BizContent = &v
    return s
}
