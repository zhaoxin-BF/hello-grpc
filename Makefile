.PHONY: buf
buf:
	buf generate

build:
	cd proto && \
	for file in $$(find . -type f -name "buf.gen.yaml" | tr -s "\n" " ");do \
		echo "build " $${file}; \
		path=$${file%/*}; \
		echo $${path}; \
		buf generate --template $$file --path $$path; \
	done
