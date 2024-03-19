package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalbaseDoctorSyncnewRequest struct {
    /*
        医生下架     */
    TopChannelDoctorSyncDTO  *domain.AlibabaAlihealthMedicalbaseDoctorSyncnewTopChannelDoctorSyncDTO `json:"top_channel_doctor_sync_d_t_o" required:"true" `
}

func (s *AlibabaAlihealthMedicalbaseDoctorSyncnewRequest) SetTopChannelDoctorSyncDTO(v domain.AlibabaAlihealthMedicalbaseDoctorSyncnewTopChannelDoctorSyncDTO) *AlibabaAlihealthMedicalbaseDoctorSyncnewRequest {
    s.TopChannelDoctorSyncDTO = &v
    return s
}

func (req *AlibabaAlihealthMedicalbaseDoctorSyncnewRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TopChannelDoctorSyncDTO != nil) {
        paramMap["top_channel_doctor_sync_d_t_o"] = util.ConvertStruct(*req.TopChannelDoctorSyncDTO)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalbaseDoctorSyncnewRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}