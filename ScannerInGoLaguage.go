package main
 
import (
	"fmt"
	"math"
	"sort"
)
 
func main() {
 
	var noOfTimesRain, noOfWetCloths, noOfTimesWeCanCollectCloths int
	var maximumRainInterval int = -1
	var ans int = 0
	timeTakenPerCloths := 0
	var rainSeconds = make([]int, 1000)
 
	fmt.Scanf("%d%d%d", &noOfTimesRain, &noOfWetCloths, &noOfTimesWeCanCollectCloths)
	// fmt.Printf("%v , %v , %v", noOfTimesRain, noOfWetCloths, noOfTimesWeCanCollectCloths)
 
	for i := 0; i < noOfTimesRain; i++ {
		fmt.Scanf("%d", &rainSeconds[i])
	}
 
	sort.Ints(rainSeconds[:noOfTimesRain])
	for i := 0; i < noOfTimesRain-1; i++ {
		maximumRainInterval = int(math.Max(float64(maximumRainInterval), float64(rainSeconds[i+1]-rainSeconds[i])))
	}
 
	for i := 0; i < noOfWetCloths; i++ {
		fmt.Scanf("%d", &timeTakenPerCloths)
		if timeTakenPerCloths <= maximumRainInterval {
			ans++
		}
	}
 
	fmt.Printf("%v", ans)
}
