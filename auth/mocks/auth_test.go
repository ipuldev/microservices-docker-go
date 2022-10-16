package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/briankliwon/microservices-docker-go/auth/pkg/models"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
)

var (
	token        string
	clientID     string
	clientSecret string
)

type UserData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseToken struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Scope        string `json:"scope"`
	Token_type   string `json:"token_type"`
}

func TestSignUp(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	var resp *http.Response

	err = pool.Retry(func() error {
		body := &UserData{
			Username: "UserTest",
			Password: "UserTest123",
			Email:    "test@gmail.com",
		}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Log("Marshaling json error...")
			return err
		}
		responseBody := bytes.NewBuffer(jsonBody)
		resp, err = http.Post("http://localhost/api/auth/signup", "application/json", responseBody)
		if err != nil {
			t.Log("Post Data have some problem...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "Post http error...")
	bodyParsed, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err, "Read response body ...")
	defer resp.Body.Close()
	var responseMapping models.HttpResponse
	err = json.Unmarshal(bodyParsed, &responseMapping)
	require.NoError(t, err, "Parse response body ...")
	clientID = responseMapping.OauthData.ClientID
	clientSecret = responseMapping.OauthData.ClientSecret
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
}

func TestLogin(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	var resp *http.Response

	err = pool.Retry(func() error {
		body := &UserData{
			Username: "UserTest",
			Password: "UserTest123",
		}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Log("Marshaling json error...")
			return err
		}
		responseBody := bytes.NewBuffer(jsonBody)
		resp, err = http.Post("http://localhost/api/auth/login", "application/json", responseBody)
		if err != nil {
			t.Log("Post Data have some problem...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "Post http error...")
	bodyParsed, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err, "Read response body ...")
	defer resp.Body.Close()
	var responseMapping models.HttpResponse
	err = json.Unmarshal(bodyParsed, &responseMapping)
	require.NoError(t, err, "Parse response body ...")
	clientID = responseMapping.OauthData.ClientID
	clientSecret = responseMapping.OauthData.ClientSecret
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
}

func TestToken(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	var resp *http.Response
	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost/api/auth/token?grant_type=client_credentials&scope=read&client_id=", clientID, "&client_secret=", clientSecret))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	var responseMapping ResponseToken
	err = json.Unmarshal(body, &responseMapping)
	require.NoError(t, err, "Parse response body ...")
	token = responseMapping.Access_token
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
}

func TestAuthorize(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	var resp *http.Response
	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost/api/auth/authorize?access_token=", token))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	_, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
}
