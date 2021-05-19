package usecase

import (
	"fmt"
	// "log"

	"github.com/alibug/go-identity-manager/domain"
	"github.com/alibug/go-identity-utils/status"
	"github.com/casbin/casbin/v2"
)

type rbacInDomainUsecase struct {
	enforcer *casbin.Enforcer
}

// NewRBACinDomainUsecase -
func NewRBACinDomainUsecase(enforcer *casbin.Enforcer) domain.RBACinDomainUseCase {
	return &rbacInDomainUsecase{enforcer: enforcer}
}

func (r *rbacInDomainUsecase) AddRoleForUserInDomainUC(ur domain.UserRoleInDomain) (bool, error) {
	if len(r.enforcer.GetFilteredNamedPolicy("p", 0, ur.GetRole(), ur.GetDomain())) == 0 {
		return false, fmt.Errorf("%w: role %s not in domain %s", status.ErrNotFound, ur.GetRole(), ur.GetDomain())
	}
	return r.enforcer.AddRoleForUserInDomain(ur.GetUser(), ur.GetRole(), ur.GetDomain())
}

func (r *rbacInDomainUsecase) DeleteRoleForUserInDomainUC(ur domain.UserRoleInDomain) (bool, error) {
	return r.enforcer.DeleteRoleForUserInDomain(ur.GetUser(), ur.GetRole(), ur.GetDomain())
}

func (r *rbacInDomainUsecase) DeleteRolesForUserInDomainUC(ur domain.UserRoleInDomain) (bool, error) {
	return r.enforcer.DeleteRolesForUserInDomain(ur.GetUser(), ur.GetDomain())
}

func (r *rbacInDomainUsecase) GetDomainsForUserUC(name string) ([]string, error) {
	return r.enforcer.GetDomainsForUser(name)
}

func (r *rbacInDomainUsecase) GetRolesForUserInDomainUC(name string, domain string) []string {
	return r.enforcer.GetRolesForUserInDomain(name, domain)
}

func (r *rbacInDomainUsecase) GetRolesInDomainsForUserUC(name string) (map[string][]string, error) {
	domains, err := r.enforcer.GetDomainsForUser(name)
	if err != nil {
		return nil, err
	}
	result := make(map[string][]string)
	for _, domain := range domains {
		roles := r.enforcer.GetRolesForUserInDomain(name, domain)
		result[domain] = roles
	}
	return result, nil
}
