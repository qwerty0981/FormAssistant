package form

import (
   "errors"
   "net/url"
   "net/http"
   "strings"
)

type googleForm struct{
   url string
   endpoints map[string]string
}

func GoogleForm(initialUrl string) googleForm{
   f := googleForm{url: initialUrl, endpoints: make(map[string]string)}
   return f
}

func (gf googleForm) AddEndpoint(name, endpoint string) {
   gf.endpoints[name] = endpoint
}

func (gf googleForm) Endpoints() map[string]string {
   return gf.endpoints
}

func (gf googleForm) Post(values map[string]string) error {
   for key, _ := range values {
      _, ok := gf.endpoints[key]
      if ok == false {
         return errors.New("Key (" + key + ") not found in endpoints")
      }
   }

   httpClient := http.Client{}

   form := url.Values{}
   for k, v := range values {
      form.Add(gf.endpoints[k], v)
   }

   req, _ := http.NewRequest("POST", gf.url, strings.NewReader(form.Encode()))
   req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
   httpClient.Do(req)
   return nil
}
