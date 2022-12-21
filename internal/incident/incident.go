package incident

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"sort"
	"status_page/internal/utils"
)

// IncidentData - Структура для хранения данных об инцидентах
type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

// InstanceIncident -  Инициации экземпляра данных
func InstanceIncident() ([]IncidentData, error) {
	data, err := getIncidentData()
	if err != nil {
		return data, err
	}
	return Validate(data), nil
}

// getIncidentData - Функция отправляет запрос к API для получения данных о системе истории инцидентов.
func getIncidentData() ([]IncidentData, error) {
	var incidentData []IncidentData

	response, err := http.Get(utils.ConfigData.IncidentServicePath)
	if err != nil {
		return incidentData, errors.New("Не удалось отправить запрос к серверу о системе истории инцидентов")
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return incidentData, errors.New("Ошибка получения данных с сервера о системе истории инцидентов")
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if json.Unmarshal(body, &incidentData) != nil {
		return incidentData, errors.New("Ошибка при чтении данных с сервера о системе истории инцидентов")
	}

	sort.SliceStable(incidentData, func(i, j int) bool {
		return incidentData[i].Status < incidentData[j].Status
	})

	return incidentData, nil
}

// Validate - Функция валидации данных
func Validate(data []IncidentData) []IncidentData {
	statusAllow := [2]string{"active", "closed"}
	var result []IncidentData

	for _, item := range data {
		if !utils.IsExist(statusAllow[:], item.Status) {
			continue
		}
		result = append(result, item)
	}

	return result
}
