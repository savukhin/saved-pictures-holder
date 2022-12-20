import pytest
import requests
import time

BASE_URL = 'http://localhost:3000'
TEST_URL = BASE_URL + '/v1/api/'

@pytest.mark.parametrize('username, password, email, confirm_password, status_code', [
    ('goodname1', 'small', 'correct@email.com', 'small', 400),
    ('goodname1', 'passwordCorrect', 'notAnEmail', 'passwordCorrect', 400),
    ('goodname1', 'passwordCorrect', '', 'passwordCorrect', 400),
    ('bad', 'passwordCorrect', 'correct@email.com', 'passwordCorrect', 400),
    ('goodname1', 'passwordCorrect', 'correct@email.com', 'passwordCorrect', 200),
])
def test_register(username, password, email, confirm_password, status_code):
    time.sleep(5)
    """Test that registration works."""
    r = requests.post(TEST_URL + 'register', json={'username': username, 'password': password, 'email': email, 'confirm_password': confirm_password})

    assert r.status_code == status_code


def test_auth():
    """Test that authentication works."""
    test_username = "mytestusername"
    test_password = 'mytestpassword'

    # Create a user

    r = requests.post(TEST_URL + 'register', json={'username': test_username, 'password': test_password, 'email': 'mymail@gmail.com', 'confirm_password': test_password})

    assert r.status_code == 200

    # Login with the user

    r = requests.post(TEST_URL + 'login', json={'username': test_username, 'password': test_password})

    assert r.status_code == 200

    # Get the token

    token = r.json()['token']
    headers = {'Authorization': 'Bearer ' + token}

    # Use the token to access a protected route

    r = requests.get(TEST_URL + 'protected', headers=headers)

    assert r.status_code == 200

    # Delete the user

    r = requests.delete(TEST_URL + 'user', headers=headers)

    assert r.status_code == 200

    # Try to access the protected route again

    r = requests.get(TEST_URL + 'protected', headers=headers)

    assert r.status_code == 401

