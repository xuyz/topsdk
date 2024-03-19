package response

import (
    "topsdk/ability1960/domain"
)

type AlibabaAlihealthMedicalbaseHosSyncnewResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        和三方交互最外层model对象
    */
    Result  domain.AlibabaAlihealthMedicalbaseHosSyncnewTopResultModel `json:"result,omitempty" `
}
