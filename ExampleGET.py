import requests

ip_address = "1.1.1.1"
url = f"https://open-ip-api.vercel.app/api?ip={ip_address}"

response = requests.get(url)

if response.status_code == 200:
    data = response.json()
    print(f"IP: {data['ip']}\nCountry: {data['country']}\nCity: {data['city']}")
else:
    print(f"Error: {response.status_code}")
