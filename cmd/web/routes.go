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

	// Grammar
	mux.Post("/api/correct-text", dynamicMiddleware.ThenFunc(CorrectHandler)) //http://localhost:4000/api/correct-text?text="менин атым Алим"

	// Synonym
	mux.Post("/api/synonym", dynamicMiddleware.ThenFunc(synonymHandler))

	// Clients
	mux.Post("/api/create/client", dynamicMiddleware.ThenFunc(app.signupClient))
	mux.Post("/api/auth/client", dynamicMiddleware.ThenFunc(app.loginClient))

	// Progress
	mux.Post("/api/progress/create", dynamicMiddleware.ThenFunc(app.createProgress))

	mux.Get("/api/progress/get", standardMiddleware.ThenFunc(app.getProgress)) // http://localhost:4000/api/progress/get?id=123

	mux.Del("/api/progress/delete", dynamicMiddleware.ThenFunc(app.deleteProgress)) // http://localhost:4000/api/progress/delete?id=123

	mux.Patch("/api/progress/update-tests/:id", dynamicMiddleware.ThenFunc(app.updateProgressTests))       // http://localhost:4000/api/progress/update-tests/17
	mux.Patch("/api/progress/update-films/:id", dynamicMiddleware.ThenFunc(app.updateProgressFilms))       // http://localhost:4000/api/progress/update-films/17
	mux.Patch("/api/progress/update-meetings/:id", dynamicMiddleware.ThenFunc(app.updateProgressMeetings)) // http://localhost:4000/api/progress/update-meetings/17

	// Video
	mux.Post("/api/video/upload/:id", dynamicMiddleware.ThenFunc(app.fileUpload))
	mux.Get("/api/videos/get", standardMiddleware.ThenFunc(app.showAllVideos))
	mux.Post("/api/video/likes", standardMiddleware.ThenFunc(app.incrementLike)) // http://localhost:4000/api/video/likes?video_id=1&client_id=8
	mux.Get("/api/video/sum_likes", standardMiddleware.ThenFunc(app.getVideosWithLikes))

	return standardMiddleware.Then(mux)
}
