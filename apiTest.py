import requests

base_url = "http://localhost:8080/api/v1/"

def post(api, data):
	global base_url
	r = requests.post(base_url + api, data)
	print(r.status_code, r.text)

def get(api):
	global base_url
	print('GET:', base_url + api)
	r = requests.get(base_url + api)
	print(r.status_code, r.text)
	print()

def put(api, data):
	global base_url
	r = requests.put(base_url + api, data)
	print(r.status_code, r.text)

def delete(api):
	global base_url
	r = requests.delete(base_url + api)
	print(r.status_code, r.text)


def user_Test():
	

	get("users")
	# post("user", data={"Name": "admin"})
	# put("user/4", data={"Name": "Master"})
	# get("user/2")
	# delete("user/1")
	get("users")

def article_Test():
	get("articles")
	# post("article", data={"Title": "world", "Body": "AAA"})
	# put("article/1", data={"Title": "world", "Body": "BBB", "CategoryID": 1})
	# get("article/2")
	# delete("article/2")
	get("articles")

def category_Test():
	
	get("categories")
	post("category", data={"Name": "C++"})
	# put("category/1", data={"Name": "java"})
	# get("category/1")
	# delete("category/2")
	get("categories")

if __name__ == '__main__':
	user_Test()
	# article_Test()
	# category_Test()