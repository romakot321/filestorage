import requests
import io

account1 = {'username': 'user', 'password': '123'}
account2 = {'username': 'admin', 'password': '123'}
file_content = 'diadfjadspgpiasd'
file_buffer = io.StringIO(file_content)
base_url = "http://localhost:8080/"

auth_resp = requests.post(base_url + 'auth/login', json=account1)
assert auth_resp.status_code == 200, auth_resp.text
token = auth_resp.json()['token']
auth_header_1 = {"Token": token}

auth_resp = requests.post(base_url + 'auth/login', json=account2)
assert auth_resp.status_code == 200, auth_resp.text
token = auth_resp.json()['token']
auth_header_2 = {"Token": token}

file_create_resp = requests.post(base_url + 'files/', files={'file': file_buffer}, headers=auth_header_1)
assert file_create_resp.status_code == 201, file_create_resp.text
created_filename = file_create_resp.json()["Filename"]
print("Created", created_filename)

files_list_resp = requests.get(base_url + 'files/', headers=auth_header_1)
assert files_list_resp.status_code == 200, files_list_resp.text
files = files_list_resp.json()
assert created_filename in [i['Filename'] for i in files], "created file not in list"

file_resp = requests.get(base_url + 'files/' + created_filename, headers=auth_header_1)
assert file_resp.status_code == 200, file_resp.text

file_resp = requests.get(base_url + 'files/' + created_filename, headers=auth_header_2)
assert file_resp.status_code == 401, file_resp.text

file_buffer.seek(0)
file_create_resp = requests.post(base_url + 'files/', files={'file': file_buffer}, headers=auth_header_2)
assert file_create_resp.status_code == 201, file_create_resp.text
created_filename_2 = file_create_resp.json()["Filename"]
print("Created", created_filename_2)

files_list_resp = requests.get(base_url + 'files/', headers=auth_header_2)
assert files_list_resp.status_code == 200, files_list_resp.text
files = files_list_resp.json()
print("User 2 files", files)

file_resp = requests.get(base_url + 'files/' + created_filename_2, headers=auth_header_2)
assert file_resp.status_code == 200, file_resp.text
print(file_resp.text)
