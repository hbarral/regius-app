package handlers

import (
	"net/http"
	"net/url"
	"regius-app/data"
	"time"

	"github.com/CloudyKit/jet/v6"
	"gitlab.com/hbarral/regius"
	"gitlab.com/hbarral/regius/filesystems"
	"gitlab.com/hbarral/regius/filesystems/miniofilesystem"
)

type Handlers struct {
	App    *regius.Regius
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering", err)
	}
}

func (h *Handlers) ListFS(w http.ResponseWriter, r *http.Request) {
	var fs filesystems.FS
	var list []filesystems.Listing

	fsType := ""
	if r.URL.Query().Get("fs-type") != "" {
		fsType = r.URL.Query().Get("fs-type")
	}

	curPath := "/"
	if r.URL.Query().Get("cur_path") != "" {
		curPath = r.URL.Query().Get("cur_path")
		curPath, _ = url.QueryUnescape(curPath)
		if curPath == "/" {
			curPath = ""
		}
	}

	if fsType != "" {
		switch fsType {
		case "MINIO":
			f := h.App.FileSystems["MINIO"].(miniofilesystem.Minio)
			fs = &f
			fsType = "MINIO"
		}

		l, err := fs.List(curPath)
		if err != nil {
			h.App.ErrorLog.Println(err)
			return
		}

		list = l
	}

	vars := make(jet.VarMap)
	vars.Set("list", list)
	vars.Set("fs_type", fsType)
	vars.Set("cur_path", curPath)
	err := h.render(w, r, "list-fs", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return
	}
}
