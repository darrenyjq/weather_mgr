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


def _migrate(offset, limit):
    with con.cursor() as cursor:
        sql = "select user_id,account_type,account_name,balance,app_name,has_rewarded_first_login,create_time,update_time from coin_account  limit %d,%d;" % (offset, limit)
        cursor.execute(sql)

        result = cursor.fetchall()
        reslen = len(result)

        for r in result:
            # print(r)
            user_id = r[0]
            app_name = r[4]
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
            # print(r)

            sql = "select * from coin_account_%.3d  where user_id=%d and app_name='%s'" % (user_id % 256, user_id, app_name)
            cursor.execute(sql)
            result2 = cursor.fetchone()
            if result2 is None or len(result2) == 0:
                table = "coin_account_%.3d" % (user_id % 256)
                sql = "insert {table} set user_id = %s, account_type = %s, account_name = %s, balance = %s,app_name = %s,  has_rewarded_first_login = %s, created_at=%s,updated_at=%s".format(table=table)
                cursor.execute(sql, r)
                con.commit()
                # print("insert ok,uid:%d,table:%s" % (user_id, table))
            else:
                print("has exist, uid:%d" % (user_id))

        return reslen


def migrate():
    offset = 0
    limit = 100
    while True:
        reslen = _migrate(offset, limit)
        if reslen < limit:
            break
        else:
            offset += limit


if __name__ == '__main__':
    migrate()
    con.close()
