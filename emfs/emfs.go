package emfs

import (
	"io/ioutil"
)

var (
	resourcesFS IEmbedFS
	libsFS      IEmbedFS
)

type IEmbedFS interface {
	ReadFile(name string) ([]byte, error)
}

func SetLibsFS(lib IEmbedFS) {
	libsFS = lib
}

func SetResourcesFS(resource IEmbedFS) {
	resourcesFS = resource
}

func GetLibsFS() IEmbedFS {
	return libsFS
}

func GetResourcesFS() IEmbedFS {
	return resourcesFS
}

func GetResources(file string) ([]byte, error) {
	if GetResourcesFS() != nil {
		return GetResourcesFS().ReadFile(file)
	} else {
		return ioutil.ReadFile(file)
	}
}

func GetLibs(fileName string) ([]byte, error) {
	if GetLibsFS() != nil {
		return GetLibsFS().ReadFile(fileName)
	} else {
		return ioutil.ReadFile(fileName)
	}
}

func SetEMFS(libs IEmbedFS, resources IEmbedFS) {
	SetLibsFS(libs)
	SetResourcesFS(resources)
}
