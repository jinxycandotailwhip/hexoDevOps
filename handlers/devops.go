package handlers

import (
	"encoding/json"
	"hexoDevOps/impl"
	"hexoDevOps/requests"
	"io/ioutil"
	"log"
	"net/http"
)

func Regenerate(w http.ResponseWriter, r *http.Request) {
	err := impl.HexoCommandGenerate()
	if err != nil {
		log.Fatal("read body failed", err)
		return
	}
	log.Println("hexo generate succeed")
}

func NewPage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("read body failed", err)
		return
	}
	pageName := requests.DevOpsPostReq{}
	err = json.Unmarshal(body, &pageName)
	if err != nil {
		log.Fatal("unmarshal body failed", err)
		return
	}
	err = impl.HexoCommandNewPage(pageName.Data.(string))
	if err != nil {
		log.Fatal("hexo new page failed", err)
		return
	}
	log.Println("hexo new page succeed")
}
