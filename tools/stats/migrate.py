
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
import logging
reload(sys)
logging.basicConfig(filename='check.log',level=logging.DEBUG,format='%(asctime)s %(message)s', datefmt='%m/%d/%Y %I:%M:%S %p')
sys.setdefaultencoding('utf8')

con = pymysql.connect(host="121.52.250.37",user="bbbb", passwd="bbbb", db="pgd_mig")



def _migrate(offset, limit):
    with con.cursor() as cursor:
        sql = "select user_id,total_reward_coins,total_consume_coins from coin_trade_total_stats limit %d,%d;" % (offset, limit)
        cursor.execute(sql)

        result = cursor.fetchall()
        reslen = len(result)

        for r in result:
            # print(r)
            user_id = r[0]
            old_data = {
                "user_id": r[0],
                "total_reward_coins": r[1],
                "total_consume_coins": r[2]
            }
            
            # print(r)
            sql = "select user_id,total_reward_coins,total_consume_coins from coin_trade_total_stats_%.3d where user_id=%s" % (user_id % 256, user_id)

            cursor.execute(sql)
            r2 = cursor.fetchone()
            table = "coin_trade_total_stats_%.3d" % (user_id % 256)
            if r2 is None or len(r2) == 0:
                try:
                    sql = "insert {table} set user_id = %s, total_reward_coins = %s, total_consume_coins = %s".format(table=table)
                    cursor.execute(sql, r)
                    con.commit()
                    logging.info("insert uid:%s", user_id)
                except Exception as e:
                    logging.exception("insert uid err:%s", user_id)
            else:
                new_data = {
                    "user_id": r2[0],
                    "total_reward_coins": r2[1],
                    "total_consume_coins": r2[2]
                }

                if old_data != new_data:
                    sql = "update {table} set total_reward_coins ={total_reward_coins}, total_consume_coins ={total_consume_coins} where user_id={user_id} ".format(table=table,user_id=user_id,total_reward_coins=old_data['total_reward_coins'],total_consume_coins=old_data['total_consume_coins'])
                    cursor.execute(sql)
                    con.commit()
                    logging.info("update uid:%s",user_id)

                #print("has exist, uid:%d" % (user_id))

        return reslen


def migrate():
    offset = 0
    limit = 10000
    while True:
        reslen = _migrate(offset, limit)
        if reslen < limit:
            break
        else:
            offset += limit
        logging.info("progress:%d"%(offset))
    logging.info("progress:%d"%(offset))

if __name__ == '__main__':
    migrate()
    con.close()