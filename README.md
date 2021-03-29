## Fiber Gorm
golang / postman / newman / k6

---

##### Start service
```
make docker.up
```

##### Newman
[https://github.com/postmanlabs/newman#installation](https://github.com/postmanlabs/newman#installation)
```
newman run ./api/Fiber-Gorm.postman_collection.json \
    -e ./api/Fiber-Gorm.postman_environment.json
```

##### Postman to k6
[https://k6.io/blog/load-testing-with-postman-collections](https://k6.io/blog/load-testing-with-postman-collections)
```
postman-to-k6 Fiber-Gorm.postman_collection.json \
    -e Fiber-Gorm.postman_environment.json -o k6-script.js
```
or
```
make newman
```

##### K6
[https://k6.io/docs/getting-started/installation](https://k6.io/docs/getting-started/installation)
```
k6 run k6-script.js
```

#### Stop service
```
make docker.down
```