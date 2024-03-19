package ability1960

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability1960/request"
    "topsdk/ability1960/response"
    "topsdk/util"
)

type Ability1960 struct {
    Client *topsdk.TopClient
}

func NewAbility1960(client *topsdk.TopClient) *Ability1960{
    return &Ability1960{client}
}

/*
    三方评论信息同步
*/
func (ability *Ability1960) AlibabaAlihealthMedicalbaseThirdEvaluateSync(req *request.AlibabaAlihealthMedicalbaseThirdEvaluateSyncRequest) (*response.AlibabaAlihealthMedicalbaseThirdEvaluateSyncResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medicalbase.third.evaluate.sync",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalbaseThirdEvaluateSyncResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalbaseThirdEvaluateSync error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    阿里健康新挂号数据回传
*/
func (ability *Ability1960) AlibabaAlihealthMedicalRegistrationSyncnew(req *request.AlibabaAlihealthMedicalRegistrationSyncnewRequest) (*response.AlibabaAlihealthMedicalRegistrationSyncnewResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medical.registration.syncnew",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalRegistrationSyncnewResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalRegistrationSyncnew error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    阿里健康支付宝挂号记录回传接口
*/
func (ability *Ability1960) AlibabaAlihealthMedicalRegistrationSync(req *request.AlibabaAlihealthMedicalRegistrationSyncRequest) (*response.AlibabaAlihealthMedicalRegistrationSyncResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medical.registration.sync",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalRegistrationSyncResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalRegistrationSync error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    第三方订单同步
*/
func (ability *Ability1960) AlibabaAlihealthMedicalbaseThirdOrderSync(req *request.AlibabaAlihealthMedicalbaseThirdOrderSyncRequest) (*response.AlibabaAlihealthMedicalbaseThirdOrderSyncResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medicalbase.third.order.sync",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalbaseThirdOrderSyncResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalbaseThirdOrderSync error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    阿里将康支付宝挂号数据医院回传接口
*/
func (ability *Ability1960) AlibabaAlihealthMedicalHospitalSync(req *request.AlibabaAlihealthMedicalHospitalSyncRequest) (*response.AlibabaAlihealthMedicalHospitalSyncResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medical.hospital.sync",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalHospitalSyncResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalHospitalSync error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    阿里健康预约挂号科室同步接口
*/
func (ability *Ability1960) AlibabaAlihealthMedicalDepartmentSync(req *request.AlibabaAlihealthMedicalDepartmentSyncRequest) (*response.AlibabaAlihealthMedicalDepartmentSyncResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medical.department.sync",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalDepartmentSyncResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalDepartmentSync error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    阿里健康预约挂号医生同步接口
*/
func (ability *Ability1960) AlibabaAlihealthMedicalDoctorSync(req *request.AlibabaAlihealthMedicalDoctorSyncRequest) (*response.AlibabaAlihealthMedicalDoctorSyncResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medical.doctor.sync",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalDoctorSyncResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalDoctorSync error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    直连医院上传接口
*/
func (ability *Ability1960) AlibabaAlihealthMedicalbaseHosSyncnew(req *request.AlibabaAlihealthMedicalbaseHosSyncnewRequest) (*response.AlibabaAlihealthMedicalbaseHosSyncnewResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medicalbase.hos.syncnew",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalbaseHosSyncnewResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalbaseHosSyncnew error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    互联网医院批量导入接口
*/
func (ability *Ability1960) AlibabaAlihealthMedicalbaseHospitalSync(req *request.AlibabaAlihealthMedicalbaseHospitalSyncRequest) (*response.AlibabaAlihealthMedicalbaseHospitalSyncResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medicalbase.hospital.sync",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalbaseHospitalSyncResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalbaseHospitalSync error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    直连科室上传
*/
func (ability *Ability1960) AlibabaAlihealthMedicalbaseDeptSyncnew(req *request.AlibabaAlihealthMedicalbaseDeptSyncnewRequest) (*response.AlibabaAlihealthMedicalbaseDeptSyncnewResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medicalbase.dept.syncnew",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalbaseDeptSyncnewResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalbaseDeptSyncnew error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    直连医生上传
*/
func (ability *Ability1960) AlibabaAlihealthMedicalbaseDoctorSyncnew(req *request.AlibabaAlihealthMedicalbaseDoctorSyncnewRequest) (*response.AlibabaAlihealthMedicalbaseDoctorSyncnewResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.medicalbase.doctor.syncnew",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthMedicalbaseDoctorSyncnewResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthMedicalbaseDoctorSyncnew error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据健康ID获取支付宝ID
*/
func (ability *Ability1960) AlibabaAlihealthDocbaseUserinfoAlipayidGet(req *request.AlibabaAlihealthDocbaseUserinfoAlipayidGetRequest) (*response.AlibabaAlihealthDocbaseUserinfoAlipayidGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability1960 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.alihealth.docbase.userinfo.alipayid.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAlihealthDocbaseUserinfoAlipayidGetResponse{}
    if(err != nil){
        log.Println("alibabaAlihealthDocbaseUserinfoAlipayidGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
