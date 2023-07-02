package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type Product struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Photo       string `json:"photo"`
}

func main() {
    url := "https://enstack.gigalixirapp.com//api/products"

    // Create a Product instance
    product := Product{
        Name:        "Adidas",
        Description: "Adidas is a popular streetfood",
        Photo:       "https://enstack.pages.dev/a/adidas.jpg",
    }

    // Convert the Product instance to JSON
    jsonData, err := json.Marshal(product)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Send the POST request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        fmt.Println("Product created successfully.")
    } else {
        fmt.Println("Failed to create product. Status code:", resp.StatusCode)
    }

    // Read the response body
    body := new(bytes.Buffer)
    _, _ = body.ReadFrom(resp.Body)
    fmt.Println("Response Body:", body.String())
}
