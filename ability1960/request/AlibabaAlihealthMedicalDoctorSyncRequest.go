package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalDoctorSyncRequest struct {
    /*
        接口入参     */
    SaveRequest  *domain.AlibabaAlihealthMedicalDoctorSyncCommonRequest4Top `json:"save_request,omitempty" required:"false" `
}

func (s *AlibabaAlihealthMedicalDoctorSyncRequest) SetSaveRequest(v domain.AlibabaAlihealthMedicalDoctorSyncCommonRequest4Top) *AlibabaAlihealthMedicalDoctorSyncRequest {
    s.SaveRequest = &v
    return s
}

func (req *AlibabaAlihealthMedicalDoctorSyncRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SaveRequest != nil) {
        paramMap["save_request"] = util.ConvertStruct(*req.SaveRequest)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalDoctorSyncRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}