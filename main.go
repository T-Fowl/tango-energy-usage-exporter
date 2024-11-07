package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)
import "log"
import "net/http"
import "net/http/cookiejar"
import "github.com/PuerkitoBio/goquery"

type ApiResponse[D bool | string] struct {
	Data D `json:"d"`
}

type AccountData struct {
	UserName                   string `json:"UserName"`
	CityName                   string `json:"CityName"`
	StateName                  string `json:"StateName"`
	Country                    string `json:"Country"`
	CountryCode                string `json:"CountryCode"`
	UserID                     string `json:"UserID"`
	BPNumber                   string `json:"BPNumber"`
	UserType                   string `json:"UserType"`
	Name                       string `json:"Name"`
	Address                    string `json:"Address"`
	AccountNumber              string `json:"AccountNumber"`
	LoginToken                 string `json:"LoginToken"`
	MeterType                  string `json:"MeterType"`
	CustomerType               string `json:"CustomerType"`
	AddressType                string `json:"AddressType"`
	HomeInfoStatus             string `json:"HomeInfoStatus"`
	LanguageCode               string `json:"LanguageCode"`
	DashboardOption            string `json:"DashboardOption"`
	GraphMode                  string `json:"GraphMode"`
	STATUS                     string `json:"STATUS"`
	EmailID                    string `json:"EmailID"`
	CustomerNo                 string `json:"CustomerNo"`
	UtilityAccountNumber       string `json:"UtilityAccountNumber"`
	IsShowGallon               string `json:"IsShowGallon"`
	IsShowHCF                  string `json:"IsShowHCF"`
	IsEnableHideShow           string `json:"IsEnableHideShow"`
	IsExternalPaymentLink      string `json:"IsExternalPaymentLink"`
	ExternalPaymentLink        string `json:"ExternalPaymentLink"`
	IsExternalPowerRateLink    string `json:"IsExternalPowerRateLink"`
	ExternalPowerRateLink      string `json:"ExternalPowerRateLink"`
	IsExternalWaterRateLink    string `json:"IsExternalWaterRateLink"`
	ExternalWaterRateLink      string `json:"ExternalWaterRateLink"`
	IsExternalGasRateLink      string `json:"IsExternalGasRateLink"`
	ExternalGasRateLink        string `json:"ExternalGasRateLink"`
	TimeZoneId                 string `json:"TimeZoneId"`
	Offset                     string `json:"Offset"`
	RegistrationTermCond       string `json:"RegistrationTermCond"`
	RegistrationPrivacyPol     string `json:"RegistrationPrivacyPol"`
	DefaultPaymentType         string `json:"DefaultPaymentType"`
	IsModernStyle              string `json:"IsModernStyle"`
	UptoDecimalPlaces          string `json:"UptoDecimalPlaces"`
	RoleId                     string `json:"RoleId"`
	Latitude                   string `json:"Latitude"`
	Longitude                  string `json:"Longitude"`
	ZipCode                    string `json:"ZipCode"`
	IsDefaultAccount           string `json:"IsDefaultAccount"`
	PrimaryContactNumber       string `json:"PrimeryContactNumber"`
	FirstName                  string `json:"FirstName"`
	LastName                   string `json:"LastName"`
	AlternateContactNumber     string `json:"AlternateContactNumber"`
	DOB                        string `json:"DOB"`
	AddressID                  string `json:"AddressID"`
	CustomerTypeDesc           string `json:"CustomerTypeDesc"`
	DashboardView              string `json:"DashboardView"`
	TemplateTypeIdHomeBusiness string `json:"TemplateTypeId_HomeBusiness"`
	ModuleIdHomeBusiness       string `json:"ModuleId_HomeBusiness"`
	ResponseGuId               string `json:"ResponseGuId"`
	PaymentMode                string `json:"PaymentMode"`
	IsCSRFirstLogin            string `json:"IsCSRFirstLogin"`
	IsFirstLogin               string `json:"IsFirstLogin"`
	SSN                        string `json:"SSN"`
	IsEmailIdAsUserName        string `json:"IsEmailIdAsUserName"`
	MobilePhoneType            string `json:"MobilePhoneType"`
	HomePhoneType              string `json:"HomePhoneType"`
	ContactType                string `json:"ContactType"`
	IsTwoFactor                string `json:"IsTwoFactor"`
	SessionToken               string `json:"SessionToken"`
	SaltValue                  string `json:"SaltValue"`
	TimeZoneName               string `json:"TimeZoneName"`
	CreatedBy                  string `json:"CreatedBy"`
	LoginCount                 string `json:"LoginCount"`
	GroupId                    string `json:"GroupId"`
	AddressSubType             string `json:"AddressSubType"`
	RefreshToken               string `json:"RefreshToken"`
}

