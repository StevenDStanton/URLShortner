package main

import (
	"log"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53targets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type AwsCdkStackProps struct {
	awscdk.StackProps
}

func NewAwsCdkStack(scope constructs.Construct, id string, props *AwsCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	domainName := os.Getenv("DOMAIN_NAME")
	subdomainName := os.Getenv("SUBDOMAIN_NAME")

	// Create the Lambda function
	lambdaFunction := awslambda.NewFunction(stack, jsii.String("urlshortnergo"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("../build/function.zip"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("main"),
		FunctionName: jsii.String("urlshortnergo"),
		Environment:  &map[string]*string{
			// Add necessary environment variables if needed
		},
	})

	// Lookup the existing Route 53 hosted zone
	hostedZone := awsroute53.HostedZone_FromLookup(stack, jsii.String("HostedZone"), &awsroute53.HostedZoneProviderProps{
		DomainName: jsii.String(domainName),
	})

	// Request an SSL certificate
	certificate := awscertificatemanager.NewCertificate(stack, jsii.String("Certificate"), &awscertificatemanager.CertificateProps{
		DomainName: jsii.String(subdomainName),
		Validation: awscertificatemanager.CertificateValidation_FromDns(hostedZone),
	})

	// Create the API Gateway with the custom domain name
	api := awsapigateway.NewRestApi(stack, jsii.String("UrlShortenerApi"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			AllowMethods: jsii.Strings("POST", "GET", "PUT", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
		DeployOptions: &awsapigateway.StageOptions{
			LoggingLevel: awsapigateway.MethodLoggingLevel_INFO,
		},
		CloudWatchRole: jsii.Bool(true),
		DomainName: &awsapigateway.DomainNameOptions{
			Certificate: certificate,
			DomainName:  jsii.String(subdomainName),
		},
	})

	integration := awsapigateway.NewLambdaIntegration(lambdaFunction, nil)

	// Define Routes
	api.Root().AddResource(jsii.String("healthCheck"), nil).AddMethod(jsii.String("GET"), integration, nil)
	api.Root().AddResource(jsii.String("shorten"), nil).AddMethod(jsii.String("PUT"), integration, nil)
	api.Root().AddMethod(jsii.String("GET"), integration, nil)

	// Configure Route 53 A record
	awsroute53.NewARecord(stack, jsii.String("ApiAliasRecord"), &awsroute53.ARecordProps{
		Zone:       hostedZone,
		RecordName: jsii.String("s"), // Subdomain
		Target:     awsroute53.RecordTarget_FromAlias(awsroute53targets.NewApiGateway(api)),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewAwsCdkStack(app, "urlshortnerstack", &AwsCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("AWS_ACCOUNT_ID")),
		Region:  jsii.String(os.Getenv("AWS_REGION")),
	}
}
