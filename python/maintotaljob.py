import requests
from bs4 import BeautifulSoup
from mongoengine import *


def connect_mongo():
    # connect(host="mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766-Final-Project?retryWrites=true&w=majority")
    connect(host="mongodb+srv://IS766:66L8E4nVUyLBzWkn@is766cluster0.wvp1c.mongodb.net/IS766FinalProject?retryWrites=true&w=majority")
    # connect(host="mongodb+srv://is766:2hxK81IxuIiVbEL4@is766cluster0.7orlx.mongodb.net/IS766-Final-Project?retryWrites=true&w=majority")
    return


connect_mongo()


class Jobs(Document):
    jobsource = StringField()
    position = StringField()
    company = StringField()
    salary = StringField()
    location = StringField()
    urllink = StringField(required=True)


def scape_jobbkk():
    r = requests.Session()
    for x in range(1, 2):
        r = requests.get(
            f'https://www.jobbkk.com/%E0%B8%AB%E0%B8%B2%E0%B8%87%E0%B8%B2%E0%B8%99/%E0%B9%84%E0%B8%AD%E0%B8%97%E0%B8%B5/{x}')
        soup = BeautifulSoup(r.text, 'html.parser')
        for jobbkk in soup.find_all('div', 'jobbkk-list-company'):
            jobsource = 'jobbkk'
            try:
                position = jobbkk.find_all('h6', 'applying')[0].text
            except IndexError:
                position = 'n/a'
            try:
                company = jobbkk.find_all('a', 'hover-work work-com')[0].text
            except IndexError:
                company = 'n/a'
            try:
                salary2 = jobbkk.find_all(
                    'div', 'col-md-12 col-sm-12 col-xs-12 list-company-salary')[0].text.strip('\n')
                salary = salary2.split("\n")[0][17:len(salary2.split("\n")[0])]
                location = salary2.split("\n")[1]
            except IndexError:
                salary = 'n/a'
            try:
                urllink = jobbkk.find_all('a', href=True)[0].attrs['href']
                job = Jobs.objects(urllink=urllink)
                if len(job) != 0:
                    continue
            except IndexError:
                urllink = 'n/a'
            jobbkkDoc = Jobs(urllink=urllink)
            jobbkkDoc.jobsource = jobsource
            jobbkkDoc.position = position
            jobbkkDoc.company = company
            jobbkkDoc.salary = salary
            jobbkkDoc.location = location
            jobbkkDoc.save()


def scape_jobdbs():
    jobdbslist = []
    for x in range(1, 2):
        r = requests.get(
            f'https://th.jobsdb.com/th/th/jobs/%E0%B8%87%E0%B8%B2%E0%B8%99%E0%B9%84%E0%B8%AD%E0%B8%97%E0%B8%B5/{x}')
        soup = BeautifulSoup(r.text, 'html.parser')
        for jobdbs in soup.find_all('div', 'sx2jih0 zcydq856 zcydq8f6 zcydq8n zcydq80 zcydq85u'):
            jobsource = 'jobdbs'
            try:
                position = jobdbs.find_all('span', 'sx2jih0')[0].text
            except IndexError:
                position = 'n/a'
            try:
                company = jobdbs.find_all(
                    'span', 'sx2jih0 zcydq84u _18qlyvc0 _18qlyvc1x _18qlyvc1 _18qlyvca')[0].text
            except IndexError:
                company = 'n/a'
            try:
                salary = jobdbs.find_all(
                    'span', 'sx2jih0 zcydq84u _18qlyvc0 _18qlyvc1x _18qlyvc3 _18qlyvc7')[1].text
            except IndexError:
                salary = 'n/a'
            try:
                location = jobdbs.find_all(
                    'span', 'sx2jih0 zcydq84u zcydq80 iwjz4h0')[0].text
            except IndexError:
                location = 'n/a'
            try:
                urllink2 = jobdbs.find_all('a', href=True)[0].attrs['href']
                urllink = 'https://th.jobsdb.com'+urllink2
                job = Jobs.objects(urllink=urllink)
                if len(job) != 0:
                    continue
            except IndexError:
                urllink = 'n/a'
            jobdbsDoc = Jobs(urllink=urllink)
            jobdbsDoc.jobsource = jobsource
            jobdbsDoc.position = position
            jobdbsDoc.company = company
            jobdbsDoc.salary = salary
            jobdbsDoc.location = location
            jobdbsDoc.save()