type ValidateLoginRequest struct {
	Username             string      `json:"username"`
	Password             string      `json:"password"`
	RememberMe           bool        `json:"rememberme"`
	CalledFrom           string      `json:"calledFrom"`
	ExternalLoginId      string      `json:"ExternalLoginId"`
	LoginMode            string      `json:"LoginMode"`
	UtilityAccountNumber string      `json:"utilityAcountNumber"`
	Token                interface{} `json:"token"`
	IsEdgeBrowser        bool        `json:"isEdgeBrowser"`
}

type UsageRequest struct {
	UsageOrGeneration string `json:"UsageOrGeneration"`
	Type              string `json:"Type"`
	Mode              string `json:"Mode"`
	StrDate           string `json:"strDate"`
	HourlyType        string `json:"hourlyType"`
	SeasonId          int    `json:"SeasonId"`
	WeatherOverlay    int    `json:"weatherOverlay"`
	UsageYear         string `json:"usageyear"`
	MeterNumber       string `json:"MeterNumber"`
	DateFromDaily     string `json:"DateFromDaily"`
	DateToDaily       string `json:"DateToDaily"`
	IsNonAmi          string `json:"IsNonAmi"`
}

type Usage struct {
	ObjUsageGenerationResultSetOne []struct {
		CityName    string `json:"CityName"`
		StateName   string `json:"StateName"`
		CountryName string `json:"CountryName"`
		FromDate    string `json:"FromDate"`
		ToDate      string `json:"ToDate"`
		Zipcode     string `json:"Zipcode"`
	} `json:"objUsageGenerationResultSetOne"`
	ObjUsageGenerationResultSetTwo []struct {
		UsageDate          string      `json:"UsageDate"`
		Hourly             string      `json:"Hourly"`
		ValueCost          float64     `json:"ValueCost"`
		UsageValue         float64     `json:"UsageValue"`
		DemandValue        int         `json:"DemandValue"`
		UsageColorCode     interface{} `json:"UsageColorCode"`
		WeatherUsageDate   string      `json:"WeatherUsageDate"`
		DemandColorCode    string      `json:"DemandColorCode"`
		UsageAttribute1    string      `json:"UsageAttribute1"`
		UsageAttribute2    string      `json:"UsageAttribute2"`
		AmiMeterNumber     string      `json:"Ami_MeterNumber"`
		MeterNumber        string      `json:"MeterNumber"`
		MeterNumber1       string      `json:"Meter_Number"`
		CountRecord        int         `json:"countRecord"`
		RatePlanDetailId   int         `json:"RatePlanDetail_Id"`
		Minute             int         `json:"Minute"`
		Year               int         `json:"Year"`
		Month              int         `json:"Month"`
		Day                int         `json:"Day"`
		HourData           int         `json:"hourData"`
		RatePlanDetailName string      `json:"RatePlanDetailName"`
		ColorCode          string      `json:"colorCode"`
		Description        string      `json:"Description"`
	} `json:"objUsageGenerationResultSetTwo"`
	ObjUsageGenerationResultSetSolar []interface{} `json:"objUsageGenerationResultSet_Solar"`
	ObjUsageGenerationResultSetThree []struct {
		FromDate  string `json:"FromDate"`
		ToDate    string `json:"ToDate"`
		FromDate1 string `json:"From_Date"`
		ToDate1   string `json:"To_Date"`
	} `json:"objUsageGenerationResultSetThree"`
	ListUsageColor     []interface{} `json:"listUsageColor"`
	ListUsageColordata []struct {
		Name      string `json:"name"`
		ColorCode string `json:"colorCode"`
	} `json:"listUsageColordata"`
	ListUsageRangeValue []interface{} `json:"listUsageRangeValue"`
	GetTentativeData    []struct {
		Skey                   int         `json:"Skey"`
		AccountNumber          int         `json:"AccountNumber"`
		UsageDate              string      `json:"UsageDate"`
		SoFar                  float64     `json:"SoFar"`
		ExpectedUsage          float64     `json:"ExpectedUsage"`
		PeakLoad               float64     `json:"PeakLoad"`
		LoadFactor             float64     `json:"LoadFactor"`
		Average                float64     `json:"Average"`
		Highest                float64     `json:"Highest"`
		SoFarColorCode         interface{} `json:"SoFarColorCode"`
		ExpectedUsageColorCode string      `json:"ExpectedUsageColorCode"`
		PeakLoadColorCode      string      `json:"PeakLoadColorCode"`
		LoadFactorColorCode    string      `json:"LoadFactorColorCode"`
		AverageColorCode       interface{} `json:"AverageColorCode"`
		HighestColorCode       interface{} `json:"HighestColorCode"`
		UpToDecimalPlaces      int         `json:"UpToDecimalPlaces"`
		IsOnlyAMI              bool        `json:"IsOnlyAMI"`
		UsageCycle             string      `json:"UsageCycle"`
	} `json:"getTentativeData"`
	GetConfigData  []interface{} `json:"getConfigData"`
	ObjHighestData []struct {
		UsageDate    string  `json:"UsageDate"`
		Value        float64 `json:"Value"`
		UsageValue   float64 `json:"UsageValue"`
		FromDate     string  `json:"FromDate"`
		ToDate       string  `json:"ToDate"`
		AvgCountData int     `json:"AvgCountData"`
	} `json:"objHighestData"`
}

