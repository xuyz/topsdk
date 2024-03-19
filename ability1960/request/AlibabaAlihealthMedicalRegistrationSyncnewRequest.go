package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalRegistrationSyncnewRequest struct {
    /*
        接口入参     */
    SaveRequest  *domain.AlibabaAlihealthMedicalRegistrationSyncnewCommonRequest4Top `json:"save_request,omitempty" required:"false" `
}

func (s *AlibabaAlihealthMedicalRegistrationSyncnewRequest) SetSaveRequest(v domain.AlibabaAlihealthMedicalRegistrationSyncnewCommonRequest4Top) *AlibabaAlihealthMedicalRegistrationSyncnewRequest {
    s.SaveRequest = &v
    return s
}

func (req *AlibabaAlihealthMedicalRegistrationSyncnewRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SaveRequest != nil) {
        paramMap["save_request"] = util.ConvertStruct(*req.SaveRequest)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalRegistrationSyncnewRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}