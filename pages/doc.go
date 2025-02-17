// WARNING -- THIS PACKAGE SHOULD NOT CONTAIN GLOBAL STATE!!!
// The metaprogram uses this package to generate code, and global state may not
// be initialized properly.
//
// I know this isn't ideal, but this is a limitation of Golang's metaprogramming
// capability.
package pages