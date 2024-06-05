# Oktetodo

A todo list application showing how you can use the Okteto Dagger module to create Preview Environments for your repos.

## Repo Secrets

1. `$TOKEN` - GitHub token to give permissions to comment on the PR
1. `$OKTETO_TOKEN` - Needed for Dagger to talk to your Okteto instance

## Steps to Run The Application With Okteto

1. Create the following [Okteto secrets](https://www.okteto.com/docs/cloud/secrets/) with your AWS account values.
```
AWS_ACCESS_KEY_ID
AWS_REGION
AWS_SECRET_ACCESS_KEY
```
Note: Make sure this user has the permissions to create S3 buckets.


1. Get an [access token from Pulumi](https://www.pulumi.com/docs/pulumi-cloud/access-management/access-tokens/#creating-personal-access-tokens) and add it as an Okteto secret.
```
PULUMI_ACCESS_TOKEN
```


1. Deploy the application from the UI using the Git URL (https://github.com/okteto/todolist-pulumi-s3) option or clone the repo and run `okteto deploy` or `okteto up`.
