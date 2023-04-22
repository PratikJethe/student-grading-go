package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	studenList := []student{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if lineNumber > 1 {

			// fmt.Println(scanner.Text())

			seperatedString := strings.Split(scanner.Text(), ",")
			ts1, _ := strconv.Atoi(seperatedString[3])
			ts2, _ := strconv.Atoi(seperatedString[4])
			ts3, _ := strconv.Atoi(seperatedString[5])
			ts4, _ := strconv.Atoi(seperatedString[6])
			student := student{firstName: seperatedString[0], lastName: seperatedString[1], university: seperatedString[2], test1Score: ts1, test2Score: ts2, test3Score: ts3, test4Score: ts4}
			studenList = append(studenList, student)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return studenList
}

func calculateGrade(students []student) []studentStat {
	studentStatsList := []studentStat{}
	for _, student := range students {
		avg := float32(student.test1Score+student.test2Score+student.test3Score+student.test4Score) / 4

		var grade Grade
		switch {
		case avg >= 70:
			grade = A
		case avg >= 50 && avg <= 70:
			grade = B
		case avg >= 35 && avg <= 50:
			grade = C
		default:
			grade = F

		}
		studentStatsList = append(studentStatsList, studentStat{
			student:    student,
			finalScore: avg,
			grade:      grade,
		})
	}

	return studentStatsList
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var topper studentStat
	for _, gradedStudent := range gradedStudents {

		if topper.finalScore < gradedStudent.finalScore {
			topper = gradedStudent
		}
	}
	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {

	studentsPerUnversityMap := map[string][]studentStat{}
	topperPerUnversityMap := map[string]studentStat{}

	for _, studentStat := range gs {
		studentsPerUnversityMap[studentStat.university] = append(studentsPerUnversityMap[studentStat.university], studentStat)
	}

	for university, studentStat := range studentsPerUnversityMap {
		topperPerUnversityMap[university] = findOverallTopper(studentStat)
	}

	return topperPerUnversityMap
}
