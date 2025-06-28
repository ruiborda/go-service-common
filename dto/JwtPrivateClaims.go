package dto

type JwtPrivateClaims struct {
	Email         string   `json:"email"`
	Roles         []string `json:"roles"`
	PermissionIds []int    `json:"permissionIds"`
}
