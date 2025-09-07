# ThinkMoney Go Challenge
   
## Description
The purpose of this task was to create a solution to the supermarket checkout problem

### Features
- **Unit Pricing**: The "normal" price per item
- **Discount Pricing**: The discounted price per item when a threshold has been reached

## Running and Testing
This will need cloning down from the git repo like so

```bash
git clone https://github.com/SamWilco-ai/ThinkMoney.git
```

once cloned 

```bash
cd cd ThinkMoney
```

### Running
To simply run the project and view the desired total sum of all products simply run in the ThinkMoney directory where the main.go file is located

```bash
go run .
```

Presently it is already a predefined list of items with an unexpeceted item in the bagging area that gets logged out


### Testing

To test the project you can run

```bash
go test ./...
```

I have already included the coverage profile but this can also be generated again (just in case you dont believe me) by running
```bash
go test -v -coverprofile cover.out ./
go tool cover -html cover.out -o cover.html
```

This will generate the coverage profile which is then subsequently converted into a more readble HTML format

#### Test Cases

Tests were split into the two main pieces of functionality, the scan function and the total price function. I tried to cover as many test cases as possible without going overboard.
