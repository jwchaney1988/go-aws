package main

// Import the required packages
import (
	"github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"

    "fmt"
    "os"
)

func main() {
    // Load credentials session from shared config, set the default Region and tell it which profile to use
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String("us-east-1"),
        Credentials: credentials.NewSharedCredentials("", "a-pal-preprod"),
    })

    // Create new EC2 client
    ec2Svc := ec2.New(sess)

    // Call to get detailed information on each instance
    result, err := ec2Svc.DescribeInstances(nil)
    // Check if an error was returned
    if err != nil {
        // Print the error and exit the app
        fmt.Println("Error", err)
        os.Exit(1)
    } else {
        // Print the EC2 response data
        fmt.Println("Success", result)
    }
}