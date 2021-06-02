package cdkstat

import (
	"bufio"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"log"
	"os"
)

const stackNamesFile = "stacks.csv"

// GetStatus get States of all Cfn Stacks
func GetStatus() *(cloudformation.DescribeStacksOutput) {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	client := cloudformation.NewFromConfig(cfg)
	input := &cloudformation.DescribeStacksInput{}

	resp, _ := client.DescribeStacks(context.TODO(), input)
	return resp
}

// Read saves Stack Names from file
func ReadStacks() *[]string {
	stackNames := make([]string, 10)
	file, err := os.Open(stackNamesFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stackNames = append(stackNames, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return &stackNames
}
