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
con = pymysql.connect(host="121.52.250.37", user="bbbb", passwd="bbbb", db="pgd_mig")
# con = pymysql.connect(host="pgd-mysql01.corp.aaaa", user="pgd_mig_rw", passwd="0U30FiDL4ZJ0jayk", db="pgd_mig")


def create_tables():
    cur = con.cursor()
    for i in range(256):
        sql = """
        CREATE TABLE IF NOT EXISTS `coin_account_%.3d`(
            `id` bigint(64) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
            `user_id` bigint(64) unsigned NOT NULL comment '用户 id',
            `app_name` varchar(64) NOT NULL comment 'app 名称',
            `balance` int(32) NOT NULL DEFAULT 0 comment '金币余额',
            `account_type` varchar(32) NOT NULL DEFAULT '' comment '账号类型, PHONE,OAUTH2_GOOGLE...',
            `account_name` varchar(128) NOT NULL DEFAULT '' comment '账号名称' ,
            `has_rewarded_first_login` boolean NOT NULL DEFAULT 0,
            `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP comment '数据创建的时间',
            `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '数据更新的时间',
            PRIMARY KEY (`id`),
            UNIQUE KEY `uk_user_id_app_name` (`user_id`,`app_name`),
            KEY `idx_account_name_account_type` (`account_name`,`account_type`)
        ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '金币账号, 按 uid对256取模';
        """
        print("create table %d " % (i))
        cur.execute(sql % (i))


def delete_table():
    cur = con.cursor()
    sql = "drop table %s"
    for i in range(256):
        table = 'coin_account_%.3d' % (i)
        cur.execute(sql % table)
        print("delete table %s " % table)


if __name__ == '__main__':
    create_tables()
    con.close()
