module github.com/aws/aws-sdk-go-v2/example/service/s3/listObjects

go 1.15

require (
	github.com/aws/aws-sdk-go-v2/config v0.3.0
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v0.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v0.30.0
	github.com/aws/aws-sdk-go-v2 v0.30.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v0.3.1
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v0.1.2
	github.com/aws/aws-sdk-go-v2/service/sts v0.30.0
	github.com/aws/aws-sdk-go-v2/credentials v0.1.5
	github.com/aws/aws-sdk-go-v2/ec2imds v0.1.5
)

replace github.com/aws/aws-sdk-go-v2/config => ../../../../config/

replace github.com/aws/aws-sdk-go-v2/service/s3 => ../../../../service/s3/

replace github.com/aws/aws-sdk-go-v2 => ../../../../

replace github.com/aws/aws-sdk-go-v2/credentials => ../../../../credentials/

replace github.com/aws/aws-sdk-go-v2/ec2imds => ../../../../ec2imds/

replace github.com/aws/aws-sdk-go-v2/service/sts => ../../../../service/sts/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../../../service/internal/s3shared/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../../../service/internal/presigned-url/
