// Code generated by goctl. DO NOT EDIT.
package types

type Detailrequest struct {
	UserID int64 `json:"user_id"`
}

type Detailresponse struct {
	Message  string `json:"message"`
	UserName string `json:"username"`
	Gender   int    `json:"gender,options=0|1|2,default=0"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	AccessToken  string `json:"accesstoken"`
	AccessExpire int    `json:"accessexpire"` //过期时间
	RefreshAfter int    `json:"refreshafter"`
}

type SignupRequest struct {
	UserName   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Gender     int    `json:"gender,options=0|1|2,default=0"`
}

type SignupResponse struct {
	Message string `json:"message"`
}
