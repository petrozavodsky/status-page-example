package sms

import (
	"sort"
	"status_page/internal/utils"
	"strings"
)

// SMSData - Структура для хранения данных системы SMS
type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

// InstanceSms - Инициации экземпляра данных
func InstanceSms() [][]SMSData {
	return Sorted(Validate(utils.ReadFileCsv(utils.ConfigData.SmsDataPath)))
}

// Validate - Функция валидации данных
func Validate(data [][]string) []SMSData {
	result := make([]SMSData, 0)

	for _, line := range data {
		row := strings.Split(line[0], ";")

		switch true {
		case len(row) != 4:
			continue
		case !utils.IsExist(utils.ConfigData.CountryCode, row[0]):
			continue
		case !utils.IsExist(utils.ConfigData.Providers, row[3]):
			continue
		default:
			var newSmsData SMSData
			newSmsData.Country = row[0]
			newSmsData.Bandwidth = row[1]
			newSmsData.ResponseTime = row[2]
			newSmsData.Provider = row[3]
			result = append(result, newSmsData)
		}
	}

	return result
}

// Sorted - Функция сортировки данных и создания двух списков
func Sorted(sms []SMSData) [][]SMSData {
	result := make([][]SMSData, 0)
	smsDataSortedByCountryName := make([]SMSData, 0)
	smsDataSortedByProviderName := make([]SMSData, 0)

	for _, item := range sms {
		item.Country = utils.ConfigData.CountryCodes[item.Country]
		smsDataSortedByCountryName = append(smsDataSortedByCountryName, item)
		smsDataSortedByProviderName = append(smsDataSortedByProviderName, item)
	}

	sort.SliceStable(smsDataSortedByCountryName, func(i, j int) bool {
		return smsDataSortedByCountryName[i].Country < smsDataSortedByCountryName[j].Country
	})

	sort.SliceStable(smsDataSortedByProviderName, func(i, j int) bool {
		return smsDataSortedByProviderName[i].Provider < smsDataSortedByProviderName[j].Provider
	})

	result = append(result, smsDataSortedByProviderName)

	result = append(result, smsDataSortedByCountryName)

	return result
}
