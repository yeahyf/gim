package service

import (
	"context"
	"gim/internal/logic/cache"
	"gim/internal/logic/dao"
	"gim/internal/logic/model"
)

type appService struct{}

var AppService = new(appService)

// Get 注册设备,先从cache中获取，如果没有获取到，则从数据库中获取，获取后将其放入到cahce中
func (*appService) Get(ctx context.Context, appId int64) (*model.App, error) {
	//从Reids中获取用户设备信息
	app, err := cache.AppCache.Get(appId)
	if err != nil {
		//return app, nil
		return nil,err
	}
	if app != nil {
		return app, nil
	}

	app, err = dao.AppDao.Get(appId)
	if err != nil {
		return app, nil
	}

	if app != nil {
		err = cache.AppCache.Set(app)
		if err != nil {
			return app, nil
		}
	}

	return app, nil
}
