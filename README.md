# google-token-verifier

Golang port of [OAuth2Client.prototype.verifyIdToken](https://github.com/google/google-auth-library-nodejs/blob/master/lib/auth/oauth2client.js) from [google-auth-library-nodejs](https://github.com/google/google-auth-library-nodejs)

_Verifique google itoken sin realizar una solicitud http a la API tokeninfo._

### Descargar
```sh
    go get github.com/pzentenoe/google-token-verifier
```

### Modo de uso

```go

import (
    "github.com/pzentenoe/google-token-verifier"
)

v := googleAuthIDTokenVerifier.Verifier{}
aud := "xxxxxx-yyyyyyy.apps.googleusercontent.com"
err := v.VerifyIDToken(TOKEN, []string{
    aud,
})
if err == nil {
    claimSet, err := googleAuthIDTokenVerifier.Decode(TOKEN)
    // claimSet.Iss,claimSet.Email ... (See claimset.go)
}
```

### Caracteristicas

- Fetch public key from www.googleapis.com/oauth2/v3/certs
- Respect cache-control in response from www.googleapis.com/oauth2/v3/certs
- JWT Parser
- Check Signature 
- Check IssueTime, ExpirationTime with ClockSkew
- Check Issuer
- Check Audience

### Dependencias

- golang.org/x/oauth2/jws

### Referencias

- http://stackoverflow.com/questions/36716117/validating-google-sign-in-id-token-in-go#
- https://github.com/GoogleIdTokenVerifier/GoogleIdTokenVerifier
