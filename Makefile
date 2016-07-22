default:
	@go get github.com/SMAWK/smawk-bot

makecert:
	@openssl req -newkey rsa:2048 -sha256 -nodes -keyout smawk_key.pem -x509 -days 365 -out smawk_cert.pem -subj "/C=US/ST=South Carolina/L=Lexington/O=My Simple Things/CN=mysimplethings.xyz"
