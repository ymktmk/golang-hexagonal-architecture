## golang-hexagonal-architecture

### アーキテクチャ & 使用ライブラリ

- ヘキサゴナルアーキテクチャ
- gorilla mux(Router)
- gorm(ORM)

### 動作確認

コンテナを立ち上げる(DBに初期データなども投入)
```
docker-compose up
```

ユーザーを取得できるか確認する
```
curl http://localhost:8000/users
```

### 参考文献

https://techblog.zozo.com/entry/zozo-microservice-conventions-in-golang
