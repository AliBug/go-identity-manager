package usecase

import (
	"log"

	"github.com/alibug/go-identity-manager/domain"
	"github.com/casbin/casbin/v2"
)

type roleManagerUsecase struct {
	enforcer *casbin.Enforcer
}

// NewRoleManagerUsecase -
func NewRoleManagerUsecase(enforcer *casbin.Enforcer) domain.RoleManagerUseCase {
	return &roleManagerUsecase{enforcer: enforcer}
}

func (r *roleManagerUsecase) GetPoliciesInDomain(domain string) [][]string {
	return r.enforcer.GetFilteredNamedPolicy("p", 1, domain)
}

func (r *roleManagerUsecase) GetPolicies() [][]string {
	return r.enforcer.GetNamedPolicy("p")
}

func (r *roleManagerUsecase) AddNamedPolicy(policy domain.Policy) (bool, error) {
	return r.enforcer.AddNamedPolicy("p", policy.GetPolcy())
}

func (r *roleManagerUsecase) RemoveNamedPolicy(policy domain.Policy) (bool, error) {
	return r.enforcer.RemoveNamedPolicy("p", policy.GetPolcy())
}

func (r *roleManagerUsecase) RemoveFilteredNamedPolicy(domainName string) (bool, error) {
	log.Println("RemoveFilteredNamedPolicy:", domainName)
	// 1、 清空指定域名下所有 Role 分配
	// _, err := r.enforcer.RemoveFilteredNamedPolicy("g", 2, domainName)

	_, err := r.enforcer.RemoveFilteredGroupingPolicy(2, domainName)

	// ⚠️ 此处有错误， 错误不明
	log.Println("通过 RemoveFilteredGroupingPolicy")

	if err != nil {
		log.Println("del err", err)
		return false, err
	}
	log.Println("it ok!")
	// 2、清空指定域名下所有 policy 分配s
	_, err = r.enforcer.RemoveFilteredPolicy(1, domainName)
	if err != nil {
		log.Println("del 🌞 err", err)
		return false, err
	}

	return true, nil
}

func (r *roleManagerUsecase) DeleteUser(userID string) (bool, error) {
	return r.enforcer.DeleteUser(userID)
}
