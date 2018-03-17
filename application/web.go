package application

import (
	webfw "github.com/gin-gonic/gin"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"encoding/json"
	"../useraccount"
)

func AuthMiddleware(ctx *webfw.Context) {
	log.Printf("Authentication required request for %s", ctx.Request.RequestURI)

	for k, v := range ctx.Request.Header {
		log.Printf("Header field %q, Value %q\n", k, v)
	}

	resp := map[string]string{
		"error": "Auth required",
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)

	ctx.Next()
}

func AdminRoleMiddleware(ctx *webfw.Context) {
	log.Printf("Admin rights required request for %s", ctx.Request.RequestURI)

	for k, v := range ctx.Request.Header {
		log.Printf("Header field %q, Value %q\n", k, v)
	}

	resp := map[string]string{
		"error": "Auth required",
	}

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)

	ctx.Next()
}

func JsonRequiredMiddleware(ctx *webfw.Context) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")

	var err error

	log.Printf("Checking required JSON request for %s", ctx.Request.RequestURI)

	data, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		return
	}

	if len(data) == 0 {
		data = []byte("{}")
	}

	var js json.RawMessage

	err = json.Unmarshal(data, &js)

	if err != nil {
		log.Printf("Invalid JSON")
		return
	}

	ctx.Next()
}

func NewApplication(engine *webfw.Engine) {
	engine.LoadHTMLGlob("**/*.html")

	engine.Use(webfw.Logger())
	engine.Use(webfw.Recovery())

	engine.NoRoute(DefaultPageNotFound)

	// Default redirect root to default language
	engine.GET("", DefaultRedirect)

	// Everything is multilingual by default
	base := engine.Group(":language")

	// Frontend (HTML + JS):
	base.GET("/", useraccount.Home)
	base.GET("/login", useraccount.Login)
	base.GET("/register", useraccount.Register)
	base.GET("/reset-password", useraccount.ResetPassword)
	base.GET("/reset-password/:token", useraccount.ResetPasswordConfirm)

	loggedIn := base.Group("")
	loggedIn.GET("/logout", useraccount.Logout)


	admin := base.Group("/admin")
	admin.Use(AdminRoleMiddleware)
	admin.GET("", useraccount.Home)


	// Web socket
	base.GET("/ws", WebSocketUpgrader)

	// Backend (API), where actually everything happens:
	apiUrlBase := base.Group("/api/v1")
	apiUrlBase.Use(JsonRequiredMiddleware)

	public := apiUrlBase.Group("/")
	public.GET("login", useraccount.ApiLogin)
	public.GET("register", useraccount.ApiRegister)
	public.GET("password/reset", useraccount.ApiResetPassword)
	public.GET("password/confirm/:token", useraccount.ApiResetPasswordConfirm)

	authNeeded := apiUrlBase.Group("/")
	authNeeded.Use(AuthMiddleware)
	authNeeded.GET("logout", useraccount.ApiLogout)

}

func DefaultRedirect(ctx *webfw.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "/en/")
}

func DefaultPageNotFound(ctx *webfw.Context) {
	ctx.HTML(http.StatusNotFound, "notfound.html", webfw.H{})
	ctx.Redirect(http.StatusNotFound, "/en/")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Upgrade HTTP(s) -> WebSocket
func WebSocketUpgrader(ctx *webfw.Context) {
	wshandler(ctx.Writer, ctx.Request)
}


func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("Failed to upgrade connection")
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		conn.WriteMessage(t, msg)

	}

}
