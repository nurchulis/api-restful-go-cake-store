
## API Restful - Golang
No Just For Test, i build this for learn and will sharing.

Spesification and Feature
- Build With Golang 1.6 - 1.9
- Support Auto Migration, with SQL Script
- Develop Not use ORM 
- Enjoy structure module

Feature Endpoint
- Create Or Insert Data
- Get List and Get Detail Data
- Update Data By ID
- Delete Data By ID


## Installation

For Installation u can build in docker images if installed with docker

```bash
sudo docker image build -t api-cloud-cake:latest .   
docker container run -p 7000:7000 --env-file ./.env --name api-cloud-cake api-cloud-cake
```

    
## Quick Test

Endpoint Get List :
- Get List Data ✅
- Get List Data With Empty Data Handled ✅

Endpoint Get Detail
- Get Detail Data ✅
- Get Detail Data With Empty Data Handled ✅

Enpoint Create / Insert Data
- Successful to Insert Data With Json Data ✅
- Will Failed If data type on parameter not correct (example : set rating with string value) ✅

Endpoint Update Data 
- Successful to Update Data With Json Data ✅
- Will Failed If data type on parameter not correct (example : set rating with string value) ✅
- Will Failed Handled If id data is been or already deleted ✅

