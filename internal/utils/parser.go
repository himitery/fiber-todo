package utils

import (
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func StringToPGText(value string) pgtype.Text {
	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func StringToPGUuid(value string) pgtype.UUID {
	data := []byte(value)

	var bytes [16]byte
	copy(bytes[:], data)
	return pgtype.UUID{
		Bytes: bytes,
		Valid: true,
	}
}

func UuidToPGUuid(value uuid.UUID) pgtype.UUID {
	data, err := value.MarshalBinary()
	if err != nil {
		return pgtype.UUID{
			Bytes: [16]byte{},
			Valid: false,
		}
	}

	var bytes [16]byte
	copy(bytes[:], data)
	return pgtype.UUID{
		Bytes: bytes,
		Valid: true,
	}
}

func PGUuidToUuid(value pgtype.UUID) uuid.UUID {
	res, _ := uuid.Parse(fmt.Sprintf("%x", value.Bytes))

	return res

}

func parseUuid(value string) (res [16]byte, err error) {
	if len(value) != 36 {
		return res, fmt.Errorf("can't parse UUID %v", value)
	}

	buf, err := hex.DecodeString(value)
	if err != nil {
		return res, fmt.Errorf("can't parse UUID %v", value)
	}

	copy(res[:], buf)
	return res, nil
}
