module github.com/aws/aws-sdk-go-v2/service/s3control

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.16.5
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.12
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.6
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.6
	github.com/aws/smithy-go v1.11.3
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
