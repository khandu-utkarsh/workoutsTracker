package handlers

import (
	"net/http"
	"workout_app_backend/internal/middleware"
	"workout_app_backend/internal/models"
)

// AuthHandler handles HTTP requests related to authentication.
type AuthHandler struct {
	authMiddleware *middleware.AuthMiddleware
	userModel      *models.UserModel
}

// GetAuthHandlerInstance creates a new AuthHandler instance.
func GetAuthHandlerInstance(authMiddleware *middleware.AuthMiddleware, userModel *models.UserModel) *AuthHandler {
	return &AuthHandler{
		authMiddleware: authMiddleware,
		userModel:      userModel,
	}
}

// GoogleLoginRequest represents the request body for Google login
type GoogleLoginRequest struct {
	GoogleToken string `json:"googleToken"`
}

// GoogleLogin handles POST /api/auth/google
func (h *AuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	logRequest("GoogleLogin")
	if err := validateHTTPMethod(r, http.MethodPost); err != nil {
		handleHTTPError(w, err)
		return
	}

	var req GoogleLoginRequest
	if err := decodeJSONBody(r, &req); err != nil {
		handleHTTPError(w, err)
		return
	}

	if req.GoogleToken == "" {
		respondWithError(w, "Google token is required", http.StatusBadRequest)
		return
	}

	// 1. Validate Google JWT token
	googleClaims, err := h.authMiddleware.ValidateGoogleJWT(req.GoogleToken)
	if err != nil {
		respondWithError(w, "Invalid Google token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	ctx := r.Context()

	// 2. Find or create user in database
	user, err := h.userModel.GetByEmail(ctx, googleClaims.Email)
	if err != nil {
		if err == models.ErrUserNotFound {
			// Create new user
			newUser := &models.User{
				Email: googleClaims.Email,
			}

			userID, err := h.userModel.Create(ctx, newUser)
			if err != nil {
				respondWithError(w, "Failed to create user", http.StatusInternalServerError)
				return
			}

			// Get the created user
			user, err = h.userModel.Get(ctx, userID)
			if err != nil {
				respondWithError(w, "Failed to get created user", http.StatusInternalServerError)
				return
			}
		} else {
			respondWithError(w, "Database error", http.StatusInternalServerError)
			return
		}
	}

	// 3. Create JWT token with user information
	token, err := h.authMiddleware.CreateJWT(int(user.ID), user.Email, googleClaims.Name, googleClaims.Picture, googleClaims.Sub)
	if err != nil {
		respondWithError(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	// 4. Set secure HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "trackbot_auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true, // HTTPS only in production
		SameSite: http.SameSiteStrictMode,
		MaxAge:   5 * 60, // 5 minutes
		Path:     "/",
	})

	//!Respond with the user from the conext
	userCtx := &middleware.UserContext{
		UserID:  int(user.ID),
		Email:   user.Email,
		Name:    googleClaims.Name,
		Picture: googleClaims.Picture,
	}

	respondWithJSON(w, http.StatusOK, userCtx)
}

// Logout handles POST /api/auth/logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	logRequest("Logout")

	if err := validateHTTPMethod(r, http.MethodPost); err != nil {
		handleHTTPError(w, err)
		return
	}

	// Clear the authentication cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "trackbot_auth_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true, // true in production (HTTPS), false in dev (HTTP)
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1, // Expire immediately
		Path:     "/",
	})

	respondWithJSON(w, http.StatusOK, map[string]bool{"success": true})
}

// Me handles GET /api/auth/me
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	logRequest("Me")

	if err := validateHTTPMethod(r, http.MethodGet); err != nil {
		handleHTTPError(w, err)
		return
	}

	// Get user from context (set by auth middleware)
	user, ok := middleware.GetUserFromContext(r.Context())
	handlerLogger.Println("user obtained from context: ", user) //! Logging the user.
	if !ok {
		respondWithError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
