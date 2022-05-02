from typing import Optional
from unicodedata import name
# from urllib.request import Request
from fastapi import FastAPI, Request
from pydantic import BaseModel, EmailStr
import requests
import json
from maintotaljob import *
from shoppee import *

app = FastAPI()


class Createjobin(BaseModel):
    position: str
    company: str
    urllink: str
    status: str


class UserIn(BaseModel):
    username: str
    password: str
    email: EmailStr
    full_name: Optional[str] = None


class Updatejobout(BaseModel):
    status: str
    urllink: str


# list เฉพาะงานที่ต้องการหา


@app.get("/position/{position}")
# GET http://127.0.0.1:8080/position/Support
async def read_jobbyPosition(position: str):
    r = requests.get('http://127.0.0.1:8080/position/'+position)
    return json.loads(r.content)


@app.get("/company/{company}")
# GET http://127.0.0.1:8080/company/ABC
async def search_by_company(company: str):
    r = requests.get('http://127.0.0.1:8080/company/'+company)
    return json.loads(r.content)


@app.get("/location/{location}")
# GET http://127.0.0.1:8080/location/BKK
async def search_by_location(location: str):
    r = requests.get('http://127.0.0.1:8080/location/'+location)
    return json.loads(r.content)

# list งานทั้งหมดที่ user สมัครไว้ในระบบ


@app.get("/alljob")
# GET http://127.0.0.1:8080/alljob
async def read_all_job():
    r = requests.get('http://127.0.0.1:8080/alljob')
    return json.loads(r.content)

# แสดง job  ทั้งหมดของ User ตาม Email


@app.get("/userjob")
# GET http://127.0.0.1:8080/userjob/test77@email.com
async def read_user_job(Authorization: str):

    r = requests.get('http://127.0.0.1:8080/userjob',
                     headers={"Authorization": Authorization})
    return json.loads(r.content)

# เลือกจบเพื่อบันทึกลงฐานข้อมูล


@app.post("/userjob")
def Create_job(userjob: Createjobin, Authorization: str):
    try:
        r = requests.post("http://127.0.0.1:8080/userjob", json={
            "position": userjob.position,
            "company": userjob.company,
            "urllink": userjob.urllink,
            "status": userjob.status,
        }, headers={"Authorization": Authorization})
    except requests.exceptions.HTTPError as err:
        return err
    res_json = json.loads(r.content)
    return res_json

# เปลี่ยนสถานนะของงาน


@app.put("/userjob")
def Update_job(Updatejob: Updatejobout, Authorization: str):
    try:
        r = requests.put('http://127.0.0.1:8080/userjob', json={
            "urllink": Updatejob.urllink,
            "status": Updatejob.status,
        }, headers={"Authorization": Authorization})
    except requests.exceptions.HTTPError as err:
        return err
    res_json = json.loads(r.content)
    return res_json

# delete เฉพาะงานที่ต้องการหา


@app.delete("/userjob/{uid}")
def read_jobbyPosition(uid: str, Authorization: str):
    r = requests.delete('http://127.0.0.1:8080/userjob/' +
                        uid, headers={"Authorization": Authorization})
    return json.loads(r.content)


@app.get("/register")
# GET http://127.0.0.1:8080/register/kaewks@gmail.com
async def get_user(Authorization: str):
    r = requests.get('http://127.0.0.1:8080/register',
                     headers={"Authorization": Authorization})
    return json.loads(r.content)


class Registerin(BaseModel):
    firstname: str = "Kaew"
    lastname: str = "KS"
    email: str = "kaewks@gmail.com"
    password: str = "123456"


@app.post("/register/")
def create_user(register: Registerin):
    r = requests.post("http://127.0.0.1:8080/register", json={
        "firstname": register.firstname,
        "lastname": register.lastname,
        "email": register.email,
        "password": register.password,
    })

    res_json = json.loads(r.content)

    return res_json


@app.post("/totaljob")
def totaljob():
    jobbkklist = scape_jobbkk()
    jobdbslist = scape_jobdbs()
    jobthailist = scape_jobthai()
    return {
        "jobbkklist": jobbkklist,
        "jobbdslist": jobdbslist,
        "jobthailist": jobthailist
    }


@app.delete("/deletetotaljob")
def delete_totaljob():
    Jobs.drop_collection()


@app.post("/shopee/{keyword}")
def shopee(keyword: str):
    scape_shopee(keyword)
    return
