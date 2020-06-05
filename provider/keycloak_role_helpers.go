package provider

import "github.com/mrparkers/terraform-provider-keycloak/keycloak"

func getMapOfRealmAndClientRoles(keycloakClient *keycloak.KeycloakClient, realmId string, roleIds []string) (map[string][]*keycloak.Role, error) {
	roles := make(map[string][]*keycloak.Role)

	for _, roleId := range roleIds {
		role, err := keycloakClient.GetRole(realmId, roleId)
		if err != nil {
			return nil, err
		}

		if role.ClientRole {
			roles[role.ClientId] = append(roles[role.ClientId], role)
		} else {
			roles["realm"] = append(roles["realm"], role)
		}
	}

	return roles, nil
}

func removeRoleFromSlice(slice []*keycloak.Role, index int) []*keycloak.Role {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func removeDuplicateRoles(one, two *map[string][]*keycloak.Role) {
	for k := range *one {
		for i1 := 0; i1 < len((*one)[k]); i1++ {
			s1 := (*one)[k][i1]

			for i2 := 0; i2 < len((*two)[k]); i2++ {
				s2 := (*two)[k][i2]

				if s1.Id == s2.Id {
					(*one)[k] = removeRoleFromSlice((*one)[k], i1)
					(*two)[k] = removeRoleFromSlice((*two)[k], i2)

					i1--
					break
				}
			}
		}
	}
}
