# Sample usage of AWS SSM Parameter Store and Secret Manager


More details can be found [here](https://j-kapuscik2.medium.com/how-to-store-secretes-on-aws-3e2f5881cd5b)

Create Simple System Manager Parameter Store secrets:
```
aws ssm put-parameter --cli-input-json \
   '{
        "Name": "/SampleApp/uat/db/postgres/password", 
        "Value": "secret-password", 
        "Type": "SecureString",
        "Tags": [
            {
            "Key": "env",
            "Value": "UAT"
            },
            {
            "Key": "app",
            "Value": "SampleApp"
            }
         ]
    }'
    
aws ssm put-parameter --cli-input-json \
   '{
        "Name": "/SampleApp/uat/db/postgres/endpoint", 
        "Value": "secret-enpoint.net", 
        "Type": "SecureString",
        "Tags": [
            {
            "Key": "env",
            "Value": "UAT"
            },
            {
            "Key": "app",
            "Value": "SampleApp"
            }
         ]
    }'
    
aws ssm get-parameters-by-path --path "/SampleApp/uat/db/postgres" --with-decryption
```

Create Secret Manager Secrets

```
aws secretsmanager create-secret  \
    --name /SampleApp/uat/db/postgres/password \
    --secret-string "secret-password" \
    --tags Key=env,Value=UAT Key=app,Value=SampleApp

aws secretsmanager create-secret  \
    --name /SampleApp/uat/db/postgres/endpoint \
    --secret-string "secret-enpoint.net" \
    --tags Key=env,Value=UAT Key=app,Value=SampleApp
    
aws secretsmanager list-secrets \
    --filters Key=name,Values=/SampleApp/uat/db/postgres
```
