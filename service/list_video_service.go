package service

import (
	"Refactor_xilixili/model"
	"Refactor_xilixili/serializer"
)

type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

func (service *ListVideoService) List() serializer.Response {
	//println("1")
	var videos []model.Video
	var total int64

	if service.Limit == 0 {
		service.Limit = 6
	}

	err := model.DB.Model(model.Video{}).Count(&total).Error //gorm中查video这张表，并把count*给total
	if err != nil {
		return serializer.Response{
			//如果要处理需要到api那里给c
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(), //打出错误信息
		}
	}
	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&videos).Error; err != nil { //查找start到start+limit的视频
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	//println("2")
	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
