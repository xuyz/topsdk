package response

import (
    "topsdk/ability1960/domain"
)

type AlibabaAlihealthMedicalbaseThirdOrderSyncResponse struct {

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
    Result  domain.AlibabaAlihealthMedicalbaseThirdOrderSyncResult `json:"result,omitempty" `
}
