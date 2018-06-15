package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"crawler/engine"
	"errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		//must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : Got item #%d : %v", itemCount, item)
			itemCount++
			err := save(client, item, index)
			if err != nil {
				log.Printf("Item Saver: error saving item %v : %v",
					item, err)
			}

		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {
	if item.Type == "" {
		return errors.New("must suppply Type")
	}
	indexService := client.Index().
		Index(index). //database
		Type(item.Type). //table
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
