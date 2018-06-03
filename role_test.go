package harvest

import "testing"

func testRole(t *testing.T) *Role {
	a := testAPI()
	roleResponse := mockResponse("roles", "role-example.json")
	a.BaseURL = roleResponse.URL
	role, err := a.GetRole(1782974, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if role.Name != "Founder" {
		t.Errorf("Incorrect Role Name '%s'", role.Name)
	}
	if len(role.UserIDs) != 1 {
		t.Errorf("Incorrect number of User IDs '%v'", len(role.UserIDs))
	}
	if role.UserIDs[0] != 8083365 {
		t.Errorf("Incorrect User ID '%v'", role.UserIDs[0])
	}
	return role
}

func TestGetRole(t *testing.T) {
	role := testRole(t)
	if role == nil {
		t.Fatal("testRole() returned nil instead of role")
	}
	if role.ID != 1782974 {
		t.Errorf("Incorrect role ID '%v'", role.ID)
	}
}

func TestGetRoles(t *testing.T) {
	a := testAPI()
	roleResponse := mockResponse("roles", "roles-example.json")
	a.BaseURL = roleResponse.URL
	roles, err := a.GetRoles(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(roles) != 3 {
		t.Errorf("Incorrect number of roles '%v'", len(roles))
	}
	if roles[2].Name != "Designer" {
		t.Errorf("Incorrect Role Name '%s'", roles[2].Name)
	}
	if roles[0].UserIDs[0] != 8083365 {
		t.Errorf("Incorrect User ID '%v'", roles[0].UserIDs[0])
	}
}
