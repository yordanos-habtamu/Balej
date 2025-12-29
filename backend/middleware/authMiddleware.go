package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/yordanos-habtamu/GormWithGraphql/auth/helpers"
)


type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}


var OperationRoleConfig = map[string]string{
	"Users": "admin",
	"RemoveUsers":"admin",
	"AddJobOffer": "admin",
	
}


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle CORS preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "failed to read request body"})
			return
		}
		// Restore the body for downstream handlers
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Parse the GraphQL request
		var gqlRequest GraphQLRequest
		if err := json.Unmarshal(body, &gqlRequest); err != nil {
			respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid GraphQL request"})
			return
		}

		// Determine the operation name
		operationName := gqlRequest.OperationName
		if operationName == "" {
			// Fallback: Extract operation name from query if not provided
			query := strings.TrimSpace(gqlRequest.Query)
			if strings.HasPrefix(query, "query") {
				// Extract operation name from query, e.g., "query GetUsers { ... }"
				parts := strings.Fields(query)
				if len(parts) > 1 && parts[0] == "query" {
					operationName = parts[1]
				}
			} else if strings.HasPrefix(query, "mutation") {
				// Extract operation name from mutation, e.g., "mutation LoginUser { ... }"
				parts := strings.Fields(query)
				if len(parts) > 1 && parts[0] == "mutation" {
					operationName = parts[1]
				}
			}
		}

		// Skip authentication for public operations (e.g., LoginUser)
		if operationName == "LoginUser" || strings.Contains(gqlRequest.Query, "mutation LoginUser") {
			next.ServeHTTP(w, r)
			return
		}

		// Check if the operation requires a specific role
		requiredRole, isProtected := OperationRoleConfig[operationName]
		if !isProtected {
			// If the operation is not in the role config, allow it (public operation)
			next.ServeHTTP(w, r)
			return
		}

		// Apply authentication for protected operations
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing auth header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		userID, role, err := helpers.ParseToken(tokenStr)
		if err != nil {
			respondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			return
		}

		// Check if the user has the required role
		if role != requiredRole {
			respondWithJSON(w, http.StatusForbidden, map[string]string{"error": "forbidden: insufficient role"})
			return
		}

		// Add user_id and role to context
		ctx := context.WithValue(r.Context(), "user_id", userID)
		ctx = context.WithValue(ctx, "role", role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper function to send JSON error responses
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}