// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"fmt"
	"io"
	"strconv"
)

type AddSubscriptionInput struct {
	// Rss地址
	URL string `json:"url"`
	// 订阅名称
	DisplayName string `json:"displayName"`
	// 刮削器类型
	Scraper *ScraperEnum `json:"scraper,omitempty"`
	// 刮削器Id
	ScraperID *string `json:"scraperId,omitempty"`
}

type ConfigResult struct {
	User *UserConfigResult `json:"user"`
}

type Mutation struct {
}

type ParseAcgSubscriptionInput struct {
	URL    string     `json:"url"`
	Source SourceEnum `json:"source"`
}

type ParseAcgSubscriptionResult struct {
	Title  string   `json:"title"`
	Files  []string `json:"files"`
	Season int      `json:"season"`
}

type Query struct {
}

type ScrapeSearchInput struct {
	Title    string          `json:"title"`
	Scraper  ScraperEnum     `json:"scraper"`
	Language ScraperLanguage `json:"language"`
}

type ScrapeSearchResult struct {
	Scraper       ScraperEnum                 `json:"scraper"`
	ID            string                      `json:"Id"`
	Title         string                      `json:"title"`
	OriginalTitle string                      `json:"originalTitle"`
	FirstAirDate  string                      `json:"firstAirDate"`
	Overview      string                      `json:"overview"`
	Poster        string                      `json:"poster"`
	Backdrop      string                      `json:"backdrop"`
	Seasons       []*ScrapeSearchSeasonResult `json:"seasons"`
}

type ScrapeSearchSeasonResult struct {
	SeasonID int    `json:"seasonId"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Poster   string `json:"poster"`
}

type UserConfigInput struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserConfigResult struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type APIStatusEnum string

const (
	APIStatusEnumSuccess                APIStatusEnum = "SUCCESS"
	APIStatusEnumInternalServerError    APIStatusEnum = "INTERNAL_SERVER_ERROR"
	APIStatusEnumNotFound               APIStatusEnum = "NOT_FOUND"
	APIStatusEnumUnauthorized           APIStatusEnum = "UNAUTHORIZED"
	APIStatusEnumUserCredentialsError   APIStatusEnum = "USER_CREDENTIALS_ERROR"
	APIStatusEnumCancel                 APIStatusEnum = "CANCEL"
	APIStatusEnumTimeout                APIStatusEnum = "TIMEOUT"
	APIStatusEnumThirdPartyAPIError     APIStatusEnum = "THIRD_PARTY_API_ERROR"
	APIStatusEnumBadRequest             APIStatusEnum = "BAD_REQUEST"
	APIStatusEnumFormValidationError    APIStatusEnum = "FORM_VALIDATION_ERROR"
	APIStatusEnumForbidden              APIStatusEnum = "FORBIDDEN"
	APIStatusEnumDatabaseMigrationError APIStatusEnum = "DATABASE_MIGRATION_ERROR"
)

var AllAPIStatusEnum = []APIStatusEnum{
	APIStatusEnumSuccess,
	APIStatusEnumInternalServerError,
	APIStatusEnumNotFound,
	APIStatusEnumUnauthorized,
	APIStatusEnumUserCredentialsError,
	APIStatusEnumCancel,
	APIStatusEnumTimeout,
	APIStatusEnumThirdPartyAPIError,
	APIStatusEnumBadRequest,
	APIStatusEnumFormValidationError,
	APIStatusEnumForbidden,
	APIStatusEnumDatabaseMigrationError,
}

func (e APIStatusEnum) IsValid() bool {
	switch e {
	case APIStatusEnumSuccess, APIStatusEnumInternalServerError, APIStatusEnumNotFound, APIStatusEnumUnauthorized, APIStatusEnumUserCredentialsError, APIStatusEnumCancel, APIStatusEnumTimeout, APIStatusEnumThirdPartyAPIError, APIStatusEnumBadRequest, APIStatusEnumFormValidationError, APIStatusEnumForbidden, APIStatusEnumDatabaseMigrationError:
		return true
	}
	return false
}

func (e APIStatusEnum) String() string {
	return string(e)
}

func (e *APIStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = APIStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ApiStatusEnum", str)
	}
	return nil
}

func (e APIStatusEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ScraperEnum string

const (
	ScraperEnumTmdb ScraperEnum = "TMDB"
)

var AllScraperEnum = []ScraperEnum{
	ScraperEnumTmdb,
}

func (e ScraperEnum) IsValid() bool {
	switch e {
	case ScraperEnumTmdb:
		return true
	}
	return false
}

func (e ScraperEnum) String() string {
	return string(e)
}

func (e *ScraperEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ScraperEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ScraperEnum", str)
	}
	return nil
}

func (e ScraperEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ScraperLanguage string

const (
	ScraperLanguageZhHans ScraperLanguage = "Zh_HANS"
	ScraperLanguageZhHant ScraperLanguage = "Zh_HANT"
	ScraperLanguageJa     ScraperLanguage = "JA"
	ScraperLanguageEn     ScraperLanguage = "EN"
)

var AllScraperLanguage = []ScraperLanguage{
	ScraperLanguageZhHans,
	ScraperLanguageZhHant,
	ScraperLanguageJa,
	ScraperLanguageEn,
}

func (e ScraperLanguage) IsValid() bool {
	switch e {
	case ScraperLanguageZhHans, ScraperLanguageZhHant, ScraperLanguageJa, ScraperLanguageEn:
		return true
	}
	return false
}

func (e ScraperLanguage) String() string {
	return string(e)
}

func (e *ScraperLanguage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ScraperLanguage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ScraperLanguage", str)
	}
	return nil
}

func (e ScraperLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SourceEnum string

const (
	SourceEnumBangumi SourceEnum = "BANGUMI"
)

var AllSourceEnum = []SourceEnum{
	SourceEnumBangumi,
}

func (e SourceEnum) IsValid() bool {
	switch e {
	case SourceEnumBangumi:
		return true
	}
	return false
}

func (e SourceEnum) String() string {
	return string(e)
}

func (e *SourceEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SourceEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SourceEnum", str)
	}
	return nil
}

func (e SourceEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
