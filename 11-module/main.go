package main

func main() {

}

// [lecture] -> Library management in Go is based around three concpets: repositories, modules, and packages. Repository is a place in a version control system where the source code for a project is stored. A module is a bundle of Go source code that's distributed and versioned as a single unit. They consist of one or more packages, which are directories of source code. Package give a module organization and strucure.

//! If you put a module with a nonunique name into a source code repository, it cannot be imported by another module.

// [lecture] -> Before using code from packages outside the standard library, you need to make sure that you have a properly created module. Every module has a globally unique identifier. This is not unique to Go. Java defines globally unique package declarations by using the reverse domain name convention (com.companyname.projectname.library). In Go, this name is called a module path. It is usually based on the repository where the module is stored.

// ? For example, the command line GOTOOLCHAIN=go1.18 go build will build your Go program with Go 1.18, downloading it if necessary. If both the GOTOOLCHAIN environment variable and the toolchain directive are set, the value assigned to GOTOOLCHAIN is used.If you haven’t seen it before, the %p verb in the fmt formatting language returns the memory location of a pointer. 

// [lecture] -> Require Directive: The require directives are present only if your module has dependencies. They list the modules that your module depends on and the minimum version required for each one.The first require section lists the direct dependencies of your module. The second one lists the dependencies of the dependencies of your module. Go uses capitalization to determine whether a package-level identifier is visible outside the package where it is declared. An identifier whose name starts with an uppercase letter is exported. Conversely, an identifier whose name starts with a lowercase letter or underscore can be accessed only from within the package where it is declared.

// [lecture] -> Creating and accessing a package: As a rule of thumb, you should make the name of the package match the directory that contains the package. package main is the starting point of your go application. Package name should be descriptive. In situation of circular dependencies, you should try to move the closely related packages into a single package.

//? Well documented Function
// Convert converts the value of one currency to another.
//
// It has two parameters: a Money instance with the value to convert,
// and a string that represents the currency to convert to. Convert returns
// the converted currency and any errors encountered from unknown or unconvertible
// currencies.
//
// If an error is returned, the Money instance is set to the zero value.
//
// Supported currencies are:
// - USD - US Dollar
// - CAD - Canadian Dollar
// - EUR - Euro
// - INR - Indian Rupee
//
// More information on exchange rates can be found at [Investopedia].
//
// [Investopedia]: https://www.investopedia.com/terms/e/exchangerate.asp
func Convert(from Money, to string) (Money, error) {
// ...
}