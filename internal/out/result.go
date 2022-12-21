package out

import (
	billing_mod "status_page/internal/billing"
	email_mod "status_page/internal/email"
	incident_mod "status_page/internal/incident"
	mms_mod "status_page/internal/mms"
	sms_mod "status_page/internal/sms"
	support_mod "status_page/internal/support"
	"status_page/internal/utils"
	voicecall_mod "status_page/internal/voicecall"
)

// ResultT - Структура с данными для вывода
type ResultT struct {
	Status bool        `json:"status"`
	Data   *ResultSetT `json:"data"`
	Error  string      `json:"error"`
}

// ResultSetT - Структура данных всех сервисов
type ResultSetT struct {
	SMS       [][]sms_mod.SMSData                `json:"sms"`
	MMS       [][]mms_mod.MMSData                `json:"mms"`
	VoiceCall []voicecall_mod.VoiceData          `json:"voice_call"`
	Email     map[string][][]email_mod.EmailData `json:"email"`
	Billing   billing_mod.BillingData            `json:"billing"`
	Support   []int                              `json:"support"`
	Incidents []incident_mod.IncidentData        `json:"incident"`
}

// GetResultData - Функция для сбора данных в структуру
func GetResultData() ResultT {
	sms := sms_mod.InstanceSms()
	mms, errMms := mms_mod.InstanceMms()
	voice := voicecall_mod.InstanceVoiceCall()
	email := email_mod.InstanceEmail()
	billing := billing_mod.InstancBilling()
	support, errSupport := support_mod.InstanceSupport()
	incident, errIncident := incident_mod.InstanceIncident()

	if errMms != nil || errSupport != nil || errIncident != nil {
		return ResultT{false, nil, utils.ErrorToString(errMms, errSupport, errIncident)}
	}

	return ResultT{
		true,
		&ResultSetT{
			sms,
			mms,
			voice,
			email,
			*billing,
			support,
			incident,
		},
		"",
	}
}
