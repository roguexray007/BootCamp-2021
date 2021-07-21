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
	numberOfHours int
	numberOfDays int
}
func (freelancer FreelancerEmployee) findSalary() int{
	return freelancer.basicSalary * freelancer.numberOfHours * freelancer.numberOfDays
}
//function to calculate total Salary at The End of Month for all types of Employees
func totalSalary(salaries ...int) int{
	var totalSalaryAtMonthEnd int = 0
	for _,totalSalaryOfEveryCategory := range salaries{
		totalSalaryAtMonthEnd += totalSalaryOfEveryCategory
	}
	return totalSalaryAtMonthEnd
}
func main(){
	var typeOfEmployee Employee
	fulltime := FullTimeEmployee{500,28}
	typeOfEmployee = fulltime //typeOfEmployee fulltime implements Employee
	FullTimeEmployeeSalary := typeOfEmployee.findSalary()
	contractor := OnContractEmployee{100,28}
	typeOfEmployee = contractor //typeOfEmployee contractor implements Employee
	OnContractEmployeeSalary := typeOfEmployee.findSalary()
	freelancer := FreelancerEmployee{10,10,28}
	typeOfEmployee = freelancer //typeOfEmployee freelancer implements Employee
	FreelancerEmployeeSalary := typeOfEmployee.findSalary()
	fmt.Println(totalSalary(FullTimeEmployeeSalary,OnContractEmployeeSalary,FreelancerEmployeeSalary))
}