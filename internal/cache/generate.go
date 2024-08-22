package cache

//go:generate sh -c "rm -rf mocks && mkdir -- mocks"
//go:generate minimock -i ChatCache -o ./mocks/ -s "_minimock.go"
