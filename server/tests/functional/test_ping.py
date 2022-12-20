import pytest
import requests
import time

BASE_URL = 'http://localhost:3000'
TEST_URL = BASE_URL + '/v1/api'

def test_ping():
    """Test that the server is running."""
    r = requests.get(BASE_URL)

    assert r.status_code == 200
