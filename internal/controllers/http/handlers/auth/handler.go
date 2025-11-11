package auth

// authmw "web/internal/controllers/http/middleware/auth"
// authsvc "web/internal/service/auth_service"

type AuthHandler struct {
	// svc    authsvc.Service
	// driver *authmw.Driver
}

// svc authsvc.Service, driver *authmw.Driver
func NewHandler() *AuthHandler {
	return &AuthHandler{} // svc: svc, driver: driver
}
