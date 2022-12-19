
## API Restful - Golang
![alt text](https://raw.githubusercontent.com/nurchulis/api-restfull-go-cake-store/feature/crud/banner_resftull.png)

No Just For Test, i build this for learn and will sharing.

Spesification and Feature
- Build With Golang 1.6 - 1.9
- Support Auto Migration, with SQL Script
- Editable Test Case For Unit Testing With Json List
- Develop Not use ORM 
- Enjoy structure module

Feature Endpoint
- Create Or Insert Data
- Get List and Get Detail Data
- Update Data By ID
- Delete Data By ID (Only Hard Delete)

Structure

    .
    ├── ...
    ├── config                    
    │   ├── config.go
    ├── migration
    │   ├── command
    │   ├──── migrationfile.sql  
    ├── models
    │   ├── cake.go
    ├── service
    │   ├── migration
    │   ├──── migration.go  
    │   ├── query
    │   ├──── cakes_query.go  
    ├── utils
    │   ├── res.go
    ├── main.go       
    └── ...
    
## Endpoint Documentation
https://api.postman.com/collections/5630104-a05b1f25-5d83-46a8-82d2-adb64561e3a7?access_key=PMAT-01GMK6GW4K2SRGWT26YPJ030XS




## Installation

For Installation u can build in docker images if installed with docker

```bash
sudo docker image build -t api-cloud-cake:latest .   
docker container run -p 7000:7000 --env-file ./.env --name api-cloud-cake api-cloud-cake
```
and well running in your container
![alt text](https://i.postimg.cc/m2WyG0jH/image.png)

## Access Demo Endpoint in Here:
http://116.193.190.246:7000 (soon on deployment)

## Quick Test

Endpoint Get List :
- Get List Data ✅
- Get List Data With Empty Data Handled ✅

![alt text](https://i.postimg.cc/nLpWcX1q/image.png)

Endpoint Get Detail
- Get Detail Data by Data ID ✅
- Get Detail Data With Empty Data Handled ✅

![alt text](https://i.postimg.cc/8PhQpgSQ/image.png)

Enpoint Create / Insert Data
- Successful to Insert Data With Json Data ✅
- Will Failed If data type on parameter not correct (example : set rating with string value) ✅

![alt text](https://i.postimg.cc/nhH5ns2B/image.png)

Endpoint Update Data 
- Successful to Update Data With Json Data ✅
- Will Failed If data type on parameter not correct (example : set rating with string value) ✅
- Will Failed Handled If id data is been or already deleted ✅

![alt text](https://i.postimg.cc/NjzxZLzg/image.png)

Endpoint Deleted Data 
- Successful to Delete Data by data ID ✅
- Will Failed Handled If id data is been or already deleted or null ✅

![alt text](https://i.postimg.cc/mD8k61mj/image.png)

## Unit Test

```bash
    go test -v
```
For List Testing Dynamic Setup on test.json, so u can set test case random not changes code

![alt text](https://i.postimg.cc/CMvHbp8B/image.png)

Will Running Like this

![alt text](https://i.postimg.cc/bJqmJNjd/image.png)


Notes Need Improvment
because time is running out
- Handle Check On Auto migration if table exitst dont running query file
- Add Seeders
- For POST Unit Test Need Improve For payload data because hardcode
- etc

## Tech Stack

**Code:** Golang