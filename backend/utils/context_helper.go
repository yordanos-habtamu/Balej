package utils

import (
	"context"
	"fmt"
)

// GetUserIDFromContext retrieves the user ID injected by the AuthMiddleware
func GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value("user_id").(uint)
	if !ok {
		return 0, fmt.Errorf("unauthorized: missing user_id in context")
	}
	return userID, nil
}

// GetRoleFromContext retrieves the user role injected by the AuthMiddleware
func GetRoleFromContext(ctx context.Context) (string, error) {
	role, ok := ctx.Value("role").(string)
	if !ok {
		return "", fmt.Errorf("unauthorized: missing role in context")
	}
	return role, nil
}
