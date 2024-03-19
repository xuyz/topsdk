package request


type AlibabaAlihealthMedicalbaseHospitalSyncRequest struct {
    /*
        是否需要用户授权     */
    IsAuth  *string `json:"is_auth" required:"true" `
    /*
        服务项列表     */
    Functions  *string `json:"functions" required:"true" `
    /*
        主院区纬度     */
    Lat  *string `json:"lat" required:"true" `
    /*
        主院区经度     */
    Lon  *string `json:"lon" required:"true" `
    /*
        主院区地址     */
    HosAddress  *string `json:"hos_address" required:"true" `
    /*
        主院区的联系电话     */
    Telephone  *string `json:"telephone,omitempty" required:"false" `
    /*
        院区名称     */
    RegionName  *string `json:"region_name" required:"true" `
    /*
        是否公立医院（Y／N）     */
    IsPublic  *string `json:"is_public" required:"true" `
    /*
        标签     */
    ServiceInfo  *string `json:"service_info,omitempty" required:"false" `
    /*
        自定义科室     */
    Special  *string `json:"special,omitempty" required:"false" `
    /*
        生活号或者服务窗url     */
    ServiceWindowUrl  *string `json:"service_window_url,omitempty" required:"false" `
    /*
        医院简介url     */
    DescriptionUrl  *string `json:"description_url,omitempty" required:"false" `
    /*
        是否支持医保（Y/N）     */
    IsInsurance  *string `json:"is_insurance,omitempty" required:"false" `
    /*
        医院等级     */
    Grade  *string `json:"grade" required:"true" `
    /*
        综合(general)、专科（special）     */
    Category  *string `json:"category" required:"true" `
    /*
        医院简称     */
    ShortName  *string `json:"short_name,omitempty" required:"false" `
    /*
        医院pid     */
    Pid  *string `json:"pid,omitempty" required:"false" `
    /*
        机构编码     */
    UnifyCode  *string `json:"unify_code,omitempty" required:"false" `
    /*
        所在城市code     */
    CityCode  *string `json:"city_code" required:"true" `
    /*
        营业执照上的医院全称     */
    HosName  *string `json:"hos_name" required:"true" `
    /*
        公司名称     */
    CompanyName  *string `json:"company_name" required:"true" `
    /*
        支付宝BD的姓名     */
    AliInterfaceMan  *string `json:"ali_interface_man" required:"true" `
    /*
        邮箱地址     */
    Email  *string `json:"email" required:"true" `
    /*
        联系人     */
    TechnicalMan  *string `json:"technical_man" required:"true" `
    /*
        联系手机     */
    Phone  *string `json:"phone" required:"true" `
    /*
        单医院（main）／ 平台（platform）     */
    HosType  *string `json:"hos_type" required:"true" `
    /*
        isv库里面的hosCode     */
    IsvHosCode  *string `json:"isv_hos_code" required:"true" `
    /*
        投放阵地alipay aliyy uc quark     */
    DeliveryChannel  *string `json:"delivery_channel,omitempty" required:"false" `
}

func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsAuth(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.IsAuth = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetFunctions(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Functions = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetLat(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Lat = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetLon(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Lon = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosAddress(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.HosAddress = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetTelephone(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Telephone = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetRegionName(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.RegionName = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsPublic(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.IsPublic = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetServiceInfo(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.ServiceInfo = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetSpecial(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Special = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetServiceWindowUrl(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.ServiceWindowUrl = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetDescriptionUrl(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.DescriptionUrl = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsInsurance(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.IsInsurance = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetGrade(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Grade = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCategory(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Category = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetShortName(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.ShortName = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetPid(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Pid = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetUnifyCode(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.UnifyCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCityCode(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.CityCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosName(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.HosName = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCompanyName(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.CompanyName = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetAliInterfaceMan(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.AliInterfaceMan = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetEmail(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Email = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetTechnicalMan(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.TechnicalMan = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetPhone(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.Phone = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosType(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.HosType = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsvHosCode(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.IsvHosCode = &v
    return s
}
func (s *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetDeliveryChannel(v string) *AlibabaAlihealthMedicalbaseHospitalSyncRequest {
    s.DeliveryChannel = &v
    return s
}

func (req *AlibabaAlihealthMedicalbaseHospitalSyncRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.IsAuth != nil) {
        paramMap["is_auth"] = *req.IsAuth
    }
    if(req.Functions != nil) {
        paramMap["functions"] = *req.Functions
    }
    if(req.Lat != nil) {
        paramMap["lat"] = *req.Lat
    }
    if(req.Lon != nil) {
        paramMap["lon"] = *req.Lon
    }
    if(req.HosAddress != nil) {
        paramMap["hos_address"] = *req.HosAddress
    }
    if(req.Telephone != nil) {
        paramMap["telephone"] = *req.Telephone
    }
    if(req.RegionName != nil) {
        paramMap["region_name"] = *req.RegionName
    }
    if(req.IsPublic != nil) {
        paramMap["is_public"] = *req.IsPublic
    }
    if(req.ServiceInfo != nil) {
        paramMap["service_info"] = *req.ServiceInfo
    }
    if(req.Special != nil) {
        paramMap["special"] = *req.Special
    }
    if(req.ServiceWindowUrl != nil) {
        paramMap["service_window_url"] = *req.ServiceWindowUrl
    }
    if(req.DescriptionUrl != nil) {
        paramMap["description_url"] = *req.DescriptionUrl
    }
    if(req.IsInsurance != nil) {
        paramMap["is_insurance"] = *req.IsInsurance
    }
    if(req.Grade != nil) {
        paramMap["grade"] = *req.Grade
    }
    if(req.Category != nil) {
        paramMap["category"] = *req.Category
    }
    if(req.ShortName != nil) {
        paramMap["short_name"] = *req.ShortName
    }
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.UnifyCode != nil) {
        paramMap["unify_code"] = *req.UnifyCode
    }
    if(req.CityCode != nil) {
        paramMap["city_code"] = *req.CityCode
    }
    if(req.HosName != nil) {
        paramMap["hos_name"] = *req.HosName
    }
    if(req.CompanyName != nil) {
        paramMap["company_name"] = *req.CompanyName
    }
    if(req.AliInterfaceMan != nil) {
        paramMap["ali_interface_man"] = *req.AliInterfaceMan
    }
    if(req.Email != nil) {
        paramMap["email"] = *req.Email
    }
    if(req.TechnicalMan != nil) {
        paramMap["technical_man"] = *req.TechnicalMan
    }
    if(req.Phone != nil) {
        paramMap["phone"] = *req.Phone
    }
    if(req.HosType != nil) {
        paramMap["hos_type"] = *req.HosType
    }
    if(req.IsvHosCode != nil) {
        paramMap["isv_hos_code"] = *req.IsvHosCode
    }
    if(req.DeliveryChannel != nil) {
        paramMap["delivery_channel"] = *req.DeliveryChannel
    }
    return paramMap
}

func (req *AlibabaAlihealthMedicalbaseHospitalSyncRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}