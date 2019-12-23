package manager

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"main/cart-service/repository"
	"main/cart-service/structs"
	"net/http"
	"strconv"
)

type Manager struct {
	Config structs.Config
	Router *mux.Router
	Repo   repository.IRepository
}

func (m *Manager) HandleRequest(method, route string, handle structs.HandlerFunc) {
	m.Router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		var res structs.IMessage

		body, err := ioutil.ReadAll(r.Body)
		userId, err := strconv.ParseInt(r.Header.Get("X-User-Id"), 10, 64)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(401)
			return
		}

		client := structs.Client{
			UserId: userId,
		}

		res, err = handle(client, body)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			return
		}

		_, err = fmt.Fprint(w, res.GetJson())
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(500)
			return
		}
	}).Methods(method)
}

func (m *Manager) Run() {
	m.Router.Use(m.checkAuthorization)
	m.HandleRequest("POST", "/addToCart", m.AddToCartHandler)
	m.HandleRequest("GET", "/getCart", m.GetCartHandler)
	m.HandleRequest("DELETE", "/cart", m.ClearCartHandler)

	srv := &http.Server{
		Handler: m.Router,
		Addr:    ":" + strconv.FormatInt(m.Config.Port, 10),
	}

	log.Fatal(srv.ListenAndServe())
}

func (m *Manager) checkAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("X-Auth")
		userId := "10" // TODO: here will be getting userId form auth

		if auth != "null" {
			r.Header.Set("X-User-Id", userId)
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(401)
		return
	})
}
