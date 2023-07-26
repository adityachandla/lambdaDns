GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o $1 cmd/$1/main.go
echo "Compilation Done"
zip $1.zip $1
echo "Created zip file"
aws s3 cp $1.zip s3://codebucketac/$1.zip
echo "Uploaded file"
rm $1 $1.zip
echo "Updating lambda"
aws lambda update-function-code --function-name $1 --s3-bucket codebucketac --s3-key $1.zip
