package usecase

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"poly_news/internal/config"
)

type NewsLetter struct {
	smtpClient *smtp.Client
	conn       *tls.Conn
}

func NewNewsLetter(config config.SmtpConfig) (*NewsLetter, error) {
	// SMTP 연결 설정 (TLS 사용)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false, // 인증서를 검증하도록 설정
		ServerName:         config.Host,
	}

	// SMTP 서버 연결
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), tlsConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect: %s", err.Error()))
	}

	client, err := smtp.NewClient(conn, config.Host)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to create SMTP client: %s", err.Error()))
	}

	auth := smtp.PlainAuth("", config.AuthEmail, config.AuthPassword, config.AuthEmail)
	if err := client.Auth(auth); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to authenticate: %s", err.Error()))
	}

	return &NewsLetter{
		smtpClient: client,
		conn:       conn,
	}, nil
}

func (n *NewsLetter) Close() {
	if n.conn != nil {
		n.conn.Close()
	}
	if n.smtpClient != nil {
		n.smtpClient.Quit()
	}
}
