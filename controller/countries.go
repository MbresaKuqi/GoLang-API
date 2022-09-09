package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"example.com/m/models"
	"example.com/m/storage"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var db *gorm.DB

func GetCountries(c echo.Context) error {
	countries, _ := GetRepoCountries()
	return c.JSON(http.StatusOK, countries)
}

func GetRepoCountries() ([]models.Country, error) {
	db := storage.GetDBInstance()
	countries := []models.Country{}

	if err := db.Find(&countries).Error; err != nil {
		return nil, err
	}

	return countries, nil
}

func CreateCountry(c echo.Context) (err error) {

	countryName := c.Param("countryName")
	response, err := http.Get("https://restcountries.com/v2/name/" + countryName + "?fields=name,capital,languages,languages")
	if response.StatusCode == 404 {
		return c.JSON(http.StatusNotFound, "Country not found")
	} else {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		var result []models.PublicCountry
		jsonErr := json.Unmarshal(data, &result)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		input := models.Country{
			Name:     result[0].Name,
			Capital:  result[0].Capital,
			Language: result[0].Language[0].Name,
		}
		db := storage.GetDBInstance()

		if err := db.Create(&input).Error; err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, input)
	}
}

func DeleteCountry(c echo.Context) (err error) {
	countryName := c.Param("countryName")
	db := storage.GetDBInstance()
	if err := db.Where("name = ?", countryName).Find(&models.Country{}).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Record not found")
	} else {
		db.Where("name = ?", countryName).Delete(&models.Country{})
		return c.JSON(http.StatusOK, countryName+" has been deleted from database")
	}
}