def scape_jobthai():
    for x in range(1, 2):
        r = requests.get(
            f'https://www.jobthai.com/%E0%B8%AB%E0%B8%B2%E0%B8%87%E0%B8%B2%E0%B8%99/%E0%B8%87%E0%B8%B2%E0%B8%99%E0%B8%84%E0%B8%AD%E0%B8%A1%E0%B8%9E%E0%B8%B4%E0%B8%A7%E0%B9%80%E0%B8%95%E0%B8%AD%E0%B8%A3%E0%B9%8C-it-%E0%B9%82%E0%B8%9B%E0%B8%A3%E0%B9%81%E0%B8%81%E0%B8%A3%E0%B8%A1%E0%B9%80%E0%B8%A1%E0%B8%AD%E0%B8%A3%E0%B9%8C/{x}')
        soup = BeautifulSoup(r.text, 'html.parser')
        for jobthai in soup.find_all('div', 'ant-row msklqa-8 hQCzHL'):
            jobsource = 'jobthai'
            try:
                position = jobthai.find_all('h2', 'ohgq7e-0 frNqfE')[0].text
            except IndexError:
                position = 'n/a'
            try:
                company = jobthai.find_all('h2', 'ohgq7e-0 gWWIiL')[0].text
            except IndexError:
                company = 'n/a'
            try:
                salary = jobthai.find_all(
                    'span', 'ohgq7e-0 msklqa-5 gfpnRh')[0].text
            except IndexError:
                salary = 'n/a'
            try:
                location = jobthai.find_all('h3', 'ohgq7e-0 ijtKqG')[0].text
            except IndexError:
                location = 'n/a'
            try:
                urllink2 = jobthai.find_all('a', href=True)[0].attrs['href']
                urllink = 'https://www.jobthai.com' + urllink2
                job = Jobs.objects(urllink=urllink)
                if len(job) != 0:
                    continue
            except IndexError:
                urllink = 'n/a'
            jobthaiDoc = Jobs(urllink=urllink)
            jobthaiDoc.jobsource = jobsource
            jobthaiDoc.position = position
            jobthaiDoc.company = company
            jobthaiDoc.salary = salary
            jobthaiDoc.location = location
            jobthaiDoc.save()


# def delete_totaljob():
#     Totaljob.drop_collection()


# app = FastAPI()
# def get_totaljob_mongo():
#     totaljob = []
#     for rec in Totaljob.objects:
#         totaljob.append(
#             {
#                 "jobsource" : rec.jobsource,
#                 "position" : rec.position,
#                 "company" : rec.company,
#                 "salary" : rec.salary,
#                 "location" : rec.location,
#                 "urllink" : rec.urllink,
#             }
#         )
#     return totaljob

# @app.get("/totaljob")
# def get_totaljob():
#     totaljob = get_totaljob_mongo()
#     return {"list_totaljob" : totaljob}


#     menuDoc = Menu(title=menu.title)
#     menuDoc.userId = menu.userId
#     menuDoc.slug = menu.slug
#     menuDoc.summary = menu.summary
#     menuDoc.type = menu.type
#     menuDoc.createdAt = menu.createdAt
#     menuDoc.updatedAt = menu.updatedAt
#     menuDoc.content = menu.content
#     menuDoc.save()
#     create_menu_mongo(menuDoc)
#     return
        # jobbkkDoc = Totaljob(urllink=urllink)
        # jobbkkDoc.jobsource = jobsource
        # jobbkkDoc.position = position
        # jobbkkDoc.company = company
        # jobbkkDoc.salary = salary
        # jobbkkDoc.location = location
        # jobbkkDoc.save()
