+++
tags = ["c"]
categories = ["code"]
date = "2016-04-10T14:59:31+11:00"
title = "Scrapy Notes"
draft = true
+++


$ pip install scrapy
$ cat > myspider.py <<EOF
import scrapy

class BlogSpider(scrapy.Spider):
    name = 'blogspider'
    start_urls = ['https://blog.scrapinghub.com']

    def parse(self, response):
        for title in response.css('h2.entry-title'):
            yield {'title': title.css('a ::text').extract_first()}

        next_page = response.css('div.prev-post > a ::attr(href)').extract_first()
        if next_page:
            yield scrapy.Request(response.urljoin(next_page), callback=self.parse)
EOF
$ scrapy runspider myspider.py


Command "c:\apps\python3\python3.exe -u -c "import setuptools, tokenize;__file__='C:\\Users\
\Harry\\AppData\\Local\\Temp\\pip-build-juixotv0\\Twisted\\setup.py';f=getattr(tokenize, 'op
en', open)(__file__);code=f.read().replace('\r\n', '\n');f.close();exec(compile(code, __file
__, 'exec'))" install --record C:\Users\Harry\AppData\Local\Temp\pip-w4_uq4jb-record\install
-record.txt --single-version-externally-managed --compile" failed with error code 1 in C:\Us
ers\Harry\AppData\Local\Temp\pip-build-juixotv0\Twisted\
