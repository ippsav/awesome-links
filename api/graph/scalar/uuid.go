package scalar

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUIDScalar(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(id.String())))
	})
}

func UnmarshalUUIDScalar(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		id, err := uuid.Parse(v)
		return id, err
	default:
		return uuid.UUID{}, fmt.Errorf("%v is not a string", v)
	}
}
