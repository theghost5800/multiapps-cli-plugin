#!/bin/bash -eux

function build() {
    local version=$1
    local platform=$2
    local arch=$3
    local plugin_name=$4

    echo "building for $platform $arch"
    GOOS=$platform GOARCH=$arch go build \
        -ldflags "-X main.Version=${version}" \
        -o ${plugin_name}
}

function buildstatic() {
    local version=$1
    local platform=$2
    local arch=$3
    local plugin_name=$4

    echo "building static for $platform $arch"
    CGO_ENABLED=0 GOOS=$platform GOARCH=$arch go build -a -tags netgo \
        -ldflags "-w -extldflags \"-static\" -X main.Version=${version}" \
        -o ${plugin_name}
}

function movePluginsToBuildFolder() {
    local folder=$1
    mv $PLUGIN_NAME_WIN_32 $folder
    mv $PLUGIN_NAME_WIN_64 $folder
    mv $PLUGIN_NAME_LINUX_32 $folder
    mv $PLUGIN_NAME_LINUX_64 $folder
    mv $PLUGIN_NAME_OSX $folder
    mv $PLUGIN_NAME_STATIC_WIN_32 $folder
    mv $PLUGIN_NAME_STATIC_WIN_64 $folder
    mv $PLUGIN_NAME_STATIC_LINUX_32 $folder
    mv $PLUGIN_NAME_STATIC_LINUX_64 $folder
    mv $PLUGIN_NAME_STATIC_OSX $folder
}

function createBuildMetadataFiles() {
    local version=$1
    local folder=$2
	echo -n "v${version}" > ${folder}/version
	grep -Pzoa "(?s)## v${version}(.*?)##" CHANGELOG.md | grep -va "##" | tr -s '\n' '\n' > ${folder}/changelog
}

script_dir="$(dirname -- "$(realpath -- "${BASH_SOURCE[0]}")")"
cd "${script_dir}"

BUILD_FOLDER=build
PLUGIN_NAME_LINUX_64=multiapps-plugin.linux64

PLUGIN_NAME_STATIC_LINUX_64=multiapps-plugin-static.linux64

version=$(<cfg/VERSION)
build $version linux arm64 $PLUGIN_NAME_LINUX_64

buildstatic $version linux arm64 $PLUGIN_NAME_STATIC_LINUX_64

mkdir $BUILD_FOLDER -p
createBuildMetadataFiles $version $BUILD_FOLDER
movePluginsToBuildFolder $BUILD_FOLDER
