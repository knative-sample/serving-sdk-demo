package manager

import (
	"context"

	servingclient "knative.dev/serving/pkg/client/injection/client"
	kserviceinformer "knative.dev/serving/pkg/client/injection/informers/serving/v1/service"

	"knative.dev/pkg/logging"
)

func NewManager(
	ctx context.Context,
) *Manager {
	logger := logging.FromContext(ctx)
	serviceInformer := kserviceinformer.Get(ctx)

	m := &Manager{
		logger:          logger,
		servingclient:   servingclient.Get(ctx),
		serviceInformer: serviceInformer,
	}

	return m
}

func (m *Manager) Run() error {
	logger := m.logger
	logger.Infof("Manager is run")

	// 这里执行业务逻辑，比如：
	// 启动 一个 http restful api 服务，接收到请求以后触发创建、删除、查询 ksvc 等操作

	// 这是一个创建 ksvc 的例子
	if err := m.CreateKsvc(); err != nil {
		logger.Errorf("create ksvc error:%s", err.Error())
		return err
	}

	return nil
}
