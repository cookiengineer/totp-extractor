
# QRCode Extractor

This is a DIY tool to deal with camera photos of exported accounts from Google Authenticator.

Google Authenticator uses a proprietary protobuf based export that is encoded inside the QRCode
image, and there was no good tool or library that could read the actual data in the format that
you can reuse it to import it into another 2FA authenticator app.

The general issue is that you need to take a photo with a camera, and the mobile apps usually
don't allow to export the QRCode as a real image file, which means you are left with a crappy
QRCode that no normal tool (outside the xzing world) can read.

## Usage

This tool extracts a collection of `otp-migration://offline` encoded QRCodes into separate
QRCode images and JSON files. You can either use the QRCode with a tool like KeePassXC, or
you can inspect the JSON file, so that you can type in manually the base32 encoded `secret`
and 2FA details such as the algorithm, type, digits, current period etc).

```bash
go run main.go ./path/to/camera-photo-of-migration-qrcode.jpg;
# Exported files are im /tmp/qrcode-*.{png,json}
```

## License

GPL3 or WTFPL, I don't care.

