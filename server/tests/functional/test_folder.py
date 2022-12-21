import pytest
import requests
import time

BASE_URL = 'http://localhost:3000'
TEST_URL = BASE_URL + '/v1/api/'

@pytest.fixture(scope='session')
def user1_headers():
    print("Creating user1")
    """Create a user for testing"""
    test_username = "mytestusernameforfolder"
    test_password = 'mytestpasswordforfolder'

    r = requests.post(TEST_URL + 'auth/register', json={'username': test_username, 'password': test_password, 'email': 'mymail1@gmail.com', 'confirm_password': test_password})
    print(r.json())
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'auth/login', json={'username': test_username, 'password': test_password})
    assert r.status_code == 200

    header = {'Authorization': 'Bearer ' + r.json()['token']}
    return header

@pytest.fixture(scope='session')
def user2_headers():
    """Create a user for testing"""
    test_username = "mytestusernameforfolder2"
    test_password = 'mytestpasswordforfolder2'

    r = requests.post(TEST_URL + 'auth/register', json={'username': test_username, 'password': test_password, 'email': 'mymail2@gmail.com', 'confirm_password': test_password})
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'auth/login', json={'username': test_username, 'password': test_password})
    assert r.status_code == 200

    header = {'Authorization': 'Bearer ' + r.json()['token']}
    return header


@pytest.mark.parametrize('folder_name, status_code', [
    ('Для важных переговоров', 200),
    ('For my only', 200),
    ('Style', 200),
    ('Very long name for my folder (>255)' + 'a'*255, 400),
])
def test_create_folder(user1_headers, folder_name, status_code):
    """Test that registration works."""
    r = requests.post(TEST_URL + 'folders/create', json={'name': folder_name}, headers=user1_headers)

    assert r.status_code == status_code


def test_get_folders(user1_headers, user2_headers):
    """Test that registration works."""
    r = requests.post(TEST_URL + 'folders/create', json={'name': "My private folder"}, headers=user1_headers)
    assert r.status_code == 200

    print(r.json())

    folder_id = str(r.json()['ID'])

    r = requests.post(TEST_URL + 'folders/get/' + folder_id, headers=user1_headers)
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'folders/get/' + folder_id, headers=user2_headers)
    assert r.status_code == 403
