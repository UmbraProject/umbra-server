package main

import (
	"log"
  "fmt"
	"github.com/titan-x/gcm/ccs"
)

// Example demonstrating the use of CCS implementation in an application server.
func main() {
  fmt.Printf("Starting GCM CCS Server...\n")

	c, err := ccs.Connect("gcm-preprod.googleapis.com:5236", "845947285905", "AIzaSyBRAAGQ6A65FGtOJmFtpT1GLQDZHhQufQA", false)
	if err != nil {
		log.Fatalf("GCM CCS connection cannot be established")
	}

	// Send a test message. Replace "device_registration_id" with an actual GCM registration ID from a device.
	n, err := c.Send(&ccs.OutMsg{To: "dGhWxsY7nmw:APA91bGiaP5pEd0w5bm_TcmGimly75sEJZtwyLx7rB4AHiqLhvn565zzHpnsyjJ3uaRv_NZ7rHbjEKXGojZVWySBvYL4JV8yB6Z_kAo18b-a9dC1ZvRyho69D0S6_wjgXKCDdabNMUcQ", Data: map[string]string{"test_message": "GCM CCS client testing message."}})
	if err != nil {
		log.Printf("Failed to send message to CCS server with error: %v\n", err)
	}
	log.Printf("Message sent with %v bytes written to the connection\n", n)

	// Start receiving messages from the CCS server.
	for {
		log.Println("Waiting for incoming CCS messages")
		m, err := c.Receive()
		if err != nil {
			log.Printf("Incoming CCS error: %v\n", err)
		}

		go handleMessage(m)
	}
}

func handleMessage(m *ccs.InMsg) {
  log.Printf("----------------------------------------------------------")
	log.Printf("Incoming CCS message: %v\n", m)
  log.Printf("----------------------------------------------------------")
}
