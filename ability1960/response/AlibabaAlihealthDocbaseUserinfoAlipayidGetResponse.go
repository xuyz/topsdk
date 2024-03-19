package response

import (
    "topsdk/ability1960/domain"
)

type AlibabaAlihealthDocbaseUserinfoAlipayidGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        alinkappserver系统返回的通用结果类
    */
    Result  domain.AlibabaAlihealthDocbaseUserinfoAlipayidGetServiceResult `json:"result,omitempty" `
}
