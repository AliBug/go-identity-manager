package main

import (
	"log"
	"time"

	"github.com/alibug/go-identity-utils/config"
	"github.com/alibug/go-identity-utils/mongoconn"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"

	_casbinDelivery "github.com/alibug/go-identity-manager/rbac/delivery/restgin"
	_rbacUseCase "github.com/alibug/go-identity-manager/rbac/usecase"
	_usersDelivery "github.com/alibug/go-identity-manager/users/delivery/restgin"
	_usersRepo "github.com/alibug/go-identity-manager/users/repository/mongodb"
	_usersUseCase "github.com/alibug/go-identity-manager/users/usecase"
	_casbinAdapter "github.com/casbin/mongodb-adapter/v3"
)

func main() {
	mongoURL := config.ReadMongoConfig("mongo")

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
	rbacDomainUseCase := _rbacUseCase.NewRBACinDomainUsecase(enforcer)
	_casbinDelivery.NewRBACinDomainHandler(route, rbacDomainUseCase)

	roleManager := _rbacUseCase.NewRoleManagerUsecase(enforcer)
	_casbinDelivery.NewRoleManagerHandler(route, roleManager)

	// read mongo config
	conn, err := mongoconn.NewConn(mongoURL, timeDuration)
	if err != nil {
		log.Fatalf("Init DB conn fail: %v", err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = conn.Ping()
	if err != nil {
		log.Fatal("Connect to DB fail")
	}

	// New users collection
	usersColl := conn.GetColl("users")
	usersRepo := _usersRepo.NewMongoUserRepository(usersColl)
	usersUserCase := _usersUseCase.NewUsersUsecase(usersRepo, timeDuration)
	_usersDelivery.NewRoleManagerHandler(route, usersUserCase, roleManager)

	route.Run()
}
