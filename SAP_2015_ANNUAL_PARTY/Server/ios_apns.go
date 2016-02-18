package main

import (
	"fmt"
	//"github.com/bitly/go-simplejson"
	"github.com/virushuo/Go-Apns"
	"os"
	"time"
)

func main() {
	apn, err := apns.New("apns-dev.pem", "apns-dev-key-noenc.pem", "gateway.sandbox.push.apple.com:2195", 1*time.Second)
	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("connect successed!")
	go readError(apn.ErrorChan)
	token := "136c6505d4141b35d4a55365a191b0e2fa74e6cec91bcd2b3f2dd298bf1339a6"

	payload := apns.Payload{}
	payload.Aps.Alert.Body = "Congratulations!\nYou won a sport camera in the raffle!\nPlease go to the right side of the stage after the party to claim your prize or contact Ms. Karen Zhao at 18800349005."
	payload.Aps.Sound = "bingbong.aiff"
	payload.SetCustom("id", time.Now().Unix())
	payload.SetCustom("tp", 0)

	//{"id":"12345678","tp":0,"aps":{"alert":{"body":"Message content"}}}

	//js, err := simplejson.NewJson([]byte(`{}`))
	//js.Set("aps", "alert")
	//	js.Set("aps", "badge")
	//	js.Set("badge", 2)
	//	js.Set("alert", "body")
	//	js.Set("alert", "action-loc-key")
	//body, _ := js.String()
	//fmt.Println(string(js))

	//body, _ := js.String()
	//payload.Aps.Alert.Body = body

	notification := apns.Notification{}
	notification.DeviceToken = token
	notification.Identifier = 0
	notification.Payload = &payload
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)

	apn.Close()
}

func readError(errorChan <-chan error) {
	for {
		apnerror := <-errorChan
		fmt.Println(apnerror.Error())
	}
}
