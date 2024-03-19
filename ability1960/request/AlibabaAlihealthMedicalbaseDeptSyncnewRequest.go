package request

import (
        "topsdk/ability1960/domain"
        "topsdk/util"
    )

type AlibabaAlihealthMedicalbaseDeptSyncnewRequest struct {
    /*
        科室状态同步     */
    TopChannelDeptSyncDTO  *domain.AlibabaAlihealthMedicalbaseDeptSyncnewTopChannelDeptSyncDto `json:"top_channel_dept_sync_d_t_o" required:"true" `
}

func (s *AlibabaAlihealthMedicalbaseDeptSyncnewRequest) SetTopChannelDeptSyncDTO(v domain.AlibabaAlihealthMedicalbaseDeptSyncnewTopChannelDeptSyncDto) *AlibabaAlihealthMedicalbaseDeptSyncnewRequest {
    s.TopChannelDeptSyncDTO = &v
    return s
}

func (req *AlibabaAlihealthMedicalbaseDeptSyncnewRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TopChannelDeptSyncDTO != nil) {
        paramMap["top_channel_dept_sync_d_t_o"] = util.ConvertStruct(*req.TopChannelDeptSyncDTO)
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalbaseDeptSyncnewRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}