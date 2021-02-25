// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: cootek.pgd.weather_mgr.proto

package weather_mgr

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type WeatherReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lng         float64 `protobuf:"fixed64,1,opt,name=lng,proto3" json:"lng,omitempty"`
	Lat         float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
	CityCode    string  `protobuf:"bytes,3,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	TestGroup   string  `protobuf:"bytes,4,opt,name=test_group,json=testGroup,proto3" json:"test_group,omitempty"`       // 实验组
	WeatherType string  `protobuf:"bytes,5,opt,name=weather_type,json=weatherType,proto3" json:"weather_type,omitempty"` // 查看天气的类型
	Uid         uint64  `protobuf:"varint,6,opt,name=uid,proto3" json:"uid,omitempty"`                                   // ass
	ApiType     string  `protobuf:"bytes,7,opt,name=api_type,json=apiType,proto3" json:"api_type,omitempty"`             // api类型
}

func (x *WeatherReq) Reset() {
	*x = WeatherReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherReq) ProtoMessage() {}

func (x *WeatherReq) ProtoReflect() protoreflect.Message {
	mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherReq.ProtoReflect.Descriptor instead.
func (*WeatherReq) Descriptor() ([]byte, []int) {
	return file_cootek_pgd_weather_mgr_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherReq) GetLng() float64 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *WeatherReq) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *WeatherReq) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *WeatherReq) GetTestGroup() string {
	if x != nil {
		return x.TestGroup
	}
	return ""
}

func (x *WeatherReq) GetWeatherType() string {
	if x != nil {
		return x.WeatherType
	}
	return ""
}

func (x *WeatherReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *WeatherReq) GetApiType() string {
	if x != nil {
		return x.ApiType
	}
	return ""
}

type HourlyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*HourlyStyle `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *HourlyResp) Reset() {
	*x = HourlyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HourlyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HourlyResp) ProtoMessage() {}

func (x *HourlyResp) ProtoReflect() protoreflect.Message {
	mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HourlyResp.ProtoReflect.Descriptor instead.
func (*HourlyResp) Descriptor() ([]byte, []int) {
	return file_cootek_pgd_weather_mgr_proto_rawDescGZIP(), []int{1}
}

func (x *HourlyResp) GetList() []*HourlyStyle {
	if x != nil {
		return x.List
	}
	return nil
}

type DailyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*DailyStyle `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *DailyResp) Reset() {
	*x = DailyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyResp) ProtoMessage() {}

func (x *DailyResp) ProtoReflect() protoreflect.Message {
	mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyResp.ProtoReflect.Descriptor instead.
func (*DailyResp) Descriptor() ([]byte, []int) {
	return file_cootek_pgd_weather_mgr_proto_rawDescGZIP(), []int{2}
}

func (x *DailyResp) GetList() []*DailyStyle {
	if x != nil {
		return x.List
	}
	return nil
}

type HourlyStyle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        int64  `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	Temperature string `protobuf:"bytes,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Skycon      string `protobuf:"bytes,3,opt,name=skycon,proto3" json:"skycon,omitempty"`
}

func (x *HourlyStyle) Reset() {
	*x = HourlyStyle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HourlyStyle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HourlyStyle) ProtoMessage() {}

func (x *HourlyStyle) ProtoReflect() protoreflect.Message {
	mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HourlyStyle.ProtoReflect.Descriptor instead.
func (*HourlyStyle) Descriptor() ([]byte, []int) {
	return file_cootek_pgd_weather_mgr_proto_rawDescGZIP(), []int{3}
}

func (x *HourlyStyle) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *HourlyStyle) GetTemperature() string {
	if x != nil {
		return x.Temperature
	}
	return ""
}

func (x *HourlyStyle) GetSkycon() string {
	if x != nil {
		return x.Skycon
	}
	return ""
}

type DailyStyle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date           int64   `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	TemperatureMax string  `protobuf:"bytes,2,opt,name=temperature_max,json=temperatureMax,proto3" json:"temperature_max,omitempty"`
	TemperatureMin string  `protobuf:"bytes,3,opt,name=temperature_min,json=temperatureMin,proto3" json:"temperature_min,omitempty"`
	Skycon         string  `protobuf:"bytes,4,opt,name=skycon,proto3" json:"skycon,omitempty"`
	Sunset         string  `protobuf:"bytes,5,opt,name=sunset,proto3" json:"sunset,omitempty"`
	Sunrise        string  `protobuf:"bytes,6,opt,name=sunrise,proto3" json:"sunrise,omitempty"`
	Aqi            float64 `protobuf:"fixed64,7,opt,name=aqi,proto3" json:"aqi,omitempty"`
}

