package handler

import (
	"encoding/json"
	// "io"
	"net/http"
	"w3/d4/config"
	"w3/d4/entity"
	"w3/d4/utils"
	"strings"
	"fmt"
	"github.com/labstack/echo/v4"
)

func GetCities(c echo.Context, cfg *config.Config) error {
	newUrl := cfg.RajaOngkirURL + "/city"

	headers := map[string]string{
		"key": cfg.RajaOngkirAPIKey,
	}

	res, err := utils.RequestGET(newUrl, headers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed parsing response from server"})
	}

	var result map[string]interface{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "error while unmarshalling response"})
	}

	rajaOngkirData := result["rajaongkir"].(map[string]interface{})
	results := rajaOngkirData["results"].([]interface{})
	cities := make([]entity.City, len(results))
	for i, r := range results {
		cityData := r.(map[string]interface{})
		cities[i] = entity.City{
			CityID:     cityData["city_id"].(string),
			ProvinceID: cityData["province_id"].(string),
			Province:   cityData["province"].(string),
			Type:       cityData["type"].(string),
			CityName:   cityData["city_name"].(string),
			PostalCode: cityData["postal_code"].(string),
		}
	}

	return c.JSON(http.StatusOK, cities)
}

func GetOngkir(c echo.Context, cfg *config.Config) error {
    newUrl := cfg.RajaOngkirURL + "/cost"
    headers := map[string]string{
        "key":          cfg.RajaOngkirAPIKey,
        "content-type": "application/x-www-form-urlencoded",
    }

    origin := c.FormValue("origin")
    destination := c.FormValue("destination")
    weight := c.FormValue("weight")
    courier := c.FormValue("courier")

    payload := strings.NewReader(fmt.Sprintf("origin=%s&destination=%s&weight=%s&courier=%s", origin, destination, weight, courier))

    res, err := utils.RequestPOST(newUrl, headers, payload)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed parsing response from server"})
    }

    var result map[string]interface{}
    err = json.Unmarshal(res, &result)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "error while unmarshalling response"})
    }

    rajaOngkirData := result["rajaongkir"].(map[string]interface{})
    results := rajaOngkirData["results"].([]interface{})

    return c.JSON(http.StatusOK, results)
}