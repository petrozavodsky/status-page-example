package support

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"status_page/internal/utils"
)

// SupportData - Структура для хранения данных системы Support
type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// InstanceSupport - Инициации экземпляра данных
func InstanceSupport() ([]int, error) {
	data, err := getSupportData()
	if err != nil {
		return []int{0, 0}, err
	}

	return Validate(data), nil
}

// getSupportData - Функция отправки запроса к API и получения ответа
func getSupportData() ([]SupportData, error) {
	var supportData []SupportData

	response, err := http.Get(utils.ConfigData.SupportServicePath)

	if err != nil {
		return supportData, errors.New("Не удалось отправить запрос к серверу support")
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return supportData, errors.New("Ошибка при получении данных с сервера support")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if json.Unmarshal(body, &supportData) != nil {
		return supportData, errors.New("Ошибка при чтении данных с сервера support")
	}

	return supportData, nil
}

// Validate - Функция валидации данных
func Validate(data []SupportData) []int {
	result := make([]int, 0)
	var totalTopic, load, averageTime int

	for _, item := range data {
		totalTopic += item.ActiveTickets
	}

	switch {
	case totalTopic < 9:
		load = 1
	case totalTopic <= 16:
		load = 2
	default:
		load = 3
	}

	averageTime = int((float64(60) / float64(18)) * float64(totalTopic))
	result = append(result, load)
	result = append(result, averageTime)

	return result

}
