
# QRCode Extractor

## The Problem

One day, your smartphone is going to be remotely wiped all of a sudden. You are losing Apps by
the second, you panic, and you get another Smartphone to take a quick camera photo of the Export
Dialog of the Google Authenticator App to save your 2FA keys and tokens; because without them,
you can access nothing.

Then you realize that the Google Authenticator's Migration format, the `otp-migration://offline...`
URL encoded inside the QR Code is actually incompatible with any other 2FA App.

What do you do? You use this QR Code Extractor!


## Features

- Splits up Camera Photos of Google Authenticator 2FA Export QR Codes
- Generates single QR Codes for (re-)import into another 2FA App
- Decodes the Protobuf encoded parameters into a nice JSON, in case your 2FA App is as shitty as Microsoft Authenticator.

## Libraries

- This tool uses [goxzing](https://github.com/makiuchi-d/gozxing) to read and decode camera photos of QR code images.


## Usage

This tool generates both a QR Code PNG image and a JSON file for reimport to `/tmp`.

You can then use the QR code images containing an `otpauth://<type>/<seed>?parameters...`
URL to scan them in with another 2FA App.

Alternatively you can take a look at the JSON to manually type in the `base32` encoded secret
and to change the parameters like algorithm, type, digits or the current period, correctly.

```bash
go run main.go ./path/to/camera-photo-of-migration-qrcode.jpg;
# Exported files are im /tmp/qrcode-*.{png,json}
```

## License

GPL3 or WTFPL, I don't care.

