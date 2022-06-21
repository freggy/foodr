REWE Shop API documentation
===========================

This document describes the REWE online shop API. Endpoints are extracted from the decompiled Android app. The base endpoint currently is `https://mobile-api.rewe.de`.
I discovered this endpoint in a [GitHub project that tracks discounts](https://github.com/foo-git/rewe-discounts/blob/master/rewe_discounts/rewe_discounts.py#L186).

To make a functional order we do not need to use all endpoints the API provides. In the section below only the needed calls are described. 

TODO: describe endpoints in open api format.

Endpoints
---------

### Product

#### `GET /products/ean/<ean>`

Gets a product by its EAN.

**Response body**
```
curl https://mobile-api.rewe.de/products/ean/7613035890008
{
  "items": [
    {
      "articleId": "7613035890008",
      "productId": "3244310",
      "description": "Für alle, die chinesische Nudeln lieben, sind Maggi Magic Asia Gebratene Nudeln Huhn genau das Richtige. Genieße kross angebratene asiatische Nudeln mit Huhn und einer leckeren Gemüseeinlage mit Paprika, Karotten, und Zwiebeln \n\nSchnelle Küche: Gebratene Nudeln mit Huhn\n\nEinfach 375ml Wasser in der Pfanne erwärmen, den Beutelinhalt einrühren und aufkochen. Das Ganze ca. 5 Minuten weiter offen kochen, bis die Flüssigkeit von den gebratenen Nudeln aufgenommen worden ist. Beim Kochen bitte die Asia-Nudeln gelegentlich umrühren. Zum Schluss noch 1 El Öl in die Pfanne geben und die Nudeln unter regelmäßigem Wenden knusprig anbraten. Fertig!\n\nOb nach einem langen Tag im Büro, zum Abendessen oder beim Entspannen auf der Couch – Maggi Magic Asia Gebratene Nudeln mit Huhn schmecken immer. Mit wenigen Handgriffen kann man sich die leckeren asiatischen Nudeln mit Huhn zubereiten und den Gang zum Asia-Imbiss sparen.\n\nMaggi Magic Asia Instant Nudel Snack Huhn:\n- Ergibt zwei leckere Portionen.\n- Gebratene Nudeln mit Huhn.\n- Bringt asiatische Abwechslung in die Küche.\n- Bequeme und schnelle Zubereitung.\n- Das braucht man dazu: eine Pfanne, einen Pfannenwender Wasser und Öl.",
      "categoryIds": [
        "391",
        "392",
        "489"
      ],
      "ean": "7613035890008",
      "name": "Maggi Magic Asia Gebratene Nudeln Huhn 121g"
    }
  ],
  "resultsSource": "REWE_GTIN"
}
```

#### `GET /mobile/products/search`

**Params**
* `searchTerm` - **required**
* `page`
* `objectsPerPage`

Searches for product by given term.

**Response body**

```
curl https://mobile-api.rewe.de/mobile/products/search?searchTerm=wasser
{
  "products": [
    {
      "title": "Vio Mineralwasser Still 6x1l",
      "depositLabel": "(MEHRWEG)",
      "price": "",
      "imageURL": "https://img.rewe-static.de/8067207/29954700_digital-image.png",
      "rawValues": {
        "productId": "8067207",
        "orderLimit": 99,
        "hasVariants": false,
        "showMarketPlaceBanner": false
      },
      "tr": {
        "id": "8067207",
        "price": "not_set",
        "name": "Vio Mineralwasser Still 6x1l",
        "merchantType": "not_set",
        "merchantName": "not_set",
        "savingsAmount": "0.00",
        "savingsPercent": "0.00"
      }
    },
    {
      "title": "Vio Mineralwasser Medium 6x1,5l",
      "depositLabel": "(EINWEG)",
      "price": "",
      "imageURL": "https://img.rewe-static.de/1964856/21766006_digital-image.png",
      "rawValues": {
        "productId": "1964856",
        "orderLimit": 99,
        "hasVariants": false,
        "showMarketPlaceBanner": false
      },
      "tr": {
        "id": "1964856",
        "price": "not_set",
        "name": "Vio Mineralwasser Medium 6x1.5l",
        "merchantType": "not_set",
        "merchantName": "not_set",
        "savingsAmount": "0.00",
        "savingsPercent": "0.00"
      }
    }
  ]
}
```

### Basket

### Order
