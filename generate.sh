# build
cd _generate || exit
go build -o generate.out
cd ..

# generate
./_generate/generate.out -pkg "emoji" -o "emoji.go"
