package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type WhoisResponse struct {
	DomainName     string           `json:"domain_name"`
	Registrar      string           `json:"registrar"`
	RegistrarUrl   URLArray         `json:"registrar_url"`
	WhoisServer    string           `json:"whois_server"`
	UpdatedDate    UpdatedDateArray `json:"updated_date"`
	CreationDate   int64            `json:"creation_date"`
	ExpirationDate int64            `json:"expiration_date"`
	NameServers    []string         `json:"name_servers"`
	Emails         string           `json:"emails"`
	Org            string           `json:"org"`
}

type URLArray []string

func (u *URLArray) UnmarshalJSON(data []byte) error {
	var singleURL string
	if err := json.Unmarshal(data, &singleURL); err == nil {
		*u = URLArray{singleURL}
		return nil
	}

	var urls []string
	if err := json.Unmarshal(data, &urls); err != nil {
		return err
	}
	*u = URLArray(urls)
	return nil
}

type UpdatedDateArray []int64

func (u *UpdatedDateArray) UnmarshalJSON(data []byte) error {
	var singleDate int64
	if err := json.Unmarshal(data, &singleDate); err == nil {
		*u = UpdatedDateArray{singleDate}
		return nil
	}

	var dates []int64
	if err := json.Unmarshal(data, &dates); err != nil {
		return err
	}
	*u = UpdatedDateArray(dates)
	return nil
}

func (w *WhoisResponse) GetCreationDate() time.Time {
	return time.Unix(w.CreationDate, 0)
}

func (w *WhoisResponse) GetExpirationDate() time.Time {
	return time.Unix(w.ExpirationDate, 0)
}

func (w *WhoisResponse) GetUpdatedDates() []time.Time {
	updatedDates := make([]time.Time, len(w.UpdatedDate))
	for i, unixTime := range w.UpdatedDate {
		updatedDates[i] = time.Unix(unixTime, 0)
	}
	return updatedDates
}

func main() {

	resp := api_req(get_domain())
	defer resp.Body.Close()
	result := json_process(resp)
	print_whois(result)
	// print_json(json_process(resp))

}

func get_domain() string {
	var domain string
	fmt.Print("Enter domain(example: google.com): ")
	fmt.Scanf("%s", &domain)
	return domain

}

func api_req(domain string) *http.Response {
	// https://api-ninjas.com/api/whois?domain=google.com

	api_url := "https://api.api-ninjas.com/v1/whois?"
	params := url.Values{
		"domain": {domain},
	}

	req, err := http.NewRequest(http.MethodGet, api_url+params.Encode(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Api-Key", os.Getenv("API_KEY"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

func json_process(resp *http.Response) *WhoisResponse {

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result WhoisResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	return &result
}

func print_json(resp *WhoisResponse) {
	pretty, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(pretty))
	// fmt.Println(resp)
}

func print_whois(resp *WhoisResponse) {
	fmt.Println("Domain Name:", resp.DomainName)
	fmt.Println("Registrar:", resp.Registrar)
	fmt.Println("Registrar URL:", resp.RegistrarUrl)
	fmt.Println("Whois Server:", resp.WhoisServer)
	fmt.Println("Creation Date:", resp.GetCreationDate())
	fmt.Println("Expiration Date:", resp.GetExpirationDate())
	fmt.Println("Updated Dates:")
	for _, updatedDate := range resp.GetUpdatedDates() {
		fmt.Println("  -", updatedDate)
	}
	fmt.Println("Name Servers:", resp.NameServers)
	fmt.Println("Emails:", resp.Emails)
	fmt.Println("Org:", resp.Org)
}
