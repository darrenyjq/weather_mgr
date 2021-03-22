# encoding: utf-8
from pymysql._compat import text_type
from pymysql import cursors
import pymysql
import urllib2
import csv
import datetime
import requests
import logging
import sys

reload(sys)

sys.setdefaultencoding('utf8')

FORMAT = '%(levelname)s  %(message)s'

# con = pymysql.connect(host="cn_db", port=3308, user="root", passwd="root", db="pgd_mig")
con = pymysql.connect(host="121.52.250.37", user="cootek", passwd="cootek", db="pgd_mig")
# con = pymysql.connect(host="pgd-mysql01.corp.cootek.com", user="pgd_mig_rw", passwd="0U30FiDL4ZJ0jayk", db="pgd_mig")


def _check(offset, limit):
    with con.cursor() as cursor:
        sql = "select user_id,account_type,account_name,balance,app_name,has_rewarded_first_login,create_time,update_time from coin_account  limit %d,%d;" % (offset, limit)
        cursor.execute(sql)

        result = cursor.fetchall()
        reslen = len(result)
        for r in result:
            user_id = r[0]
            rl = list(r)
            rl[6] = r[6].strftime("%Y-%m-%d %H:%M:%S")
            rl[7] = r[7].strftime("%Y-%m-%d %H:%M:%S")
            for i, x in enumerate(rl):
                if x is None:
                    if i in (3, 5):
                        rl[i] = 0
                    else:
                        rl[i] = ""
            r = tuple(rl)
            old_data = dict(
                account_type=r[1],
                account_name=r[2],
                balance=r[3],
                app_name=r[4],
                has_rewarded_first_login=r[5],
                create_time=r[6],
                update_time=r[7],
            )
            if not old_data["account_type"]:
                old_data["account_type"] = ""
            if not old_data["account_name"]:
                old_data["account_name"] = ""

            sql = "select user_id,account_type,account_name,balance,app_name,has_rewarded_first_login,created_at,updated_at from coin_account_%.3d  where user_id=%d and app_name='%s'" % (user_id % 256, user_id, old_data["app_name"])
            cursor.execute(sql)
            r2 = cursor.fetchone()
            con.commit()
            if not r2:
                print("not in new table, uid: %d" % (user_id))
            else:
                new_data = dict(
                    account_type=r2[1],
                    account_name=r2[2],
                    balance=r2[3],
                    app_name=r2[4],
                    has_rewarded_first_login=r2[5],
                    create_time=r2[6].strftime("%Y-%m-%d %H:%M:%S"),
                    update_time=r2[7].strftime("%Y-%m-%d %H:%M:%S"),
                )
                if old_data != new_data:
                    print(old_data)
                    print(new_data)
                    break
        return reslen


def check():
    offset = 0
    limit = 100
    while True:
        reslen = _check(offset, limit)
        if reslen < limit:
            print("end")
            break
        else:
            offset += limit


if __name__ == '__main__':
    check()
    con.close()
