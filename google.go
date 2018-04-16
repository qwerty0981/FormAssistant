package form

type googleForm struct{
   url string
   endpoints map[string]string
}

func GoogleForm(initialUrl string) googleForm{
   f := googleForm{url: initialUrl, endpoints: make(map[string]string)}
   return f
}

func (gf googleForm) addEndpoint(name, endpoint string) {
   gf.endpoints[name] = endpoint
}
