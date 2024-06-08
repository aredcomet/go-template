package database

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var ErrorMap = map[string]string{
	"unq-cmp_member-company-company_id-user_id": "User already a member of this company",
}

func IsPostgresError(err error) string {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		if msg, ok := ErrorMap[pgError.ConstraintName]; ok {
			return msg
		}
		return ""
	}
	return ""
}
