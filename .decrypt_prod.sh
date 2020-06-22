
# Decrypt .ENVIRONMENT.env file

gpg --quiet --batch --yes --decrypt --passphrase="$PROD_ENV_SECRET_PASSPHRASE"  --output ./.dev.env ./.dev.env.gpg

