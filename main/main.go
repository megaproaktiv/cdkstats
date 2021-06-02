package main

import (
	"github.com/megaproaktiv/cdkstat"
	"fmt"
)

func main() {
	LOCALONLY := "LOCAL_ONLY"

	remoteStacks := cdkstat.GetStatus()
	localCDKStackNames := cdkstat.ReadStacks()
	remoteStackNames := make([]string, 5)

	fmt.Printf("%-32s %-32s %-32s \n", "Name", "Status", "Description")
	fmt.Printf("%-32s %-32s %-32s \n", "----", "------", "-----------")
	// Remote State
	for i := range remoteStacks.Stacks {
		stack := remoteStacks.Stacks[i]
		remoteStackNames = append(remoteStackNames, *stack.StackName)
		name := FixedLengthString(*stack.StackName)
		status := FixedLengthString(string(stack.StackStatus))
		description := "-"
		if stack.Description != nil{
			description = FixedLengthString(string(*stack.Description))
		}
		if contains(localCDKStackNames, *stack.StackName) {
			fmt.Printf("%s %s %s\n", name, status, description)
		}
	}
	// Local only
	status := FixedLengthString(LOCALONLY)
	for _, nameLocal := range *localCDKStackNames {
		name := FixedLengthString(*&nameLocal)
		if !contains(&remoteStackNames, nameLocal) {
			fmt.Printf("%s %s\n", name, status)
		}
	}

}

// FixedLengthString some formatting
func FixedLengthString(str string) string {
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
