package voicecall

import (
	"status_page/internal/utils"
	"strings"
)

// VoiceData - Структура для хранениния данных о звонках
type VoiceData struct {
	Country             string
	CurrentLoad         int
	AverageResponseTime int
	Provider            string
	ConnectionStability float32
	PurityCommunication int
	TTFB                int
	MedianCallDuration  int
}

// InstanceVoiceCall - Инициации экземпляра данных
func InstanceVoiceCall() []VoiceData {
	return Validate(utils.ReadFileCsv(utils.ConfigData.VoiceDataPath))
}

// Validate - Функция валидации данных
func Validate(data [][]string) []VoiceData {
	result := make([]VoiceData, 0)

	for _, line := range data {
		row := strings.Split(line[0], ";")
		switch true {
		case len(row) != 8:
			continue
		case !utils.IsExist(utils.ConfigData.CountryCode, row[0]):
			continue
		case !utils.IsExist(utils.ConfigData.ProvidersCall, row[3]):
			continue
		default:
			var newVoiceData VoiceData
			newVoiceData.Country = row[0]
			newVoiceData.CurrentLoad = utils.ToInt(row[1])
			newVoiceData.AverageResponseTime = utils.ToInt(row[2])
			newVoiceData.Provider = row[3]
			newVoiceData.ConnectionStability = utils.ToFloat32(row[4])
			newVoiceData.PurityCommunication = utils.ToInt(row[5])
			newVoiceData.TTFB = utils.ToInt(row[6])
			newVoiceData.MedianCallDuration = utils.ToInt(row[7])
			result = append(result, newVoiceData)
		}
	}

	return result
}
