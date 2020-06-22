
# Decrypt .ENVIRONMENT.env file

gpg --quiet --batch --yes --decrypt --passphrase="$PROD_ENV_SECRET_PASSPHRASE" \ --output ./.prod.env ./.prod.env.gpg

