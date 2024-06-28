package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regius-app/data"
	"time"

	"github.com/CloudyKit/jet/v6"
	"gitlab.com/hbarral/regius"
	"gitlab.com/hbarral/regius/filesystems"
	"gitlab.com/hbarral/regius/filesystems/miniofilesystem"
	"gitlab.com/hbarral/regius/filesystems/s3filesystem"
	"gitlab.com/hbarral/regius/filesystems/sftpfilesystem"
	"gitlab.com/hbarral/regius/filesystems/webdavfilesystem"
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

		case "SFTP":
			f := h.App.FileSystems["SFTP"].(sftpfilesystem.SFTP)
			fs = &f
			fsType = "SFTP"

		case "WebDAV":
			f := h.App.FileSystems["WebDAV"].(webdavfilesystem.WebDAV)
			fs = &f
			fsType = "WebDAV"
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
	}
}

func (h *Handlers) UploadToFS(w http.ResponseWriter, r *http.Request) {
	fsType := r.URL.Query().Get("type")

	vars := make(jet.VarMap)
	vars.Set("fs_type", fsType)

	err := h.render(w, r, "upload", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostUploadToFS(w http.ResponseWriter, r *http.Request) {
	fieldName, err := getFileToUpload(r, "form-file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uploadType := r.Form.Get("upload-type")

	switch uploadType {
	case "MINIO":
		fs := h.App.FileSystems["MINIO"].(miniofilesystem.Minio)
		err = fs.Put(fieldName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "SFTP":
		fs := h.App.FileSystems["SFTP"].(sftpfilesystem.SFTP)
		err = fs.Put(fieldName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "WebDAV":
		fs := h.App.FileSystems["WebDAV"].(webdavfilesystem.WebDAV)
		err = fs.Put(fieldName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	case "S3":
		fs := h.App.FileSystems["S3"].(s3filesystem.S3)
		err = fs.Put(fieldName, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	h.App.Session.Put(r.Context(), "flash", "File uploaded!")
	http.Redirect(w, r, "/files/upload?type="+uploadType, http.StatusSeeOther)
}

func getFileToUpload(r *http.Request, fieldName string) (string, error) {
	_ = r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile(fieldName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	destination, err := os.Create(fmt.Sprintf("./tmp/%s", handler.Filename))
	if err != nil {
		return "", err
	}
	defer destination.Close()

	_, err = io.Copy(destination, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("./tmp/%s", handler.Filename), nil
}

func (h *Handlers) DeleteFromFS(w http.ResponseWriter, r *http.Request) {
	var fs filesystems.FS
	fsType := r.URL.Query().Get("fs_type")
	file := r.URL.Query().Get("file")

	switch fsType {
	case "MINIO":
		f := h.App.FileSystems["MINIO"].(miniofilesystem.Minio)
		fs = &f

	case "SFTP":
		f := h.App.FileSystems["SFTP"].(sftpfilesystem.SFTP)
		fs = &f

	case "WebDAV":
		f := h.App.FileSystems["WebDAV"].(webdavfilesystem.WebDAV)
		fs = &f
	}

	deleted := fs.Delete([]string{file})
	if deleted {
		h.App.Session.Put(r.Context(), "flash", fmt.Sprintf("%s was deleted!", file))
	} else {
		h.App.Session.Put(r.Context(), "flash", "File not deleted!")
	}

	http.Redirect(w, r, "/list-fs?fs-type="+fsType, http.StatusSeeOther)
}
