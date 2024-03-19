package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalbaseThirdOrderSyncRequest struct {
    /*
        请求参数     */
    OrderRequest  *domain.AlibabaAlihealthMedicalbaseThirdOrderSyncMedicalBaseTopRequestDTO `json:"order_request" required:"true" `
}

func (s *AlibabaAlihealthMedicalbaseThirdOrderSyncRequest) SetOrderRequest(v domain.AlibabaAlihealthMedicalbaseThirdOrderSyncMedicalBaseTopRequestDTO) *AlibabaAlihealthMedicalbaseThirdOrderSyncRequest {
    s.OrderRequest = &v
    return s
}

func (req *AlibabaAlihealthMedicalbaseThirdOrderSyncRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderRequest != nil) {
        paramMap["order_request"] = util.ConvertStruct(*req.OrderRequest)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalbaseThirdOrderSyncRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}