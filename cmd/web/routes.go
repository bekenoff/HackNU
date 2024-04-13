package main

import (
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	mux.Post("/api/create-module-info", dynamicMiddleware.ThenFunc(app.createModuleInfo))
	mux.Get("/api/get-module-info/:id", dynamicMiddleware.ThenFunc(app.getModuleInfo)) //http://localhost:4001/api/get-module-info/2
	mux.Put("/api/update-module-info", dynamicMiddleware.ThenFunc(app.updateModuleInfo))
	mux.Del("/api/delete-module-info/:id", dynamicMiddleware.ThenFunc(app.deleteModuleInfo))

	return standardMiddleware.Then(mux)
}
