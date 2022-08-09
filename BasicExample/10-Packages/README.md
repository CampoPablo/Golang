#Package

A package is a collection of files which declare constants, types, variables and functions belonging to the package.

All source files of a package are located in a single directory. Each source file begins with a package clause defining the package to which it belongs. Constants, types, variables and functions belonging to a package are accessible in all files of that package. The elements of the package can be exported to other packages when needed.

All source files of a package are stored in a single directory. The name of the directory and the package does not have to match, but is is a common practise that both share the same name.

In order to use elements of a package within another package, those elements must be exported. In Go, the exporting is performed by capitalizing the names of the elements.

Packages serve the following purpose:

- organize code for reuse
- prevent name conflicts
- speed up compiler by allowing recompilation of smaller chunks of a program

Packages are further organized into modules. A module is a collection of Go packages stored in a file tree with a go.mod file at its root. The module path is the prefix path that is used to import all packages of that module.