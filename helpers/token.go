package helpers

import(
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"os"
)

var mySigningKey = []byte(os.Getenv("MY_JWT_TOKEN"))

// GenerateJWT func
func GenerateJWT(userName string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["authorized"] = true
    claims["client"] = userName
    claims["exp"] = time.Now().Add(time.Hour * 30).Unix()

    tokenString, err := token.SignedString(mySigningKey)

    if err != nil {
        fmt.Errorf("Something Went Wrong: %s", err.Error())
        return "", err
    }

    return tokenString, nil
}