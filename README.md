# unit_test

go install github.com/golang/mock/mockgen@v1.6.0

go get github.com/golang/mock/gomock


mockgen -source=path/to/your/interface.go -destination=path/to/your/mock/mock_interface.go -package=mock


