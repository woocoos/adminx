package validates

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func IsURLEncoded() func(s string) error {
	return func(s string) error {
		if !uRLEncodedRegex.MatchString(s) {
			return errors.New("value is invalid url encoded")
		}
		return nil
	}
}

func IsHTMLEncoded() func(s string) error {
	return func(s string) error {
		if !hTMLEncodedRegex.MatchString(s) {
			return errors.New("value is invalid html encoded")
		}
		return nil
	}
}

func IsHTML() func(s string) error {
	return func(s string) error {
		if !hTMLRegex.MatchString(s) {
			return errors.New("value is invalid html format")
		}
		return nil
	}
}

// IsMac is the validation function for validating if the field's value is a valid MAC address.
func IsMac() func(s string) error {
	return func(s string) error {
		_, err := net.ParseMAC(s)
		if err != nil {
			return errors.New("value is invalid mac address")
		}
		return nil
	}
}

// IsCIDRv4 is the validation function for validating if the field's value is a valid v4 CIDR address.
func IsCIDRv4() func(s string) error {
	return func(s string) error {
		ip, _, err := net.ParseCIDR(s)
		if err == nil && ip.To4() != nil {
			return nil
		}
		return errors.New("value is invalid v4 CIDR address")
	}
}

// IsCIDRv6 is the validation function for validating if the field's value is a valid v6 CIDR address.
func IsCIDRv6() func(s string) error {
	return func(s string) error {
		ip, _, err := net.ParseCIDR(s)
		if err == nil && ip.To4() == nil {
			return nil
		}
		return errors.New("value is invalid v6 CIDR address")
	}
}

// IsCIDR is the validation function for validating if the field's value is a valid v4 or v6 CIDR address.
func IsCIDR() func(s string) error {
	return func(s string) error {
		_, _, err := net.ParseCIDR(s)
		if err == nil {
			return nil
		}
		return errors.New("value is invalid v4 or v6 CIDR address")
	}
}

// IsIPv4 is the validation function for validating if a value is a valid v4 IP address.
func IsIPv4() func(s string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip != nil && ip.To4() != nil {
			return nil
		}
		return errors.New("value is invalid v4 IP address")
	}
}

// IsIPv6 is the validation function for validating if the field's value is a valid v6 IP address.
func IsIPv6() func(s string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip != nil && ip.To4() == nil {
			return nil
		}
		return errors.New("value is invalid v6 IP address")
	}
}

// IsIP is the validation function for validating if the field's value is a valid v4 or v6 IP address.
func IsIP() func(s string) error {
	return func(s string) error {
		ip := net.ParseIP(s)
		if ip != nil {
			return nil
		}
		return errors.New("value is invalid v4 or v6 IP address")
	}
}

// IsSSN is the validation function for validating if the field's value is a valid SSN.
func IsSSN() func(s string) error {
	return func(s string) error {
		if len(s) != 11 {
			return errors.New("value is invalid SSN")
		}
		if sSNRegex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid SSN")
	}
}

