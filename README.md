# OAuth2Server

## 메소드

1. /credentials에 가서 client_id와 client_secret을 받습니다.

```ini
http://domain.tld/credentials
```

2. /token?grant_type에 적정한 scope와 승인 타입, clientId, clientSecret을 넣어줍니다.  
   이때 clientId, clientSecret은 1단계에서 생성한 값입니다.  
   `P.S. 대괄호는 제거해야 합니다.`

```ini
http://domain.tld/token?grant_type=client_credentials&client_id=[client_id]&client_secret=[client_secret]&scope=all
```

3. protected에다가 발급받은 토큰을 입력합니다.  
   `P.S. 대괄호는 제거해야 합니다.`

```ini
http://domain.tld/protected?access_token=[TOKEN]
```
