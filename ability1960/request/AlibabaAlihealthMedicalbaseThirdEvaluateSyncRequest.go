package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest struct {
    /*
        请求参数     */
    EvaluateRequest  *domain.AlibabaAlihealthMedicalbaseThirdEvaluateSyncMedicalBaseTopRequestDTO `json:"evaluate_request" required:"true" `
}

func (s *AlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest) SetEvaluateRequest(v domain.AlibabaAlihealthMedicalbaseThirdEvaluateSyncMedicalBaseTopRequestDTO) *AlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest {
    s.EvaluateRequest = &v
    return s
}

func (req *AlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.EvaluateRequest != nil) {
        paramMap["evaluate_request"] = util.ConvertStruct(*req.EvaluateRequest)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}