// IsLongitude is the validation function for validating if the field's value is a valid longitude coordinate.
func IsLongitude() func(s interface{}) error {
	return func(s interface{}) error {
		var v string
		switch reflect.TypeOf(s).Kind() {
		case reflect.String:
			v = s.(string)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v = strconv.FormatInt(s.(int64), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v = strconv.FormatUint(s.(uint64), 10)
		case reflect.Float32:
			v = strconv.FormatFloat(s.(float64), 'f', -1, 32)
		case reflect.Float64:
			v = strconv.FormatFloat(s.(float64), 'f', -1, 64)
		default:
			return errors.New(fmt.Sprintf("Bad field type %T", s))
		}
		if longitudeRegex.MatchString(v) {
			return nil
		}
		return errors.New("value is invalid longitude coordinate")
	}
}

// IsLatitude is the validation function for validating if the field's value is a valid latitude coordinate.
func IsLatitude() func(s interface{}) error {
	return func(s interface{}) error {
		var v string
		switch reflect.TypeOf(s).Kind() {
		case reflect.String:
			v = s.(string)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v = strconv.FormatInt(s.(int64), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			v = strconv.FormatUint(s.(uint64), 10)
		case reflect.Float32:
			v = strconv.FormatFloat(s.(float64), 'f', -1, 32)
		case reflect.Float64:
			v = strconv.FormatFloat(s.(float64), 'f', -1, 64)
		default:
			return errors.New(fmt.Sprintf("Bad field type %T", s))
		}
		if latitudeRegex.MatchString(v) {
			return nil
		}
		return errors.New("value is invalid latitude coordinate")
	}
}

// IsDataURI is the validation function for validating if the field's value is a valid data URI.
func IsDataURI() func(s string) error {
	return func(s string) error {
		uri := strings.SplitN(s, ",", 2)
		if len(uri) != 2 {
			return errors.New("value is invalid data URI")
		}
		if !dataURIRegex.MatchString(uri[0]) {
			return errors.New("value is invalid data URI")
		}
		if base64Regex.MatchString(uri[1]) {
			return nil
		}
		return errors.New("value is invalid data URI")
	}
}

// IsPrintableASCII is the validation function for validating if the field's value is a valid printable ASCII character.
func IsPrintableASCII() func(s string) error {
	return func(s string) error {
		if printableASCIIRegex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid printable ASCII character")
	}
}

// IsASCII is the validation function for validating if the field's value is a valid ASCII character.
func IsASCII() func(s string) error {
	return func(s string) error {
		if aSCIIRegex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid ASCII character")
	}
}

// IsUUID5 is the validation function for validating if the field's value is a valid v5 UUID.
func IsUUID5() func(s string) error {
	return func(s string) error {
		if uUID5Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid v5 UUID")
	}
}

// IsUUID4 is the validation function for validating if the field's value is a valid v4 UUID.
func IsUUID4() func(s string) error {
	return func(s string) error {
		if uUID4Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid v4 UUID")
	}
}

// IsUUID3 is the validation function for validating if the field's value is a valid v3 UUID.
func IsUUID3() func(s string) error {
	return func(s string) error {
		if uUID3Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid v3 UUID")
	}
}

// IsUUID is the validation function for validating if the field's value is a valid UUID of any version.
func IsUUID() func(s string) error {
	return func(s string) error {
		if uUIDRegex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid UUID of any version")
	}
}

// IsUUID5RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v5 UUID.
func IsUUID5RFC4122() func(s string) error {
	return func(s string) error {
		if uUID5RFC4122Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid RFC4122 v5 UUID")
	}
}

// IsUUID4RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v4 UUID.
func IsUUID4RFC4122() func(s string) error {
	return func(s string) error {
		if uUID4RFC4122Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid RFC4122 v4 UUID")
	}
}

// IsUUID3RFC4122 is the validation function for validating if the field's value is a valid RFC4122 v3 UUID.
func IsUUID3RFC4122() func(s string) error {
	return func(s string) error {
		if uUID3RFC4122Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid RFC4122 v3 UUID")
	}
}

// IsUUIDRFC4122 is the validation function for validating if the field's value is a valid RFC4122 UUID of any version.
func IsUUIDRFC4122() func(s string) error {
	return func(s string) error {
		if uUIDRFC4122Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid RFC4122 UUID of any version")
	}
}

// IsMD4 is the validation function for validating if the field's value is a valid MD4.
func IsMD4() func(s string) error {
	return func(s string) error {
		if md4Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid MD4")
	}
}

// IsMD5 is the validation function for validating if the field's value is a valid MD5.
func IsMD5() func(s string) error {
	return func(s string) error {
		if md5Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid MD5")
	}
}

// IsSHA256 is the validation function for validating if the field's value is a valid SHA256.
func IsSHA256() func(s string) error {
	return func(s string) error {
		if sha256Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid SHA256")
	}
}

// IsSHA384 is the validation function for validating if the field's value is a valid SHA384.
func IsSHA384() func(s string) error {
	return func(s string) error {
		if sha384Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid SHA384")
	}
}

// IsSHA512 is the validation function for validating if the field's value is a valid SHA512.
func IsSHA512() func(s string) error {
	return func(s string) error {
		if sha512Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid SHA512")
	}
}

// IsBase64 is the validation function for validating if the current field's value is a valid base 64.
func IsBase64() func(s string) error {
	return func(s string) error {
		if base64Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid base 64")
	}
}

// IsBase64URL is the validation function for validating if the current field's value is a valid base64 URL safe string.
func IsBase64URL() func(s string) error {
	return func(s string) error {
		if base64URLRegex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid base64 URL safe string")
	}
}

// IsURI is the validation function for validating if the current field's value is a valid URI.
func IsURI() func(s string) error {
	return func(s string) error {
		// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
		// emulate browser and strip the '#' suffix prior to validation. see issue-#237
		if i := strings.Index(s, "#"); i > -1 {
			s = s[:i]
		}
		if len(s) == 0 {
			return errors.New("value is invalid URI")
		}
		_, err := url.ParseRequestURI(s)
		if err != nil {
			return errors.New("value is invalid URI")
		}
		return nil
	}
}

// IsURL is the validation function for validating if the current field's value is a valid URL.
func IsURL() func(s string) error {
	return func(s string) error {
		var i int
		// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
		// emulate browser and strip the '#' suffix prior to validation. see issue-#237
		if i = strings.Index(s, "#"); i > -1 {
			s = s[:i]
		}
		if len(s) == 0 {
			return errors.New("value is invalid URL")
		}
		_url, err := url.ParseRequestURI(s)
		if err != nil || _url.Scheme == "" {
			return errors.New("value is invalid URL")
		}
		return nil
	}
}

// IsE164 is the validation function for validating if the current field's value is a valid e.164 formatted phone number.
func IsE164() func(s string) error {
	return func(s string) error {
		if e164Regex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid e.164 formatted phone number")
	}
}

// IsEmail is the validation function for validating if the current field's value is a valid email address.
func IsEmail() func(s string) error {
	return func(s string) error {
		if emailRegex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid email address")
	}
}

// IsNumber is the validation function for validating if the current field's value is a valid number.
func IsNumber() func(s interface{}) error {
	return func(s interface{}) error {
		switch reflect.TypeOf(s).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			return nil
		case reflect.String:
			if numberRegex.MatchString(s.(string)) {
				return nil
			}
			return errors.New("value is invalid number")
		default:
			return errors.New("value is invalid number")
		}
	}
}

// IsNumeric is the validation function for validating if the current field's value is a valid numeric value.
func IsNumeric() func(s interface{}) error {
	return func(s interface{}) error {
		switch reflect.TypeOf(s).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			return nil
		case reflect.String:
			if numericRegex.MatchString(s.(string)) {
				return nil
			}
			return errors.New("value is invalid numeric")
		default:
			return errors.New("value is invalid numeric")
		}
	}
}

