package api

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
	"wbTech/internal/db"
	"wbTech/internal/domain"
)

const orderKey = "order"

type Api struct {
	rtr *chi.Mux
	csh *db.Cache
	srv *http.Server
}

func NewApi(ctx context.Context, csh *db.Cache) *Api {
	api := Api{}
	newCtx := ctx
	api.Init(newCtx, csh)
	return &api
}

func (a *Api) Init(ctx context.Context, cache *db.Cache) {
	a.csh = cache
	a.rtr = chi.NewRouter()
	a.rtr.Route("/orders", func(r chi.Router) {
		r.Route("/{OrderUid}", func(r chi.Router) {
			r.Use(a.getOrderByID)
			r.Get("/", a.Index)
		})
	})
	err := http.ListenAndServe(":8080", a.rtr)
	if err != nil {
		log.Fatal("Error listen and serve", err)
	}
	fmt.Println("server start on :8080")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit from api server")
			return
		default:

		}
	}
	return
}

func (a *Api) Index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	order, ok := ctx.Value(orderKey).(*domain.NewOrder)
	if !ok {
		log.Printf("%v: getOrder(): ошибка приведения интерфейса к типу *domain.NewOrder\n", ok)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity) // 422
		return
	}
	t, err := template.ParseFiles("ui/templates/order.html")
	if err != nil {
		log.Fatal("Error parse file html:", err)
		http.Error(w, "Error", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "order.html", order)
	if err != nil {
		log.Printf("%v Error execute file html:\n", err)
		return
	}
}

func (a *Api) getOrderByID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		orderIDstr := chi.URLParam(request, "OrderUid")
		c, err := a.csh.Get(orderIDstr)
		if err != nil {
			panic(err.Error())
		}
		ctx := context.WithValue(request.Context(), orderKey, c)
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
