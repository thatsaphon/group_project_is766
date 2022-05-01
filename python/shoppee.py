from turtle import delay
from unicodedata import name
import selenium
import bs4
from selenium import webdriver
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.chrome.options import Options
from bs4 import BeautifulSoup
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
import time
import firebase_admin
from firebase_admin import credentials
from firebase_admin import firestore
from pydantic import BaseModel, HttpUrl
from fastapi import FastAPI


def connectfirebase():
    cred = firebase_admin.credentials.Certificate(
        'is766-project-firebase-adminsdk-wpaoi-f4d93f5107.json')
    firebase_admin.initialize_app(cred)


connectfirebase()


def save(collection_id, data):
    db = firestore.client()
    db.collection(collection_id).document().set(data)


def delete(collection_id):
    db = firestore.client()
    db.collection(collection_id).document().delete()


def scape_shopee(keyword: str):
    options = Options()
    options.headless = False
    options.add_experimental_option('excludeSwitches', ['enable-logging'])
    driver = webdriver.Chrome(
        executable_path=r"chromedriver.exe", chrome_options=options)
    driver.get('https://shopee.co.th/')
    thai_button = driver.find_element(
        By.XPATH, '/html/body/div[2]/div[1]/div[1]/div/div[3]/div[1]/button')
    thai_button.click()
    close_button = driver.execute_script(
        'return document.querySelector("shopee-banner-popup-stateful").shadowRoot.querySelector("div.shopee-popup__close-btn")')
    close_button.click()
    search = driver.find_element(
        By.XPATH, '/html/body/div[1]/div/div[2]/div[2]/div/div[1]/div[1]/div/form/input')
    search.send_keys(keyword)
    search.send_keys(Keys.ENTER)
    driver.execute_script("document.body.style.zoom='10%'")
    time.sleep(5)

    data = driver.page_source
    soup = BeautifulSoup(data, 'html.parser')
    for shopee in soup.find_all('div', 'KMyn8J'):
        try:
            product_name = shopee.find_all(
                'div', 'ie3A+n bM+7UW Cve6sh')[0].text
        except IndexError:
            product_name = 'n/a'
        try:
            product_price = shopee.find_all(
                'div', 'vioxXd pw1xTt rVLWG6')[0].text
        except IndexError:
            product_price = 'n/a'
        shopee_list = {"product_name": product_name,
                       "product_price": product_price}

        save(
            collection_id="shopee_product",
            data=shopee_list)
