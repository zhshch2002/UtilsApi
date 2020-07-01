# UtilsApi
一些有帮助的API服务。
```shell script
docker pull zhshch/utils-api
```

```yaml
version: "3"

services:
    app:
        image: zhshch/utils-api
        restart: on-failure
        ports:
            - 4000:4000
```
已经部署的Demo [https://api.imagician.net/](https://api.imagician.net/v1/time)

## 获取IP
* `/v1/ip`返回请求者的IP地址
* [https://api.imagician.net/v1/ip](https://api.imagician.net/v1/ip)

## 获取时间
* `/v1/time`返回当前时间
* [https://api.imagician.net/v1/time](https://api.imagician.net/v1/time)

## 代理请求
* `/v1/proxy?url=https%3A%2F%2Fwiki.imagician.net%2Fsitemap.xml`代理请求并返回结果（注意，只响应Body和content-type）
* **url** 目标源地址
* [https://api.imagician.net/v1/proxy?url=https%3A%2F%2Fwiki.imagician.net%2Fsitemap.xml](https://api.imagician.net/v1/proxy?url=https%3A%2F%2Fwiki.imagician.net%2Fsitemap.xml)

## 代理请求网页
* `/v1/htmlProxy?url=https%3A%2F%2Fimagician.net%2F`代理请求网页，将自动替换链接、脚本、样式文件、图像的链接为绝对地址
* **url** 目标源地址
* [https://api.imagician.net/v1/htmlProxy?url=https%3A%2F%2Fimagician.net%2F](https://api.imagician.net/v1/htmlProxy?url=https%3A%2F%2Fimagician.net%2F)

## rss转json
* `/v1/rss?url=https%3A%2F%2Fimagician.net%2Ffeed%2F`解析rss并转换为json
* **url** 目标源地址
* [https://api.imagician.net/v1/rss?url=https%3A%2F%2Fimagician.net%2Ffeed%2F](https://api.imagician.net/v1/rss?url=https%3A%2F%2Fimagician.net%2Ffeed%2F)

## 代理请求图片
* `/v1/img?url=https%3A%2F%2Fwiki.imagician.net%2Flogo.png`代理请求图片
* **url** 目标源地址
* **rotate** 旋转角度 Ex:`rotate=180&url=https%3A%2F%2Fwiki.imagician.net%2Flogo.png`
* **quality** 图像质量 Ex:`quality=80&url=https%3A%2F%2Fwiki.imagician.net%2Flogo.png`
* **flip** 上下翻转 Ex:`flip&url=https%3A%2F%2Fwiki.imagician.net%2Flogo.png`
* **flop** 左右翻转 Ex:`flop&url=https%3A%2F%2Fwiki.imagician.net%2Flogo.png`
* [https://api.imagician.net/v1/img?url=https%3A%2F%2Fwiki.imagician.net%2Flogo.png](https://api.imagician.net/v1/img?url=https%3A%2F%2Fwiki.imagician.net%2Flogo.png)

## 网站favicon
* `/v1/favicon?url=https%3A%2F%2Fimagician.net%2F`根据url解析favicon地址
* **url** 目标源地址
* 兼容**代理请求图片 /v1/img**的图片操作参数
* [https://api.imagician.net/v1/favicon?url=https%3A%2F%2Fimagician.net%2F](https://api.imagician.net/v1/favicon?url=https%3A%2F%2Fimagician.net%2F)

