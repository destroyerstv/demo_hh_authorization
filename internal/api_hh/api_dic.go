package api_hh

type IDictionaries interface {
	GetApplicantCommentsOrder() IArrayValues
	GetApplicantNegotiationStatus() IArrayValues
	GetApplicantCommentAccessType() IArrayValues
	GetBusinessTripReadiness() IArrayValues
	GetCurrency() IArrayValues
	GetDriverLicenseTypes() IArrayValues
	GetEducationLevel() IArrayValues
	GetEmployerActiveVacanciesOrder() IArrayValues
	GetEmployerArchivedVacanciesOrder() IArrayValues
	GetEmployerHiddenVacanciesOrder() IArrayValues
	GetEmployerRelation() IArrayValues
	GetEmployerType() IArrayValues
	GetEmployment() IArrayValues
	GetExperience() IArrayValues
	GetGender() IArrayValues
	GetJobSearchStatusesApplicant() IArrayValues
	GetJobSearchStatusesEmployer() IArrayValues
	GetLanguageLevel() IArrayValues
	GetMessagingStatus() IArrayValues
	GetNegotiationsOrder() IArrayValues
	GetNegotiationsParticipantType() IArrayValues
	GetNegotiationsState() IArrayValues
	GetPhoneCallStatus() IArrayValues
	GetPreferredContactType() IArrayValues
	GetRelocationType() IArrayValues
	GetResumeAccessType() IArrayValues
	GetResumeContactsSiteType() IArrayValues
	GetResumeHiddenFields() IArrayValues
	GetResumeModerationNote() IArrayValues
	GetResumeSearchExperiencePeriod() IArrayValues
	GetResumeSearchFields() IArrayValues
	GetResumeSearchLabel() IArrayValues
	GetResumeSearchLogic() IArrayValues
	GetResumeSearchOrder() IArrayValues
	GetResumeSearchRelocation() IArrayValues
	GetResumeStatus() IArrayValues
	GetSchedule() IArrayValues
	GetTravelTime() IArrayValues
	GetVacancyBillingType() IArrayValues
	GetVacancyCluster() IArrayValues
	GetVacancyLabel() IArrayValues
	GetVacancyNotProlongedReason() IArrayValues
	GetVacancyRelation() IArrayValues
	GetVacancySearchFields() IArrayValues
	GetVacancySearchOrder() IArrayValues
	GetVacancyType() IArrayValues
	GetWorkingDays() IArrayValues
	GetWorkingTimeIntervals() IArrayValues
	GetWorkingTimeModes() IArrayValues
}

type IArrayValues interface {
	GetAllValues() map[string]string
	GetIdByName(name string) string
	GetNameById(id string) string
}

type Values struct {
	Id   string
	Name string
}

type ArrayValues struct {
	values []*Values
}

func (av *ArrayValues) GetAllValues() map[string]string {
	m := make(map[string]string)

	for _, value := range av.values {
		m[value.Id] = value.Name
	}

	return m
}

func (av *ArrayValues) GetIdByName(name string) string {

	for _, value := range av.values {
		if value.Name == name {
			return value.Id
		}
	}

	return ""
}

func (av *ArrayValues) GetNameById(id string) string {

	for _, value := range av.values {
		if value.Id == id {
			return value.Name
		}
	}

	return ""
}

