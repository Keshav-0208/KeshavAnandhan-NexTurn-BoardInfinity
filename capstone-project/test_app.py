import requests

BASE_URL = "http://127.0.0.1:5000"

def add_test_users():
    test_users = [
        {"name": "Keshav", "email": "keshav@example.com"},
        {"name": "Dinesh", "email": "dinesh@example.com"},
        {"name": "Agnus", "email": "agnus@example.com"},
        {"name": "Ravi", "email": "ravi@example.com"},
        {"name": "Sita", "email": "sita@example.com"}
    ]
    
    for user in test_users:
        response = requests.post(f"{BASE_URL}/users", json=user)
        assert response.status_code == 201

def test_add_users():
    add_test_users()
    response = requests.get(f"{BASE_URL}/users")
    assert response.status_code == 200
    data = response.json()
    added_users = {user[1] for user in data}
    assert "Keshav" in added_users
    assert "Dinesh" in added_users
    assert "Agnus" in added_users

def test_get_users():
    response = requests.get(f"{BASE_URL}/users")
    assert response.status_code == 200
    assert isinstance(response.json(), list)

def test_get_user_by_id():
    response = requests.get(f"{BASE_URL}/users/11")
    assert response.status_code == 200
    assert response.json()[1] == "Charlie"

def test_add_user_without_email():
    response = requests.post(f"{BASE_URL}/users", json={"name": "NoEmailUser"})
    assert response.status_code == 400

def test_invalid_method():
    response = requests.delete(f"{BASE_URL}/users")
    assert response.status_code in [400, 405]
