package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

type CdkGoStackProps struct {
	awscdk.StackProps
}

func NewCdkGoStack(scope constructs.Construct, id string, props *CdkGoStackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	// Create IAM User for CI/CD with access to deploy Lambda functions
	ciCdUser := awsiam.NewUser(stack, jsii.String("GoLambdaUser"), &awsiam.UserProps{
		UserName: jsii.String("go-lambda-user"),
	})

	// Create access keys for the CI/CD user
	accessKey := awsiam.NewAccessKey(stack, jsii.String("CICDAccessKey"), &awsiam.AccessKeyProps{
		User: ciCdUser,
	})

	// Grant CI/CD user permissions to manage Lambda functions
	ciCdUser.AddManagedPolicy(awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AWSLambda_FullAccess")))

	// Print the access key and secret key for the user
	awscdk.NewCfnOutput(stack, jsii.String("AccessKeyId"), &awscdk.CfnOutputProps{
		Value: accessKey.AccessKeyId(),
	})

	awscdk.NewCfnOutput(stack, jsii.String("SecretAccessKey"), &awscdk.CfnOutputProps{
		Value: accessKey.SecretAccessKey().UnsafeUnwrap(),
	})

	// Create a Lambda function with Amazon Linux 2 runtime
	lambdaFunction := awslambda.NewFunction(stack, jsii.String("MyLambdaFunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("MyAmazonLinux2Function"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Code:         awslambda.Code_FromAsset(jsii.String("./../build/bootstrap.zip"), nil),
		Handler:      jsii.String("handler"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(30)),
	})

	// Print the Lambda function ARN
	awscdk.NewCfnOutput(stack, jsii.String("LambdaFunctionARN"), &awscdk.CfnOutputProps{
		Value: lambdaFunction.FunctionArn(),
	})

	// Create an API Gateway for the Lambda function
	api := awsapigateway.NewLambdaRestApi(stack, jsii.String("MyAPIGateway"), &awsapigateway.LambdaRestApiProps{
		Handler:     lambdaFunction,
		RestApiName: jsii.String("MyAPIGateway"),
		DeployOptions: &awsapigateway.StageOptions{
			StageName: jsii.String("prod"),
		},
	})

	// Print the API Gateway URL
	awscdk.NewCfnOutput(stack, jsii.String("APIGatewayURL"), &awscdk.CfnOutputProps{
		Value: api.Url(),
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkGoStack(app, "CdkGoStack", &CdkGoStackProps{
		StackProps: awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	account := os.Getenv("AWS_ACCOUNT_ID")
	region := os.Getenv("AWS_REGION")

	if account == "" || region == "" {
		panic("AWS_ACCOUNT_ID and AWS_REGION must be set in the environment variables")
	}

	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
