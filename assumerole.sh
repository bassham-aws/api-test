unset  AWS_SESSION_TOKEN

echo $AWS_ACCESS_KEY_ID
echo $AWS_SECRET_ACCESS_KEY

temp_role=$(aws sts assume-role \
                    --role-arn "arn:aws:iam::183375172283:role/ecr_admin" \
                    --role-session-name "circleci")

export AWS_ACCESS_KEY_ID=$(echo $temp_role | jq .Credentials.AccessKeyId | xargs)
export AWS_SECRET_ACCESS_KEY=$(echo $temp_role | jq .Credentials.SecretAccessKey | xargs)
export AWS_SESSION_TOKEN=$(echo $temp_role | jq .Credentials.SessionToken | xargs)

echo $AWS_ACCESS_KEY_ID
echo $AWS_SECRET_ACCESS_KEY
echo $AWS_SESSION_TOKEN

login = "$(aws ecr get-login --no-include-email)"
echo ${login}
${login}
docker push ${AWS_ACCOUNT}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com/api-test:${CIRCLE_TAG}
