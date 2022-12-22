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

@pytest.fixture(scope='session')
def some_folder_id(user1_headers):
    r = requests.post(TEST_URL + 'folders/create', json={'name': 'mytestfolderforpicture'}, headers=user1_headers)
    assert r.status_code == 200

    folder_id = r.json()['id']
    return folder_id

def test_create_picture(some_folder_id, user1_headers, sample_picture_1, sample_picture_2, sample_picture_3):
    folder_id = str(some_folder_id)

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture', files=sample_picture_1)
    assert r.status_code == 403

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture', files=sample_picture_1, headers=user1_headers)
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture',  files=sample_picture_2, headers=user1_headers)
    assert r.status_code == 200

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture',  files=sample_picture_3, headers=user1_headers)
    assert r.status_code == 200


def test_get_pictures(some_folder_id, user1_headers):
    folder_id = str(some_folder_id)

    r = requests.get(TEST_URL + 'folders/' + folder_id + '/pictures', headers=user1_headers, params={'offset': 0, 'limit': 1})
    assert r.status_code == 200
    print(r.json())
    assert len(r.json()['pictures']) == 1

    r = requests.get(TEST_URL + 'folders/' + folder_id + '/pictures', headers=user1_headers, params={'offset': 1, 'limit': 1})
    assert r.status_code == 200
    assert len(r.json()['pictures']) == 1

    r = requests.get(TEST_URL + 'folders/' + folder_id + '/pictures', headers=user1_headers, params={'offset': 0, 'limit': 4})
    assert r.status_code == 200
    assert len(r.json()['pictures']) == 3

    r = requests.get(TEST_URL + 'folders/' + folder_id + '/pictures', headers=user1_headers, params={'offset': 10, 'limit': 5})
    assert r.status_code == 200
    assert len(r.json()['pictures']) == 0

def test_update_picture(some_folder_id, user1_headers, user2_headers):
    folder_id = str(some_folder_id)

    r = requests.get(TEST_URL + 'folders/' + folder_id + '/pictures', headers=user1_headers, params={'offset': 0, 'limit': 1})
    assert r.status_code == 200
    picture_id = str(r.json()['pictures'][0]['id'])

    r = requests.post(TEST_URL + 'picture/' + picture_id + "/update", 
        json={'title': 'title 1', 'description': 'description 1' }, 
        headers=user2_headers
    )
    assert r.status_code == 403

    r = requests.post(
        TEST_URL + 'picture/' + picture_id + "/update", 
        json={'title': 'title 1', 'description': 'description 1' }, 
        headers=user1_headers
    )
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'picture/' + picture_id, headers=user2_headers)
    assert r.status_code == 403

    r = requests.get(TEST_URL + 'picture/' + picture_id, headers=user1_headers)
    print(r.text)
    assert r.status_code == 200
    assert r.json()['title'] == 'title 1'
    assert r.json()['description'] == 'description 1'


def test_delete_picture(some_folder_id, user1_headers, user2_headers, sample_picture_1):
    folder_id = str(some_folder_id)

    r = requests.post(TEST_URL + 'folders/' + folder_id + '/create-picture', files=sample_picture_1, headers=user1_headers)
    assert r.status_code == 200
    print(r.text)
    picture_id = str(r.json()['id'])

    r = requests.get(TEST_URL + 'picture/' + picture_id, headers=user1_headers)
    assert r.status_code == 200

    r = requests.delete(TEST_URL + 'picture/' + picture_id, headers=user2_headers)
    assert r.status_code == 403

    r = requests.delete(TEST_URL + 'picture/' + picture_id, headers=user1_headers)
    assert r.status_code == 200

    r = requests.get(TEST_URL + 'picture/' + picture_id, headers=user1_headers)
    assert r.status_code == 410
