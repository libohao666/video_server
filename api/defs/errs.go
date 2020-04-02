package defs

type Err struct {
    Error string `json:"error"`
    ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
    HttpSC int
    Error Err
}

var {
    ErrorRequestBodyParaseFailed = ErrResponse{HttpSC: 400, Err{Error: "Request body is not correct", ErrorCode :"001"}}
    ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed.", ErrorCode :"002"}}
}