func CreateClient() (*http.Client, error) {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Jar: cookieJar}

	return client, nil
}

func GetDocument(client *http.Client, url string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status Code Error: %d (%s)", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func UnmarshalApiResponse[T any](response *ApiResponse[string]) (*T, error) {
	data := new(T)
	err := json.NewDecoder(strings.NewReader(response.Data)).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SendApiRequest[R bool | string](client *http.Client, url string, reqBody string, token *string) (*ApiResponse[R], error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if token != nil {
		req.Header.Add("csrftoken", *token)
		req.Header.Add("isajax", "1")
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	api := ApiResponse[R]{}

	err = json.Unmarshal(resBody, &api)
	if err != nil {
		return nil, err
	}

	return &api, nil
}

const UrlLoginPage = "https://my.tangoenergy.com/portal/"
const UrlLoginUpdateState = "https://my.tangoenergy.com/Portal/default.aspx/updateState"
const LoginValidateLogin = "https://my.tangoenergy.com/Portal/default.aspx/validateLogin"
const UrlDashboard = "https://my.tangoenergy.com/Portal/Dashboard.aspx"
const UrlUsagePage = "https://my.tangoenergy.com/Portal/usage.aspx"
const UrlUsageLoadUsage = "https://my.tangoenergy.com/Portal/Usages.aspx/LoadUsage"

func Authenticate(client *http.Client, req ValidateLoginRequest) (*AccountData, error) {
	_, err := GetDocument(client, UrlLoginPage)
	if err != nil {
		return nil, err
	}

	_, err = SendApiRequest[bool](client, UrlLoginUpdateState, "", nil)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	validateLoginResponse, err := SendApiRequest[string](client, LoginValidateLogin, string(marshal), nil)
	if err != nil {
		return nil, err
	}
	accountsData, err := UnmarshalApiResponse[[]AccountData](validateLoginResponse)
	if err != nil {
		return nil, err
	}

	dashboard, err := GetDocument(client, UrlDashboard)
	if err != nil {
		return nil, err
	}

	if dashboard.Find(".godomornig").Length() < 1 {
		return nil, fmt.Errorf("could not find dashboard greeting implying login was unsuccessful")
	}

	// TODO: Just returning the first one?
	return &(*accountsData)[0], nil
}

type MeterDetailsResponse struct {
	MeterDetails []struct {
		MeterType       string      `json:"MeterType"`
		MeterNumber     string      `json:"MeterNumber"`
		IsAMI           bool        `json:"IsAMI"`
		Status          int         `json:"Status"`
		Address         interface{} `json:"Address"`
		Meterattribute1 interface{} `json:"Meterattribute1"`
		MeterNumber1    string      `json:"Meter_Number"`
	} `json:"MeterDetails"`
	HighUsage []struct {
		MeterType          string `json:"MeterType"`
		HighUsageColorCode string `json:"HighUsageColorCode"`
	} `json:"HighUsage"`
	MeterTypeDetails []struct {
		MeterType string `json:"MeterType"`
	} `json:"MeterTypeDetails"`
}

func GetMeters(client *http.Client) (*MeterDetailsResponse, error) {
	log.Printf("Getting meters\n")

	page, err := GetDocument(client, UrlUsagePage)
	if err != nil {
		return nil, err
	}

	token, exists := page.Find("#hdnCSRFToken").First().Attr("value")
	if !exists {
		token = ""
	}

	apiResponse, err := SendApiRequest[string](client, "https://my.tangoenergy.com/Portal/Usages.aspx/BindMultiMeter", "{\"MeterType\":\"E\"}", &token)
	if err != nil {
		return nil, err
	}

	response, err := UnmarshalApiResponse[MeterDetailsResponse](apiResponse)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetUsage(client *http.Client, request UsageRequest) (*Usage, error) {
	log.Printf("Getting Usage For Meter %s on date %s for type %s\n", request.MeterNumber, request.StrDate, request.Type)

	reqMarshalled, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	page, err := GetDocument(client, UrlUsagePage)
	if err != nil {
		return nil, err
	}

	token, exists := page.Find("#hdnCSRFToken").First().Attr("value")
	if !exists {
		token = ""
	}

	response, err := SendApiRequest[string](client, UrlUsageLoadUsage, string(reqMarshalled), &token)
	if err != nil {
		return nil, err
	}

	usage, err := UnmarshalApiResponse[Usage](response)
	if err != nil {
		return nil, err
	}

	return usage, nil
}

func CreateAuthenticatedClient(username string, password string) (*http.Client, *AccountData, error) {
	log.Printf("Creating authenticated client")

	client, err := CreateClient()
	if err != nil {
		return nil, nil, err
	}

	account, err := Authenticate(client, ValidateLoginRequest{
		Username:             username,
		Password:             password,
		RememberMe:           false,
		CalledFrom:           "LN",
		ExternalLoginId:      "",
		LoginMode:            "1",
		UtilityAccountNumber: "",
		Token:                nil,
		IsEdgeBrowser:        false,
	})
	if err != nil {
		return nil, nil, err
	}

	return client, account, nil
}

func Run(cfg *config, start time.Time, end time.Time) {
	client, account, err := CreateAuthenticatedClient(cfg.Username, cfg.Password)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Logged in as %s at %s\n", account.Name, account.Address)

	meters, err := GetMeters(client)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(os.Stdout)

	err = w.Write([]string{"MeterNumber", "RatePlanDetailName", "UsageDate", "Hourly", "UsageValue"})
	if err != nil {
		log.Fatalln("Failed to write csv header")
	}

	for d := start; d.After(end) != true; d = d.AddDate(0, 0, 1) {
		time.Sleep(1 * time.Second)
		usage, err := GetUsage(client, UsageRequest{
			UsageOrGeneration: "1",
			Type:              "K",
			Mode:              "MI", // MI, H, D, M
			StrDate:           d.Format("02-January-06"),
			HourlyType:        "H",
			SeasonId:          0,
			WeatherOverlay:    0,
			UsageYear:         "",
			MeterNumber:       meters.MeterDetails[0].MeterNumber,
			DateFromDaily:     "",
			DateToDaily:       "",
			IsNonAmi:          "true",
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, usageElement := range usage.ObjUsageGenerationResultSetTwo {
			//fmt.Printf("%s,%s,%s,%s,%fkwh\n",
			//	usageElement.AmiMeterNumber,
			//	usageElement.RatePlanDetailName,
			//	usageElement.UsageDate,
			//	usageElement.Hourly,
			//	usageElement.UsageValue,
			//)

			err := w.Write([]string{
				usageElement.AmiMeterNumber,
				usageElement.RatePlanDetailName,
				usageElement.UsageDate,
				usageElement.Hourly,
				fmt.Sprintf("%fkwh", usageElement.UsageValue),
			})
			if err != nil {
				log.Fatalln("Failed to write csv line")
			}
		}

		w.Flush()
		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
	}
}

type config struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func loadConfig() *config {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/tango-energy-usage-exporter")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Could not loadConfig config file: %q", err)
	}

	config := &config{}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Could not loadConfig config file: %q", err)
	}

	return config
}

func CreateLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
}

func main() {
	CreateLogger()
	config := loadConfig()

	start := time.Date(2023, time.February, 20, 0, 0, 0, 0, time.Local)
	end := time.Now().AddDate(0, 0, -1)

	Run(config, start, end)
}
