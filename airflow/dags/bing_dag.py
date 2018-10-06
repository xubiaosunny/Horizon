# -*- coding: utf-8 -*-
import airflow
from airflow.models import DAG
from airflow.operators.python_operator import PythonOperator

import requests
import json
import time
import datetime
import os
import logging

from common.config import default_args

BING_HOST = "https://cn.bing.com"
WALLPAPER_JSON_URL = '/HPImageArchive.aspx?format=js&idx={idx}&n={n}'
image_dir = os.path.join(os.path.abspath(os.path.dirname(__file__)), '../../dist/bing')

dag = DAG(
    dag_id='Bing', default_args=default_args,
    schedule_interval='0 2 * * *',
)


def get_wallpaper():
    res = requests.get(BING_HOST + WALLPAPER_JSON_URL.format(idx=0, n=8))
    res_dict = json.loads(res.content.decode(encoding='utf-8'))

    for image in res_dict['images']:
        im_res = requests.get(BING_HOST + image['url'])
        date = image['enddate']
        image_name = os.path.join(image_dir, date + '.jpg')

        if os.path.exists(image_name):
            logging.info('[~] %s already exists' % date)
            continue

        with open(image_name, 'wb') as f:
            f.write(im_res.content)
            logging.info('[+] %s download successful' % date)

        time.sleep(1)


def bing_spider(**kwargs):
    get_wallpaper()


bing_spider_task = PythonOperator(
    task_id='bing_spider',
    provide_context=True,
    python_callable=bing_spider,
    dag=dag)
