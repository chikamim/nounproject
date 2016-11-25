# nounproject
nounproject api client

## example
    package main

    import (
      "fmt"

      "github.com/chikamim/nounproject"
    )

    const (
      consumerKey    = "..."
      consumerSecret = "..."
    )

    func main() {
      client := nounproject.NewClient(consumerKey, consumerSecret)
      res, err := client.GetIcons("star", true, &nounproject.Pagination{Limit: 10, Offset: 0, Page: 0})
      if err != nil {
        panic(err)
      }

      for _, icon := range res {
        fmt.Println(icon)
        icon.DownloadPreview(".")
      }
      
      // try these functions
      //res, err := client.GetRecentUploads(&nounproject.Pagination{Limit: 3, Offset: 2, Page: 1})
      //res, err := client.GetCollection("1")
      //res, err := client.GetCollectionIcons("star", nil)
      //res, err := client.GetCollections(nil)
      //res, err := client.GetUsage()
      //res, err := client.GetUserCollections(434089, "")
      //res, err := client.GetUserUploads("9738729691", nil)
      //res, err := client.GetUserCollections(434089, "")
    }
