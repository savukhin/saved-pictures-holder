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
    r = requests.post(TEST_URL + 'folders/create', json={'name': folder_name}, headers=user1_headers)

    assert r.status_code == status_code


def test_get_folders(user1_headers, user2_headers):
    folder_id = 4
    init_count = 3

    r = requests.get(TEST_URL + 'folders/get/all', headers=user1_headers)
    assert r.status_code == 200
    assert len(r.json()["folders"]) == init_count

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 404

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user2_headers)
    assert r.status_code == 404

    """Create a folder for testing"""

    r = requests.post(TEST_URL + 'folders/create', json={'name': "My private folder"}, headers=user1_headers)
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user2_headers)
    assert r.status_code == 403

    r = requests.get(TEST_URL + 'folders/get/all', headers=user1_headers)
    assert r.status_code == 200
    assert len(r.json()["folders"]) == init_count + 1
    assert r.json()["folders"][3]['ID'] == folder_id
    assert r.json()["folders"][3]['Name'] == "My private folder"

    r = requests.delete(TEST_URL + 'folders/delete/' + str(folder_id), headers=user2_headers)
    assert r.status_code == 403

    r = requests.delete(TEST_URL + 'folders/delete/' + str(folder_id), headers=user1_headers)
    print(r.text)
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 410

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user2_headers)
    assert r.status_code == 403

    r = requests.get(TEST_URL + 'folders/get/all', headers=user1_headers)
    assert r.status_code == 200
    assert len(r.json()["folders"]) == init_count

def test_folder_update(user1_headers):
    r = requests.post(TEST_URL + 'folders/create', json={'name': "Name 1"}, headers=user1_headers)
    assert r.status_code == 200
    assert r.json()["Name"] == "Name 1"

    folder_id = r.json()["ID"]

    r = requests.put(TEST_URL + 'folders/update/' + str(folder_id), json={'name': "Name 2"}, headers=user1_headers)
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 200
    assert r.json()["Name"] == "Name 2"

    r = requests.put(TEST_URL + 'folders/update/' + str(folder_id), json={'name': "Name 3"}, headers=user1_headers)
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 200
    assert r.json()["Name"] == "Name 3"

    r = requests.delete(TEST_URL + 'folders/delete/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 200

    r = requests.put(TEST_URL + 'folders/update/' + str(folder_id), json={'name': "Name 4"}, headers=user1_headers)
    assert r.status_code == 410

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 410


def test_folder_delete(user1_headers):
    r = requests.post(TEST_URL + 'folders/create', json={'name': "Name 1"}, headers=user1_headers)
    assert r.status_code == 200
    assert r.json()["Name"] == "Name 1"

    folder_id = r.json()["ID"]

    r = requests.delete(TEST_URL + 'folders/delete/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 410

    r = requests.delete(TEST_URL + 'folders/delete/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 410

    r = requests.get(TEST_URL + 'folders/get/' + str(folder_id), headers=user1_headers)
    assert r.status_code == 410
