import datetime
import re
import time

import pymysql
import requests
from lxml import etree
from requests.exceptions import RequestException

def call_page(url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.text
        return None
    except RequestException:
        return None


def find_longest_str(str_list):
    '''
    找到列表中字符串最长的位置索引
    先获取列表中每个字符串的长度，查找长度最大位置的索引值即可
    '''
    num_list = [len(one) for one in str_list]
    index_num = num_list.index(max(num_list))
    return str_list[int(index_num)]


def remove_block(items):
    new_items = []
    for it in items:
        f = "".join(it.split())
        new_items.append(f)
    return new_items

# 正则和lxml混用
def parse_html(html):  # 正则专门有反爬虫的布局设置，不适合爬取表格化数据！
    big_list = []
    selector = etree.HTML(html)
    try:

        title = selector.xpath('//*[@id="wrapper"]/div/div/div[1]/article/div[1]/div/h2/a/text()')
        f_title= remove_block(title)
        # 两种类型链接解析
        link = selector.xpath('//*[@id="wrapper"]/div/div/div[1]/article/div[1]/div/h2/a/@href')
        pulish_time = selector.xpath('//*[@id="wrapper"]/div/div/div[1]/article/div[2]/div[1]/span[2]/text()')
        reading_times = selector.xpath('//*[@id="wrapper"]/div/div/div[1]/article/div[2]/div[2]/span/span/text()')


        for i1,i2,i3,i4 in zip(f_title,link,pulish_time,reading_times):
            big_list.append((i1,'https://studygolang.com/'+i2,i3,i4))
        return big_list
    except  ValueError:
        pass




def insertDB(content):
    connection = pymysql.connect(host='127.0.0.1', port=3306, user='root', password='123456', db='LL',
                                 charset='utf8mb4', cursorclass=pymysql.cursors.DictCursor)

    cursor = connection.cursor()
    try:
        cursor.executemany('insert into golang_china (title,link,p_time,reading_times) values (%s,%s,%s,%s)', content)
        connection.commit()
        connection.close()
        print('向MySQL中添加数据成功！')
    except TypeError :
        pass

if __name__ == '__main__':
    for PageNum in range(541,1273):
        url = 'https://studygolang.com/articles?p='+str(PageNum)
        html = call_page(url)
        content = parse_html(html)
        insertDB(content)
        time.sleep(0.5)
        print(PageNum)



# title,link,p_time,reading_times


# create table golang_china(
# id int not null primary key auto_increment,
# title text,
# link varchar(50),
# p_time varchar(25),
# reading_times int
# ) engine=InnoDB  charset=utf8;

# drop table golang_china;

#







