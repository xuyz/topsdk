package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalRegistrationSyncRequest struct {
    /*
        接口入参     */
    SaveRequest  *domain.AlibabaAlihealthMedicalRegistrationSyncCommonRequest4Top `json:"save_request,omitempty" required:"false" `
}

func (s *AlibabaAlihealthMedicalRegistrationSyncRequest) SetSaveRequest(v domain.AlibabaAlihealthMedicalRegistrationSyncCommonRequest4Top) *AlibabaAlihealthMedicalRegistrationSyncRequest {
    s.SaveRequest = &v
    return s
}

func (req *AlibabaAlihealthMedicalRegistrationSyncRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SaveRequest != nil) {
        paramMap["save_request"] = util.ConvertStruct(*req.SaveRequest)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalRegistrationSyncRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}