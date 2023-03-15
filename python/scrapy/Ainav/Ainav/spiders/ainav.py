import scrapy
from scrapy import Selector, Request
from Ainav.items import AinavItem

class AinavSpider(scrapy.Spider):
    name = "ainav"
    allowed_domains = ["www.ainavpro.com"]
    start_urls = ["http://www.ainavpro.com"]

    def parse(self, response):
        sel = Selector(response)
        list_items = sel.css('div.url-body')
        for list_item in list_items:
            ainav_item = AinavItem()
            ainav_item['title'] = list_item.css('strong::text').extract_first()
            ainav_item['info'] = list_item.css('a > div > div > div > p::text').extract_first()
            ainav_item['url'] = list_item.xpath('div/div/a/@href').get()
            ainav_item['logo'] = list_item.xpath('a/div/div/div/img[@data-src]/@data-src').get()
            goto_url = list_item.css('div > div > a::attr(href)').extract_first()
            if goto_url is not None:
                yield Request(
                    url=response.urljoin(goto_url), callback=self.parse_target,
                    cb_kwargs={'item': ainav_item}
                )
            else:
                self.logger.warning("goto_url is None for list_item: %s", list_item)

    def parse_target(self, response, **kwargs):
        ainav_item = kwargs['item']
        sel = Selector(response)
        ainav_item['target'] = sel.css('span.loading-url::text').extract_first()
        yield ainav_item