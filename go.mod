module github.com/alibug/go-identity-manager

go 1.16

require (
	github.com/alibug/go-identity-utils v0.1.6
	github.com/casbin/casbin/v2 v2.30.2
	github.com/casbin/mongodb-adapter/v3 v3.2.0
	github.com/gin-gonic/gin v1.7.1
	go.mongodb.org/mongo-driver v1.5.2
)

replace github.com/casbin/mongodb-adapter/v3 => github.com/alibug/go-casbin-mongodb-adapter/v3 v3.2.1
