#!/usr/bin/env bash
#
# This is a helper script that uses the AWS CLI configuration to construct an
# AWS ProviderConfig.

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

cd "$(dirname "$0")"

set -e -o pipefail

AWS_PROFILE=${AWS_PROFILE:=default} # retrieve profile's credentials, save it under 'default' profile, and base64 encode it
BASE64ENCODED_AWS_ACCOUNT_CREDS=$(echo "[default]\naws_access_key_id = $(aws configure get aws_access_key_id --profile $AWS_PROFILE)\naws_secret_access_key = $(aws configure get aws_secret_access_key --profile $AWS_PROFILE)" | base64  | tr -d "\n")

cat providerconfig.yaml | sed "s/<REPLACEME>/${BASE64ENCODED_AWS_ACCOUNT_CREDS}/g" | kubectl apply -f -
