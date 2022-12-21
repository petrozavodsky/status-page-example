package mms

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"sort"
	"status_page/internal/utils"
)

// MMSData - Структура для хранения данных
type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

// InstanceMms -  Инициации экземпляра данных
func InstanceMms() ([][]MMSData, error) {
	data, err := getMmsData()

	if err != nil {
		var res [][]MMSData
		return res, err
	}

	return Sorted(Validate(data)), nil
}

// getMmsData - Функция отправки запроса к API и получения ответа
func getMmsData() ([]MMSData, error) {

	var mmsData []MMSData

	response, err := http.Get(utils.ConfigData.MmsServicePath)
	if err != nil {
		return mmsData, errors.New("Не удалось отправить запрос к серверу о состоянии системы MMS")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Printf("Ошибка получения данных с сервера о состоянии системы MMS")
		return mmsData, errors.New("Ошибка получения данных о mms с сервера")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if json.Unmarshal(body, &mmsData) != nil {
		fmt.Printf("Ошибка при чтении данных о mms")
		return mmsData, errors.New("Ошибка при чтении данных о mms")
	}

	return mmsData, nil
}

// Validate - Функция валидации данных
func Validate(data []MMSData) []MMSData {
	var result []MMSData

	for _, item := range data {
		if !utils.IsExist(utils.ConfigData.CountryCode, item.Country) || !utils.IsExist(utils.ConfigData.Providers, item.Provider) {
			continue
		}
		result = append(result, item)
	}

	return result
}

// Sorted - Функция сортировки данных и создания двух списков
func Sorted(mms []MMSData) [][]MMSData {
	result := make([][]MMSData, 0)
	mmsDataSortedByCountryName := make([]MMSData, 0)
	mmsDataSortedByProviderName := make([]MMSData, 0)

	for _, item := range mms {
		item.Country = utils.ConfigData.CountryCodes[item.Country]
		mmsDataSortedByCountryName = append(mmsDataSortedByCountryName, item)
		mmsDataSortedByProviderName = append(mmsDataSortedByProviderName, item)
	}

	sort.SliceStable(mmsDataSortedByCountryName, func(i, j int) bool {
		return mmsDataSortedByCountryName[i].Country < mmsDataSortedByCountryName[j].Country
	})

	sort.SliceStable(mmsDataSortedByProviderName, func(i, j int) bool {
		return mmsDataSortedByProviderName[i].Provider < mmsDataSortedByProviderName[j].Provider
	})

	result = append(result, mmsDataSortedByProviderName)
	result = append(result, mmsDataSortedByCountryName)

	return result
}
