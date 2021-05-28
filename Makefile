binname = pslite

build_osx: create_osx_icon
	go build -ldflags "-s -w" -o pslite.app/Contents/MacOS/$(binname) .

create_osx_icon:
	chmod u+x ./script/png2icns.sh
	./script/png2icns.sh ps.png
	mv ps.icns ./pslite.app/Contents/Resources

macopen: build_osx
	open ./pslite.app

macapp: build_osx
	cp -rf pslite.app /Applications