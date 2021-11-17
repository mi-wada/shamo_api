ENDPOINT_URL=http://localhost:8000

# create table `shamo_user`
aws dynamodb create-table \
  --table-name shamo_user \
  --attribute-definitions \
    AttributeName=id,AttributeType=S \
  --key-schema \
    AttributeName=id,KeyType=HASH \
  --endpoint-url $ENDPOINT_URL \
  --provisioned-throughput \
    ReadCapacityUnits=10,WriteCapacityUnits=5

# create table `shamo_room`
aws dynamodb create-table \
  --table-name shamo_room \
  --attribute-definitions \
    AttributeName=id,AttributeType=S \
  --key-schema \
    AttributeName=id,KeyType=HASH \
  --endpoint-url $ENDPOINT_URL \
  --provisioned-throughput \
    ReadCapacityUnits=10,WriteCapacityUnits=5

# create table `shamo_room_member`
aws dynamodb create-table \
  --table-name shamo_room_member \
  --attribute-definitions \
    AttributeName=room_id,AttributeType=S \
    AttributeName=user_id,AttributeType=S \
  --key-schema \
    AttributeName=room_id,KeyType=HASH \
    AttributeName=user_id,KeyType=RANGE \
  --endpoint-url $ENDPOINT_URL \
  --provisioned-throughput \
    ReadCapacityUnits=10,WriteCapacityUnits=5

# create table `shamo_payment`
aws dynamodb create-table \
  --table-name shamo_payment \
  --attribute-definitions \
    AttributeName=id,AttributeType=S \
    AttributeName=created_at,AttributeType=S \
    AttributeName=room_id,AttributeType=S \
  --key-schema \
    AttributeName=id,KeyType=HASH \
  --global-secondary-indexes \
    "[{
      \"IndexName\": \"room_id-created_at-index\",
      \"KeySchema\": [
        {\"AttributeName\": \"room_id\", \"KeyType\": \"HASH\"},
        {\"AttributeName\": \"created_at\", \"KeyType\": \"RANGE\"}
      ],
      \"Projection\": {
        \"ProjectionType\": \"ALL\"
      },
      \"ProvisionedThroughput\": {
        \"ReadCapacityUnits\": 10,
        \"WriteCapacityUnits\": 5
      }
     }]" \
  --endpoint-url $ENDPOINT_URL \
  --provisioned-throughput \
    ReadCapacityUnits=10,WriteCapacityUnits=5
