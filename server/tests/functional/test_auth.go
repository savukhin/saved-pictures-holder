package functional_test

import (
	"testing"
)

const (
	BASE_URL = "http://localhost:3000"
	API_URL  = BASE_URL + "/v1/api/"
)

func TestAuth(t *testing.T) {
	// Test that authentication works.
	// test_username := "mytestusername"
	// test_password := "mytestpassword"

	// Create a user

	// var form url.Values
	// form = map[string][]string{"username": test_username, "password": test_password, "email": "mymail@gmail.com", "confirm_password": test_password}

	// response, err := http.Post(API_URL+"register", "application/json", strings.NewReader(form.Encode()))

	// assert.Equal(t, 200, r.StatusCode)

	// Login with the user

	// r = requests.post(TEST_URL + 'login', json={'username': test_username, 'password': test_password})

	// assert r.status_code == 200

	// Get the token

	// token = r.json()['token']
	// headers = {'Authorization': 'Bearer ' + token}

	// # Use the token to access a protected route

	// r = requests.get(TEST_URL + 'protected', headers=headers)
	// print(r.json())

	// assert r.status_code == 200

	// # Delete the user

	// r = requests.delete(TEST_URL + 'user', headers=headers)

	// assert r.status_code == 200

	// # Try to access the protected route again

	// r = requests.get(TEST_URL + 'protected', headers=headers)

	// assert r.status_code == 401
}
