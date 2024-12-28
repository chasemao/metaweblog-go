package main

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type GetUsersBlogsReq struct {
	AppKey   string
	UserName string
	Password string
}

type Blog struct {
	Blogid   string `xml:"blogid"`
	URL      string `xml:"url"`
	BlogName string `xml:"blogName"`
	IsAdmin  bool   `xml:"isAdmin"`
	XMLRPC   string `xml:"xmlrpc"`
}

type GetUsersBlogsResp struct {
	Blogs []Blog
}

func (a *apiImpl) GetUsersBlogs(r *http.Request, req *GetUsersBlogsReq, resp *GetUsersBlogsResp) error {
	fmt.Println("GetUsersBlogs called")
	fmt.Println(spew.Sdump(req))
	if err := auth(req.UserName, req.Password); err != nil {
		fmt.Println(err)
		return err
	}
	resp.Blogs = []Blog{
		{
			Blogid:   "",
			URL:      getConfig().BlogURL,
			BlogName: getConfig().BlogTitle,
			IsAdmin:  true,
		},
	}
	fmt.Println(spew.Sdump(resp))
	fmt.Println("GetUsersBlogs success")
	return nil
}
