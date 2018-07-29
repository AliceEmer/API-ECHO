package controllers

import (
	"fmt"
	"net/http"

	"github.com/AliceEmer/API2/models"
	"github.com/labstack/echo"
)

//GET

func (cn *Controller) GetAllPersons(c echo.Context) error {
	pers, err := cn.allPersons()

	if err != nil {
		fmt.Println("ERROR QUERY ALL PERSONS")
		return err
	}

	if len(pers) == 0 {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "VIDE",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*models.Person{
		"people": pers,
	})
}

func (cn *Controller) GetPerson(c echo.Context) error {

	id := c.Param("id")
	pers, err := cn.personByID(id)

	if err != nil || len(pers) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person with this ID",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*models.Person{
		"people": pers,
	})

}

func (cn *Controller) GetAddress(c echo.Context) error {

	id := c.Param("id")

	fmt.Printf("id : %v", id)

	adds, err := cn.addressByID(id)

	if err != nil || len(adds) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No person with this ID or this person has no address",
		})
	}

	return c.JSON(http.StatusOK, map[string][]*models.Address{
		"address": adds,
	})

}

//POST
func (cn *Controller) CreatePerson(c echo.Context) error {
	person := models.Person{}
	if err := c.Bind(&person); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cn.addPerson(&person)
	return c.JSON(http.StatusOK, map[string]string{
		"firstname": person.Firstname,
		"Lastname":  person.Lastname,
	})
}

func (cn *Controller) CreateAddress(c echo.Context) error {

	id := c.Param("id")
	address := models.Address{}
	if err := c.Bind(&address); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	cn.addAddress(&address, id)
	return c.JSON(http.StatusOK, map[string]string{
		"city":      address.City,
		"state":     address.State,
		"person_id": id,
	})
}

//DELETE
func (cn *Controller) DeletePerson(c echo.Context) error {
	id := c.Param("id")
	cn.dropPerson(id)
	return c.JSON(http.StatusOK, "Person deleted")
}

func (cn *Controller) DeleteAddress(c echo.Context) error {
	id := c.Param("id")
	cn.dropAddress(id)
	return c.JSON(http.StatusOK, "Address deleted")
}

//comment on lit result de Exec ?
//Struct vraiment utile ? creation de struc à chaque fois qu'on ne réutilise plus