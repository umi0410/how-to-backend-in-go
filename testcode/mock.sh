# 이런 식으로 각각의 code generation tool 이용 가능
mockery --dir . --name "UserRepository" --inpackage
mockgen -package main -destination gomock.go -source repository.go --mock_names UserRepository=GomockRepository