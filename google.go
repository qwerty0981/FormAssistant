package form

type googleForm struct{
   url string
   endpoints map[string]string
}

func GoogleForm(initialUrl string) googleForm{
   f := googleForm{url: initialUrl}
   return f
}


