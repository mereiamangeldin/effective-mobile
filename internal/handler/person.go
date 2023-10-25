package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mereiamangeldin/effective-mobile-test/api"
	"github.com/mereiamangeldin/effective-mobile-test/internal/entity"
	"io"
	"log"
	"net/http"
)

var ageUrl string = "https://api.agify.io/?name=%s"
var genderUrl string = "https://api.genderize.io/?name=%s"
var nationUrl string = "https://api.nationalize.io/?name=%s"

func StringToUint(s string) (uint, error) {
	var num uint
	_, err := fmt.Sscanf(s, "%d", &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func handleError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, &api.Error{
		Message: err.Error(),
	})
}

func fetchData(ctx *gin.Context, url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code is not OK")
	}

	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &target); err != nil {
		return err
	}

	return nil
}

func (h *Handler) addPerson(ctx *gin.Context) {
	h.logger.Info("calling a handler of adding a new person")
	var req api.PersonInput

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	var personAge api.PersonAge
	err = fetchData(ctx, fmt.Sprintf(ageUrl, req.Name), &personAge)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	var personGender api.PersonGender
	err = fetchData(ctx, fmt.Sprintf(genderUrl, req.Name), &personGender)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	var personNation api.PersonNationality
	err = fetchData(ctx, fmt.Sprintf(nationUrl, req.Name), &personNation)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}

	person := &entity.Person{
		Name:        req.Name,
		Surname:     req.Surname,
		Patronymic:  req.Patronymic,
		Age:         personAge.Age,
		Gender:      personGender.Gender,
		Nationality: personNation.Country[0].CountryID,
	}
	h.logger.Debugf("adding a new person: %s", person)

	h.logger.Info("waiting response from service level")
	err = h.srvs.AddPerson(ctx, person)

	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	h.logger.Info("new person is successfully added")
	ctx.JSON(http.StatusCreated, api.Ok{Message: "Person is successfully added"})
}

func (h *Handler) updatePerson(ctx *gin.Context) {
	h.logger.Info("calling a handler of updating a person")
	id := ctx.Param("person_id")
	personId, err := StringToUint(id)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	var person entity.Person

	err = ctx.ShouldBindJSON(&person)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	h.logger.Debugf("updating a person with id: %d", personId)

	h.logger.Info("waiting response from service level")
	err = h.srvs.UpdatePerson(ctx, personId, person)

	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	h.logger.Info("person is successfully updated")
	ctx.JSON(http.StatusOK, api.Ok{Message: "person is successfully updated"})
}

func (h *Handler) getPerson(ctx *gin.Context) {
	h.logger.Info("calling a handler of getting a person")
	id := ctx.Param("person_id")
	personId, err := StringToUint(id)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	h.logger.Info("waiting response from service level")
	person, err := h.srvs.GetPerson(ctx, personId)

	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	h.logger.Info("person is successfully returned")
	ctx.JSON(http.StatusOK, person)
}

func (h *Handler) deletePerson(ctx *gin.Context) {
	h.logger.Info("calling a handler of deleting a new person")
	id := ctx.Param("person_id")
	personId, err := StringToUint(id)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	h.logger.Debugf("deleting a person with id: %d", personId)

	h.logger.Info("waiting response from service level")
	err = h.srvs.DeletePerson(ctx, personId)

	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	h.logger.Info("person is successfully deleted")
	ctx.JSON(http.StatusOK, api.Ok{Message: "person is successfully deleted"})
}

func (h *Handler) getPeople(ctx *gin.Context) {
	h.logger.Info("calling a handler of getting a people")
	gender := ctx.DefaultQuery("gender", "")
	limit_ := ctx.DefaultQuery("limit", "5")
	limit, err := StringToUint(limit_)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	offset_ := ctx.DefaultQuery("offset", "0")
	offset, err := StringToUint(offset_)
	if err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}
	peopleFilter := api.PeopleFilter{Gender: gender}
	pagination := api.Pagination{Limit: limit, Offset: offset}
	h.logger.Info("waiting response from service level")
	people, err := h.srvs.GetPeople(ctx, peopleFilter, pagination)

	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}
	h.logger.Info("people are successfully returned")
	ctx.JSON(http.StatusOK, people)
}
