package services

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHashFormat        = errors.New("invalid hash format")
	ErrIncompatibleArgonVersion = errors.New("incompatible argon version")
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type PasswordService struct {
	params params
}

func NewPasswordService() *PasswordService {
	return &PasswordService{
		params: params{
			memory:      64 * 1024,
			iterations:  3,
			parallelism: 2,
			saltLength:  16,
			keyLength:   32,
		},
	}
}

func (s *PasswordService) Hash(password string) (string, *entities.Exception) {
	salt, err := s.generateRandomBytes(s.params.saltLength)

	if err != nil {
		return "", entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	hash := argon2.IDKey([]byte(password), salt, s.params.iterations, s.params.memory, s.params.parallelism, s.params.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, s.params.memory, s.params.iterations, s.params.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (s *PasswordService) Compare(plainPassword, hash string) (bool, *entities.Exception) {
	p, salt, decodedHash, err := s.decodeHash(hash)
	if err != nil {
		return false, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	passwordHash := argon2.IDKey([]byte(plainPassword), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(decodedHash, passwordHash) == 1 {
		return true, nil
	}

	return false, nil
}

func (*PasswordService) decodeHash(encodedHash string) (*params, []byte, []byte, error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHashFormat
	}

	var version int
	_, err := fmt.Scanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}

	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleArgonVersion
	}

	p := &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func (*PasswordService) generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
