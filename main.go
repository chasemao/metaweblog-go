package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
)

type apiImpl struct{}

func main() {
	addr := flag.String("a", "localhost:1314", "Specify host addr, default localhost:1314")
	flag.Parse()

	initConfig()

	svr := rpc.NewServer()
	codec := xml.NewCodec()
	codec.RegisterAlias("blogger.getUsersBlogs", "apiImpl.GetUsersBlogs")
	codec.RegisterAlias("metaWeblog.newPost", "apiImpl.NewPost")
	codec.RegisterAlias("metaWeblog.newMediaObject", "apiImpl.NewMediaObject")
	svr.RegisterCodec(codec, "text/xml")
	svr.RegisterService(&apiImpl{}, "apiImpl")

	http.Handle("/metaweblog", newSvrWrapper(svr))

	log.Printf("Starting XML-RPC server on %s/metaweblog", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
