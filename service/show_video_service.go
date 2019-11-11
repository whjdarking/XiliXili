package service

import (
	"xilixili/model"
	"xilixili/serializer"
)

type ShowVideoService struct {
}
func (service *ShowVideoService) Show(id string) serializer.Response{
    var video model.Video
	//全局单例DB
    err := model.DB.Where("id = ?", id).Find(&video).Error
    if err != nil {
		return serializer.Response{
			//如果要处理code需要到api那里给c
			Code: 404,
			Msg:  "不存在",
			Error: err.Error(), //打出错误信息
		}
	}

	video.AddView() //计数

	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildVideo(video),
		Msg:   "",
		Error: "",
	}
}