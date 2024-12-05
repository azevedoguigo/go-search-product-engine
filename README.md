# Go Search Product Engine

## About the project
This project is a search engine specialized in searching for products for websites such as e-commerce. My goal with this project is initially focused on learning how this type of implementation works, but if you want to contribute, feel free to send a pull request and contribute to the engine.


To use the project, you need to connect to your database and when running for the first time, all products will be indexed in Elasticsearch.

You might be thinking that this current approach of indexing items only at project startup, right? I agree. But the project is still very experimental and should certainly not be implemented in this way in an e-commerce. My noob suggestion is that for each new product added to the database, a new indexing is performed. This way, new products will be available for search as soon as they are added to the store.

## Feature
The term entered by the user during the search will be used to search in the product name and description fields, considering that the name field has a weight three times greater than description.

## Demo

Request: http://localhost:8080/search?q=testing

Response: </br>
![image](https://github.com/user-attachments/assets/c1f9fff6-e3e7-44bb-b0c6-d77cbdd785ea)
</br>
(Note that the first item is the one that contains the search term as the product name.)


Request: http://localhost:8080/search?q=star-wars

Response: </br>
![image](https://github.com/user-attachments/assets/2250cb44-a62c-4d41-846f-19237873c030)
</br>
(Multiple items with similar names.)

Request: http://localhost:8080/search?q=275mm

Response: </br>
![image](https://github.com/user-attachments/assets/8bf96221-5e8a-4d84-a539-ff56f35acfbd)
</br>
(Multiple items with similar term in description.)
