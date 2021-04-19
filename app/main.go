package main

import (
	"log"
	"time"

	"github.com/alibug/go-identity-utils/config"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	_casbinDelivery "github.com/alibug/go-user-casbin/rbac/delivery/restgin"
	_rbacInDomainRepo "github.com/alibug/go-user-casbin/rbac/repository/rbacdomain"
	_rbacInDomainUseCase "github.com/alibug/go-user-casbin/rbac/usecase"
	_casbinAdapter "github.com/casbin/mongodb-adapter/v3"
)

func main() {
	mongoURL := config.ReadMongoConfig("mongo")
	log.Println("url:", mongoURL)

	timeDuration := 100 * time.Second

	adapter, err := _casbinAdapter.NewAdapter(mongoURL, timeDuration)

	if err != nil {
		log.Fatalf("Init mongo adapter fail: %s", err)
	}

	casbinModel := config.ReadCasbinFilePath("casbin")
	enforcer, err := casbin.NewEnforcer(casbinModel, adapter)

	if err != nil {
		log.Fatalf("Init Casbin enforcer fail: %s", err)
	}

	route := gin.Default()
	rbacDomainRepo := _rbacInDomainRepo.NewCasbinRepository(enforcer)
	rbacDomainUseCase := _rbacInDomainUseCase.NewRBACinDomainUsecase(rbacDomainRepo, timeDuration)
	_casbinDelivery.NewRBACinDomainHandler(route, rbacDomainUseCase)
	route.Run()
}
