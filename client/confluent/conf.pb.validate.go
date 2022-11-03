// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: client/confluent/conf.proto

package confluent

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on SASL with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *SASL) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Mechanisms

	// no validation rules for User

	// no validation rules for Password

	return nil
}

// SASLValidationError is the validation error returned by SASL.Validate if the
// designated constraints aren't met.
type SASLValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SASLValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SASLValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SASLValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SASLValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SASLValidationError) ErrorName() string { return "SASLValidationError" }

// Error satisfies the builtin error interface
func (e SASLValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSASL.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SASLValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SASLValidationError{}

// Validate checks the field values on SSL with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *SSL) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CaLocation

	// no validation rules for CaPem

	return nil
}

// SSLValidationError is the validation error returned by SSL.Validate if the
// designated constraints aren't met.
type SSLValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SSLValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SSLValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SSLValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SSLValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SSLValidationError) ErrorName() string { return "SSLValidationError" }

// Error satisfies the builtin error interface
func (e SSLValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSSL.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SSLValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SSLValidationError{}

// Validate checks the field values on ConsumerConf with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConsumerConf) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BootstrapServers

	// no validation rules for SecurityProtocol

	if v, ok := interface{}(m.GetSasl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConsumerConfValidationError{
				field:  "Sasl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSsl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConsumerConfValidationError{
				field:  "Ssl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Group

	return nil
}

// ConsumerConfValidationError is the validation error returned by
// ConsumerConf.Validate if the designated constraints aren't met.
type ConsumerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsumerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsumerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsumerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsumerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsumerConfValidationError) ErrorName() string { return "ConsumerConfValidationError" }

// Error satisfies the builtin error interface
func (e ConsumerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsumerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsumerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsumerConfValidationError{}

// Validate checks the field values on ProducerConf with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProducerConf) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BootstrapServers

	// no validation rules for SecurityProtocol

	if v, ok := interface{}(m.GetSasl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProducerConfValidationError{
				field:  "Sasl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetSsl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProducerConfValidationError{
				field:  "Ssl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RequestRequiredAcks

	// no validation rules for Acks

	// no validation rules for Partitioner

	// no validation rules for RequestTimeoutMs

	// no validation rules for Group

	// no validation rules for Topic

	return nil
}

// ProducerConfValidationError is the validation error returned by
// ProducerConf.Validate if the designated constraints aren't met.
type ProducerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProducerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProducerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProducerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProducerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProducerConfValidationError) ErrorName() string { return "ProducerConfValidationError" }

// Error satisfies the builtin error interface
func (e ProducerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProducerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProducerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProducerConfValidationError{}
