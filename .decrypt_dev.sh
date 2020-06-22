#!bin/sh

# Decrypt .ENVIRONMENT.env file

gpg --quiet --batch --yes --decrypt --passphrase="$DEV_ENV_SECRET_PASSPHRASE" \ --output .prod.env .prod.env.gpg

