# -*- coding: utf-8 -*-
import airflow
from airflow.models import DAG
from airflow.operators.python_operator import PythonOperator

from common.config import default_args
from bing.cn_bing import bing_spider


dag = DAG(
    dag_id='Bing', default_args=default_args,
    schedule_interval='0 2 * * *',
)

bing_spider_task = PythonOperator(
    task_id='bing_spider',
    provide_context=True,
    python_callable=bing_spider,
    dag=dag)
