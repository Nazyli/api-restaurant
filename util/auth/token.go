package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken . . .
func CreateToken(uid string, scopes *string) (token *Token, err error) {
	// duration set 3600 seconds
	duration := (time.Hour * 1).Seconds()

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["scope"] = scopes
	claims["sub"] = uid
	claims["iat"] = time.Now().Unix()                    //Token create
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := tk.SignedString([]byte(os.Getenv("API_SECRET")))
	token = &Token{
		TokenType: "Bearer",
		Duration:  duration,
		Token:     accessToken,
	}
	return token, err

}

// TokenValid . . .
func TokenValid(r *http.Request, scopes ...string) (errsMsg string) {
	tokenString := ExtractToken(r)
	if len(tokenString) == 0 {
		return "Missing authorization header"
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		return "Token has been expired"
	}
	if err != nil {
		log.Println(err)
		return "Invalid Token"
	}
	isExist := cekScopes(claims, scopes...)
	if !isExist {
		return "Invalid Scope"
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return ""
}
func cekScopes(claims jwt.MapClaims, scopes ...string) (isExist bool) {
	isExist = false
	cekScopes, isClaims := claims["scope"].(string)
	if !isClaims {
		log.Println("error invalid Scope")
		return isExist
	}
	tokenScopes := strings.Split(cekScopes, ",")
	i := make(map[string]bool, len(tokenScopes))
	// create hashmap boolean scope yg dibutuhkan
	for _, j := range tokenScopes {
		i[j] = true
	}
	//jika scope ada di dalam hashmap return true
	for _, j := range scopes {
		_, isExist = i[j]
		if isExist {
			return true
		}
	}
	return isExist
}

// ExtractToken . . .
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
func IsAdmin(r *http.Request) (string, bool) {
	scope := "su"
	uid := ""
	tokenString := ExtractToken(r)
	if len(tokenString) == 0 {
		return "", false
	}
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		return "", false
	}
	if err != nil {
		log.Println(err)
		return "", false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid = fmt.Sprintf("%s", claims["sub"])
	}
	isExist := cekScopes(claims, scope)
	// jika bukan admin return false
	if !isExist {
		return uid, false
	}
	// jika admin return true
	return uid, true
}

// ExtractTokenID . . .
func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["sub"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	_, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	// fmt.Println(string(b))
}