// IsBoolean is the validation function for validating if the current field's value is a valid boolean value or can be safely converted to a boolean value.
func IsBoolean() func(s interface{}) error {
	return func(s interface{}) error {
		switch reflect.TypeOf(s).Kind() {
		case reflect.Bool:
			return nil
		case reflect.String:
			_, err := strconv.ParseBool(s.(string))
			if err == nil {
				return nil
			}
			return errors.New("value is invalid boolean")
		default:
			return errors.New("value is invalid boolean")
		}
	}
}

// excludesRune is the validation function for validating that the field's value does not contain the rune specified within the param.
//func excludesRune(es string) func(s string) error {
//	return func(s string) error {
//		if containsRune(es)(s) == nil {
//			return nil
//		}
//		return errors.New("value does not contain the rune specified within the param")
//	}
//}

//// excludesAll is the validation function for validating that the field's value does not contain any of the characters specified within the param.
//func excludesAll(es string) func(s string) error {
//	return !containsAny(s, es)
//}
//
//// excludes is the validation function for validating that the field's value does not contain the text specified within the param.
//func excludes(es string) func(s string) error {
//	return !contains(s, es)
//}
//
//// containsRune is the validation function for validating that the field's value contains the rune specified within the param.
//func containsRune(es string) func(s string) error {
//	return func(s string) error {
//		r, _ := utf8.DecodeRuneInString(es)
//		if strings.ContainsRune(s, r) {
//			return nil
//		}
//		return errors.New("value contains the rune specified within the param")
//	}
//}
//
//// containsAny is the validation function for validating that the field's value contains any of the characters specified within the param.
//func containsAny(fl FieldLevel) bool {
//	return strings.ContainsAny(fl.Field().String(), fl.Param())
//}
//
//// contains is the validation function for validating that the field's value contains the text specified within the param.
//func contains(fl FieldLevel) bool {
//	return strings.Contains(fl.Field().String(), fl.Param())
//}
//
//// startsWith is the validation function for validating that the field's value starts with the text specified within the param.
//func startsWith(fl FieldLevel) bool {
//	return strings.HasPrefix(fl.Field().String(), fl.Param())
//}
//
//// endsWith is the validation function for validating that the field's value ends with the text specified within the param.
//func endsWith(fl FieldLevel) bool {
//	return strings.HasSuffix(fl.Field().String(), fl.Param())
//}
//
//// startsNotWith is the validation function for validating that the field's value does not start with the text specified within the param.
//func startsNotWith(fl FieldLevel) bool {
//	return !startsWith(fl)
//}
//
//// endsNotWith is the validation function for validating that the field's value does not end with the text specified within the param.
//func endsNotWith(fl FieldLevel) bool {
//	return !endsWith(fl)
//}

