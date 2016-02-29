package main

import (
  "os"
  "github.com/codegangsta/cli"

  "fmt"
  "log"

  "gopkg.in/yaml.v2"

  "io/ioutil"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/ec2"
)

func ListHosts(region string) {
  // Create an EC2 service object in the "us-west-2" region
  // Note that you can also configure your region globally by
  // exporting the AWS_REGION environment variable
  svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

  // Call the DescribeInstances Operation
  resp, err := svc.DescribeInstances(nil)
  if err != nil {
    panic(err)
  }

  // resp has all of the response data, pull out instance IDs:
  fmt.Println("> Number of reservation sets: ", len(resp.Reservations))
  for idx, res := range resp.Reservations {
    fmt.Println("  > Number of instances: ", len(res.Instances))
    for _, inst := range resp.Reservations[idx].Instances {
      fmt.Println("    - Instance ID: ", *inst.InstanceId)
    }
  }
}

// Opens and parses the given YML file:
func ParseDemoFile(path string) (Demo) {
  demo := Demo{}

  data, err := ioutil.ReadFile(path)
  check(err)

  err = yaml.Unmarshal([]byte(data), &demo)
  if err != nil {
    log.Fatalf("error: %v", err)
  }

  fmt.Printf("--- Debug: %v\n", demo.SecurityGroups["docker-ucp-controller"].InboundRules[0].Protocol)

  return demo
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

type Demo struct {
  Provider string `yaml:"provider"`

  SecurityGroups map[string]struct{
    Description string `yaml:"description"`
    InboundRules []struct {
      Port int `yaml:"port"`
      Protocol string `yaml:"protocol"`
      Ip string `yaml:"ip"`
    } `yaml:"inbound"`
  } `yaml:"security-groups"`

  Hosts map[string]struct{
    ImageId string `yaml:"image-id"`
    InstanceType string `yaml:"instance-type"`
    Count int `yaml:"count"`
    SecurityGroups []string `yaml:"security-groups"`
  } `yaml:"hosts"`
}

func ListSecurityGroups(region string) {
  // Create an EC2 service object in the "us-west-2" region
  // Note that you can also configure your region globally by
  // exporting the AWS_REGION environment variable
  svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

  // Call the DescribeInstances Operation
  resp, err := svc.DescribeSecurityGroups(nil)
  if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
  	// Message from an error.
  	fmt.Println(err.Error())
  	return
  }

  // Pretty-print the response data.
  fmt.Println(resp)
  fmt.Println("> Number of security groups: ", len(resp.SecurityGroups))
}

func main() {
  app := cli.NewApp()

  var region string

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "region, r",
      Value:       "us-east-1",
      Usage:       "region where the hosts should be created",
      Destination: &region,
    },
  }

  app.Name = "docker-4-production-demo"
  app.Usage = "fight the loneliness!"
  app.Action = func(c *cli.Context) {
    demoFilePath := "rancher/amazon"

    if c.NArg() > 0 {
      demoFilePath = c.Args()[0]
    }

    // ListHosts(region)
    // ListSecurityGroups(region)
    ParseDemoFile(demoFilePath)
  }

  app.Run(os.Args)
}
