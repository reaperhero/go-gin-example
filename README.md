# go-gin-example

#

- 获取token
http://127.0.0.1:8000/auth?username=test&password=test123

```
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjMiLCJleHAiOjE2MDQ1NTA1OTUsImlzcyI6InJlYXBlcmhlcm8ifQ.JFu6Eapq0IUYsPvir6SnLbsxnC4KuWlKPf78ONw3gkE"
    },
    "msg": "ok"
}
```