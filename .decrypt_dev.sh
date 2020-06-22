
# Decrypt .ENVIRONMENT.env file

gpg --quiet --batch --yes --decrypt --passphrase="$DEV_ENV_SECRET_PASSPHRASE"  --output ./.dev.env ./.dev.env.gpg

