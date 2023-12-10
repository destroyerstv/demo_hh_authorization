package api_hh

import "encoding/json"

type TUserArea struct {
	Id string `json:"id"`
}

type TCitizenship struct {
	Id string `json:"id"`
}

type TContactValue struct {
	Email     string
	City      string
	Country   string
	Formatted string
	Number    string
}

func (tcv TContactValue) MarshalJSON() ([]byte, error) {
	ret := map[string]string{}

	if tcv.Email != "" {
		ret["email"] = tcv.Email

	} else {
		ret["city"] = tcv.City
		ret["country"] = tcv.Country
		ret["formatted"] = tcv.Formatted
		ret["number"] = tcv.Number
	}

	return json.Marshal(ret)
}

type TContact struct {
	Comment          *string       `json:"area,omitempty"`
	NeedVerification *bool         `json:"need_verification,omitempty"`
	Preferred        bool          `json:"preferred,omitempty"`
	Type             Values        `json:"type"`
	Value            TContactValue `json:"value"`
	Verified         *bool         `json:"verified,omitempty"`
}

type TAdditionalAttestation struct {
	Name         string
	Organization string
	Result       string
	Year         string
}

type TElementary struct {
	Id   string
	Name string
}

type TPrimary struct {
	Name           string
	NameId         *string `json:"name_id,omitempty"`
	Organization   *string `json:"organization,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	Result         *string `json:"result,omitempty"`
	ResultId       *string `json:"result_id,omitempty"`
	Year           int
}

type TEducation struct {
	Additional  []*TAdditionalAttestation `json:"additional,omitempty"`
	Attestation []*TAdditionalAttestation `json:"attestation,omitempty"`
	Elementary  []*TElementary            `json:"elementary,omitempty"`
	Level       *Values
	Primary     []*TPrimary
}

type TExpirensArea struct {
	Id   string
	Name string
	Url  string
}

type TLogoUrls struct {
	X90      string `json:"90,omitempty"`
	X240     string `json:"240,omitempty"`
	Original string `json:"original"`
}

type TEmployer struct {
	AlternateUrl string `json:"alternate_url"`
	Id           string
	LogoUrls     *TLogoUrls `json:"logo_urls,omitempty"`
	Name         string
	Url          string
}

type TExperience struct {
	Area        *TExpirensArea `json:"area,omitempty"`
	Company     string
	CompanyId   *string `json:"company_id,omitempty"`
	CompanyUrl  *string `json:"company_url,omitempty"`
	Description string
	Employer    *TEmployer `json:"employer,omitempty"`
	End         *string    `json:"end,omitempty"`
	Industries  *Values    `json:"industries,omitempty"`
	Position    string
	Start       string
}

type TLanguage struct {
	Id    string
	Name  string
	Level Values
}

type TPhoto struct {
	Medium string
	Small  string
	Id     string
}

type TPortfolio struct {
	Id     string
	Medium string
	Small  string
}

type TGender struct {
	Id string
}
