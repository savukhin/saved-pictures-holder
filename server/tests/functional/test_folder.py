import pytest
import requests
import time

BASE_URL = 'http://localhost:3000'
TEST_URL = BASE_URL + '/v1/api/'

@pytest.fixture
def user_headers():
    """Create a user for testing"""
    test_username = "mytestusernameforfolder"
    test_password = 'mytestpasswordforfolder'

    r = requests.post(TEST_URL + 'auth/register', json={'username': test_username, 'password': test_password, 'email': 'mymail@gmail.com', 'confirm_password': test_password})
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'auth/login', json={'username': test_username, 'password': test_password})
    assert r.status_code == 200

    header = {'Authorization': 'Bearer ' + r.json()['token']}
    return header


@pytest.mark.parametrize('folder_name, status_code', [
    ('Для важных переговоров', 200),
    ('For my only', 200),
    ('Style', 200),
])
def test_create_folder(user_headers, folder_name, status_code):
    """Test that registration works."""
    r = requests.post(TEST_URL + 'folders/create', json={'name': folder_name}, headers=user_headers)
    print(r.json())

    assert r.status_code == status_code

