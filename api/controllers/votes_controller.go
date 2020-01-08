package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/auth"
	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/models"
	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/responses"
	"github.com/Muhammad-Tounsi/Vote-Go-Vue/api/utils/formaterror"
	"github.com/gorilla/mux"
)

func (server *Server) Createvote(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vote := models.Vote{}
	err = json.Unmarshal(body, &vote)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	vote.Prepare()
	err = vote.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != vote.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	voteCreated, err := vote.SaveVote(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, voteCreated.ID))
	responses.JSON(w, http.StatusCreated, voteCreated)
}

func (server *Server) Getvotes(w http.ResponseWriter, r *http.Request) {

	vote := models.Vote{}
	fmt.Print(vote)
	votes, err := vote.FindAllVotes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, votes)
}

func (server *Server) Getvote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	vote := models.Vote{}

	voteReceived, err := vote.FindVoteByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, voteReceived)
}

func (server *Server) Updatevote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the vote id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the vote exist
	vote := models.Vote{}
	err = server.DB.Debug().Model(models.Vote{}).Where("id = ?", pid).Take(&vote).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Vote not found"))
		return
	}

	fmt.Printf("User id: %v", uid)
	fmt.Printf(" Vote id: %v", vote.ID)
	// If a user attempt to update a vote not belonging to him
	if uid != vote.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data voteed
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	voteUpdate := models.Vote{}
	err = json.Unmarshal(body, &voteUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	voteUpdate.AuthorID = uid
	//Also check if the request user id is equal to the one gotten from token
	if uid != voteUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	voteUpdate.Prepare()
	err = voteUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Print(voteUpdate)
	voteUpdated, err := voteUpdate.UpdateAVote(server.DB, pid)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, voteUpdated)
}


func (server *Server) Addvote(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	// Check if the vote id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	
	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the vote exist
	vote := models.Vote{}
	err = server.DB.Debug().Model(models.Vote{}).Where("id = ?", pid).Take(&vote).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Vote not found"))
		return
	}

	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	fmt.Printf("User id: %v", user.ID)
	fmt.Printf(" Vote id: %v", vote.ID)

	server.DB.Model(&vote).Association("Users").Append(&user);

	responses.JSON(w, http.StatusOK, vote)
}


func (server *Server) Deleteuservote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid vote id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the vote exist
	vote := models.Vote{}
	err = server.DB.Debug().Model(models.Vote{}).Where("id = ?", pid).Take(&vote).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Vote not found"))
		return
	}

	user := models.User{}
	err = server.DB.Debug().Model(models.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("User not found"))
		return
	}

	fmt.Printf("User id: %v", user.ID)
	fmt.Printf(" Vote id: %v", vote.ID)

	server.DB.Model(&vote).Association("Users").Delete(&user);

	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}



func (server *Server) Deletevote(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid vote id given to us?
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the vote exist
	vote := models.Vote{}
	err = server.DB.Debug().Model(models.Vote{}).Where("id = ?", pid).Take(&vote).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this vote?
	if uid != vote.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = vote.DeleteAVote(server.DB, pid, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
