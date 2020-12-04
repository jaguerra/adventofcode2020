package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	policies := getPolicies()
	fmt.Println(getValidPasswordsV1(policies))
	fmt.Println(getValidPasswordsV2(policies))
}

func getValidPasswordsV1(policies []policy) int {
	validPasswords := 0
	for _, policy := range policies {
		if policy.IsValid() == true {
			validPasswords = validPasswords + 1
		}
	}
	return validPasswords
}

func getValidPasswordsV2(policies []policy) int {
	validPasswords := 0
	for _, policy := range policies {
		if policy.IsValid2() == true {
			validPasswords = validPasswords + 1
		}
	}
	return validPasswords
}

type policy struct {
	min, max int
	letter   string
	password string
}

func (p policy) String() string {
	return fmt.Sprintf("min: %d max:%d letter:%s password:%s", p.min, p.max, p.letter, p.password)
}

func (p policy) IsValid() bool {
	count := 0
	for _, char := range p.password {
		if char == rune(p.letter[0]) {
			count = count + 1
		}
	}
	if count >= p.min && count <= p.max {
		return true
	} else {
		return false
	}
}

func (p policy) IsValid2() bool {
	count := 0
	if p.password[p.min-1] == p.letter[0] {
		count = count + 1
	}
	if p.password[p.max-1] == p.letter[0] {
		count = count + 1
	}
	return count == 1
}

func getPolicies() []policy {
	file, _ := os.Open("2.input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	policies := make([]policy, 0)

	for {
		success := scanner.Scan()
		if success == false {
			break
		}
		policy := policy{}
		_, _ = fmt.Sscanf(scanner.Text(), "%d-%d %1s:%s", &policy.min, &policy.max, &policy.letter, &policy.password)
		policies = append(policies, policy)
	}
	return policies
}
