# migrate
```shell script
curl -X GET -H "Content-Type: application/json" localhost:8000/api/v1/migrate
```

# create demo article
```shell script
curl -X POST -H "Content-Type: application/json" -d '{
"title":"新しい起業アイディア",
"overview":"これは飲食業界における画期的なアイディアです",
"problem":"飲食業界における人材不足",
"customerSegment":"飲食店オーナー",
"solution": "画期的な技術を提供します",
"uniqueValue":"OOがユニークな強みです",
"channel":"instagram",
"keyMetric":"月1000万の使用",
"unfairAdvantage":"AIプロフェッショナルを抱えていること",
"costStructure":"広告費、人件費",
"revenueStream":"ユーザーの月間使用料",
"elevatorPitch":"30秒でこのアイディアを伝えると?",
"remarks":"備考",
"status":"ステータス"}' localhost:8000/api/v1/article
```

# read demo all article
```shell script
curl -X GET -H "Content-Type: application/json" localhost:8000/api/v1/article
```

# update demo article
```shell script
curl -X PATCH -H "Content-Type: application/json" -d '{
"title":"新しい起業アイディアNEW!",
"overview":"これは飲食業界における画期的なアイディアです",
"problem":"飲食業界における人材不足",
"customerSegment":"飲食店オーナー",
"solution": "画期的な技術を提供します",
"uniqueValue":"OOがユニークな強みです",
"channel":"instagram",
"keyMetric":"月1000万の使用",
"unfairAdvantage":"AIプロフェッショナルを抱えていること",
"costStructure":"広告費、人件費",
"revenueStream":"ユーザーの月間使用料",
"elevatorPitch":"30秒でこのアイディアを伝えると?",
"remarks":"備考",
"status":"ステータス"}' localhost:8000/api/v1/article
```