Project to explore GitHub Actions, AWS CDK, and AWS Lambda.

Base app is written in Go and simply outputs an IP address of a function visitor.

Local Lambda development with Amazon Sam
- [install](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)
- run `sam local start-api` (on mac os `DOCKER_HOST=unix://$HOME/.docker/run/docker.sock sam local start-api`)
- visit http://127.0.0.1:3000/hello

Creates Infrastructure with Amazon CDK
- npm install -g aws-sdk
- run `make build-and-zip` to create a binary. It's required to create a Lambda function
- run `cd cdk`
- run `cdk bootstrap`
- run `AWS_ACCOUNT_ID=1111 AWS_REGION=us-east1 cdk deploy`. It will eventually output values required to run a pipeline
- run `cdk destroy` to delete created resources

CI/CD with GitHub Actions
- pipeline config is located in .github/workflows
- pipeline requires the following secrets to work: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_LAMBDA_ARN
- secrets are added on github.com/user/repository/settings/secrets/actions under Repository secrets