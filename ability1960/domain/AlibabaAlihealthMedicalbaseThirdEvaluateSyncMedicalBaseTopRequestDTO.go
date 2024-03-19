package domain


type AlibabaAlihealthMedicalbaseThirdEvaluateSyncMedicalBaseTopRequestDTO struct {
    /*
        评论业务数据     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseThirdEvaluateSyncMedicalBaseTopRequestDTO) SetBizContent(v string) *AlibabaAlihealthMedicalbaseThirdEvaluateSyncMedicalBaseTopRequestDTO {
    s.BizContent = &v
    return s
}
