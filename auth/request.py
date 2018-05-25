import requests
url = 'https://master.example.com:8443/oapi/v1'
headers = {'Authorization': 'Bearer Lp2VY3i0bBeENa7kxNWt4ApGuVHRGJ4vat96sjhGqis'}
r=requests.get(url, headers=headers, verify=False)
with open('req.json','w') as f:
  for chunk in r.iter_content(1024):
	f.write(chunk)
f.close()
