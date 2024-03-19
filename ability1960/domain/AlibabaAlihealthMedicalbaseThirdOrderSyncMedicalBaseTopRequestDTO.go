package domain


type AlibabaAlihealthMedicalbaseThirdOrderSyncMedicalBaseTopRequestDTO struct {
    /*
        订单业务数据     */
    BizContent  *string `json:"biz_content,omitempty" `

}

func (s *AlibabaAlihealthMedicalbaseThirdOrderSyncMedicalBaseTopRequestDTO) SetBizContent(v string) *AlibabaAlihealthMedicalbaseThirdOrderSyncMedicalBaseTopRequestDTO {
    s.BizContent = &v
    return s
}
