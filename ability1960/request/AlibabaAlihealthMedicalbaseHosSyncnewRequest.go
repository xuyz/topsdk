package request


type AlibabaAlihealthMedicalbaseHosSyncnewRequest struct {
    /*
        212     */
    BizContent  *string `json:"biz_content" required:"true" `
}

func (s *AlibabaAlihealthMedicalbaseHosSyncnewRequest) SetBizContent(v string) *AlibabaAlihealthMedicalbaseHosSyncnewRequest {
    s.BizContent = &v
    return s
}

func (req *AlibabaAlihealthMedicalbaseHosSyncnewRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizContent != nil) {
        paramMap["biz_content"] = *req.BizContent
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalbaseHosSyncnewRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}