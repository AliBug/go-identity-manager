module github.com/alibug/go-user-casbin

go 1.16

require (
	github.com/alibug/go-identity-utils v0.1.2
	github.com/casbin/casbin/v2 v2.27.0
	github.com/casbin/mongodb-adapter/v3 v3.2.0
	github.com/gin-gonic/gin v1.7.1
)

replace github.com/casbin/mongodb-adapter/v3 => github.com/AliBug/mongodb-adapter/v3 v3.2.1
