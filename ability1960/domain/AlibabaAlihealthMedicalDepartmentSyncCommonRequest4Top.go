package domain


type AlibabaAlihealthMedicalDepartmentSyncCommonRequest4Top struct {
    /*
        hosNo:渠道医院ID,deptNo:渠道科室ID,deptName:科室名称,status:状态     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalDepartmentSyncCommonRequest4Top) SetBizContent(v string) *AlibabaAlihealthMedicalDepartmentSyncCommonRequest4Top {
    s.BizContent = &v
    return s
}
