package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"{{ module }}/domain"
)

type exampleHandler struct {
	router   chi.Router

	exampleService domain.ExampleService
}

func newExampleHandler() *exampleHandler {
	h := &exampleHandler{router: chi.NewRouter()}
	h.router.Get("/", h.handleGetAll())
	h.router.Get("/{exampleId}", h.handleGet())
	h.router.Post("/", h.handleCreate())
	h.router.Put("/", h.handleUpdate())
	h.router.Delete("/{exampleId}", h.handleDelete())
	return h
}

func (h *exampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	txn := h.newRelic.StartTransaction(r.Method + " " + r.RequestURI)
	defer txn.End()

	w = txn.SetWebResponse(w)
	txn.SetWebRequestHTTP(r)

	h.router.ServeHTTP(w, r)
}

func (h *exampleHandler) handleGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		efc, err := h.exampleService.Examples()
		if err != nil {
			respondErr(w, http.StatusBadRequest, err)
			return
		}

		respond(w, http.StatusOK, efc)
		return
	}
}

func (h *exampleHandler) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "exampleId"))

		if err != nil {
			respondErr(w, http.StatusBadRequest, err)
			return
		}
		example, err := h.exampleService.Example(id)
		if err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}
		respond(w, http.StatusOK, example)
		return
	}
}

func (h *exampleHandler) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var example domain.Example
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			respondErr(w, http.StatusBadRequest, err)
			return
		}

		if err := json.Unmarshal(body, &example); err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}

		createdExample, err := h.exampleService.CreateExample(example)
		if err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}

		respond(w, http.StatusOK, createdEFC)
		return
	}
}

func (h *exampleHandler) handleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var example domain.Example
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			respondErr(w, http.StatusBadRequest, err)
			return
		}

		if err := json.Unmarshal(body, &example); err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}

		updatedExample, err := h.exampleService.UpdateExample(example)
		if err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}

		respond(w, http.StatusOK, updatedEFC)
		return
	}
}

func (h *exampleHandler) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "exampleId"))

		if err != nil {
			respondErr(w, http.StatusBadRequest, err)
			return
		}
		err = h.exampleService.DeleteExample(id)
		if err != nil {
			respondErr(w, http.StatusInternalServerError, err)
			return
		}
		respond(w, http.StatusOK, "EFC Deleted with id: "+string(id))
		return
	}
}