type Dictionaries struct {
	ApplicantCommentsOrder         []*Values `json:"applicant_comment_access_type"`
	ApplicantNegotiationStatus     []*Values `json:"applicant_comments_order"`
	ApplicantCommentAccessType     []*Values `json:"applicant_negotiation_status"`
	BusinessTripReadiness          []*Values `json:"business_trip_readiness"`
	Currency                       []*Values `json:"currency"`
	DriverLicenseTypes             []*Values `json:"driver_license_types"`
	EducationLevel                 []*Values `json:"education_level"`
	EmployerActiveVacanciesOrder   []*Values `json:"employer_active_vacancies_order"`
	EmployerArchivedVacanciesOrder []*Values `json:"employer_archived_vacancies_order"`
	EmployerHiddenVacanciesOrder   []*Values `json:"employer_hidden_vacancies_order"`
	EmployerRelation               []*Values `json:"employer_relation"`
	EmployerType                   []*Values `json:"employer_type"`
	Employment                     []*Values `json:"employment"`
	Experience                     []*Values `json:"experience"`
	Gender                         []*Values `json:"gender"`
	JobSearchStatusesApplicant     []*Values `json:"job_search_statuses_applicant"`
	JobSearchStatusesEmployer      []*Values `json:"job_search_statuses_employer"`
	LanguageLevel                  []*Values `json:"language_level"`
	MessagingStatus                []*Values `json:"messaging_status"`
	NegotiationsOrder              []*Values `json:"negotiations_order"`
	NegotiationsParticipantType    []*Values `json:"negotiations_participant_type"`
	NegotiationsState              []*Values `json:"negotiations_state"`
	PhoneCallStatus                []*Values `json:"phone_call_status"`
	PreferredContactType           []*Values `json:"preferred_contact_type"`
	RelocationType                 []*Values `json:"relocation_type"`
	ResumeAccessType               []*Values `json:"resume_access_type"`
	ResumeContactsSiteType         []*Values `json:"resume_contacts_site_type"`
	ResumeHiddenFields             []*Values `json:"resume_hidden_fields"`
	ResumeModerationNote           []*Values `json:"resume_moderation_note"`
	ResumeSearchExperiencePeriod   []*Values `json:"resume_search_experience_period"`
	ResumeSearchFields             []*Values `json:"resume_search_fields"`
	ResumeSearchLabel              []*Values `json:"resume_search_label"`
	ResumeSearchLogic              []*Values `json:"resume_search_logic"`
	ResumeSearchOrder              []*Values `json:"resume_search_order"`
	ResumeSearchRelocation         []*Values `json:"resume_search_relocation"`
	ResumeStatus                   []*Values `json:"resume_status"`
	Schedule                       []*Values `json:"schedule"`
	TravelTime                     []*Values `json:"travel_time"`
	VacancyBillingType             []*Values `json:"vacancy_billing_type"`
	VacancyCluster                 []*Values `json:"vacancy_cluster"`
	VacancyLabel                   []*Values `json:"vacancy_label"`
	VacancyNotProlongedReason      []*Values `json:"vacancy_not_prolonged_reason"`
	VacancyRelation                []*Values `json:"vacancy_relation"`
	VacancySearchFields            []*Values `json:"vacancy_search_fields"`
	VacancySearchOrder             []*Values `json:"vacancy_search_order"`
	VacancyType                    []*Values `json:"vacancy_type"`
	WorkingDays                    []*Values `json:"working_days"`
	WorkingTimeIntervals           []*Values `json:"working_time_intervals"`
	WorkingTimeModes               []*Values `json:"working_time_modes"`
}

func (dict *Dictionaries) GetApplicantCommentsOrder() IArrayValues {
	return &ArrayValues{
		dict.ApplicantCommentsOrder,
	}
}

func (dict *Dictionaries) GetApplicantNegotiationStatus() IArrayValues {
	return &ArrayValues{
		dict.ApplicantNegotiationStatus,
	}
}
func (dict *Dictionaries) GetApplicantCommentAccessType() IArrayValues {
	return &ArrayValues{
		dict.ApplicantCommentAccessType,
	}
}
func (dict *Dictionaries) GetBusinessTripReadiness() IArrayValues {
	return &ArrayValues{
		dict.BusinessTripReadiness,
	}
}

func (dict *Dictionaries) GetCurrency() IArrayValues {
	return &ArrayValues{
		dict.Currency,
	}
}

func (dict *Dictionaries) GetDriverLicenseTypes() IArrayValues {
	return &ArrayValues{
		dict.DriverLicenseTypes,
	}
}

func (dict *Dictionaries) GetEducationLevel() IArrayValues {
	return &ArrayValues{
		dict.EducationLevel,
	}
}

func (dict *Dictionaries) GetEmployerActiveVacanciesOrder() IArrayValues {
	return &ArrayValues{
		dict.EmployerActiveVacanciesOrder,
	}
}

func (dict *Dictionaries) GetEmployerArchivedVacanciesOrder() IArrayValues {
	return &ArrayValues{
		dict.EmployerArchivedVacanciesOrder,
	}
}

func (dict *Dictionaries) GetEmployerHiddenVacanciesOrder() IArrayValues {
	return &ArrayValues{
		dict.EmployerHiddenVacanciesOrder,
	}
}

func (dict *Dictionaries) GetEmployerRelation() IArrayValues {
	return &ArrayValues{
		dict.EmployerRelation,
	}
}

func (dict *Dictionaries) GetEmployerType() IArrayValues {
	return &ArrayValues{
		dict.EmployerType,
	}
}

func (dict *Dictionaries) GetEmployment() IArrayValues {
	return &ArrayValues{
		dict.Employment,
	}
}

func (dict *Dictionaries) GetExperience() IArrayValues {
	return &ArrayValues{
		dict.Experience,
	}
}

func (dict *Dictionaries) GetGender() IArrayValues {
	return &ArrayValues{
		dict.Gender,
	}
}

func (dict *Dictionaries) GetJobSearchStatusesApplicant() IArrayValues {
	return &ArrayValues{
		dict.JobSearchStatusesApplicant,
	}
}

func (dict *Dictionaries) GetJobSearchStatusesEmployer() IArrayValues {
	return &ArrayValues{
		dict.JobSearchStatusesEmployer,
	}
}

func (dict *Dictionaries) GetLanguageLevel() IArrayValues {
	return &ArrayValues{
		dict.LanguageLevel,
	}
}

