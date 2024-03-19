package response

import (
    "topsdk/ability1960/domain"
)

type AlibabaAlihealthMedicalbaseThirdEvaluateSyncResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回信息
    */
    Result  domain.AlibabaAlihealthMedicalbaseThirdEvaluateSyncResult `json:"result,omitempty" `
}
