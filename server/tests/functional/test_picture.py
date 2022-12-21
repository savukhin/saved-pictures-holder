import requests
import pytest
import json

BASE_URL = 'http://localhost:3000'
TEST_URL = BASE_URL + '/v1/api/'

@pytest.fixture(scope='session')
def user1_headers():
    """Create a user for testing"""
    test_username = "mytestusernameforpicture"
    test_password = 'mytestpasswordforpicture'

    r = requests.post(TEST_URL + 'auth/register', json={'username': test_username, 'password': test_password, 'email': 'mypicturemail1@gmail.com', 'confirm_password': test_password})
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'auth/login', json={'username': test_username, 'password': test_password})
    assert r.status_code == 200

    header = {'Authorization': 'Bearer ' + r.json()['token']}
    return header

@pytest.fixture(scope='session')
def user2_headers():
    """Create a user for testing"""
    test_username = "mytestusernameforpicture2"
    test_password = 'mytestpasswordforpicture2'

    r = requests.post(TEST_URL + 'auth/register', json={'username': test_username, 'password': test_password, 'email': 'mypicturemail2@gmail.com', 'confirm_password': test_password})
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'auth/login', json={'username': test_username, 'password': test_password})
    assert r.status_code == 200

    header = {'Authorization': 'Bearer ' + r.json()['token'], 'Content-Type': 'multipart/form-data'}
    return header

def get_picture(picture_number: int):
    picture_path = "./Picture " + str(picture_number) + ".png"
    files = {'picture': open(picture_path, 'rb')}
    return files

@pytest.fixture(scope='session')
def sample_picture_1():
    return get_picture(1)

@pytest.fixture(scope='session')
def sample_picture_2():
    return get_picture(2)

@pytest.fixture(scope='session')
def sample_picture_3():
    return get_picture(3)

def test_create_picture(user1_headers, sample_picture_1, sample_picture_2, sample_picture_3):
    r = requests.post(TEST_URL + 'folders/create', json={'name': 'mytestfolderforpicture'}, headers=user1_headers)
    assert r.status_code == 200
    folder_id = str(r.json()['ID'])

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture', files=sample_picture_1)
    assert r.status_code == 403

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture', files=sample_picture_1, headers=user1_headers)
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture',  files=sample_picture_2, headers=user1_headers)
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture',  files=sample_picture_2, headers=user1_headers)
    assert r.status_code == 200
