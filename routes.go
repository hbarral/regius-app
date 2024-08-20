package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/hbarral/regius/filesystems/miniofilesystem"
)

func (a *application) routes() *chi.Mux {
	// middlewares
	a.use(a.Middleware.CheckRemember)

	// routes
	a.get("/users/signin", a.Handlers.UserSignIn)
	a.post("/users/signin", a.Handlers.PostUserSignIn)
	a.get("/users/signout", a.Handlers.SignOut)
	a.get("/auth/{provider}", a.Handlers.SocialSignin)
	a.get("/auth/{provider}/callback", a.Handlers.SocialCallback)
	a.get("/", a.Handlers.Home)
	a.get("/upload", a.Handlers.RegiusUpload)
	a.post("/upload", a.Handlers.PostRegiusUpload)
	a.get("/list-fs", a.Handlers.ListFS)
	a.get("/files/upload", a.Handlers.UploadToFS)
	a.post("/files/upload", a.Handlers.PostUploadToFS)
	a.get("/delete-from-fs", a.Handlers.DeleteFromFS)
	a.get("/test-minio", func(w http.ResponseWriter, r *http.Request) {
		f := a.App.FileSystems["MINIO"].(miniofilesystem.Minio)

		files, err := f.List("")
		if err != nil {
			log.Println(err)
			return
		}

		for _, file := range files {
			log.Println(file.Key)
		}
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
