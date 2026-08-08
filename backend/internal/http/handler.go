package httpapi

import (
	"encoding/json"
	"net/http"

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
	mux.HandleFunc("GET /api/health", h.health)
	return withCORS(mux)
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
		items, err = h.products.SearchVulnerable(r.Context(), search)
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
