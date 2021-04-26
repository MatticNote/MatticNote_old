package internal

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/MatticNote/MatticNote/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	PasswordHashCost = 12
	KeyPairLength    = 2048
)

var (
	ErrAccountAlreadyExists = errors.New("account already exists")
)

func CreateUser(email, username, password string, skipEmailVerify bool) error {
	var count int
	err := db.DB.QueryRow(
		context.Background(),
		"SELECT count(*) FROM \"user\" WHERE (username ILIKE $1 AND host IS NULL) OR email ILIKE $2\n",
		username,
		email,
	).Scan(&count)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if count > 0 {
		return ErrAccountAlreadyExists
	}

	tx, err := db.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	rsaKeyRaw, err := rsa.GenerateKey(rand.Reader, KeyPairLength)
	if err != nil {
		return err
	}
	rsaPrivateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsaKeyRaw),
	})
	rsaPublicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(rsaKeyRaw.Public().(*rsa.PublicKey)),
	})

	newKeyPairUuid := uuid.Must(uuid.NewRandom())
	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO signature_key(uuid, public_key, private_key) VALUES ($1, $2, $3)",
		newKeyPairUuid.String(),
		string(rsaPublicKey),
		string(rsaPrivateKey),
	)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), PasswordHashCost)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		"INSERT INTO \"user\"(username, email, password, signature_key_uuid, is_mail_verified) VALUES ($1, $2, $3, $4, $5)",
		username,
		email,
		string(hashedPassword),
		newKeyPairUuid.String(),
		skipEmailVerify,
	)
	if err != nil {
		return nil
	}

	if !skipEmailVerify {
		// TODO: email send system
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return nil
	}
	return nil
}
