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
    email: str
    status: str


class UserIn(BaseModel):
    username: str
    password: str
    email: EmailStr
    full_name: Optional[str] = None


class Updatejobout(BaseModel):
    status: str
    urllink: str


@app.get("/")
def read_root():
    return {"Hello": "World"}

# list งานทั้งหมดจากแนน


@app.get("/jobbyPosition")
# GET http://127.0.0.1:8080/jobbyPosition HTTP/1.1
async def read_job_by_position():
    r = requests.get('http://127.0.0.1:8080/jobbyPosition')
    return json.loads(r.content)

# list เฉพาะงานที่ต้องการหา


@app.get("/position/{position}")
# GET http://127.0.0.1:8080/position/Support
async def read_jobbyPosition(position: str):
    r = requests.get('http://127.0.0.1:8080/position/'+position)
    return json.loads(r.content)

# list งานทั้งหมดที่ user สมัครไว้ในระบบ


@app.get("/alluserjob")
# GET http://127.0.0.1:8080/alluserjob
async def read_jobbyPosition(Authorization: str):
    r = requests.get('http://127.0.0.1:8080/alluserjob',
                     headers={"Authorization": Authorization})
    return json.loads(r.content)

# แสดง job  ทั้งหมดของ User ตาม Email


@app.get("/userjob")
# GET http://127.0.0.1:8080/userjob/test77@email.com
async def read_jobbyPosition(Authorization: str):

    r = requests.get('http://127.0.0.1:8080/userjob',
                     headers={"Authorization": Authorization})
    return json.loads(r.content)

# เลือกจบเพื่อบันทึกลงฐานข้อมูล


@app.post("/userjob")
def Create_job(userjob: Createjobin, Authorization: str):
    r = requests.post("http://127.0.0.1:8080/userjob", json={
        "position": userjob.position,
        "company": userjob.company,
        "urllink": userjob.urllink,
        "email": userjob.email,
        "status": userjob.status,
    }, headers={"Authorization": Authorization})
    res_json = json.loads(r.content)
    return res_json

# เปลี่ยนสถานนะของงาน


@app.put("/userjob/{email}")
def Update_job(email: str, Updatejob: Updatejobout, Authorization: str):
    r = requests.put('http://127.0.0.1:8080/userjob/'+email, json={
        "urllink": Updatejob.urllink,
        "status": Updatejob.status,
    }, headers={"Authorization": Authorization})
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
    birthdate: str = "16 May 1994"
    address: str = "Home"
    education: str = "Bachelor degree"
    workExperience: Optional[str] = "work"
    email: str = "kaewks@gmail.com"
    phone: str = "+66863527384"
    password: str = "123456"


@app.post("/register/")
def create_user(register: Registerin):
    r = requests.post("http://127.0.0.1:8080/register", json={
        "firstname": register.firstname,
        "lastname": register.lastname,
        "birthdate": register.birthdate,
        "address": register.birthdate,
        "email": register.email,
        "phone": register.phone,
        "education": register.education,
        "workexperience": register.workExperience,
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
