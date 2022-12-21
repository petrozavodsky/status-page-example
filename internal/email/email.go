package email

import (
	"sort"
	"status_page/internal/utils"
	"strings"
)

// EmailData - Структура для хранения данных системы Email
type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

// InstanceEmail - Инициации экземпляра данных
func InstanceEmail() map[string][][]EmailData {
	return Sorted(Validate(utils.ReadFileCsv(utils.ConfigData.EmailDataPath)))
}

// Validate - Функция валидации данных
func Validate(data [][]string) []EmailData {
	result := make([]EmailData, 0)
	for _, line := range data {

		row := strings.Split(line[0], ";")

		switch true {
		case len(row) != 3:
			continue
		case !utils.IsExist(utils.ConfigData.CountryCode, row[0]):
			continue
		case !utils.IsExist(utils.ConfigData.ProvidersEmail, row[1]):
			continue
		default:
			var newEmailData EmailData
			newEmailData.Country = row[0]
			newEmailData.Provider = row[1]
			newEmailData.DeliveryTime = utils.ToInt(row[2])
			result = append(result, newEmailData)
		}
	}
	return result
}

// Sorted - Функция сортировки данных и создания двух списков
func Sorted(emailData []EmailData) map[string][][]EmailData {
	result := make(map[string][][]EmailData)
	sort.SliceStable(emailData, func(i, j int) bool {
		return emailData[i].Country > emailData[j].Country
	})
	uniqueCountryList := getCountry(emailData)

	for _, country := range uniqueCountryList {
		res := make([]EmailData, 0)
		for _, data := range emailData {
			if country == data.Country {
				res = append(res, data)
			}
		}
		sort.SliceStable(res, func(i, j int) bool {
			return res[i].DeliveryTime < res[j].DeliveryTime
		})

		top := 3

		if top > len(res) {
			top = len(res)
		}

		bottom := top

		topProviders := res[:top]
		bottomProviders := res[len(res)-bottom:]

		sort.SliceStable(bottomProviders, func(i, j int) bool {
			return bottomProviders[i].DeliveryTime > bottomProviders[j].DeliveryTime
		})

		result[country] = [][]EmailData{topProviders, bottomProviders}
	}

	return result
}

// getCountry - Получение списка стра
func getCountry(data []EmailData) []string {
	result := make([]string, 0)

	for _, item := range data {
		result = append(result, item.Country)
	}

	return uniqueCountry(result)
}

// uniqueCountry - Функция фильитрации повторящихся эллементов
func uniqueCountry(array []string) []string {
	keys := make(map[string]bool)
	result := make([]string, 0)

	for _, item := range array {
		if _, value := keys[item]; !value {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}
