package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type RatingModule struct{
	numberOfStudent int
	lowestRating int
	highestRating int
	mutex sync.Mutex
}
func (rating_module *RatingModule) takeRating(student *sync.WaitGroup,totalRating *uint64){
	defer student.Done()
	rating_module.mutex.Lock()
	rand.Seed(time.Now().UnixNano())
	currentStudentRating := uint64(rating_module.lowestRating) +
				uint64(rand.Intn(rating_module.highestRating - rating_module.lowestRating + 1))
	fmt.Println(currentStudentRating)
	*totalRating += currentStudentRating
	rating_module.mutex.Unlock()

}
func (rating_module *RatingModule) calculateAverageRating(totalRating uint64) float64{
	return float64(totalRating/uint64(rating_module.numberOfStudent))
}
func main(){
	var student sync.WaitGroup
	rating_module := RatingModule{}
	fmt.Println("Enter number of students, lowest integer rating and highest integer rating :")
	fmt.Scanf("%d %d %d",&rating_module.numberOfStudent,&rating_module.lowestRating,&rating_module.highestRating)
	var totalRating uint64//atomic counter for accumulating Rating by each student
	for currentStudent := 1; currentStudent <= rating_module.numberOfStudent; currentStudent++{
		student.Add(1)
		go rating_module.takeRating(&student,&totalRating)
	}
	student.Wait()
	fmt.Println("Total Rating after ending process: ",totalRating)
	rating_module.calculateAverageRating(totalRating)
	fmt.Println("Total Average Rating after ending process: ",rating_module.calculateAverageRating(totalRating))

}
