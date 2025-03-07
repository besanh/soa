# Curl below
I provided some curls. Please replace http://localhost:8000 with your address


1. Get Product categories

```
curl --location 'http://localhost:8000/v1/product-categories' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT'
```

2. Insert Product Category

```
curl --location 'http://localhost:8000/v1/product-categories' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT' \
--header 'Content-Type: application/json' \
--data '{
    "product_category_name": "Clothing",
    "status": "active"
}'
```

3. Get Supplier

```
curl --location --request GET 'http://localhost:8000/v1/suppliers' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT' \
--header 'Content-Type: application/json' \
--data '{
    "supplier_name": "Anh Le",
    "status": "active"
}'
```

4. Insert supplier

```
curl --location 'http://localhost:8000/v1/suppliers' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT' \
--header 'Content-Type: application/json' \
--data '{
    "supplier_name": "Anh Le",
    "status": "active"
}'
```

5. Get products

Scrolling

```
curl --location 'http://localhost:8000/v1/products/scroll' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT'
```

Normal
```
curl --location 'http://localhost:8000/v1/products/scroll' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT'
```

6. Insert products

```
curl --location 'http://localhost:8000/v1/products' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT' \
--header 'Content-Type: application/json' \
--data '{
    "supplier_id": "20dcd6f5-3ea5-41a6-a89d-59494b1e3170",
    "product_category_id": "e2bb9114-6532-40c1-a9ec-ebc92274dc72",
    "product_name": "Test",
    "product_reference": "Test123",
    "status": "available",
    "price": 10000,
    "stock_location": "Ho Chi Minh city",
    "quantity": 10
}'
```

7. Get distance

```
curl --location 'http://localhost:8000/v1/distance' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT'
```

8. Statistics

Statistics product per category

```
curl --location 'http://localhost:8000/api/statistics/products-per-category' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT'
```

Statistics product per supplier

```
curl --location 'http://localhost:8000/api/statistics/products-per-supplier' \
--header 'Authorization: 8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT'
```