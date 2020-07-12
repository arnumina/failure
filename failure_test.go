/*
#######
##         ____     _ __
##        / __/__ _(_) /_ _________
##       / __/ _ `/ / / // / __/ -_)
##      /_/  \_,_/_/_/\_,_/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package failure_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/arnumina/failure"
)

func TestNew(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{nil, ""},
		{fmt.Errorf(""), ""},
		{fmt.Errorf("foo"), "foo"},
		{failure.New(io.EOF), "EOF"},
	}

	for i, tt := range tests {
		got := failure.New(tt.err)
		if got.Error() != tt.want {
			t.Errorf("[%02d] => got: %s, want %s", i, got, tt.want)
		}
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		err   error
		key   string
		value interface{}
		want  string
	}{
		{nil, "foo", "bar", `foo="bar"`},
		{fmt.Errorf(""), "foo", 123, "foo=123"},
		{fmt.Errorf("bar"), "foo", true, "foo=true >>> bar"},
		{failure.New(io.EOF), "foo", nil, "foo=<nil> >>> EOF"},
		{failure.New(io.EOF).Set("k", "v"), "foo", nil, `foo=<nil> >>> k="v" >>> EOF`},
		{nil, "foo", failure.New(io.EOF).Error(), `foo="EOF"`},
		{nil, "foo", failure.New(io.EOF).Set("k", "v").Error(), `foo="k=\"v\" >>> EOF"`},
	}

	for i, tt := range tests {
		got := failure.New(tt.err).Set(tt.key, tt.value)
		if got.Error() != tt.want {
			t.Errorf("[%02d] => got: %s, want %s", i, got, tt.want)
		}
	}
}

func TestSetf(t *testing.T) {
	tests := []struct {
		err    error
		key    string
		format string
		args   []interface{}
		want   string
	}{
		{nil, "foo", "%s", []interface{}{"bar"}, `foo="bar"`},
		{fmt.Errorf(""), "foo", "%d", []interface{}{123}, `foo="123"`},
		{fmt.Errorf("bar"), "foo", "%v", []interface{}{true}, `foo="true" >>> bar`},
		{failure.New(io.EOF), "foo", "%v", []interface{}{nil}, `foo="<nil>" >>> EOF`},
		{failure.New(io.EOF).Set("k", "v"), "foo", "%v", []interface{}{nil}, `foo="<nil>" >>> k="v" >>> EOF`},
		{nil, "foo", "%s", []interface{}{failure.New(io.EOF)}, `foo="EOF"`},
		{nil, "foo", "%s", []interface{}{failure.New(io.EOF).Set("k", "v")}, `foo="k=\"v\" >>> EOF"`},
	}

	for i, tt := range tests {
		got := failure.New(tt.err).Setf(tt.key, tt.format, tt.args...)
		if got.Error() != tt.want {
			t.Errorf("[%02d] => got: %s, want %s", i, got, tt.want)
		}
	}
}

func TestMsgf(t *testing.T) {
	tests := []struct {
		err    error
		format string
		args   []interface{}
		want   string
	}{
		{nil, "%s", []interface{}{"bar"}, "bar"},
		{nil, "%d", []interface{}{123}, "123"},
		{nil, "%v", []interface{}{false}, "false"},
		{failure.New(io.EOF), "%s", []interface{}{"bar"}, "bar >>> EOF"},
		{fmt.Errorf(""), "%d", []interface{}{123}, "123"},
		{failure.New(io.EOF).Set("k", "v"), "%v", []interface{}{false}, `false >>> k="v" >>> EOF`},
	}

	for i, tt := range tests {
		got := failure.New(tt.err).Msgf(tt.format, tt.args...)
		if got.Error() != tt.want {
			t.Errorf("[%02d] => got: %s, want %s", i, got, tt.want)
		}
	}
}

func TestSetMsg(t *testing.T) {
	tests := []struct {
		err   error
		key   string
		value interface{}
		msg   string
		want  string
	}{
		{nil, "foo", "bar", "read error", `read error: foo="bar"`},
		{fmt.Errorf(""), "foo", 123, "read error", "read error: foo=123"},
		{failure.New(io.EOF), "foo", true, "read error", "read error: foo=true >>> EOF"},
	}

	for i, tt := range tests {
		got := failure.New(tt.err).Set(tt.key, tt.value).Msg(tt.msg)
		if got.Error() != tt.want {
			t.Errorf("[%02d] => got: %s, want %s", i, got, tt.want)
		}
	}
}

func TestUnwrap(t *testing.T) {
	f := failure.New(failure.New(io.EOF))
	if !errors.Is(f, io.EOF) {
		t.Error("Unwrap error")
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
