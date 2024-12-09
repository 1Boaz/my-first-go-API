package handelers

import (
	"encoding/json"
	"net/http"

	"github.com/1Boaz/my-first-go-API/api/api"
	"github.com/1Boaz/my-first-go-API/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request)  {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil{
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetalis *tools.CoinDetails
	tokenDetalis = (*database).GetUserCoins(params.Username)
	if tokenDetalis == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetalis),
		httpCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}