package godog

const humanAgeRatio int = 7

// Years functions allows you to receive a human number of ages and will return the equivalent age for a dog.
//
// Note: Don't input a negative number or you'll have a problem!
func Years(humanYears int){
  return humanYears * humanAgeRatio
}
