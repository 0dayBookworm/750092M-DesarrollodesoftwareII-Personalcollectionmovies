package domain

// Estructura peticion.
type LoginRequest struct {
    Username string `form:"Username"`
    Password string `form:"Password"`
}
// Estructura respuesta.
type LoginResponse struct {
    Status string
    Message string
}