import requests
import json
import time
import datetime
import os
import logging

BING_HOST = "https://cn.bing.com"
WALLPAPER_JSON_URL = '/HPImageArchive.aspx?format=js&idx={idx}&n={n}'
image_dir = os.path.join(os.path.abspath(os.path.dirname(__file__)), '../../../dist/bing')


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
            save_copyright(image)
        
        time.sleep(1)


def save_copyright(image):
    import sqlite3
    conn = sqlite3.connect(os.path.join(image_dir, 'bing.db'))
    cursor = conn.cursor()
    try:
        cursor.execute('''create table cn_bing (
                            date varchar(25) primary key,
                            copyright varchar(1000)
                            )''')
    except Exception:
        pass

    cursor.execute('insert into cn_bing (date, copyright) values (?, ?)', (image['enddate'], image['copyright']))
    # cursor.execute("SELECT * FROM cn_bing;")
    # print(cursor.fetchall())

    cursor.close()
    conn.commit()
    conn.close()



def bing_spider(**kwargs):
    get_wallpaper()


if __name__ == "__main__":
    bing_spider()
