package main

import "fmt"

func isSalad(ingredients []string, servingDevice string) bool {
  salad := false
  if servingDevice != "bowl" && servingDevice != "plate" {
    return false
  }

  for _,v := range ingredients {
    if v == "lettuce" || v == "kale" || v == "field greens" || v == "spinach"  {
      salad = true
    }
    // Bread implies sandwich (ie BLT, Cheeseburger)
    if v == "bread" || v == "crust" || v == "pita" || v == "wrap" || v == "tortilla" {
      return false
    }
  }

  if !salad {
    return false
  }

  if len(ingredients) >= 3 {
    salad = true
  } else {
    return false
  }

  return salad
}

func main() {
  ingredients := []string{"tomato", "mayo", "spinach"}
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
    if !isSalad(ingredients, servingDevice) {
      print(" not")
    }
    print(" a salad.\n")
  }
