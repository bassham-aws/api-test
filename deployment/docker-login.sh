unset  AWS_SESSION_TOKEN

temp_role=$(aws sts assume-role \
                    --role-arn "arn:aws:iam::${AWS_ACCOUNT}:role/ecr_admin" \
                    --role-session-name "circleci")

export AWS_ACCESS_KEY_ID=$(echo $temp_role | jq .Credentials.AccessKeyId | xargs)
export AWS_SECRET_ACCESS_KEY=$(echo $temp_role | jq .Credentials.SecretAccessKey | xargs)
export AWS_SESSION_TOKEN=$(echo $temp_role | jq .Credentials.SessionToken | xargs)

login="$(aws ecr get-login --no-include-email)"
${login}
