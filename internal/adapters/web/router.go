package web

import (
	"net/http"

	"github.com/smithshelke/flur/internal/adapters/web/blocks"
	"github.com/smithshelke/flur/internal/adapters/web/connection"
	"github.com/smithshelke/flur/internal/adapters/web/user"
	"github.com/smithshelke/flur/internal/adapters/web/workflow"
)

func (a *App) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /health", health)
	core := a.Core

	workerflowController := workflow.NewWorkflowController(core.GraphAPI)
	AddRoute(router, "POST /workflow", workerflowController.Create)

	userController := user.NewUserController(core.UserAPI)
	AddRoute(router, "GET /users", userController.List)
	AddRoute(router, "POST /users", userController.Create)

	blocksController := blocks.NewBlocksController(core.GraphAPI)
	AddRoute(router, "POST /blocks", blocksController.Create)

	connectionController := connection.NewConnectionController(core.GraphAPI)
	AddRoute(router, "POST /connections", connectionController.Create)
}

func AddRoute(router *http.ServeMux, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(pattern, Chain(
		handler,
		LoggingMiddleware(),
		CommonMiddleware(),
	))
}
