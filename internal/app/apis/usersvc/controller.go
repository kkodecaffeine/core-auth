package usersvc

import (
	"net/http"
	"strconv"
	"time"

	"core-auth/internal/pkg/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

var temp = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

// Controller type definition
type Controller struct {
	usecase user.UseCase
}

// NewController returns new controller instance.
func NewController(e *gin.Engine, au user.UseCase) Controller {
	nc := Controller{au}

	e.POST("/sign-in", nc.SignIn)

	return nc
}

var ACCESS_SECRET = viper.GetString(`token.ACCESS_SECRET`)
var REFRESH_SECRET = viper.GetString(`token.REFRESH_SECRET`)

func (nc *Controller) SignIn(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	// compare the user from the request, with the one we defined:
	// TODO. DB에 값이 있는지 확인
	if temp.Username != u.Username || temp.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	// 토큰 발급
	ts, err := CreateToken(temp.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// saveErr := CreateAuth(user.ID, ts)
	// if saveErr != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	// }

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func CreateToken(userid uint64) (td TokenDetails, err error) {
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.AccessUuid + "++" + strconv.Itoa(int(userid))

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(ACCESS_SECRET))
	if err != nil {
		return
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(REFRESH_SECRET))
	if err != nil {
		return
	}

	return td, nil
}
