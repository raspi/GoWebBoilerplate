package useraccount

import (
	webfw "github.com/gin-gonic/gin"
	"net/http"
)

// User's home
func Home(ctx *webfw.Context) {
	ctx.HTML(http.StatusOK, "home", webfw.H{})
}

// Register new user account
func Register(ctx *webfw.Context) {
	ctx.HTML(http.StatusOK, "register", webfw.H{})
}

// Login to user account
func Login(ctx *webfw.Context) {
	ctx.HTML(http.StatusOK, "login", webfw.H{})
}

// Log out user
func Logout(ctx *webfw.Context) {
	ctx.HTML(http.StatusOK, "logout", webfw.H{})
}

// Reset password for user account
func ResetPassword(ctx *webfw.Context) {
	ctx.HTML(http.StatusOK, "reset_password", webfw.H{})
}

// Reset password for user account
func ResetPasswordConfirm(ctx *webfw.Context) {
	ctx.HTML(http.StatusOK, "reset_password", webfw.H{})
}
