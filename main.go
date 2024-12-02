package main

import (
	"e-commerce/service/infra/log"
	"e-commerce/service/infra/repo"
	"e-commerce/service/infra/routine"
	"e-commerce/service/kit"
	"e-commerce/service/kit/client"
)

const SERVICE_NAME = "retail"

func init() {

}

func main() {
	log.Info("Run cms-center service ...")

	// init db, ebus, ...
	repo.Init()
	defer repo.Close()

	//trace.NewZipKinTracer(SERVICE_NAME + "#" + os.Getenv("RUN_TIME"))
	// init other svc clients
	client.StartClients()
	log.Info("Start clients done")
	kit.StartMainCenterGrpcService()
	log.Info("Start main service done")
	routine.LoopUntilExit()
	log.Info("Exit. Thanks Bye!")
}
