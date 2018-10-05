import airflow
import datetime


default_args = {
    'owner': 'xubiao',
    'depends_on_past': False,
    'start_date': airflow.utils.dates.days_ago(1),
    'email': ['xubiaosunny@example.com'],
    # 'email_on_failure': False,
    # 'email_on_retry': False,
    # 'retries': 1,
    # 'retry_delay': datetime.timedelta(minutes=5),
    # 'queue': 'bash_queue',
    # 'pool': 'backfill',
    # 'priority_weight': 10,
    # 'end_date': datetime(2016, 1, 1),
}