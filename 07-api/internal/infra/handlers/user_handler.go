package handlers

import (
	"encoding/json"
	"github.com/giovannymassuia/learning-go/07-api/internal/dto"
	"github.com/giovannymassuia/learning-go/07-api/internal/entity"
	"github.com/giovannymassuia/learning-go/07-api/internal/infra/database"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB       database.UserInterface
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwtExpiresIn int) *UserHandler {
	return &UserHandler{UserDB: userDB, JwtExpiresIn: jwtExpiresIn}
}

// GetJwt godoc
// @Summary		Get JWT token
// @Description	Get JWT token
// @Tags 		users
// @Accept		json
// @Produce 	json
// @Param 		request 	body 		dto.GetJwtInput 	true 	"user credentials"
// @Success 	200 		{object} 	dto.GetJwtOutput
// @Failure 	400 		{object} 	Error
// @Failure 	500 		{object} 	Error
// @Router		/users/jwt 	[post]
func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)

	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		e := Error{Message: "user not found"}
		json.NewEncoder(w).Encode(e)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJwtOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// CreateUser godoc
// @Summary		Create a new user
// @Description	Create a new user
// @Tags 		users
// @Accept		json
// @Produce 	json
// @Param 		request 	body 		dto.CreateUserInput 	true 	"user request"
// @Success 	201
// @Failure 	500 		{object} 	Error
// @Router		/users 		[post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = h.UserDB.FindByEmail(user.Email)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		e := Error{Message: "email already exists"}
		json.NewEncoder(w).Encode(e)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		e := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(e)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
