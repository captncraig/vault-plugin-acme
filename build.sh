export EXE=plugins/acme.exe
export VAULT_ADDR=http://127.0.0.1:8200
go build -o $EXE
SHASUM=$(go run plugins/sum.go $EXE)
echo $SHASUM
vault write sys/plugins/catalog/acme-plugin \
    sha_256="$SHASUM" \
    command="acme.exe"

vault write sys/plugins/reload/backend plugin=acme-plugin