package keycloak

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/vitalmzzz/devsecopshub/internal/repository"
)

func GetToken() repository.Token {

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	grant_type := os.Getenv("GRANT_TYPE")
	client_id := os.Getenv("CLIENT_ID")

	payload := io.MultiReader(
		strings.NewReader("username="),
		strings.NewReader(username),
		strings.NewReader("&"),
		strings.NewReader("password="),
		strings.NewReader(password),
		strings.NewReader("&"),
		strings.NewReader("grant_type="),
		strings.NewReader(grant_type),
		strings.NewReader("&"),
		strings.NewReader("client_id="),
		strings.NewReader(client_id),
	)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://server/auth/realms/cicd/protocol/openid-connect/token", payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var Bearer repository.Token
	json.Unmarshal(body, &Bearer)

	return Bearer
}
