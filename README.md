# sec2env

`sec2env` is a very simple program that retrieves a secret from AWS
secrets manager, parses the secret value as a json and prints all
contained key, value pairs to standard output.

The purpose is to `eval` the resulting string in a shell and export
all key, value pairs to the process environment.

Is this safe/secure? Depends on how much one can trust one's own
secrets :)

## Build

Running `go build -o sec2env` will output the `sec2env` executable.

## Example

Let's assume that we have stored a secret in aws region `eu-central-1`
with the name `foo` and we have stored the following key, value pairs
in the secret: key `alpha` has value `1` and key `beta` has value 2.

Running `sec2env -name foo -region eu-central-1` will output the
following text:

```
export alpha=1
export beta=2
```

If we wish to export these variables into the current environment, we
need to do `eval $(sec2env -name foo -region eu-central-1)` in the
shell.

In case of an error, the program will return an exit value of 1 and
fail silently so as not to break downstream output processing).

## Access to AWS Secrets Manager

The program uses the [standard AWS
guidelines](https://aws.amazon.com/blogs/security/a-new-and-standardized-way-to-manage-credentials-in-the-aws-sdks/)
in order to authenticate with AWS secrets manager.

For production, the best practice is to define an IAM role for the
underlying resource (e.g. ec2 instance).

Given an ec2 instance and a secret with name `foo` in region
`eu-central-1` and account `111111111111`, the following example
defines an IAM role with the secret access policy in yaml form. The
role then needs to be attached to the ec2 instance.

```yaml
AWSTemplateFormatVersion: '2010-09-09'

Description: >-
  An example of defining an IAM rolw with a secret access policy

Resources:
  SecretsAccessPolicy:
    Type: AWS::IAM::Policy
    Properties:
      PolicyName: "secret-access"
      PolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Action:
              - secretsmanager:GetSecretValue
            Resource:
              - "arn:aws:secretsmanager:eu-central-1:111111111111:secret:foo"
      Roles:
        - !Ref ServerRole

  ServerRole:
    Type: AWS::IAM::Role
    Properties:
      RoleName: "server-role"
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Principal:
              Service:
                - "ec2.amazonaws.com"
            Action:
              - "sts:AssumeRole"
```
