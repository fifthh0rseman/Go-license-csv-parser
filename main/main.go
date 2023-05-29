package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	sourceFilename := "licenses.csv"
	source, err := os.Open(sourceFilename)
	if err != nil {
		fmt.Printf("Error while opening %s: %v", sourceFilename, err)
	}
	data, err := csv.NewReader(source).ReadAll()
	if err != nil {
		fmt.Printf("Error while reading csv from %s: %v", sourceFilename, err)
	}
	dest, err := os.Create("result.txt")

	//pwd, err := os.Getwd()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//writeString(dest, fmt.Sprintf("Project: %s\n", pwd))

	licenseAndProjects := make(map[string][]string)
	for _, line := range data {
		licenseAndProjects[line[2]] = append(licenseAndProjects[line[2]], line[0])
	}
	writeString(dest, fmt.Sprintf("Licenses in project: \n%s\n\n", stringArrayToString(Keys(licenseAndProjects))))
	for k, v := range licenseAndProjects {
		tmp := fmt.Sprintf("License: %s, projects: %s\n\n", k, stringArrayToString(v))
		writeString(dest, tmp)
	}

	writeString(dest, "Source\tVendor-path\tLicense\n\n")
	for _, line := range data {
		writeString(dest, fmt.Sprintf("%s\t%s\t%s\n", line[0], line[1], line[2]))
	}
	closeFile(source)
	closeFile(dest)
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		return
	}
}

func writeString(file *os.File, toWrite string) {
	_, err := file.WriteString(toWrite)
	if err != nil {
		return
	}
}

func stringArrayToString(array []string) string {
	res := ""
	for i, str := range array {
		res += str
		if i != len(array) {
			res += "\n"
		}
	}
	return res
}

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}
