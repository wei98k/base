import scrapy


class TitleSpider(scrapy.Spider):
    name = "title"
    allowed_domains = ["c.biancheng.net"]
    start_urls = ["http://c.biancheng.net/"]

    def parse(self, response):
        result = response.xpath('/html/head/title/text()').extract_first()
        print('-' * 60 )
        print(result)
        print('-' * 60)