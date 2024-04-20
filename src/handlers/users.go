package handlers

import (
	"api/src/common/request"
	"api/src/common/responses"
	"api/src/database"
	"api/src/entities"
	"api/src/repositories"
	"fmt"
	"net/http"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user entities.User
	if err := request.ProcessBody(r, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(true); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusCreated, user)
}

// GetUsers gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	users, err := repository.GetAll()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, users)
}

// GetUser gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user, err := repository.Get(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.Json(w, http.StatusOK, user)
}

// UpdateUser updates a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	userID, err := request.GetId(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var user entities.User
	if err = request.ProcessBody(r, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(false); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	rowsAffected, err := repository.Update(userID, user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		error := fmt.Errorf("there is no user with id %d", userID)
		responses.Error(w, http.StatusNotFound, error)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Updated %d rows", rowsAffected))
}

// DeleteUser deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := request.GetId(r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.OpenConnection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	rowsAffected, err := repository.Delete(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		error := fmt.Errorf("there is no user with id %d", userID)
		responses.Error(w, http.StatusNotFound, error)
		return
	}

	responses.Json(w, http.StatusOK, fmt.Sprintf("Deleted %d rows", rowsAffected))
}
