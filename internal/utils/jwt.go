package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = "EyIJSvCKbmlZbm9ysTpWkNaMAl122Sf7"

var TokenRemeberDuration = 86400 * 7 // 一定时间内免登录, todo: 后续需要放入配置文件中

// jwt payload
type MyJwtClaims struct {
	UserId uint `json:"userId"`
	jwt.StandardClaims
}

// 重写jwt校验
// @see StandardClaims.Valid()
// token 过期规则： 若token已过期，检查签发时间， 若是当前时间在签发时间的一定时间内（比如7天内免登录），
// 则说明在有效期间使用过，则重新签发一个token, 不需要重新登录， 若是token已过期，签发时间也大于某个值，需要重新登录
func (c MyJwtClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := time.Now().Unix()

	if !c.VerifyIssuedAt(now, false) {
		vErr.Inner = fmt.Errorf("Token used before issued")
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if !c.VerifyNotBefore(now, false) {
		vErr.Inner = fmt.Errorf("token is not valid yet")
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}

	if vErr.Errors == 0 {
		return nil
	}

	return vErr
}

// 生成JWT Token
func GenerateToken(userId uint) (string, error) {
	now := time.Now().Unix()
	// payload
	claims := MyJwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// token有效期，1天， 但是在TokenRemeberDuration时间内都可以免登录， 在最长TokenRemeberDuration之内都没有登录过， 则需要重新登录
			// ExpiresAt: now + 86400,
			ExpiresAt: now + 3,
			// 签发时间
			IssuedAt: now,
			// 发行人
			// Issuer: "starfileCloud",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, err
}

// 解析校验JWT Token
func ParseToken(token string) (*MyJwtClaims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &MyJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return tokenClaim.Claims.(*MyJwtClaims), err
}

func main() {
	// s, err := GenerateToken(258)
	// fmt.Println("token: ", s, ", err:", err)

	res, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjI1OCwiZXhwIjoxNjMwMjEwOTg3LCJpYXQiOjE2MzAyMTA5ODd9.wasI23AQYCh1kBT3bcST3xDgGfx0NLmulDFJqTOiIHQ")
	fmt.Println("res: ", res, ", err:", err)
	if err == nil {
		fmt.Println(res.UserId + 3)
	}
}
