# test web application
とりあえず、FrontendとBackendの構築のみ連携は未実装  
オントロジーと文章の連携のために

## Description
* Frontend
  * React
* Backend
  * go
* Database(未実装)
  * elasticsearch

```
docker-compose build
docker-compose run react sh -c "cd react-sample && yarn install"
docker-compose up (-d)
docker-compose exec go-echo air
```
