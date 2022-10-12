package lasagna

func PreparationTime(slice []string, avgPrepTime int) int{
    var prepTime int
    if avgPrepTime == 0{
    	prepTime = 2
    } else {
    	prepTime = avgPrepTime
    }
    layers := len(slice)
    totalTime := layers * prepTime
    return totalTime
}

func Quantities(layers []string)(noodles int, sauce float64){
    var numbaONoodles int
    var numbaOSauce int

	for _, v := range layers {
        if v == "sauce"{
            numbaOSauce++
        } else if v == "noodles"{
        	numbaONoodles++
        }
    }

    noodles = numbaONoodles * 50
    sauce = float64(numbaOSauce) * 0.2
    
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendlist []string, mylist []string) []string{
    friendLength := len(friendlist) - 1
    myLength := len(mylist) - 1

    mylist[myLength] = friendlist[friendLength]
    
    return mylist
}

func ScaleRecipe(slice []float64, portions int)(scaledPortions []float64){
    
    for _, item := range slice{
        scaledPortions = append(scaledPortions, item * (float64(portions)/float64(2)))
    }
    return
}
