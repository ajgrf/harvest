package harvest

import (
	"fmt"
	"time"
)

type Role struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserIDs   []int64   `json:"user_ids"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *API) GetRole(roleID int64, args Arguments) (role *Role, err error) {
	role = &Role{}
	path := fmt.Sprintf("/roles/%v", roleID)
	err = a.Get(path, args, &role)
	return role, err
}

func (a *API) GetRoles(args Arguments) (roles []*Role, err error) {
	var rolesResponse struct {
		PagedResponse
		Roles []*Role `json:"roles"`
	}
	roles = make([]*Role, 0)
	err = a.GetPaginated("/roles", args, &rolesResponse, func() {
		for _, r := range rolesResponse.Roles {
			roles = append(roles, r)
		}
		rolesResponse.Roles = make([]*Role, 0)
	})
	return roles, err
}
