package main

import "fmt"

type Employee interface {
	findSalary() int
}

type FullTimeEmployee struct{
	basicSalary int
	numberOfDays int
}
func (fullTime FullTimeEmployee) findSalary() int{
	return fullTime.basicSalary * fullTime.numberOfDays
}
type OnContractEmployee struct {
	basicSalary int
	numberOfDays int
}
func (contractor OnContractEmployee) findSalary() int{
	return contractor.basicSalary * contractor.numberOfDays
}

type FreelancerEmployee struct {
	basicSalary int
	numberOfHoursDaily int//number of hours should be entered on a daily basis
	numberOfDays int
}
func (freelancer FreelancerEmployee) findSalary() int{
	return freelancer.basicSalary * freelancer.numberOfHoursDaily * freelancer.numberOfDays
}
//function to calculate total Salary at The End of Month for all types of Employees
func totalSalary(salaryMapOfEmployees ...Employee) int{
	var totalSalaryAtMonthEnd int = 0
	for _,totalSalaryOfEveryCategory := range salaryMapOfEmployees{
		totalSalaryAtMonthEnd += totalSalaryOfEveryCategory.findSalary()
	}
	return totalSalaryAtMonthEnd
}
func main(){
	fullTime := FullTimeEmployee{500,28}
	contractor := OnContractEmployee{100,28}
	freelancer := FreelancerEmployee{10,10,28}
	salaryMapOfEmployees := []Employee{fullTime,contractor,freelancer}
	fmt.Println(totalSalary(salaryMapOfEmployees...))
}
