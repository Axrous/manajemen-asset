package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"regexp"
	"unicode"
)

func GenerateUUID() string {
	return uuid.NewString()
}

func ContainsUppercase(s string) bool {
	for _, char := range s {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func ContainsSpecialChar(s string) bool {
	// Regular expression to match any special character
	re := regexp.MustCompile(`[!@#$%^&*()_+=\[{\]};:'",<.>/?]`)
	return re.MatchString(s)
}

// email event nya disini
func SendEmailRegister(emailTo, firstName string) {
	url := "https://api.brevo.com/v3/smtp/email"

	payloadData := struct {
		Sender struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"sender"`
		To []struct {
			Email string `json:"email"`
		} `json:"to"`
		Subject     string `json:"subject"`
		HtmlContent string `json:"htmlContent"`
	}{
		Sender: struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{
			Name:  "Stephanie Project",
			Email: "Stephanie@stephanieproject.my.id",
		},
		To: []struct {
			Email string `json:"email"`
		}{
			{
				Email: emailTo,
			},
		},
		Subject: "Welcome on aboard 🎉",
		HtmlContent: fmt.Sprintf(`
		<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Thank you for joining the Twist beta mailing list!</title>
</head>
<body>
  <table width="100%" cellpadding="0" cellspacing="0">
    <tr>
      <td align="center">
        <img src="https://blockfriend.net/images/hero-image2x.png" alt="Twist logo" width="200">
      </td>
    </tr>
    <tr>
      <td align="center">
        <h1>Dear ` + firstName + `</h1>
        <h1>Thank you for signing up on our website!</h1>
        <h1>We're thrilled to have you as a part of our community.</h1>
        <p>Twist is the communication app for teams who want to create a calmer, more organized, more productive workplace.</p>
        <p>If you have any questions or need assistance, feel free to reach out to our support team.</p>
        <p>Have questions about the project? We'd love to help! Just hit reply :)</p>
      </td>
    </tr>
    <tr>
      <td align="center">
        <p>Our Best,</p>
        <p>Hugo and the Stephanie team</p>
        <p>Made by the team across ten time zones at Doist♡ We're hiring!</p>
      </td>
    </tr>
  </table>
</body>
</html>
`),
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	payloadReader := bytes.NewReader(payloadBytes)

	//newrequest
	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", os.Getenv("MAILER_API_KEY")) //submit your api key

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(body))
}
