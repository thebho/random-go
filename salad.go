package main

import "fmt"

var saladRequirements = []string{"lettuce", "kale", "field greens", "spinach"}
var saladExclusions = []string{"bread", "crust", "pita", "wrap", "tortilla"}
var ingredientMin = 3

var ingredients = []string{"tomato", "mayo"}

func testRequirements(ingredients []string, requirements []string, exclusions []string) (valid bool) {
  for _,v := range ingredients {
    for _,vv := range requirements {
      if v == vv {
        valid = true
      }
    }

    for _,vx := range exclusions {
      if v == vx {
        return false
      }
    }
  }
  return
}

func isMeal(ingredients []string, servingDevice string) bool {
  salad := false
  if servingDevice != "bowl" && servingDevice != "plate" {
    return false
  }


  salad = testRequirements(ingredients, saladRequirements, saladExclusions)

  if !salad {
    return false
  }

  if len(ingredients) >= ingredientMin {
    salad = true
  } else {
    return false
  }

  return salad
}

func main() {

  servingDevice := "bowl"
  print("A meal with the ingredients:")

  for i,v := range ingredients {
    last := false
    if i == len(ingredients) - 1 {
      last = true
      print(" and")
    }
    print(" ")
    print(v)
    if !last && len(ingredients) > 2 {
      print(",")
    }
  }
  print(", served")
  if servingDevice == "bowl" {
    print(" in a bowl, is")
    } else {
      print(fmt.Sprintf(" on a %s, is", servingDevice))
    }
    if !isMeal(ingredients, servingDevice) {
      print(" not")
    }
    print(" a salad.\n")
  }
