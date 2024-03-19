package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalHospitalSyncRequest struct {
    /*
        top保存入参     */
    SaveRequest  *domain.AlibabaAlihealthMedicalHospitalSyncCommonRequest4Top `json:"save_request,omitempty" required:"false" `
}

func (s *AlibabaAlihealthMedicalHospitalSyncRequest) SetSaveRequest(v domain.AlibabaAlihealthMedicalHospitalSyncCommonRequest4Top) *AlibabaAlihealthMedicalHospitalSyncRequest {
    s.SaveRequest = &v
    return s
}

func (req *AlibabaAlihealthMedicalHospitalSyncRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SaveRequest != nil) {
        paramMap["save_request"] = util.ConvertStruct(*req.SaveRequest)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalHospitalSyncRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}