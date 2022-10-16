package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/briankliwon/microservices-docker-go/product/pkg/models"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
)

var (
	token        string
	clientID     string
	clientSecret string
	productID    string
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

type HttpResponse struct {
	Message   string    `json:"message"`
	OauthData Oauth2Key `json:"oauth_data"`
}

type Oauth2Key struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
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
	var responseMapping HttpResponse
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
	require.NoError(t, err, "Read response body ...")
	var responseMapping ResponseToken
	err = json.Unmarshal(body, &responseMapping)
	require.NoError(t, err, "Parse response body ...")
	token = responseMapping.Access_token
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
}
func TestInsertProduct(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	var resp *http.Response

	err = pool.Retry(func() error {
		body := &models.Product{
			Name:        "test",
			Description: "Description test",
			Price:       15000,
		}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Log("Marshaling json error...")
			return err
		}
		responseBody := bytes.NewBuffer(jsonBody)
		log.Println(fmt.Sprint("http://localhost/api/product/?access_token=", token))
		resp, err = http.Post(fmt.Sprint("http://localhost/api/product/?access_token=", token), "application/json", responseBody)
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
	var responseMapping models.InsertResponse
	err = json.Unmarshal(bodyParsed, &responseMapping)
	require.NoError(t, err, "Parse response body ...")
	productID = responseMapping.Insert_id
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
}

func TestGetProduct(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	var resp *http.Response
	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost/api/product/?access_token=", token))
		log.Println(fmt.Sprint("http://localhost/api/product/?access_token=", token))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	Body, err := io.ReadAll(resp.Body)
	log.Println(string(Body))
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
}

func TestGetDetailProduct(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	var resp *http.Response
	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost/api/product/", productID, "?access_token=", token))
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

func TestDeleteProduct(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")
	// Create client
	client := &http.Client{}

	var resp *http.Response
	err = pool.Retry(func() error {
		req, err := http.NewRequest("DELETE", fmt.Sprint("http://localhost/api/product/", productID, "?access_token=", token), nil)
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		resp, err = client.Do(req)
		if err != nil {
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.Status, "HTTP status code")
}
