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

## TODO