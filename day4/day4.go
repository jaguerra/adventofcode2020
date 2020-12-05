package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := Passports()
	fmt.Println(CountValid(passports))
	fmt.Println(CountValid2(passports))
}

func CountValid(passports []passport) int {
	validCount := 0
	for _, passport := range passports {
		if passport.Valid() {
			validCount++
		}
	}
	return validCount
}

func CountValid2(passports []passport) int {
	validCount := 0
	for _, passport := range passports {
		if passport.Valid2() {
			validCount++
		}
	}
	return validCount
}

type passport struct {
	fields map[string]string
}

func (p passport) Valid() bool {
	mandatoryFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, mandatoryField := range mandatoryFields {
		_, exists := p.fields[mandatoryField]
		if exists == false {
			return false
		}
	}
	return true
}

func (p passport) Valid2() bool {
	if p.Valid() == false {
		return false
	}

	// Validate byr
	byr, err := strconv.Atoi(p.fields["byr"])
	if err != nil {
		return false
	}
	if byr < 1920 || byr > 2002 {
		return false
	}

	// Validate iyr
	iyr, err := strconv.Atoi(p.fields["iyr"])
	if err != nil {
		return false
	}
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// Validate eyr
	eyr, err := strconv.Atoi(p.fields["eyr"])
	if err != nil {
		return false
	}
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// Validate hgt
	var heightMetric string
	var heightValue int
	fmt.Sscanf(p.fields["hgt"], "%d%s", &heightValue, &heightMetric)
	switch heightMetric {
	case "cm":
		if heightValue < 150 || heightValue > 193 {
			return false
		}
	case "in":
		if heightValue < 59 || heightValue > 76 {
			return false
		}
	default:
		return false
	}

	// Validate hcl
	matches, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, p.fields["hcl"])
	if matches == false {
		return false
	}

	// Validate ecl
	matches, _ = regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, p.fields["ecl"])
	if matches == false {
		return false
	}

	// Validate pid
	matches, _ = regexp.MatchString(`^[0-9]{9}$`, p.fields["pid"])
	if matches == false {
		return false
	}
	return true
}

func Passports() []passport {
	file, _ := os.Open("4.input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	passports := make([]passport, 0)
	pass := &passport{make(map[string]string)}

	for {
		success := scanner.Scan()
		if success == false {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			passports = append(passports, *pass)
			pass = &passport{make(map[string]string)}
		}

		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			fieldParts := strings.Split(token, ":")
			if len(fieldParts) == 2 {
				pass.fields[fieldParts[0]] = fieldParts[1]
			}
		}
	}

	if len(pass.fields) > 0 {
		passports = append(passports, *pass)
	}
	return passports
}
