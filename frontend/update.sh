cd home
npm run build
rm -rf ../../web_root
cp -r dist ../../web_root

cd ../post_html_generator
go run src/error.go src/list.go src/main.go src/metadata.go src/pandoc.go src/parser.go