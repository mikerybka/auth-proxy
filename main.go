package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/mikerybka/auth"
	"github.com/mikerybka/twilio"
)

func main() {
	backendURL, err := url.Parse(requireEnvVar("BACKEND_URL"))
	if err != nil {
		fmt.Println("invalid BACKEND_URL")
		os.Exit(1)
	}
	s := &auth.Proxy{
		DB: &auth.DB{
			Dir: requireEnvVar("AUTH_DATA_DIR"),
		},
		TwilioClient: &twilio.Client{
			AccountSID:  requireEnvVar("TWILIO_ACCOUNT_SID"),
			AuthToken:   requireEnvVar("TWILIO_AUTH_TOKEN"),
			PhoneNumber: requireEnvVar("TWILIO_PHONE_NUMBER"),
		},
		BackendURL: backendURL,
	}
	addr := ":" + requireEnvVar("PORT")
	err = http.ListenAndServe(addr, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func requireEnvVar(name string) string {
	v := os.Getenv(name)
	if v == "" {
		fmt.Println(name, "required")
		os.Exit(1)
	}
	return v
}
