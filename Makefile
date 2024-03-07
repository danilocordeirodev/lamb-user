GOARCH=arm64 
GOOS=linux 

build:
	go build -tags lambda.norpc -o ./bin/bootstrap

deploy: build
	cd bin && zip -r bootstrap.zip bootstrap
	cd bin && aws lambda update-function-code --function-name "lambUser" --zip-file fileb://bootstrap.zip --region="us-east-1"