package mpgo

import (
	"encoding/json"
	"errors"
	"bytes"
	// "log"
	"net/http"
	// "net/url"
)

type Customer struct {
	Id 					string 		 	`json:"id,omitempty"`
	Email 				string			`json:"email,omitempty"`
	First_name 	 		string			`json:"first_name,omitempty"`
	Last_name 	 		string			`json:"last_name,omitempty"`
	Phone 				Phone			`json:"phone,omitempty"`
	Identification 	 	Identification	`json:"identification,omitempty"`
	Default_address 	string			`json:"default_address,omitempty"`
	Address 	 		Address			`json:"address,omitempty"`
	Date_registered 	string			`json:"date_registered,omitempty"`
	Description 	 	string			`json:"description,omitempty"`
	Date_created 	 	string			`json:"date_created,omitempty"`
	Date_last_updated 	string			`json:"date_last_updated,omitempty"`
	Metadata 	 		interface{}		`json:"metadata,omitempty"`
	Default_card 		string			`json:"default_card,omitempty"`
	Cards 	 			[]Card			`json:"cards,omitempty"`
	Addresses 	 		[]Addresses		`json:"addresses,omitempty"`
	Live_mode 	 		bool			`json:"live_mode,omitempty"`
	Message 			string 			`json:"message,omitempty"`
	Status 				int 			`json:"status,omitempty"`
	Cause 				[]CauseE 		`json:"cause,omitempty"`  
}

type Phone struct {
	Area_code 	string 	`json:"area_code,omitempty"`
	Number 		string 	`json:"number,omitempty"`
}
type Identification struct {
	Type 		string 	`json:"type,omitempty"`
	Number 		string 	`json:"number,omitempty"`
}
type Address struct {
	Id 				string 		 	`json:"id,omitempty"`
	Zip_code 		string 		 	`json:"zip_code,omitempty"`
	Street_name 	string 		 	`json:"street_name,omitempty"`
	Street_number 	int 		 	`json:"street_number,omitempty"`
}
type Addresses struct {
	Id 				string 		 	`json:"id,omitempty"`
	Zip_code 		string 		 	`json:"zip_code,omitempty"`
	Street_name 	string 		 	`json:"street_name,omitempty"`
	Street_number 	int 		 	`json:"street_number,omitempty"`
	Phone 			string 		 	`json:"phone,omitempty"`
	Name 			string 		 	`json:"name,omitempty"`
	Floor 			string 		 	`json:"floor,omitempty"`
	Apartment 		string 		 	`json:"apartment,omitempty"`
	City 			City 		 	`json:"city,omitempty"`
	State 			State 		 	`json:"state,omitempty"`
	Country 		Country 		`json:"country,omitempty"`
	Neighborhood 	Neighborhood 	`json:"neighborhood,omitempty"`
	Municipality 	Municipality 	`json:"neighborhood,omitempty"`
	Comments 		string 		 	`json:"comments,omitempty"`
	Date_created 	string 		 	`json:"date_created,omitempty"`
	Verifications 	interface{} 	`json:"verifications,omitempty"`
}
type City struct {
	Id 		string 	`json:"id,omitempty"`
	Name 	string 	`json:"name,omitempty"`
}
type State struct {
	Id 		string 	`json:"id,omitempty"`
	Name 	string 	`json:"name,omitempty"`
}
type Country struct {
	Id 		string 	`json:"id,omitempty"`
	Name 	string 	`json:"name,omitempty"`
}
type Neighborhood struct {
	Id 		string 	`json:"id,omitempty"`
	Name 	string 	`json:"name,omitempty"`
}
type Municipality struct {
	Id 		string 	`json:"id,omitempty"`
	Name 	string 	`json:"name,omitempty"`
}


type Card struct {
	Id 					string 		 	`json:"id,omitempty"`
	Customer_id 		string 			`json:"customer_id,omitempty"`
	Expiration_month 	int 			`json:"expiration_month,omitempty"`
	Expiration_year 	int 			`json:"expiration_year,omitempty"`
	First_six_digits 	string 			`json:"first_six_digits,omitempty"`
	Last_four_digits 	string 			`json:"last_four_digits,omitempty"`
	Payment_method 		PaymentMethod 	`json:"payment_method,omitempty"`
	Security_code 		SecurityCode 	`json:"security_code,omitempty"`
	Issuer 				Issuer 			`json:"issuer,omitempty"`
	Cardholder 			Cardholder 		`json:"cardholder,omitempty"`
	Date_created 		string 			`json:"date_created,omitempty"`
	Date_last_updated 	string 			`json:"date_last_updated,omitempty"`
}
type PaymentMethod struct {
	Id 					string 	`json:"id,omitempty"`
	Name 				string 	`json:"name,omitempty"`
	Payment_type_id 	string 	`json:"payment_type_id,omitempty"`
	Thumbnail 			string 	`json:"thumbnail,omitempty"`
	Secure_thumbnail 	string 	`json:"secure_thumbnail,omitempty"`
}
type SecurityCode struct {
	Length 			int 	`json:"length,omitempty"`
	Card_location 	string 	`json:"card_location,omitempty"`
}
type Issuer struct {
	Id 		int 	`json:"id,omitempty"`
	Name 	string 	`json:"name,omitempty"`
}
type Cardholder struct {
	Name 			int 			`json:"name,omitempty"`
	identification 	interface{} 	`json:"identification,omitempty"`
}

func (c ClientMP) NewCustomer(newCustomer Customer) (Customer, error) {

	// plan_id := "fff2e1b31a2f4ce1a337dca1b28a2c1e"
	customer := Customer{}

	url := urlBase + "v1/customers?access_token=" + c.Access_token

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(newCustomer)

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return customer, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {

		return customer, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		return customer, errors.New("Plan not found")
	}

	if customer.Status != 0 {
		return customer, errors.New(customer.Message)
	}

	return customer, err

}

func (c ClientMP) GetCustomer(customer_id string) (Customer, error) {

	// plan_id := "fff2e1b31a2f4ce1a337dca1b28a2c1e"
	customer := Customer{}

	url := urlBase + "v1/customers/" + customer_id + "?access_token=" + c.Access_token

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return customer, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {

		return customer, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		return customer, errors.New("Json decode error")
	}

	if customer.Status != 0 {
		return customer, errors.New(customer.Cause[0].Description)
	}

	return customer, err

}

func (c ClientMP) AddNewCustomerCard(customer_id string, tokenCard string) (Customer, error) {

	// plan_id := "fff2e1b31a2f4ce1a337dca1b28a2c1e"
	customer := Customer{}

	type NewTokenCard struct {
		Token 	string 	`json:"token"`
	}

	url := urlBase + "v1/customers/"+ customer_id +"/cards?access_token=" + c.Access_token

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode( NewTokenCard{Token: tokenCard,} )

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return customer, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {

		return customer, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		return customer, errors.New("Json decode error")
	}

	// if customer.Status != 0 {
	// 	return customer, errors.New(customer.Cause[0].Description)
	// }

	return customer, err

}