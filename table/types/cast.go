package types

import (
	"errors"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/value"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
)

var errNilValue = errors.New("nil value")

// CastTo try cast value to destination type value
func CastTo(v Value, dst interface{}) error {
	if v == nil {
		return xerrors.WithStackTrace(errNilValue)
	}
	return value.CastTo(v, dst)
}

func ToDecimal(v Value) (*Decimal, error) {
	if valuer, isDecimalValuer := v.(value.DecimalValuer); isDecimalValuer {
		return &Decimal{
			Bytes:     valuer.Value(),
			Precision: valuer.Precision(),
			Scale:     valuer.Scale(),
		}, nil
	}
	return nil, xerrors.WithStackTrace(fmt.Errorf("value type '%s' is not decimal type", v.Type().Yql()))
}

func ListItems(v Value) ([]Value, error) {
	if vv, has := v.(interface {
		ListItems() []Value
	}); has {
		return vv.ListItems(), nil
	}
	return nil, xerrors.WithStackTrace(fmt.Errorf("cannot get list items from '%s'", v.Type().Yql()))
}

func TupleItems(v Value) ([]Value, error) {
	if vv, has := v.(interface {
		TupleItems() []Value
	}); has {
		return vv.TupleItems(), nil
	}
	return nil, xerrors.WithStackTrace(fmt.Errorf("cannot get tuple items from '%s'", v.Type().Yql()))
}

func StructFields(v Value) (map[string]Value, error) {
	if vv, has := v.(interface {
		StructFields() map[string]Value
	}); has {
		return vv.StructFields(), nil
	}
	return nil, xerrors.WithStackTrace(fmt.Errorf("cannot get struct fields from '%s'", v.Type().Yql()))
}

func DictValues(v Value) (map[Value]Value, error) {
	if vv, has := v.(interface {
		DictValues() map[Value]Value
	}); has {
		return vv.DictValues(), nil
	}
	return nil, xerrors.WithStackTrace(fmt.Errorf("cannot get dict values from '%s'", v.Type().Yql()))
}
