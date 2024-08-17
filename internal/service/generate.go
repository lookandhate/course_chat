package service

//go:generate sh -c "rm -rf mocks && mkdir -- mocks"
//go:generate minimock -i ChatService -o ./mocks/ -s "_minimock.go"
