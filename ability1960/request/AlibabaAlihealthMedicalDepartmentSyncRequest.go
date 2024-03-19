package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalDepartmentSyncRequest struct {
    /*
        接口入参     */
    SaveRequest  *domain.AlibabaAlihealthMedicalDepartmentSyncCommonRequest4Top `json:"save_request,omitempty" required:"false" `
}

func (s *AlibabaAlihealthMedicalDepartmentSyncRequest) SetSaveRequest(v domain.AlibabaAlihealthMedicalDepartmentSyncCommonRequest4Top) *AlibabaAlihealthMedicalDepartmentSyncRequest {
    s.SaveRequest = &v
    return s
}

func (req *AlibabaAlihealthMedicalDepartmentSyncRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SaveRequest != nil) {
        paramMap["save_request"] = util.ConvertStruct(*req.SaveRequest)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalDepartmentSyncRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}