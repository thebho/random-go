package main

func isSalad(ingredients []string, housing string) bool {
  salad := false
  if housing != "bowl" && housing != "plate" {
    return false
  }

  for _,v := range ingredients {
    if v == "lettuce" || v == "kale" || v == "field greens" {
      salad = true
    }
    // Bread implies sandwich (ie BLT, Cheeseburger)
    if v == "bread" {
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
  println("Salad:", isSalad([]string{"lettuce", "tomato", "brea"}, "plate"))
}
