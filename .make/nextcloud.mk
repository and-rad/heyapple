NC_APPS_DIR := /tmp/nextcloud/apps

build-nextcloud:
# copy files to output directory
	@rm -rf out/nextcloud
	@cp -r clients/nextcloud out

# copy files to NextCloud testing environment
	@mkdir -p ${NC_APPS_DIR}/${BINARY_NAME}
	@rm -rf ${NC_APPS_DIR}/${BINARY_NAME}/*
	@cp -r out/nextcloud/* ${NC_APPS_DIR}/${BINARY_NAME}
