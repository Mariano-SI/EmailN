package mail

import (
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail() error {

	dialer := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("GMAIL_APP_PASSWORD"))

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("EMAIL_USER"))
	message.SetHeader("To", "marianocarlos42@gmail.com")
	message.SetHeader("Subject", "Hello")
	htmlBody := `
	<!DOCTYPE html>
	<html lang="pt-BR">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Email Teste</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f9f9f9;
				margin: 0;
				padding: 20px;
			}
			.container {
				max-width: 600px;
				margin: 0 auto;
				background: #ffffff;
				border-radius: 8px;
				padding: 20px;
				box-shadow: 0 2px 6px rgba(0,0,0,0.1);
			}
			h1 {
				color: #4F46E5;
				text-align: center;
			}
			p {
				color: #333333;
				line-height: 1.6;
			}
			.footer {
				text-align: center;
				color: #888888;
				font-size: 12px;
				margin-top: 20px;
			}
			.button {
				display: inline-block;
				background-color: #4F46E5;
				color: #ffffff;
				padding: 10px 20px;
				text-decoration: none;
				border-radius: 5px;
				font-weight: bold;
				margin-top: 20px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>Olá, Mariano 👋</h1>
			<p>Este é um e-mail de teste enviado a partir da sua aplicação Go!</p>
			<p>Se você está lendo isso formatado, significa que o envio HTML está funcionando corretamente 🎉</p>
			<a href="https://golang.org" class="button">Visitar Golang</a>
			<div class="footer">
				<p>© 2025 EmailN System</p>
			</div>
		</div>
	</body>
	</html>
	`

	message.SetBody("text/html", htmlBody)

	fmt.Println("Chegou aqui")
	return dialer.DialAndSend(message)
}
