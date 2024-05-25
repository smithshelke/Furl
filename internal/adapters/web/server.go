package web

import (
	"fmt"
	"net/http"

	"github.com/smithshelke/flur/internal/adapters/config"
	"github.com/smithshelke/flur/internal/core"
)

type App struct {
	Config *config.Config
	Core   *core.Core
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (a *App) Initialize() {
	router := http.NewServeMux()
	a.registerRoutes(router)
	err := http.ListenAndServe(a.Config.Server.Addr, router)
	fmt.Printf("Something went wrong: %+v\n", err)
}
