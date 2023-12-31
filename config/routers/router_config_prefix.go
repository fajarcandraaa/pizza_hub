package routers

import "github.com/gorilla/mux"

type PathPrefix struct {
	V1    *mux.Router
	Chef  *mux.Router
	Menu  *mux.Router
	Order *mux.Router
}

func RouterConfigPrefix(se *Serve) *PathPrefix {
	var (
		api   = se.Router.PathPrefix("/api").Subrouter()
		v1    = api.PathPrefix("/v1").Subrouter()
		chef  = v1.PathPrefix("/chef").Subrouter()
		menu  = v1.PathPrefix("/menus").Subrouter()
		order = v1.PathPrefix("/orders").Subrouter()
	)

	result := &PathPrefix{
		V1:    v1,
		Chef:  chef,
		Menu:  menu,
		Order: order,
	}

	return result
}
