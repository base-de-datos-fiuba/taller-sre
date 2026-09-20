package httpapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"bytemarket/backend/internal/products"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db       *pgxpool.Pool
	products *products.Repository
}

func NewHandler(db *pgxpool.Pool) http.Handler {
	h := &Handler{db: db, products: products.NewRepository(db)}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/products", h.listProducts)
	mux.HandleFunc("PATCH /api/products/{id}/price", h.changePrice)
	mux.HandleFunc("GET /api/health", h.health)
	return withCORS(mux)
}

var validPrice = regexp.MustCompile(`^(0|[1-9][0-9]{0,9})(\.[0-9]{1,2})?$`)

func (h *Handler) changePrice(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, errors.New("invalid product id"))
		return
	}

	var input struct {
		Price string `json:"price"`
	}
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1024))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, errors.New("expected JSON with a price string"))
		return
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		writeError(w, http.StatusBadRequest, errors.New("expected one JSON object"))
		return
	}
	if !validPrice.MatchString(input.Price) {
		writeError(w, http.StatusBadRequest, errors.New("price must be a nonnegative decimal string with up to two fractional digits"))
		return
	}

	updated, err := h.products.ChangePrice(r.Context(), id, input.Price)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if !updated {
		writeError(w, http.StatusNotFound, errors.New("product not found"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) listProducts(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	var (
		items []products.Product
		err   error
	)
	if search == "" {
		items, err = h.products.List(r.Context())
	} else {
		items, err = h.products.Search(r.Context(), search)
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	var count int
	if err := h.db.QueryRow(r.Context(), "SELECT COUNT(*) FROM product").Scan(&count); err != nil {
		writeError(w, http.StatusServiceUnavailable, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"status": "online", "products": count})
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"status": "error", "error": err.Error()})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		next.ServeHTTP(w, r)
	})
}
