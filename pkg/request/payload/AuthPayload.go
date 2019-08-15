package request

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginToken struct {
	AccessToken string      `json:"access_token,omitempty"`
	User        interface{} `json:"user,omitempty"`
}

type LoginRequest struct {
	Data Login `json:"data"`
}

type LoginResponse struct {
	Data       LoginToken `json:"data,omitempty"`
	Message    string     `json:"message"`
	StatusCode int32      `json:"status_code"`
	Err        bool       `json:"error,omitempty"`
}

type RegisterAttributes struct {
}

type TokenInstropectionRequest struct {
	Token string `json:"token"`
}

type TokenInstropectionResponse struct {
	Issuer     string `json:"issuer,omitempty"`
	Activate   bool   `json:"active,omitempty"`
	Sub        string `json:"sub,omitempty"`
	Exp        int64  `json:"exp,omitempty"`
	UserID     int64  `json:"user_id,omitempty"`
	Error      string `json:"error,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int64  `json:"status_code,omitempty"`
}
