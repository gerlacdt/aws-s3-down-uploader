.PHONY: cmd clean

cmd:
	cd cmd/s3; go build

clean:
	rm -f cmd/s3/s3
