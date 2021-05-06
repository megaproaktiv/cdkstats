package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

const stackNamesFile = "stacks.csv"

func main() {
	LOCALONLY := "LOCAL_ONLY"

	remoteStacks := GetStatus()
	localCDKStackNames := ReadStacks()
	remoteStackNames := make([]string, 5)

	fmt.Printf("%-32s %-32s\n","Name", "Status")
	fmt.Printf("%-32s %-32s\n","----", "------")
	// Remote State
	for i := range  remoteStacks.Stacks{
		stack :=  remoteStacks.Stacks[i]
		remoteStackNames = append(remoteStackNames, *stack.StackName)
		name := FixedLengthString(*stack.StackName)
		status := FixedLengthString(string(stack.StackStatus))
		if contains( localCDKStackNames, *stack.StackName) {
			fmt.Printf("%s %s\n",name, status)
		}
	}	
	// Local only
	status := FixedLengthString(LOCALONLY)
	for _, nameLocal := range *localCDKStackNames {
		name := FixedLengthString(*&nameLocal)
		if !contains( &remoteStackNames, nameLocal) {
			fmt.Printf("%s %s\n",name, status)
		}
	}	

}


// GetStatus get States of all Cfn Stacks
func GetStatus() *(cloudformation.DescribeStacksOutput){

	cfg, err := config.LoadDefaultConfig(context.TODO())

    if err != nil {
        panic("unable to load SDK config, " + err.Error())
	}

	client := cloudformation.NewFromConfig(cfg);
	input := &cloudformation.DescribeStacksInput{}

	resp, _ := client.DescribeStacks(context.TODO(), input)
	return resp
}

// Read saves Stack Names from file
func ReadStacks() *[]string{
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


// FixedLengthString some formatting
func FixedLengthString( str string) string {
	return fmt.Sprintf("%-32s", str)
}

// does slice contain key
func contains(stacks *[]string, stack string) bool {
    for _, cdkStack := range *stacks {
		theSame := (cdkStack == stack)
		if theSame {
			return true
		}
	}
    return false
}