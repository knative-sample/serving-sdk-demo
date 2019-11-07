package app

import (
	"encoding/json"

	"log"

	"time"

	manager2 "github.com/knative-sample/serving-sdk-demo/pkg/manager"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/signals"
)

var defaultZLC = []byte(`{
  "level": "info",
  "development": false,
  "outputPaths": ["stdout"],
  "errorOutputPaths": ["stderr"],
  "encoding": "json",
  "encoderConfig": {
    "timeKey": "ts",
    "levelKey": "level",
    "nameKey": "logger",
    "callerKey": "caller",
    "messageKey": "msg",
    "stacktraceKey": "stacktrace",
    "lineEnding": "",
    "levelEncoder": "",
    "timeEncoder": "iso8601",
    "durationEncoder": "",
    "callerEncoder": ""
  }
}`)

// start edas api
func NewCommandStartServer() *cobra.Command {
	ops := &Options{}
	mainCmd := &cobra.Command{
		Short: "serving-controller",
		Long:  "serving-controller",
		RunE: func(c *cobra.Command, args []string) error {
			run(ops)
			return nil
		},
	}

	ops.SetOps(mainCmd)
	return mainCmd
}

func run(ops *Options) {
	var logger *zap.SugaredLogger
	// setup logger
	var zlc zap.Config
	if err := json.Unmarshal(defaultZLC, &zlc); err != nil {
		log.Fatalf("Unmarshal zap.Logger config error:%s ", err)
	}
	if l, err := zlc.Build(); err != nil {
		log.Fatalf("Build Logger error:%s", err)
	} else {
		logger = l.Sugar()
	}

	defer logger.Sync()

	ctx := signals.NewContext()
	ctx = logging.WithLogger(ctx, logger)
	logger.Info("logger construction succeeded")

	// setup informer
	cfg, err := sharedmain.GetConfig(ops.MasterURL, ops.Kubeconfig)
	if err != nil {
		logger.Fatal("Error building kubeconfig:", err)
	}

	logger.Infof("Registering %d clients", len(injection.Default.GetClients()))
	logger.Infof("Registering %d informer factories", len(injection.Default.GetInformerFactories()))
	logger.Infof("Registering %d informers", len(injection.Default.GetInformers()))

	// set ResyncPeriod
	ctx = controller.WithResyncPeriod(ctx, time.Second*300)
	ctx, informers := injection.Default.SetupInformers(ctx, cfg)

	// Start all of the informers and wait for them to sync.
	logger.Info("Starting informers.")
	if err := controller.StartInformers(ctx.Done(), informers...); err != nil {
		logger.Fatalw("Failed to start informers", err)
	}

	// 启动 manager，负责管理 ksvc
	manager := manager2.NewManager(ctx)
	go manager.Run()

	_, egCtx := errgroup.WithContext(ctx)
	// This will block until either a signal arrives or one of the grouped functions
	// returns an error.
	<-egCtx.Done()
}
