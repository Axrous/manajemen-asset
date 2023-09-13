package helper

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"os"
	"strconv"
	"time"
)

// init otp for email in here
func GenerateOTP() (int, error) {
	passcode := GeneratePassCode(os.Getenv("jwt_secret"))
	otpInt, err := strconv.Atoi(passcode)
	if err != nil {
		return 0, err
	}

	return otpInt, nil
}

func GeneratePassCode(utf8string string) string {
	secret := base32.StdEncoding.EncodeToString([]byte(utf8string))
	passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA512,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}
