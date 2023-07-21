package micro

import (
	"log"
	"net/http"
)

func SendUserRegistration(token string) error {
	req, _ := http.NewRequest("POST", "http://localhost:1323/users/update", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Println("Сука.")
	}
	return nil
}