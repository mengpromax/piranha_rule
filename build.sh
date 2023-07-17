#/bin/bash

if [ ! -d "target" ]; then
    mkdir target
fi

gox -osarch "darwin/amd64 darwin/arm64" -output target/{{.Arch}}/piranha_rule

for f in `find target -type d ! -name "target" ! -path "*executors*" ! -path "*rules*"`;do
    cp -r executors $f
    cp -r rules $f
    chmod +x $f/executors/*
done

if [ ! -d "target/package" ]; then
    mkdir target/package
fi

for f in `find target -type d ! -name "target" ! -path "*executors*" ! -path "*package*"  ! -path "*rules*"`;do
    tar -cvf $(dirname $f)/package/$(basename $f).tar $f
done