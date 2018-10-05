import airflow
from airflow import models, settings
from airflow.contrib.auth.backends.password_auth import PasswordUser
user = PasswordUser(models.User())
user.username = 'xxx'
user.email = 'xxx@xxx.com'
user.password = 'xxx'
user.is_superuser = True
session = settings.Session()
session.add(user)
session.commit()
session.close()