// IsJSON is the validation function for validating if the current field's value is a valid json string.
func IsJSON() func(s interface{}) error {
	return func(s interface{}) error {
		if reflect.TypeOf(s).Kind() == reflect.String {
			val := s.(string)
			if json.Valid([]byte(val)) {
				return nil
			}
		}
		return errors.New("value is invalid json string")
	}
}

// IsJWT is the validation function for validating if the current field's value is a valid JWT string.
func IsJWT() func(s string) error {
	return func(s string) error {
		if jWTRegex.MatchString(s) {
			return nil
		}
		return errors.New("value is invalid JWT string")
	}
}

// IsHostnamePort validates a <dns>:<port> combination for fields typically used for socket address.
func IsHostnamePort() func(s string) error {
	return func(s string) error {
		val := s
		host, port, err := net.SplitHostPort(val)
		if err != nil {
			return errors.New("value is invalid <dns>:<port> address")
		}
		// Port must be a iny <= 65535.
		if portNum, err := strconv.ParseInt(port, 10, 32); err != nil || portNum > 65535 || portNum < 1 {
			return errors.New("value is invalid <dns>:<port> address")
		}
		// If host is specified, it should match a DNS name
		if host != "" && hostnameRegexRFC1123.MatchString(host) {
			return nil
		}
		return errors.New("value is invalid <dns>:<port> address")
	}
}

// IsLowercase is the validation function for validating if the current field's value is a lowercase string.
func IsLowercase() func(s interface{}) error {
	return func(s interface{}) error {
		if reflect.TypeOf(s).Kind() == reflect.String {
			val := s.(string)
			if val != "" && val == strings.ToLower(val) {
				return nil
			}
		}
		return errors.New("value is invalid lowercase string")
	}
}

// IsUppercase is the validation function for validating if the current field's value is an uppercase string.
func IsUppercase() func(s interface{}) error {
	return func(s interface{}) error {
		if reflect.TypeOf(s).Kind() == reflect.String {
			val := s.(string)
			if val != "" && val == strings.ToUpper(val) {
				return nil
			}
		}
		return errors.New("value is invalid uppercase string")
	}
}

// IsTimeZone is the validation function for validating if the current field's value is a valid time zone string.
func IsTimeZone() func(s interface{}) error {
	return func(s interface{}) error {
		if reflect.TypeOf(s).Kind() == reflect.String {
			val := s.(string)
			// empty value is converted to UTC by time.LoadLocation but disallow it as it is not a valid time zone name
			if val == "" {
				return errors.New("value is invalid time zone string")
			}
			// Local value is converted to the current system time zone by time.LoadLocation but disallow it as it is not a valid time zone name
			if strings.ToLower(val) == "local" {
				return errors.New("value is invalid time zone string")
			}
			_, err := time.LoadLocation(val)
			if err == nil {
				return nil
			}
		}
		return errors.New("value is invalid time zone string")
	}
}
