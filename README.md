# Article MAker

## How to run the application

1. Make sure `GOBIN` is properly set up in your computer

2. Clone the application with `git clone https://github.com/bmulobi/Article-Maker.git`

3. Navigate to the directory you have cloned i.e Article-Maker

4. Run `make install`

5. edit `config.yml` in the root folder

> Note: I have tested the app with MySQL, haven't tested the other db drivers.
 Create your database before running the app

6. To run the application, run `make run`

7. To run the tests, run `make test`

> Note: Create your test database before running the tests

## Endpoints Description

### Create Article

```JSON
    URL - *http://localhost:<port>/api/v1/article*
    Method - POST

    Request Body:
    {
        "title": "Article title",
        "body": "Article Body",
        "publisher": {
        	"Name": "Publisher Name"
        },
        "category": {
        	"Name": "Category Name"
        }
    }
    - You can also create a new article by passing in the id of an existing category or publisher
    {
            "title": "Article title",
            "body": "Article Body",
            "publisher": {
            	"id": 1
            },
            "category": {
            	"id": 1
            }
    }

    Response:
    {
        "Value": {
            "id": 10,
            "title": "Article title",
            "body": "Article Body",
            "publisher_id": 7,
            "publisher": {
                "id": 7,
                "name": "Publisher Name"
            },
            "category_id": 15,
            "category": {
                "id": 15,
                "name": "Category Name"
            },
            "created_at": "2020-02-25T00:34:19.484446+03:00",
            "updated_at": "2020-02-25T00:34:19.484446+03:00",
            "published_at": "0001-01-01T00:00:00Z"
        },
        "Error": null,
        "RowsAffected": 1
    }
    
```

### Get Article

```JSON
    URL - *http://localhost:<port>/api/v1/article/<articleID>*
    Method - GET

    Response:
    {
        "id": 10,
        "title": "Article title",
        "body": "Article Body",
        "publisher_id": 7,
        "publisher": {
            "id": 7,
            "name": "Publisher Name"
        },
        "category_id": 15,
        "category": {
            "id": 15,
            "name": "Category Name"
        },
        "created_at": "2020-02-25T00:34:19+03:00",
        "updated_at": "2020-02-25T00:34:19+03:00",
        "published_at": "0001-01-01T00:00:00Z"
    }
```

### Get Articles

```JSON
    URL - *http://localhost:<port>/api/v1/article*
    URL - *http://localhost:<port>/api/v1/article?publisher=Publisher Name&category=Category Name&created_at=2020-02-25 00:34:19*
    Method - GET
    Query Parameters: publisher, category, created_at, published_at
    You can use the query parameters in any combination, or none
```

### Update Article

```JSON
    URL - *http://localhost:<port>/api/v1/article*
    Method - PUT
    Body:
    {
    	"id": 1,
        "title": "Updated Title",
        "body": "Updated body",
        "publisher": "Updated publisher",
        "category": "Updated category",
        "published_at": "2022-02-23 23:20:00" 
    }
```

### Delete Article

```JSON
    URL - *http://localhost:<port>/api/v1/article/<articleID*
    Method - DELETE
```