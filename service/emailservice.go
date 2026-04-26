package service

import (
	"fmt"
	"log"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"

	"github.com/wneessen/go-mail"
)

func SendEmail(emailAccountId int64, toEmailAddress string, subject string, message string) error {

	// get account
	account, err := dao.GetEmailAccountById(emailAccountId)

	if err != nil {
		log.Printf("An error has occured with trying to get the account information for sending the email : %s", err.Error())
		return err
	}

	// get the corresponding smtp config required for this account:

	emailConfig, err := dao.GetEmailConfigById(account.SmtpConfigId)
	if err != nil {
		log.Printf("An error has occured with trying to get the config for sending the email : %s", err.Error())
		return err
	}

	log.Printf("[INFO] Sending email to %s from %s via %s", toEmailAddress, account.Username, emailConfig.Host)
	m := mail.NewMsg()
	err3 := m.From(account.Username) // Username is the email address
	m.FromFormat(account.FormalName, account.Username)
	m.ReplyTo(account.Username)
	m.ReplyToFormat(account.FormalName, account.Username)
	if err3 != nil {
		fmt.Printf("[ERROR] Error in setting the from address : %s", err3.Error())
		return err
	}
	err2 := m.To(toEmailAddress)
	if err2 != nil {
		fmt.Printf("[ERROR] Error in setting the To address : %s", err2.Error())
		return err2
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextPlain, message)
	m.AddAlternativeString(mail.TypeTextHTML, "<p>"+message+"</p>")

	c, err := mail.NewClient(
		emailConfig.Host,
		mail.WithPort(emailConfig.Port),       // Gmail STARTTLS port
		mail.WithSMTPAuth(mail.SMTPAuthLogin), // Gmail supports LOGIN
		mail.WithUsername(account.Username),
		mail.WithPassword(account.Password),   // 16‑char Gmail App Password
		mail.WithTLSPolicy(mail.TLSMandatory), // STARTTLS required
	)
	if err != nil {
		log.Printf("failed to create mail client: %s", err)
		return err
	}

	if err := c.DialAndSend(m); err != nil {
		log.Printf("failed to send mail: %s", err)
		return err
	}

	log.Printf("[INFO] Email successfully send to %s", toEmailAddress)

	return nil
}

func AddEmailConfig(emailConfig model.SMTPConfig) error {
	return dao.AddEmailConfig(emailConfig)
}

func UpdateEmailConfig(emailConfig model.SMTPConfig) error {
	return dao.UpdateEmailConfig(emailConfig)
}

func DeleteEmailConfig(oid int64) error {
	return dao.DeleteEmailConfig(oid)
}

func GetAllEmailConfigs() ([]model.SMTPConfig, error) {
	return dao.GetAllEmailConfigs()
}

func AddEmailAccount(emailAccount model.EmailAccount) error {
	return dao.AddEmailAccount(emailAccount)
}

func UpdateEmailAccount(emailAccount model.EmailAccount) error {
	return dao.UpdateEmailAccount(emailAccount)
}

func DeleteEmailAccount(oid int64) error {
	return dao.DeleteEmailAccount(oid)
}

func GetAllEmailAccounts() ([]model.EmailAccount, error) {
	return dao.GetAllEmailAccounts()
}
