package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

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

func SendEmailWithOTP(emailTo string, otp string) {
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
		Subject: "Log in to Stephanie Project",
		HtmlContent: fmt.Sprintf(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Verify your login</title>
  <!--[if mso]><style type="text/css">body, table, td, a { font-family: Arial, Helvetica, sans-serif !important; }</style><![endif]-->
</head>

<body style="font-family: Helvetica, Arial, sans-serif; margin: 0px; padding: 0px; background-color: #ffffff;">
  <!-- Tambahkan style "text-align: center" pada tabel utama -->
  <table role="presentation"
    style="width: 100%; border-collapse: collapse; border: 0px; border-spacing: 0px; font-family: Arial, Helvetica, sans-serif; background-color: rgb(239, 239, 239); text-align: center;">
    <tbody>
      <tr>
        <td align="center" style="padding: 1rem 2rem; vertical-align: top; width: 100%;">
          <!-- Tambahkan style "margin: 0 auto" untuk mengatur tabel agar ditengahkan -->
          <table role="presentation" style="max-width: 600px; border-collapse: collapse; border: 0px; border-spacing: 0px; text-align: left; margin: 0 auto;">
            <tbody>
              <tr>
                <td style="padding: 40px 0px 0px;">
                  <div style="text-align: center;">
                    <div style="padding-bottom: 20px;"><img src="https://blockfriend.net/images/hero-image2x.png" alt="Company" style="width: 56px;"></div>
                  </div>
                  <div style="padding: 20px; background-color: rgb(255, 255, 255);">
                    <div style="color: rgb(0, 0, 0); text-align: center;">
                      <h1 style="margin: 1rem 0">Verification code</h1>
                      <p style="padding-bottom: 16px">Please use the verification code below to sign in.</p>
                      <p style="padding-bottom: 16px"><strong style="font-size: 130%">`+otp+`</strong></p>
                      <p style="padding-bottom: 16px">If you didn’t request this, you can ignore this email.</p>
                      <p style="padding-bottom: 16px">This code is valid for 30 minutes.</p>
                      <p style="padding-bottom: 16px">Thanks,<br>Stephanie Project team</p>
                    </div>
                  </div>
                  <div style="padding-top: 20px; color: rgb(153, 153, 153); text-align: center;">
                    <p style="padding-bottom: 16px">Made with ♥ in Indonesia</p>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </td>
      </tr>
    </tbody>
  </table>
</body>

</html>

`, otp),
	}

	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	payloadReader := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, payloadReader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("api-key", os.Getenv("MAILER_API_KEY"))

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
