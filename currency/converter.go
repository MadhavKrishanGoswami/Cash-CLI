package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/MadhavKrishanGoswami/Cash-CLI/models"
)

func Convertor(input models.Input) models.Output {
	in_int, err := strconv.Atoi(input.Input)
	if err != nil {
		log.Fatalf("Wround input could not cunvert to val to int %v", err)
	}
	rep, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/WoXy-Sensei/currency-api/main/api/%v_%v.json", input.Convert, input.Base))
	if err != nil {
		log.Fatalf("could not get api %v", err)
	}

	var apidata models.APIval
	body, err := io.ReadAll(rep.Body)

	json.Unmarshal(body, &apidata)
	var Output models.Output
	Output.Converstion = apidata.Rate
	Output.Value = apidata.Rate * float64(in_int)

	return Output
}
