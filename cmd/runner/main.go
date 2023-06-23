package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	commonv1alpha1 "github.com/octopipe/cloudx/apis/common/v1alpha1"
	"github.com/octopipe/cloudx/internal/controller/sharedinfra"
	"github.com/octopipe/cloudx/internal/execution"
	"github.com/octopipe/cloudx/internal/rpcclient"
	"github.com/octopipe/cloudx/internal/terraform"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(commonv1alpha1.AddToScheme(scheme))
}

type runnerContext struct {
	logger    *zap.Logger
	rpcClient rpcclient.Client
}

func main() {
	logger, _ := zap.NewProduction()
	_ = godotenv.Load()

	rpcClient, err := rpcclient.NewRPCClient(os.Getenv("RPC_SERVER_ADDRESS"))
	if err != nil {
		logger.Fatal("Error to connect with controllerr", zap.Error(err), zap.String("address", os.Getenv("RPC_SERVER_ADDRESS")))
	}

	terraformProvider, err := terraform.NewTerraformProvider(logger)
	if err != nil {
		panic(err)
	}

	newRunnerContext := runnerContext{
		rpcClient: rpcClient,
		logger:    logger,
	}

	currentSharedInfra, executionId, action, err := newRunnerContext.getDataFromCommandArgs()
	if err != nil {
		panic(err)
	}

	sharedInfraRef := types.NamespacedName{Name: currentSharedInfra.Name, Namespace: currentSharedInfra.Namespace}

	go newRunnerContext.startTimeout(sharedInfraRef, executionId)

	rpcRunnerFinishedArgs := &sharedinfra.RPCSetRunnerFinishedArgs{
		ExecutionId: executionId,
		Ref:         sharedInfraRef,
		Error:       "",
		Plugins:     []commonv1alpha1.PluginStatus{},
		Status:      "SUCCESS",
		FinishedAt:  time.Now().Format(time.RFC3339),
	}

	currentExecution := execution.NewExecution(logger, terraformProvider, currentSharedInfra, newRunnerContext.setRunnerStatus(*rpcRunnerFinishedArgs))
	if action == "APPLY" {
		pluginStatus, err := currentExecution.Apply()
		if err != nil {
			rpcRunnerFinishedArgs.Status = "ERROR"
			rpcRunnerFinishedArgs.Error = err.Error()
		}

		if newRunnerContext.hasErrorsInPluginExecutions(pluginStatus) {
			rpcRunnerFinishedArgs.Status = "ERROR"
			rpcRunnerFinishedArgs.Error = "Found error in plugin execution"
		}

		rpcRunnerFinishedArgs.Plugins = pluginStatus

		var reply int
		err = rpcClient.Call("RPCServer.SetRunnerFinished", rpcRunnerFinishedArgs, &reply)
		if err != nil {
			logger.Fatal("Error to call controller", zap.Error(err))
		}
	}

	if action == "DESTROY" {
		err := currentExecution.Destroy()
		if err != nil {
			rpcRunnerFinishedArgs.Status = "ERROR"
			rpcRunnerFinishedArgs.Error = err.Error()
		}
	}

	logger.Info("Finish runner execution")
}

func (c runnerContext) hasErrorsInPluginExecutions(pluginStatus []commonv1alpha1.PluginStatus) bool {
	for _, p := range pluginStatus {
		if p.Status == execution.ExecutionErrorStatus || p.Status == execution.ExecutionFailedStatus {
			return true
		}
	}

	return false
}

func (c runnerContext) setRunnerStatus(args sharedinfra.RPCSetRunnerFinishedArgs) func(plugins []commonv1alpha1.PluginStatus) error {
	return func(plugins []commonv1alpha1.PluginStatus) error {
		var reply int
		args.Plugins = plugins

		err := c.rpcClient.Call("RPCServer.SetRunnerFinished", args, &reply)
		return err
	}
}

func (c runnerContext) getDataFromCommandArgs() (commonv1alpha1.SharedInfra, string, string, error) {
	commandArgs := os.Args[1:]
	executionId := commandArgs[1]
	action := commandArgs[0]

	var rawSharedInfra string
	err := json.Unmarshal([]byte(commandArgs[2]), &rawSharedInfra)
	if err != nil {
		return commonv1alpha1.SharedInfra{}, "", "", err
	}

	fmt.Println(rawSharedInfra)

	var sharedInfra commonv1alpha1.SharedInfra
	err = json.Unmarshal([]byte(rawSharedInfra), &sharedInfra)
	if err != nil {
		return commonv1alpha1.SharedInfra{}, "", "", err
	}

	return sharedInfra, executionId, action, nil
}

func (c runnerContext) startTimeout(sharedInfraRef types.NamespacedName, executionId string) {
	time.Sleep(5 * time.Minute)

	var reply int
	args := &sharedinfra.RPCSetRunnerTimeoutArgs{
		SharedInfraRef: sharedInfraRef,
		ExecutionId:    executionId,
	}

	err := c.rpcClient.Call("RPCServer.SetRunnerTimeout", args, &reply)
	if err != nil {
		c.logger.Fatal("call rpc timeout error")
	}

	c.logger.Fatal("Runner timeout")
}
