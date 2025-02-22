package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Content struct {
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Categories  []string `xml:"categories"`
}

type NewPostReq struct {
	BlogID    string
	UserName  string
	Password  string
	Content   Content
	IsPublish bool
}

type NewPostResp struct {
	Postid string
}

func (a *apiImpl) NewPost(r *http.Request, req *NewPostReq, resp *NewPostResp) error {
	fmt.Println(spew.Sdump(req))
	if err := auth(req.UserName, req.Password); err != nil {
		fmt.Println(err)
		return err
	}
	if req.Content.Title == "" {
		fmt.Println(errInvalidTitle)
		return errInvalidTitle
	}
	if req.Content.Description == "" {
		fmt.Println(errInvalidDescription)
		return errInvalidDescription
	}

	fileName := time.Now().Format("2006-01-02") + "-" + sanitizeFileName(req.Content.Title) + ".htm"
	fp := filepath.Join(getConfig().BlogDir, fileName)
	file, err := os.Create(fp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(req.Content.Description)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(spew.Sdump(resp))
	fmt.Println("NewPost success")
	return nil
}

func sanitizeFileName(fileName string) string {
	fileName = strings.TrimSpace(fileName)
	fileName = strings.ReplaceAll(fileName, " ", "-")

	return fileName
}
