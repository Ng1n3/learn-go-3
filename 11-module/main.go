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
// func Convert(from Money, to string) (Money, error) {
// ...
//}

//[lecture] -> In addition to updating go.mod, a go.sum file is created. For each module in the dependency tree of your project, the go.sum file has one or two entries: one with the module, its version, and a hash of the module; the other with the hash of the go.mod file for the module. When committing your own modules to source control, always include up-to-date go.mod and go.sum files. Doing so specifies exactly what versions of your dependencies are being used. This enables repeatable builds; when anyone else (including your future self) builds this module, they will get the exact same binary.

// [lecture] -> go list -> lists the packages that are used in your module. go list -m -versions github.com/learning-go-book-2e/simpletax. m changes the output to list the modules instead. -versions changes go list to report on the available versions for the specified module.

//! You might find that while module A works with version v1.1.0 of module D, it does not work with version v1.2.3. What do you do then? Go’s answer is that you need to contact the module authors to fix their incompatibilities. The import compatibility rule says, “If an old package and a new package have the same import path, the new package must be backward compatible with the old package.” All minor and patch versions of a module must be backward compatible. If they aren’t, it’s a bug. In our hypothetical example, either module D needs to be fixed because it broke backward compatibility, or module A needs to be fixed because it made a faulty assumption about the behavior of module D.You might find that while module A works with version v1.1.0 of module D, it does not work with version v1.2.3. What do you do then? Go’s answer is that you need to contact the module authors to fix their incompatibilities. The import compatibility rule says, “If an old package and a new package have the same import path, the new package must be backward compatible with the old package.” All minor and patch versions of a module must be backward compatible. If they aren’t, it’s a bug. In our hypothetical example, either module D needs to be fixed because it broke backward compatibility, or module A needs to be fixed because it made a faulty assumption about the behavior of module D.

//? You’ll notice the github.com/fatih/color module is declared to depend on version v0.0.14 of github.com/mattn/go-isatty, while github.com/mattn/go-colorable depends on v0.0.12. The Go compiler selects version v0.0.14 to use, because it is the minimal version that meets all requirements. This occurs even though, as of this writing, the latest version of github.com/mattn/go-isatty is v0.0.16. Your minimal version requirement is met with v0.0.14, so that’s what is used.

/*
[lecture] -> updating compatible versions. From v1.0.0 upgrading to v1.1.0. Use go get github.com/<packagename>@1.1.0. Then run go goet -u=patch github.com/<packagename>. This takes it to v1.1.1. go get -u github.com/<packagename> to 1.2.1
*/