func (x *DailyStyle) Reset() {
	*x = DailyStyle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyStyle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyStyle) ProtoMessage() {}

func (x *DailyStyle) ProtoReflect() protoreflect.Message {
	mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyStyle.ProtoReflect.Descriptor instead.
func (*DailyStyle) Descriptor() ([]byte, []int) {
	return file_cootek_pgd_weather_mgr_proto_rawDescGZIP(), []int{4}
}

func (x *DailyStyle) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *DailyStyle) GetTemperatureMax() string {
	if x != nil {
		return x.TemperatureMax
	}
	return ""
}

func (x *DailyStyle) GetTemperatureMin() string {
	if x != nil {
		return x.TemperatureMin
	}
	return ""
}

func (x *DailyStyle) GetSkycon() string {
	if x != nil {
		return x.Skycon
	}
	return ""
}

func (x *DailyStyle) GetSunset() string {
	if x != nil {
		return x.Sunset
	}
	return ""
}

func (x *DailyStyle) GetSunrise() string {
	if x != nil {
		return x.Sunrise
	}
	return ""
}

func (x *DailyStyle) GetAqi() float64 {
	if x != nil {
		return x.Aqi
	}
	return 0
}

type TodayResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForecastKeypoint string          `protobuf:"bytes,1,opt,name=forecast_keypoint,json=forecastKeypoint,proto3" json:"forecast_keypoint,omitempty"` // 生活指数预报的详细描述，可能为空
	RainDesc         string          `protobuf:"bytes,2,opt,name=rain_desc,json=rainDesc,proto3" json:"rain_desc,omitempty"`                         // 降水描述，可能为空
	AlertDesc        string          `protobuf:"bytes,3,opt,name=alert_desc,json=alertDesc,proto3" json:"alert_desc,omitempty"`                      // 预警详细文字描述
	WarmRemind       string          `protobuf:"bytes,4,opt,name=warm_remind,json=warmRemind,proto3" json:"warm_remind,omitempty"`                   // 提示
	Humidity         string          `protobuf:"bytes,5,opt,name=humidity,proto3" json:"humidity,omitempty"`                                         // 湿度
	Comfort          string          `protobuf:"bytes,6,opt,name=comfort,proto3" json:"comfort,omitempty"`                                           // 舒适指数
	LifeSuggestion   *LifeSuggestion `protobuf:"bytes,7,opt,name=life_suggestion,json=lifeSuggestion,proto3" json:"life_suggestion,omitempty"`
	Date             int64           `protobuf:"varint,8,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *TodayResp) Reset() {
	*x = TodayResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodayResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodayResp) ProtoMessage() {}

func (x *TodayResp) ProtoReflect() protoreflect.Message {
	mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodayResp.ProtoReflect.Descriptor instead.
func (*TodayResp) Descriptor() ([]byte, []int) {
	return file_cootek_pgd_weather_mgr_proto_rawDescGZIP(), []int{5}
}

func (x *TodayResp) GetForecastKeypoint() string {
	if x != nil {
		return x.ForecastKeypoint
	}
	return ""
}

func (x *TodayResp) GetRainDesc() string {
	if x != nil {
		return x.RainDesc
	}
	return ""
}

func (x *TodayResp) GetAlertDesc() string {
	if x != nil {
		return x.AlertDesc
	}
	return ""
}

func (x *TodayResp) GetWarmRemind() string {
	if x != nil {
		return x.WarmRemind
	}
	return ""
}

func (x *TodayResp) GetHumidity() string {
	if x != nil {
		return x.Humidity
	}
	return ""
}

func (x *TodayResp) GetComfort() string {
	if x != nil {
		return x.Comfort
	}
	return ""
}

func (x *TodayResp) GetLifeSuggestion() *LifeSuggestion {
	if x != nil {
		return x.LifeSuggestion
	}
	return nil
}

func (x *TodayResp) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type LifeSuggestion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dressing string `protobuf:"bytes,1,opt,name=dressing,proto3" json:"dressing,omitempty"` // 穿衣指数
	Fishing  string `protobuf:"bytes,2,opt,name=fishing,proto3" json:"fishing,omitempty"`   // 钓鱼指数
	Flu      string `protobuf:"bytes,3,opt,name=flu,proto3" json:"flu,omitempty"`           // 感冒指数
	Sport    string `protobuf:"bytes,4,opt,name=sport,proto3" json:"sport,omitempty"`       // 运动指数
	Uv       string `protobuf:"bytes,5,opt,name=uv,proto3" json:"uv,omitempty"`             // 紫外线指数
	Travel   string `protobuf:"bytes,6,opt,name=travel,proto3" json:"travel,omitempty"`     // 旅游指数
	Airing   string `protobuf:"bytes,7,opt,name=airing,proto3" json:"airing,omitempty"`     // 晾晒
}

func (x *LifeSuggestion) Reset() {
	*x = LifeSuggestion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifeSuggestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifeSuggestion) ProtoMessage() {}

func (x *LifeSuggestion) ProtoReflect() protoreflect.Message {
	mi := &file_cootek_pgd_weather_mgr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifeSuggestion.ProtoReflect.Descriptor instead.
func (*LifeSuggestion) Descriptor() ([]byte, []int) {
	return file_cootek_pgd_weather_mgr_proto_rawDescGZIP(), []int{6}
}

func (x *LifeSuggestion) GetDressing() string {
	if x != nil {
		return x.Dressing
	}
	return ""
}

func (x *LifeSuggestion) GetFishing() string {
	if x != nil {
		return x.Fishing
	}
	return ""
}

func (x *LifeSuggestion) GetFlu() string {
	if x != nil {
		return x.Flu
	}
	return ""
}

func (x *LifeSuggestion) GetSport() string {
	if x != nil {
		return x.Sport
	}
	return ""
}

func (x *LifeSuggestion) GetUv() string {
	if x != nil {
		return x.Uv
	}
	return ""
}

func (x *LifeSuggestion) GetTravel() string {
	if x != nil {
		return x.Travel
	}
	return ""
}

func (x *LifeSuggestion) GetAiring() string {
	if x != nil {
		return x.Airing
	}
	return ""
}

var File_cootek_pgd_weather_mgr_proto protoreflect.FileDescriptor

var file_cootek_pgd_weather_mgr_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x22, 0xbc, 0x01, 0x0a, 0x0a, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70,
	0x69, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70,
	0x69, 0x54, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x0a, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x48, 0x6f, 0x75, 0x72, 0x6c,
	0x79, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x09,
	0x44, 0x61, 0x69, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b,
	0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72,
	0x2e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x5b, 0x0a, 0x0b, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x53, 0x74, 0x79, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x79, 0x63, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x79, 0x63, 0x6f, 0x6e, 0x22, 0xce,
	0x01, 0x0a, 0x0a, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x4d, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6b, 0x79, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x79, 0x63, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x75, 0x6e, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x6e,
	0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6e, 0x72, 0x69, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x71, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x71, 0x69, 0x22,
	0xb0, 0x02, 0x0a, 0x09, 0x54, 0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a,
	0x11, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61,
	0x69, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x61, 0x69, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x6d, 0x5f, 0x72,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x66, 0x6f, 0x72, 0x74, 0x12, 0x4f, 0x0a,
	0x0f, 0x6c, 0x69, 0x66, 0x65, 0x5f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e,
	0x70, 0x67, 0x64, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e,
	0x4c, 0x69, 0x66, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e,
	0x6c, 0x69, 0x66, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x66, 0x65, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x6c, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x6c, 0x75, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x75, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x69, 0x72,
	0x69, 0x6e, 0x67, 0x32, 0x84, 0x02, 0x0a, 0x0a, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4d,
	0x67, 0x72, 0x12, 0x52, 0x0a, 0x06, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x12, 0x22, 0x2e, 0x63,
	0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x48, 0x6f, 0x75, 0x72, 0x6c, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x05, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x12,
	0x22, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64,
	0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x44, 0x61, 0x69,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x05, 0x54, 0x6f, 0x64, 0x61,
	0x79, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70, 0x67, 0x64, 0x2e, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b, 0x2e, 0x70,
	0x67, 0x64, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2e, 0x54,
	0x6f, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72, 0x2f, 0x63, 0x6f, 0x6f, 0x74, 0x65, 0x6b,
	0x2f, 0x70, 0x67, 0x64, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x67, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cootek_pgd_weather_mgr_proto_rawDescOnce sync.Once
	file_cootek_pgd_weather_mgr_proto_rawDescData = file_cootek_pgd_weather_mgr_proto_rawDesc
)

func file_cootek_pgd_weather_mgr_proto_rawDescGZIP() []byte {
	file_cootek_pgd_weather_mgr_proto_rawDescOnce.Do(func() {
		file_cootek_pgd_weather_mgr_proto_rawDescData = protoimpl.X.CompressGZIP(file_cootek_pgd_weather_mgr_proto_rawDescData)
	})
	return file_cootek_pgd_weather_mgr_proto_rawDescData
}

var file_cootek_pgd_weather_mgr_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cootek_pgd_weather_mgr_proto_goTypes = []interface{}{
	(*WeatherReq)(nil),     // 0: cootek.pgd.weather_mgr.WeatherReq
	(*HourlyResp)(nil),     // 1: cootek.pgd.weather_mgr.HourlyResp
	(*DailyResp)(nil),      // 2: cootek.pgd.weather_mgr.DailyResp
	(*HourlyStyle)(nil),    // 3: cootek.pgd.weather_mgr.HourlyStyle
	(*DailyStyle)(nil),     // 4: cootek.pgd.weather_mgr.DailyStyle
	(*TodayResp)(nil),      // 5: cootek.pgd.weather_mgr.TodayResp
	(*LifeSuggestion)(nil), // 6: cootek.pgd.weather_mgr.LifeSuggestion
}
var file_cootek_pgd_weather_mgr_proto_depIdxs = []int32{
	3, // 0: cootek.pgd.weather_mgr.HourlyResp.list:type_name -> cootek.pgd.weather_mgr.HourlyStyle
	4, // 1: cootek.pgd.weather_mgr.DailyResp.list:type_name -> cootek.pgd.weather_mgr.DailyStyle
	6, // 2: cootek.pgd.weather_mgr.TodayResp.life_suggestion:type_name -> cootek.pgd.weather_mgr.LifeSuggestion
	0, // 3: cootek.pgd.weather_mgr.WeatherMgr.Hourly:input_type -> cootek.pgd.weather_mgr.WeatherReq
	0, // 4: cootek.pgd.weather_mgr.WeatherMgr.Daily:input_type -> cootek.pgd.weather_mgr.WeatherReq
	0, // 5: cootek.pgd.weather_mgr.WeatherMgr.Today:input_type -> cootek.pgd.weather_mgr.WeatherReq
	1, // 6: cootek.pgd.weather_mgr.WeatherMgr.Hourly:output_type -> cootek.pgd.weather_mgr.HourlyResp
	2, // 7: cootek.pgd.weather_mgr.WeatherMgr.Daily:output_type -> cootek.pgd.weather_mgr.DailyResp
	5, // 8: cootek.pgd.weather_mgr.WeatherMgr.Today:output_type -> cootek.pgd.weather_mgr.TodayResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cootek_pgd_weather_mgr_proto_init() }
func file_cootek_pgd_weather_mgr_proto_init() {
	if File_cootek_pgd_weather_mgr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cootek_pgd_weather_mgr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeatherReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cootek_pgd_weather_mgr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HourlyResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cootek_pgd_weather_mgr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cootek_pgd_weather_mgr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HourlyStyle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cootek_pgd_weather_mgr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DailyStyle); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cootek_pgd_weather_mgr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodayResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cootek_pgd_weather_mgr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifeSuggestion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cootek_pgd_weather_mgr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cootek_pgd_weather_mgr_proto_goTypes,
		DependencyIndexes: file_cootek_pgd_weather_mgr_proto_depIdxs,
		MessageInfos:      file_cootek_pgd_weather_mgr_proto_msgTypes,
	}.Build()
	File_cootek_pgd_weather_mgr_proto = out.File
	file_cootek_pgd_weather_mgr_proto_rawDesc = nil
	file_cootek_pgd_weather_mgr_proto_goTypes = nil
	file_cootek_pgd_weather_mgr_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WeatherMgrClient is the client API for WeatherMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeatherMgrClient interface {
	// 24小时 天气情况
	Hourly(ctx context.Context, in *WeatherReq, opts ...grpc.CallOption) (*HourlyResp, error)
	// 近一周天气情况
	Daily(ctx context.Context, in *WeatherReq, opts ...grpc.CallOption) (*DailyResp, error)
	// 当天生活指数情况
	Today(ctx context.Context, in *WeatherReq, opts ...grpc.CallOption) (*TodayResp, error)
}

type weatherMgrClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherMgrClient(cc grpc.ClientConnInterface) WeatherMgrClient {
	return &weatherMgrClient{cc}
}

func (c *weatherMgrClient) Hourly(ctx context.Context, in *WeatherReq, opts ...grpc.CallOption) (*HourlyResp, error) {
	out := new(HourlyResp)
	err := c.cc.Invoke(ctx, "/cootek.pgd.weather_mgr.WeatherMgr/Hourly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherMgrClient) Daily(ctx context.Context, in *WeatherReq, opts ...grpc.CallOption) (*DailyResp, error) {
	out := new(DailyResp)
	err := c.cc.Invoke(ctx, "/cootek.pgd.weather_mgr.WeatherMgr/Daily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherMgrClient) Today(ctx context.Context, in *WeatherReq, opts ...grpc.CallOption) (*TodayResp, error) {
	out := new(TodayResp)
	err := c.cc.Invoke(ctx, "/cootek.pgd.weather_mgr.WeatherMgr/Today", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherMgrServer is the server API for WeatherMgr service.
type WeatherMgrServer interface {
	// 24小时 天气情况
	Hourly(context.Context, *WeatherReq) (*HourlyResp, error)
	// 近一周天气情况
	Daily(context.Context, *WeatherReq) (*DailyResp, error)
	// 当天生活指数情况
	Today(context.Context, *WeatherReq) (*TodayResp, error)
}

// UnimplementedWeatherMgrServer can be embedded to have forward compatible implementations.
type UnimplementedWeatherMgrServer struct {
}

func (*UnimplementedWeatherMgrServer) Hourly(context.Context, *WeatherReq) (*HourlyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hourly not implemented")
}
func (*UnimplementedWeatherMgrServer) Daily(context.Context, *WeatherReq) (*DailyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Daily not implemented")
}
func (*UnimplementedWeatherMgrServer) Today(context.Context, *WeatherReq) (*TodayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Today not implemented")
}

func RegisterWeatherMgrServer(s *grpc.Server, srv WeatherMgrServer) {
	s.RegisterService(&_WeatherMgr_serviceDesc, srv)
}

func _WeatherMgr_Hourly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherMgrServer).Hourly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cootek.pgd.weather_mgr.WeatherMgr/Hourly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherMgrServer).Hourly(ctx, req.(*WeatherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherMgr_Daily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherMgrServer).Daily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cootek.pgd.weather_mgr.WeatherMgr/Daily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherMgrServer).Daily(ctx, req.(*WeatherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeatherMgr_Today_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WeatherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherMgrServer).Today(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cootek.pgd.weather_mgr.WeatherMgr/Today",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherMgrServer).Today(ctx, req.(*WeatherReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _WeatherMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cootek.pgd.weather_mgr.WeatherMgr",
	HandlerType: (*WeatherMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hourly",
			Handler:    _WeatherMgr_Hourly_Handler,
		},
		{
			MethodName: "Daily",
			Handler:    _WeatherMgr_Daily_Handler,
		},
		{
			MethodName: "Today",
			Handler:    _WeatherMgr_Today_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cootek.pgd.weather_mgr.proto",
}
