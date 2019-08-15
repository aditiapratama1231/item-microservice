package request

import "time"

type User struct {
	ID             int64     `json:"id"`
	MerchantID     int64     `json:"merchant_id"`
	OutletID       int64     `json:"outlet_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password,omitempty"`
	Passcode       string    `json:"passcode,omitempty"`
	Mobile         string    `json:"passcode"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Title          string    `json:"title"`
	Image          string    `json:"image"`
	Access         int8      `json:"access"`
	ApiToken       string    `json:"api_token,omitempty"`
	RememberToken  string    `json:"remember_token,omitempty"`
	TokenGenerated time.Time `json:"token_generated,omitempty"`
	TokenExpired   time.Time `json:"token_expired,omitempty"`
	Level          int8      `json:"level"`
	Status         string    `json:"status"`
	CreatedBy      int64     `json:"created_by,omitempty"`
	UpdatedBy      int64     `json:"updated_by,omitempty"`
	DeletedBy      int64     `json:"deleted_by,omitempty"`
}
