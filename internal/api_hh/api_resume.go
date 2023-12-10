package api_hh

import "encoding/json"

type IHHResume interface {
	GetJSON() (string, error)
	SetAccess() bool
	SetBirthDate() bool
	SetBusinessTripReadiness() bool
	SetCertificate() bool
	SetDriverLicenseTypes() bool
	SetEmployments() bool
	SetFirstName() bool
	SetHasVenicle() bool
	SetHiddenFields() bool
	SetLastName() bool
	SetMetro() bool
	SetMiddleName() bool
	SetPhoto() bool
	SetPortfolio() bool
	SetProfessionalRoles() bool
	SetRecommendation() bool
	SetRelocation() bool
	SetResumeLocale() bool
	SetSalary() bool
	SetSchedules() bool
	SetSite() bool
	SetSkillSet() bool
	SetSkills() bool
	SetTitle() bool
	SetTotalExperience() bool
	SetTravelTime() bool
	SetWorkTicket() bool
	SetArea() bool
	SetCitizenship() bool
	SetContact() bool
	SetEducation() bool
	SetExperience() bool
	SetGender() bool
	SetLanguage() bool
}

// type HHResume struct {
// 	Access
// 	BirthDate
// 	BusinessTripReadiness
// 	Certificate
// 	DriverLicenseTypes
// 	Employments
// 	FirstName
// 	HasVenicle
// 	HiddenFields
// 	LastName
// 	Metro
// 	MiddleName
// 	Photo
// 	Portfolio
// 	ProfessionalRoles
// 	Recommendation
// 	Relocation
// 	ResumeLocale
// 	Salary
// 	Schedules
// 	Site
// 	SkillSet
// 	Skills
// 	Title
// 	TotalExperience
// 	TravelTime
// 	WorkTicket
// 	Area
// 	Citizenship
// 	Contact
// 	Education
// 	Experience
// 	Gender
// 	Language
// }

type HHMinimalResume struct {
	Area        *TUserArea      `json:"area,omitempty"`
	Citizenship []*TCitizenship `json:"citizenship,omitempty"`
	Contact     []*TContact     `json:"contact,omitempty"`
	Education   *TEducation     `json:"education,omitempty"`
	Experience  []TExperience   `json:"experience"`
	FirstName   *string         `json:"first_name,omitempty"`
	Gender      *TGender        `json:"gender,omitempty"`
	Language    *TLanguage      `json:"language,omitempty"`
	LastName    *string         `json:"last_name,omitempty"`
	Photo       *TPhoto         `json:"photo"`
	Portfolio   *TPortfolio     `json:"portfolio"`
	Skills      *string         `json:"skills,omitempty"`
	Title       *string         `json:"title,omitempty"`
}

// GetJSON implements IHHResume.
func (mr *HHMinimalResume) GetJSON() (string, error) {
	b, err := json.Marshal(mr)
	return string(b), err
}

// SetAccess implements IHHResume.
func (mr *HHMinimalResume) SetAccess() bool {
	return false
}

// SetArea implements IHHResume.
func (mr *HHMinimalResume) SetArea() bool {
	return false
}

// SetBirthDate implements IHHResume.
func (mr *HHMinimalResume) SetBirthDate() bool {
	return false
}

// SetBusinessTripReadiness implements IHHResume.
func (mr *HHMinimalResume) SetBusinessTripReadiness() bool {
	return false
}

// SetCertificate implements IHHResume.
func (mr *HHMinimalResume) SetCertificate() bool {
	return false
}

// SetCitizenship implements IHHResume.
func (mr *HHMinimalResume) SetCitizenship() bool {
	return false
}

// SetContact implements IHHResume.
func (mr *HHMinimalResume) SetContact() bool {
	return false
}

// SetDriverLicenseTypes implements IHHResume.
func (mr *HHMinimalResume) SetDriverLicenseTypes() bool {
	return false
}

// SetEducation implements IHHResume.
func (mr *HHMinimalResume) SetEducation() bool {
	return false
}

// SetEmployments implements IHHResume.
func (mr *HHMinimalResume) SetEmployments() bool {
	return false
}

// SetExperience implements IHHResume.
func (mr *HHMinimalResume) SetExperience() bool {
	return false
}

// SetFirstName implements IHHResume.
func (mr *HHMinimalResume) SetFirstName() bool {
	return false
}

// SetGender implements IHHResume.
func (mr *HHMinimalResume) SetGender() bool {
	return false
}

// SetHasVenicle implements IHHResume.
func (mr *HHMinimalResume) SetHasVenicle() bool {
	return false
}

// SetHiddenFields implements IHHResume.
func (mr *HHMinimalResume) SetHiddenFields() bool {
	return false
}

// SetLanguage implements IHHResume.
func (mr *HHMinimalResume) SetLanguage() bool {
	return false
}

// SetLastName implements IHHResume.
func (mr *HHMinimalResume) SetLastName() bool {
	return false
}

// SetMetro implements IHHResume.
func (mr *HHMinimalResume) SetMetro() bool {
	return false
}

// SetMiddleName implements IHHResume.
func (mr *HHMinimalResume) SetMiddleName() bool {
	return false
}

// SetPhoto implements IHHResume.
func (mr *HHMinimalResume) SetPhoto() bool {
	return false
}

// SetPortfolio implements IHHResume.
func (mr *HHMinimalResume) SetPortfolio() bool {
	return false
}

// SetProfessionalRoles implements IHHResume.
func (mr *HHMinimalResume) SetProfessionalRoles() bool {
	return false
}

// SetRecommendation implements IHHResume.
func (mr *HHMinimalResume) SetRecommendation() bool {
	return false
}

// SetRelocation implements IHHResume.
func (mr *HHMinimalResume) SetRelocation() bool {
	return false
}

// SetResumeLocale implements IHHResume.
func (mr *HHMinimalResume) SetResumeLocale() bool {
	return false
}

// SetSalary implements IHHResume.
func (mr *HHMinimalResume) SetSalary() bool {
	return false
}

// SetSchedules implements IHHResume.
func (mr *HHMinimalResume) SetSchedules() bool {
	return false
}

// SetSite implements IHHResume.
func (mr *HHMinimalResume) SetSite() bool {
	return false
}

// SetSkillSet implements IHHResume.
func (mr *HHMinimalResume) SetSkillSet() bool {
	return false
}

// SetSkills implements IHHResume.
func (mr *HHMinimalResume) SetSkills() bool {
	return false
}

// SetTitle implements IHHResume.
func (mr *HHMinimalResume) SetTitle() bool {
	return false
}

// SetTotalExperience implements IHHResume.
func (mr *HHMinimalResume) SetTotalExperience() bool {
	return false
}

// SetTravelTime implements IHHResume.
func (mr *HHMinimalResume) SetTravelTime() bool {
	return false
}

// SetWorkTicket implements IHHResume.
func (mr *HHMinimalResume) SetWorkTicket() bool {
	return false
}

func NewMinimalResume() IHHResume {
	return &HHMinimalResume{}
}
