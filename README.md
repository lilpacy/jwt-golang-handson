# jwt golang implements

## format

```sh
make fmt
```

## run

```sh
export SIGNINGKEY=xxxx # or may use direnv etc
make
```

## request

```sh
curl localhost:8080/public
TOKEN=`curl localhost:8080/auth`
curl localhost:8080/private -H "Authorization:Bearer ${TOKEN}"
```

## reference
[https://qiita.com/po3rin/items/740445d21487dfcb5d9f](https://qiita.com/po3rin/items/740445d21487dfcb5d9f)
