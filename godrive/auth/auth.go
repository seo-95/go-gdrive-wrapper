package auth

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	ACCESS_TOKEN_FILE  = "token.json"
	SECRET_CREDENTIALS = "credentials.json"
	SCOPES             = "https://www.googleapis.com/auth/drive"
	EXPIRY_LAYOUT      = time.RFC3339Nano
)

type secretCredentials struct {
	Client_id                   string
	Project_id                  string
	Auth_uri                    string
	Token_uri                   string
	Auth_provider_x509_cert_url string
	Client_secret               string
	Redirect_uris               []string
}

//standard way to create const []string in Go
func getScopes() []string {
	return []string{SCOPES}
}

// GetAuth provides an http client configured with OAuth 2.0 from Google.
// This client can be used to send authorized API requests.
func GetAuth() *http.Client {

	credentials, err := getCredentialsFromFile(SECRET_CREDENTIALS)
	if err != nil {
		panic(err.Error())
	}

	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	conf := &oauth2.Config{
		ClientID:     credentials.Client_id,
		ClientSecret: credentials.Client_secret,
		RedirectURL:  credentials.Redirect_uris[1],
		Scopes:       getScopes(),
		Endpoint:     google.Endpoint,
	}

	token, err := getAccessToken(conf, ACCESS_TOKEN_FILE)
	if err != nil {
		panic(err.Error())
	}

	client := conf.Client(context.Background(), &token)
	return client

}

func getCredentialsFromFile(fname string) (secretCredentials, error) {
	dat, err := os.ReadFile(fname)
	if err != nil {
		return secretCredentials{}, err
	}
	jsonDat := map[string]interface{}{}
	if err := json.Unmarshal([]byte(dat), &jsonDat); err != nil {
		return secretCredentials{}, err
	}
	// interface{} to map[string]interface{}
	jsonCred := make(map[string]interface{})
	v := reflect.ValueOf(jsonDat["installed"])
	for _, key := range v.MapKeys() {
		jsonCred[key.String()] = v.MapIndex(key).Interface()
	}
	var credentials secretCredentials
	mapstructure.Decode(jsonCred, &credentials)
	return credentials, nil
}

func getAccessToken(conf *oauth2.Config, fname string) (oauth2.Token, error) {
	token, err := getAccessTokensFromFile(fname)
	if err != nil {
		fmt.Printf("%s file not present or corrupted. Requesting new credentials.\n", fname)
		// file not present or corrupted
		newTok, err := requestNewAccessToken(conf)
		if err != nil {
			panic(err.Error())
		}
		token = newTok
		saveTokenToFile(&token, ACCESS_TOKEN_FILE)
	}
	if isTokenExpired(&token) {
		fmt.Println("Token expired. Try to recover or requesting new one.")
		newTok, err := refreshToken(conf, token.RefreshToken)
		if err != nil {
			newTok, err = requestNewAccessToken(conf)
			if err != nil {
				return oauth2.Token{}, err
			}
		}
		token = newTok
		saveTokenToFile(&token, ACCESS_TOKEN_FILE)
	}
	return token, nil
}

func refreshToken(conf *oauth2.Config, refresh string) (oauth2.Token, error) {
	newToken, err := conf.Exchange(context.Background(), refresh)
	if err != nil {
		return oauth2.Token{}, err
	}
	return *newToken, nil
}

func getAccessTokensFromFile(fname string) (oauth2.Token, error) {
	dat, err := os.ReadFile(fname)
	if err != nil {
		return oauth2.Token{}, err
	}
	jsonDat := map[string]interface{}{}
	if err := json.Unmarshal([]byte(dat), &jsonDat); err != nil {
		return oauth2.Token{}, err
	}
	tokenMap := make(map[string]string)
	for key, val := range jsonDat {
		str := fmt.Sprintf("%v", val)
		tokenMap[key] = str
	}
	token := buildOauth2TokenFromMap(tokenMap)
	return token, nil
}

func requestNewAccessToken(conf *oauth2.Config) (oauth2.Token, error) {
	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state")
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	fmt.Println()
	fmt.Println()
	fmt.Print("Enter the authorization code in the final URL:")
	reader := bufio.NewReader(os.Stdin)
	authCode, _ := reader.ReadString('\n')
	fmt.Println()
	// Handle the exchange code to initiate a transport.
	tok, err := conf.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatal(err)
	}
	return *tok, nil
}

func isTokenExpired(token *oauth2.Token) bool {
	return time.Now().After(token.Expiry)
}

func saveTokenToFile(token *oauth2.Token, fname string) {
	file, _ := json.MarshalIndent(token, "", "")

	err := ioutil.WriteFile(fname, file, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func buildOauth2TokenFromMap(tokenMap map[string]string) oauth2.Token {
	var token oauth2.Token
	for key, value := range tokenMap {
		switch key {
		case "access_token":
			token.AccessToken = value
		case "token_type":
			token.TokenType = value
		case "refresh_token":
			token.RefreshToken = value
		case "expiry":
			token.Expiry = strToExpiry(value)
		}
	}
	return token
}

func strToExpiry(str string) time.Time {
	// expected format: 2006-01-02T15:04:05.000+02:00
	expiry, err := time.Parse(EXPIRY_LAYOUT, str)
	if err != nil {
		fmt.Println(err.Error())
	}
	return expiry
}
