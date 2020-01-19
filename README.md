# migrate
```shell script
curl -X GET -H "Content-Type: application/json" localhost:8000/api/v1/migrate
```

# create demo article
```shell script
curl -X POST -H "Content-Type: application/json" -d '{"Name":"sensuikan1973", "Age":"100"}' localhost:8000/api/v1/article
```