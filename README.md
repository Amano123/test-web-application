# test web application

とりあえず、FrontendとBackendの構築  
オントロジーと文章の連携のために

## Description

* Frontend
  * React
* Backend
  * go
* Database
  * elasticsearch

## 起動コマンド

```bash
docker-compose build
docker-compose run react sh -c "cd react-sample && yarn install"
docker-compose up (-d)
```

## 動作確認

React           : http://localhost:8080/  
Go[echo]        : http://localhost:3000/  
Go[echo][DB]    : http://localhost:3000/DB  
Elastic Search  : http://localhost:9200/  
nginx           : http://localhost:1234/ (Reactにリダイレクトされる)  
