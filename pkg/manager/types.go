package manager

import (
	"go.uber.org/zap"
	versioned "knative.dev/serving/pkg/client/clientset/versioned"
	"knative.dev/serving/pkg/client/informers/externalversions/serving/v1"
)

type Manager struct {
	logger          *zap.SugaredLogger
	serviceInformer v1.ServiceInformer
	servingclient   versioned.Interface
}
