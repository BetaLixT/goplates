package user

import "github.com/betalixt/gorr"



var (
  ROLE_INVALID_ERROR_CODE = gorr.ErrorCode{
    Code: 2000,
    Message: "USER_ROLE_INVALID",
  }
  PROVIDER_USER_EXISTS_ERROR_CODE = gorr.ErrorCode{
    Code: 2001,
    Message: "USER_PROVIDER_USER_EXISTS",
  }
)
