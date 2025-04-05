#!/bin/sh
# Utility to extract Linux Kernel signature keys

# Set key owner email
KEY_EMAIL="torvalds@kernel.org"

# Temporary file for exported PGP key
PGP_KEY_FILE="torvalds.asc"
PEM_KEY_FILE="torvalds.pem"

echo "[*] Exporting PGP key for $KEY_EMAIL..."
gpg --export --armor "$KEY_EMAIL" > "$PGP_KEY_FILE"

if [ ! -s "$PGP_KEY_FILE" ]; then
    echo "[!] Failed to export PGP key."
    exit 1
fi

echo "[*] Extracting key ID..."
KEY_ID=$(gpg --list-packets "$PGP_KEY_FILE" | awk '/keyid:/ {print $2; exit}')

if [ -z "$KEY_ID" ]; then
    echo "[!] Failed to extract key ID."
    exit 1
fi

echo "[*] Converting PGP key to PEM format..."
gpg --export "$KEY_ID" | openssl rsa -pubin -inform DER -outform PEM -out "$PEM_KEY_FILE"

if [ ! -s "$PEM_KEY_FILE" ]; then
    echo "[!] Failed to convert to PEM format."
    exit 1
fi

echo "[+] Successfully extracted and converted key to $PEM_KEY_FILE"
