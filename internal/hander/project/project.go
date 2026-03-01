package project

import (
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/puriice/httplibs/pkg/json"
	"github.com/puriice/pProject/internal/repository"
	"github.com/puriice/pProject/internal/types"
)

type Handler struct {
	repo repository.ProjectRepository
}

func NewHandler(model repository.ProjectRepository) *Handler {
	return &Handler{
		repo: model,
	}
}

func (h *Handler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /projects", h.handleProjectCreate)
	router.HandleFunc("GET /projects/id/{id}", h.handleProjectQueryByID)
	router.HandleFunc("GET /projects/name/{name}", h.handleProjectQueryByName)
	router.HandleFunc("PATCH /projects/{id}", h.handleProjectUpdating)
	router.HandleFunc("DELETE /projects/{id}", h.handleProjectDeletion)
}

func (h *Handler) handleProjectCreate(w http.ResponseWriter, r *http.Request) {
	var payload types.ProjectPayload

	err := json.ParseJSON(r, &payload)

	if err != nil {
		if errors.Is(err, json.MissingBody) {
			http.Error(w, "Missing Body", http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response, err := h.repo.CreateProject(r.Context(), &payload)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				w.WriteHeader(http.StatusConflict)
				return
			}
		} else {
			log.Print(err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.SendJSON(w, http.StatusCreated, response)
}

func (h *Handler) handleProjectQueryByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	project, err := h.repo.QueryProjectByID(r.Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	json.SendJSON(w, http.StatusOK, project)
}

func (h *Handler) handleProjectQueryByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	project, err := h.repo.QueryProjectByName(r.Context(), name)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	json.SendJSON(w, http.StatusOK, project)
}

func (h *Handler) handleProjectUpdating(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload := new(types.ProjectPayload)

	err := json.ParseJSON(r, payload)

	if err != nil {
		if errors.Is(err, json.MissingBody) {
			http.Error(w, "Missing Body", http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	err = h.repo.UpdateProject(r.Context(), id, payload)

	if err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if errors.Is(err, types.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}

func (h *Handler) handleProjectDeletion(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.repo.DeleteProject(r.Context(), id)

	if err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if errors.Is(err, types.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}
