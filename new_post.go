package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
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
		fmt.Println(invalidTitleError)
		return invalidTitleError
	}
	if req.Content.Description == "" {
		fmt.Println(invalidDescriptionError)
		return invalidDescriptionError
	}

	fileName := time.Now().Format("2006-01-02") + "-" + sanitizeFileName(req.Content.Title) + ".htm"
	fp := filepath.Join(getConfig().HugoRootDir, "content/article", fileName)
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

var re = regexp.MustCompile(`[.<>:"/\\|?*]`)

func sanitizeFileName(fileName string) string {
	safeFileName := re.ReplaceAllString(fileName, "-")
	safeFileName = strings.TrimSpace(safeFileName)
	safeFileName = strings.ReplaceAll(safeFileName, " ", "-")

	return safeFileName
}
