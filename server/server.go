package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// NewServer 返回一个新的negroni实例
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{IndentJSON: true})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// 初始化服务器路由
func initRoutes(mx *mux.Router, formatter *render.Render) {
	// 访问"/unknown" 返回501
	mx.HandleFunc("/unknown", notImplementedHandler()).Methods(http.MethodGet)

	// 接受post的表单
	mx.HandleFunc("/", postHandler(formatter)).Methods(http.MethodPost)

	// 服务静态文件
	mx.PathPrefix("/").Handler(
		http.StripPrefix("/", http.FileServer(http.Dir("./static")))).
		Methods(http.MethodGet)
}

func notImplementedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, "501 page not implemented", http.StatusNotImplemented)
	}
}

func postHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		data := struct {
			Name     string `json:"name"`
			Birthday string `json:"birthday"`
		}{req.FormValue("name"), req.FormValue("birthday")}
		formatter.JSON(w, http.StatusCreated, data)
	}
}
