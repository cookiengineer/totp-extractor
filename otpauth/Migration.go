
package otpauth;

import "github.com/dim13/otpauth/migration";



func DecodeMigration (data string) []OTP {

	var result []OTP;

	decoded, err1 := migration.Data(data);

	if err1 == nil {

		payload, err2 := migration.Unmarshal(decoded);

		if err2 == nil {

			if len(payload.OtpParameters) > 0 {

				for o := 0; o < len(payload.OtpParameters); o++ {

					var otp OTP;

					parameters := payload.OtpParameters[o];

					if parameters.Algorithm == migration.Payload_ALGORITHM_SHA1 {
						otp.Algorithm = "SHA1";
					} else if parameters.Algorithm == migration.Payload_ALGORITHM_SHA256 {
						otp.Algorithm = "SHA256";
					} else if parameters.Algorithm == migration.Payload_ALGORITHM_SHA512 {
						otp.Algorithm = "SHA512";
					} else if parameters.Algorithm == migration.Payload_ALGORITHM_MD5 {
						otp.Algorithm = "MD5";
					}

					if len(parameters.Secret) > 0 {
						otp.SetSecret(parameters.Secret);
					}

					if parameters.Type == migration.Payload_OTP_TYPE_HOTP {
						otp.Type = "hotp";
					} else if parameters.Type == migration.Payload_OTP_TYPE_TOTP {
						otp.Type = "totp";
					}

					if parameters.Counter > 0 {
						otp.Period = parameters.Counter;
					}

					if parameters.Digits == migration.Payload_DIGIT_COUNT_SIX {
						otp.Digits = 6;
					} else if parameters.Digits == migration.Payload_DIGIT_COUNT_EIGHT {
						otp.Digits = 8;
					}

					if parameters.Name != "" {
						otp.Name = parameters.Name;
					}

					if parameters.Issuer != "" {
						otp.Issuer = parameters.Issuer;
					}

					result = append(result, otp);

				}

			}

		}

	}

	return result;

}

