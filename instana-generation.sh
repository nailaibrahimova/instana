#!/bin/bash

### This script assumes you have the `java` and `wget` commands on the path

wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.0.0-beta/openapi-generator-cli-5.0.0-beta.jar \
  -O openapi-generator-cli.jar

# take peek at the different generators it offers us
GO_POST_PROCESS_FILE="gofmt -s -w" java -jar openapi-generator-cli.jar generate -i openapi.yaml -g go \
    -o pkg/instana \
    --additional-properties packageName=instana \
    --additional-properties isGoSubmodule=true \
    --additional-properties packageVersion=1.185.824 \
    --type-mappings=object=interface{} \
    --enable-post-process-file \
    --skip-validate-spec
