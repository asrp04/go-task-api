package main

import (
	"fmt"
	"time"
)

// ईमेल भेजने वाला फंक्शन, जो काम खत्म होने पर चैनल में एक मैसेज डाल देगा
func sendEmail(user string, ch chan string) {
	fmt.Printf("📧 %s के लिए ईमेल प्रोसेस शुरू...\n", user)
	time.Sleep(2 * time.Second) // काम होने में लगा समय
	
	// 🛶 चैनल के अंदर डेटा भेजना (Arrow direction: ch <- data)
	ch <- fmt.Sprintf("✅ %s का काम पूरा हो गया!", user)
}

func main() {
	startTime := time.Now()

	// 1. एक स्ट्रिंग टाइप का चैनल बनाना
	ch := make(chan string)

	// 2. बैकग्राउंड में 3 Goroutines शुरू करना और उन्हें चैनल पास करना
	go sendEmail("User A", ch)
	go sendEmail("User B", ch)
	go sendEmail("User C", ch)

	// 3. चैनल से डेटा रिसीव करना (Arrow direction: variable <- ch)
	// यह कोड तब तक 'ब्लॉक' रहेगा जब तक चैनल में डेटा नहीं आ जाता।
	// इसलिए हमें मैन्युअली time.Sleep लगाने की कोई ज़रूरत नहीं है!
	msg1 := <-ch
	msg2 := <-ch
	msg3 := <-ch

	// प्रिंट कर रहे हैं जो मैसेज चैनल से मिले
	fmt.Println(msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)

	fmt.Printf("⏱️ कुल समय लगा: %v\n", time.Since(startTime))
}
