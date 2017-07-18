package mpgo

import (
	"encoding/json"
	"errors"
	"bytes"
	"net/http"
	// "net/url"
)


type Plan struct{
	Id 					string 		 	`json:"id,omitempty"`
	Application_fee 	int				`json:"application_fee,omitempty"`
	Description 	 	string			`json:"description,omitempty"`
	External_reference 	string			`json:"external_reference,omitempty"`
	Date_created 	 	string			`json:"date_created,omitempty"`
	Last_modified 	 	string			`json:"last_modified,omitempty"`
	Auto_recurring 	 	AutoRecurring	`json:"auto_recurring,omitempty"`
	Live_mode 	 		bool			`json:"live_mode,omitempty"`
	Setup_fee 	 		float64			`json:"setup_fee,omitempty"`
	Metadata 	 		interface{}		`json:"metadata,omitempty"`
	Message 			string 			`json:"message,omitempty"`
	Status 				string 			`json:"status,omitempty"`
	Cause 				[]CauseE 		`json:"cause,omitempty"`  
}

type AutoRecurring struct {
	Frequency 			int 		`json:"frequency,omitempty"`
	Frequency_type 		string 		`json:"frequency_type,omitempty"`
	Transaction_amount 	float64 	`json:"transaction_amount,omitempty"`
	Currency_id 		string 		`json:"currency_id,omitempty"`
	Repetitions 		int 		`json:"repetitions,omitempty"`
	Debit_date 			int8 		`json:"debit_date,omitempty"`
	Free_trial 			FreeTrial 	`json:"free_trial,omitempty"`
}
type FreeTrial struct {
	Frequency 		int 	`json:"frequency,omitempty"`
	Frequency_type 	string 	`json:"frequency_type,omitempty"`
}

type Subscription struct {
	Id 					string 		 	`json:"id,omitempty"`
	Plan_id 			string			`json:"plan_id,omitempty"`
	Payer 	 			Payer			`json:"payer,omitempty"`
	Application_fee 	float64			`json:"application_fee,omitempty"`
	Description 	 	string			`json:"description,omitempty"`
	External_reference 	string			`json:"external_reference,omitempty"`
	Date_created 	 	string			`json:"date_created,omitempty"`
	Last_modified 	 	string			`json:"last_modified,omitempty"`
	Live_mode 	 		bool			`json:"live_mode,omitempty"`
	Start_date 	 		string			`json:"start_date,omitempty"`
	End_date 	 		string			`json:"end_date,omitempty"`
	Metadata 	 		interface{}		`json:"metadata,omitempty"`
	Charges_detail 	 	ChargesDetail	`json:"charges_detail,omitempty"`
	Setup_fee 			float64			`json:"setup_fee,omitempty"`
	Message 			string 			`json:"message,omitempty"`
}

type Payer struct {
	Id 					string 		 		`json:"id,omitempty"`
	Type 				string				`json:"plan_id,omitempty"`
	Email 	 			string				`json:"email,omitempty"`
	Identification 	 	Identification		`json:"identification,omitempty"`
}

type ChargesDetail struct {
	Invoices 				[]Invoices 	`json:"invoices,omitempty"`
	Last_charged_date 	 	string		`json:"last_charged_date,omitempty"`
	Last_charged_amount 	float64		`json:"last_charged_amount,omitempty"`
	Pending_charge_amount 	float64		`json:"pending_charge_amount,omitempty"`
	Pending_charge_periods 	int			`json:"pending_charge_periods,omitempty"`
	Charged_amount 	 		float64		`json:"charged_amount,omitempty"`
	Charged_periods 	 	int			`json:"charged_periods,omitempty"`
	Debt_amount 	 		float64		`json:"debt_amount,omitempty"`
	Debt_periods 	 		int			`json:"debt_periods,omitempty"`
	Next_payment_date 	 	string		`json:"next_payment_date,omitempty"`
}
type Invoices struct {
	Period 	string	`json:"period,omitempty"`
	Id 	 	string	`json:"id,omitempty"`
}

func (c ClientMP) GetPlan(plan_id string) (Plan, error) {

	// plan_id := "fff2e1b31a2f4ce1a337dca1b28a2c1e"
	plan := Plan{}

	url := urlBase + "v1/plans/"+ plan_id + "?access_token=" + c.Access_token

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return plan, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return plan, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&plan); err != nil {
		return plan, errors.New("Plan not found")
	}

	return plan, nil

}

func (c ClientMP) NewSubscription(plan_id, customer_id string) (Subscription, error) {

	// plan_id := "fff2e1b31a2f4ce1a337dca1b28a2c1e"

	type NewPayer struct{
		Id string `json:"id"`
	}
	type NewSubscription struct {
		Plan_id string 					`json:"plan_id"`
		Payer 	NewPayer 				`json:"payer"`
	}

	newSubscription := NewSubscription{ 
		Plan_id: plan_id,
		Payer: NewPayer{ Id : customer_id, },
	}

	subscription := Subscription{}

	url := "https://api.mercadopago.com.mx/subscriptions?access_token=" + c.Access_token

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(newSubscription)

	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return subscription, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return subscription, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&subscription); err != nil {
		return subscription, err
	}

	if subscription.Message != "" {
		return subscription, errors.New(subscription.Message)
	}

	return subscription, nil

}