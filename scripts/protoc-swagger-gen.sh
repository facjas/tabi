#!/usr/bin/env bash
SWAGGER_DIR=./swagger-proto

set -eo pipefail

# prepare swagger generation
mkdir -p "$SWAGGER_DIR/proto"
printf "version: v1\ndirectories:\n  - proto\n  - third_party" > "$SWAGGER_DIR/buf.work.yaml"
printf "version: v1\nname: buf.build/tabi/tabi\n" > "$SWAGGER_DIR/proto/buf.yaml"
cp ./proto/buf.gen.swagger.yaml "$SWAGGER_DIR/proto/buf.gen.swagger.yaml"

# copy existing proto files
cp -r ./proto/tabi "$SWAGGER_DIR/proto"
cp -r ./proto/ethermint "$SWAGGER_DIR/proto"

# create temporary folder to store intermediate results from `buf generate`
mkdir -p ./tmp-swagger-gen

# step into swagger folder
cd "$SWAGGER_DIR"

# create swagger files on an individual basis  w/ `buf build` and `buf generate` (needed for `swagger-combine`)
proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ -n "$query_file" ]]; then
    echo $query_file
    buf generate --template proto/buf.gen.swagger.yaml "$query_file"
  fi

  tx_file=$(find "${dir}" -maxdepth 1 \( -name 'tx.proto' -o -name 'service.proto' \))
  if [[ -n "$tx_file" ]]; then
    echo $tx_file
    buf generate --template proto/buf.gen.swagger.yaml "$tx_file"
  fi
done

proto_dirs=$(find ./third_party -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ -n "$query_file" ]]; then
    buf generate --template proto/buf.gen.swagger.yaml "$query_file"
  fi
done

cd ..

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
