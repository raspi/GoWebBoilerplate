package useraccount

import (
	webfw "github.com/gin-gonic/gin"
	"net/http"
)

// Register new user account
func ApiRegister(ctx *webfw.Context) {

	resp := map[string]string{
		"hello": "world",
	}

	ctx.JSON(http.StatusOK, resp)
}

// Login to user account
func ApiLogin(ctx *webfw.Context) {
	resp := map[string]string{
		"hello": "world",
	}

	ctx.JSON(http.StatusOK, resp)
}

// Logout to user account
func ApiLogout(ctx *webfw.Context) {
	resp := map[string]string{
		"hello": "world",
	}

	ctx.JSON(http.StatusOK, resp)
}


// Reset password
func ApiResetPassword(ctx *webfw.Context) {

	// Send email

	resp := map[string]string{
		"hello": "world",
	}

	ctx.JSON(http.StatusOK, resp)
}

// Reset password phase 2: confirm
func ApiResetPasswordConfirm(ctx *webfw.Context) {

	// Verify token

	resp := map[string]string{
		"hello": "world",
	}

	ctx.JSON(http.StatusOK, resp)
}
