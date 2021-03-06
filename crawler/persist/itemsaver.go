package persist

import (
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver Got Item #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

func save(item interface{}) {

}
