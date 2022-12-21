package utils

import (
	"encoding/csv"
	"github.com/joho/godotenv"
	//nolint
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var ConfigData = instanceAppConfig()

// AppConfig - Конфиг приложения.
type AppConfig struct {
	CountryCode         []string
	Providers           []string
	ProvidersCall       []string
	ProvidersEmail      []string
	CountryCodes        map[string]string
	BillingDataPath     string
	EmailDataPath       string
	SmsDataPath         string
	VoiceDataPath       string
	SupportServicePath  string
	MmsServicePath      string
	IncidentServicePath string
}

// getBaseDir - Получение корня проекта
func getBaseDir() string {
	baseDir, _ := os.Getwd()

	return baseDir
}

// getEnv - Получение значения переменной окружения.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func instanceAppConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Не удалось получить файл .env", err)
	}

	return &AppConfig{
		GetCountryCodes(path.Join(getBaseDir(), getEnv("ALPHA_CODES_PATH", ""))),
		GetProvidersAllow(getEnv("PROVIDERS_LINE", "")),
		GetProvidersAllow(getEnv("PROVIDERS_CALL_LINE", "")),
		GetProvidersAllow(getEnv("PROVIDERS_EMAIL_LINE", "")),
		GetCountryCode(path.Join(getBaseDir(), getEnv("ALPHA_CODES_PATH", ""))),
		path.Join(getBaseDir(), getEnv("BILLING_DATA_FILE", "")),
		path.Join(getBaseDir(), getEnv("EMAIL_DATA_FILE", "")),
		path.Join(getBaseDir(), getEnv("SMS_DATA_FILE", "")),
		path.Join(getBaseDir(), getEnv("VOICE_DATA_FILE", "")),
		getEnv("SUPPORT_API_HANDLER", ""),
		getEnv("MMS_API_HANDLER", ""),
		getEnv("INCIDENT_API_HANDLER", ""),
	}
}

// ReadCsvFileDeprecated - Чтение csv файла.
func ReadCsvFileDeprecated(fileName string) [][]string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	out := make([][]string, 0)

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {

		if len(line) > 0 {
			out = append(out, strings.Split(line, ","))
		}
	}

	return out
}

// ReadFileCsv - Чтение csv файла.
func ReadFileCsv(fileName string) [][]string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Не удалось открыть файл", err.Error())
	}
	defer file.Close()
	reader := csv.NewReader(file)

	if "old" == getEnv("CSV_PARSER", "") {
		return ReadCsvFileDeprecated(fileName)
	}

	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Не удалось прочитать файл", err.Error())
	}
	return data
}

// GetCountryCodes - Получение кодов из файла.
func GetCountryCodes(path string) []string {
	code := make([]string, 0)
	data := ReadFileCsv(path)
	for _, line := range data {
		code = append(code, line[1])
	}

	return code
}

// GetCountryCode - Получение стран  из файла.
func GetCountryCode(path string) map[string]string {
	result := make(map[string]string, 0)
	data := ReadFileCsv(path)
	for _, line := range data {
		result[line[1]] = line[0]
	}
	return result
}

// GetProvidersAllow - Получения разрещенных провайдеров.
func GetProvidersAllow(line string) []string {
	allowProviders := make([]string, 0)

	items := strings.Split(string(line), ",")

	for _, item := range items {
		allowProviders = append(allowProviders, strings.TrimSpace(item))
	}

	return allowProviders
}

func IsExist(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}

// ToInt - Функция конвертирует string в int.
func ToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

// ToFloat32 - Конвертирует string в float32.
func ToFloat32(str string) float32 {
	number, err := strconv.ParseFloat(str, 32)
	if err != nil {
		log.Fatal(err)
	}
	return float32(number)
}

// ByteToBool - Конвертирует byte в bool
func ByteToBool(b byte) bool {
	return b != 48
}

// ErrorToString - Конвертирует error в string
func ErrorToString(err ...error) string {
	var errorString string
	for _, item := range err {
		if item != nil {
			errorString += item.Error() + ", "
		}

	}
	return strings.TrimRight(errorString, ", ")
}
