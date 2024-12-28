package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Media struct {
	Name string
	Type string
	Bits []byte
}

type NewMediaObjectReq struct {
	BlogID   string
	UserName string
	Password string
	Media    Media
}

type NewMediaObjectResp struct {
	MediaResp MediaResp
}

type MediaResp struct {
	ID       string `xml:"id"`
	FileName string `xml:"file"`
	URL      string `xml:"url"`
	Type     string `xml:"type"`
}

func (a *apiImpl) NewMediaObject(r *http.Request, req *NewMediaObjectReq, resp *NewMediaObjectResp) error {
	log.Println(spew.Sdump(req))
	if err := auth(req.UserName, req.Password); err != nil {
		log.Println(err)
		return err
	}

	fileName := fmt.Sprintf("%d", time.Now().UnixMilli()) + "-" + req.Media.Name
	fp := filepath.Join(getConfig().MediaDir, fileName)
	file, err := os.Create(fp)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	_, err = file.Write(req.Media.Bits)
	if err != nil {
		log.Println(err)
		return err
	}

	resp.MediaResp = MediaResp{
		ID:       "xxx",
		FileName: fileName,
		URL:      filepath.Join(getConfig().MediaRelDirForBlogHtml, fileName),
		Type:     req.Media.Type,
	}

	log.Println(spew.Sdump(resp))
	log.Println("NewMediaObject success")
	return nil
}
