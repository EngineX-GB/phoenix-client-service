package dao

import (
	"log"
	"phoenix-client-service/datasource"
	"phoenix-client-service/model"
	"phoenix-client-service/util"
)

func AddEmailAccount(emailAccount model.EmailAccount) error {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer db.Close()
	defer tx.Rollback()
	if err != nil {
		log.Printf("An error has occured when starting the transaction : %s", err.Error())
		return err
	}
	_, err = tx.Exec(util.AddEmailAccount(emailAccount))

	if err != nil {
		log.Printf("An error has occurred when executing the transaction : %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		print("An error has occurred when trying to save an the email account: " + err.Error())
		return err
	}

	return nil
}

func UpdateEmailAccount(emailAccount model.EmailAccount) error {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer db.Close()
	defer tx.Rollback()
	if err != nil {
		log.Printf("An error has occured when starting the transaction : %s", err.Error())
		return err
	}
	_, err = tx.Exec(util.UpdateEmailAccount(emailAccount))

	if err != nil {
		log.Printf("An error has occurred when executing the transaction : %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		print("An error has occurred when trying to save an the email account: " + err.Error())
		return err
	}

	return nil
}

func GetEmailAccountById(oid int64) (model.EmailAccount, error) {
	var emailAccount model.EmailAccount
	db := datasource.Connect()
	defer db.Close()
	row := db.QueryRow(util.GetEmailAccountById(oid))
	err := row.Scan(&emailAccount.Oid, &emailAccount.Username, &emailAccount.Password, &emailAccount.FormalName, &emailAccount.SmtpConfigId)
	if err != nil {
		return emailAccount, err
	}
	return emailAccount, nil
}

func GetEmailConfigById(oid int64) (model.SMTPConfig, error) {
	var emailConfig model.SMTPConfig
	db := datasource.Connect()
	defer db.Close()
	row := db.QueryRow(util.GetEmailConfigById(oid))
	err := row.Scan(&emailConfig.Oid, &emailConfig.Name, &emailConfig.Host, &emailConfig.Port)
	if err != nil {
		return emailConfig, err
	}
	return emailConfig, nil
}

/**
 *
 * Get all email accounts
 *
 */

func GetAllEmailAccounts() ([]model.EmailAccount, error) {
	var listOfEmailAccounts = make([]model.EmailAccount, 0)
	db := datasource.Connect()
	defer db.Close()
	rows, err := db.Query(util.GetAllEmailAccounts())

	if err != nil {
		log.Printf("[ERROR] An error has occured when trying to fetch the email accounts : %s", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var emailAccount model.EmailAccount
		if err := rows.Scan(&emailAccount.Oid, &emailAccount.Username, &emailAccount.Password, &emailAccount.FormalName,
			&emailAccount.SmtpConfigId); err != nil {
			return listOfEmailAccounts, err
		}
		listOfEmailAccounts = append(listOfEmailAccounts, emailAccount)
	}
	if err = rows.Err(); err != nil {
		return listOfEmailAccounts, err
	}
	return listOfEmailAccounts, nil
}

/**
 *
 * Get all email configurations
 *
 */

func GetAllEmailConfigs() ([]model.SMTPConfig, error) {
	var listOfEmailConfigs = make([]model.SMTPConfig, 0)
	db := datasource.Connect()
	defer db.Close()
	rows, err := db.Query(util.GetAllEmailConfigs())

	if err != nil {
		log.Printf("[ERROR] An error has occured when trying to fetch the email accounts : %s", err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var emailConfig model.SMTPConfig
		if err := rows.Scan(&emailConfig.Oid, &emailConfig.Name, &emailConfig.Host, &emailConfig.Port); err != nil {
			return listOfEmailConfigs, err
		}
		listOfEmailConfigs = append(listOfEmailConfigs, emailConfig)
	}
	if err = rows.Err(); err != nil {
		return listOfEmailConfigs, err
	}
	return listOfEmailConfigs, nil
}

func AddEmailConfig(emailConfig model.SMTPConfig) error {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer db.Close()
	defer tx.Rollback()
	if err != nil {
		log.Printf("An error has occured when starting the transaction : %s", err.Error())
		return err
	}
	_, err = tx.Exec(util.AddEmailConfig(emailConfig))

	if err != nil {
		log.Printf("An error has occurred when executing the transaction : %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		print("An error has occurred when trying to save an the email config: " + err.Error())
		return err
	}

	return nil
}

func UpdateEmailConfig(emailConfig model.SMTPConfig) error {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer db.Close()
	defer tx.Rollback()
	if err != nil {
		log.Printf("An error has occured when starting the transaction : %s", err.Error())
		return err
	}
	_, err = tx.Exec(util.UpdateEmailConfig(emailConfig))

	if err != nil {
		log.Printf("An error has occurred when executing the transaction : %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		print("An error has occurred when trying to save an the email config: " + err.Error())
		return err
	}

	return nil
}

func DeleteEmailConfig(oid int64) error {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer db.Close()
	defer tx.Rollback()
	if err != nil {
		log.Printf("An error has occured when starting the transaction : %s", err.Error())
		return err
	}
	_, err = tx.Exec(util.DeleteEmailConfig(oid))

	if err != nil {
		log.Printf("An error has occurred when executing the transaction : %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		print("An error has occurred when trying to delete email config: " + err.Error())
		return err
	}

	return nil
}

func DeleteEmailAccount(oid int64) error {
	db := datasource.Connect()
	tx, err := db.Begin()
	defer db.Close()
	defer tx.Rollback()
	if err != nil {
		log.Printf("An error has occured when starting the transaction : %s", err.Error())
		return err
	}
	_, err = tx.Exec(util.DeleteEmailAccount(oid))

	if err != nil {
		log.Printf("An error has occurred when executing the transaction : %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		print("An error has occurred when trying to delete email account: " + err.Error())
		return err
	}

	return nil
}
