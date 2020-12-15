// Package tools ...
package tools

// Install will install tool with latest version
func Install(name string) {

}

// InstallWithVersion will install tool by specific version
func InstallWithVersion(name string) {

}

type tool interface {
	Install()
	Remove()
	Update()
}

type base struct {
	step    []string
	version []string
}
