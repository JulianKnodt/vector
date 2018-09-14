// A package which implements basic 3d vector functionality for
// float64s. Most functions that return a vector allow for method chaining on pointers
// with <Name>Set that will perform an operation and set the vector
// This is both for performance and readability.
// All functions that create a new vector return a pointer, that allows for immediate usage
// of *Set functions, and they should be preferred for performance reasons
package vector