func (dict *Dictionaries) GetMessagingStatus() IArrayValues {
	return &ArrayValues{
		dict.MessagingStatus,
	}
}

func (dict *Dictionaries) GetNegotiationsOrder() IArrayValues {
	return &ArrayValues{
		dict.NegotiationsOrder,
	}
}

func (dict *Dictionaries) GetNegotiationsParticipantType() IArrayValues {
	return &ArrayValues{
		dict.NegotiationsParticipantType,
	}
}

func (dict *Dictionaries) GetNegotiationsState() IArrayValues {
	return &ArrayValues{
		dict.NegotiationsState,
	}
}

func (dict *Dictionaries) GetPhoneCallStatus() IArrayValues {
	return &ArrayValues{
		dict.PhoneCallStatus,
	}
}

func (dict *Dictionaries) GetPreferredContactType() IArrayValues {
	return &ArrayValues{
		dict.PreferredContactType,
	}
}

func (dict *Dictionaries) GetRelocationType() IArrayValues {
	return &ArrayValues{
		dict.RelocationType,
	}
}

func (dict *Dictionaries) GetResumeAccessType() IArrayValues {
	return &ArrayValues{
		dict.ResumeAccessType,
	}
}

func (dict *Dictionaries) GetResumeContactsSiteType() IArrayValues {
	return &ArrayValues{
		dict.ResumeContactsSiteType,
	}
}

func (dict *Dictionaries) GetResumeHiddenFields() IArrayValues {
	return &ArrayValues{
		dict.ResumeHiddenFields,
	}
}

func (dict *Dictionaries) GetResumeModerationNote() IArrayValues {
	return &ArrayValues{
		dict.ResumeModerationNote,
	}
}

func (dict *Dictionaries) GetResumeSearchExperiencePeriod() IArrayValues {
	return &ArrayValues{
		dict.ResumeSearchExperiencePeriod,
	}
}

func (dict *Dictionaries) GetResumeSearchFields() IArrayValues {
	return &ArrayValues{
		dict.ResumeSearchFields,
	}
}

func (dict *Dictionaries) GetResumeSearchLabel() IArrayValues {
	return &ArrayValues{
		dict.ResumeSearchLabel,
	}
}

func (dict *Dictionaries) GetResumeSearchLogic() IArrayValues {
	return &ArrayValues{
		dict.ResumeSearchLogic,
	}
}

func (dict *Dictionaries) GetResumeSearchOrder() IArrayValues {
	return &ArrayValues{
		dict.ResumeSearchOrder,
	}
}

func (dict *Dictionaries) GetResumeSearchRelocation() IArrayValues {
	return &ArrayValues{
		dict.ResumeSearchRelocation,
	}
}

func (dict *Dictionaries) GetResumeStatus() IArrayValues {
	return &ArrayValues{
		dict.ResumeStatus,
	}
}

func (dict *Dictionaries) GetSchedule() IArrayValues {
	return &ArrayValues{
		dict.Schedule,
	}
}

func (dict *Dictionaries) GetTravelTime() IArrayValues {
	return &ArrayValues{
		dict.TravelTime,
	}
}

func (dict *Dictionaries) GetVacancyBillingType() IArrayValues {
	return &ArrayValues{
		dict.VacancyBillingType,
	}
}

func (dict *Dictionaries) GetVacancyCluster() IArrayValues {
	return &ArrayValues{
		dict.VacancyCluster,
	}
}

func (dict *Dictionaries) GetVacancyLabel() IArrayValues {
	return &ArrayValues{
		dict.VacancyLabel,
	}
}

func (dict *Dictionaries) GetVacancyNotProlongedReason() IArrayValues {
	return &ArrayValues{
		dict.VacancyNotProlongedReason,
	}
}

func (dict *Dictionaries) GetVacancyRelation() IArrayValues {
	return &ArrayValues{
		dict.VacancyRelation,
	}
}

func (dict *Dictionaries) GetVacancySearchFields() IArrayValues {
	return &ArrayValues{
		dict.VacancySearchFields,
	}
}

func (dict *Dictionaries) GetVacancySearchOrder() IArrayValues {
	return &ArrayValues{
		dict.VacancySearchOrder,
	}
}

func (dict *Dictionaries) GetVacancyType() IArrayValues {
	return &ArrayValues{
		dict.VacancyType,
	}
}

func (dict *Dictionaries) GetWorkingDays() IArrayValues {
	return &ArrayValues{
		dict.WorkingDays,
	}
}

func (dict *Dictionaries) GetWorkingTimeIntervals() IArrayValues {
	return &ArrayValues{
		dict.WorkingTimeIntervals,
	}
}

func (dict *Dictionaries) GetWorkingTimeModes() IArrayValues {
	return &ArrayValues{
		dict.WorkingTimeModes,
	}
}
