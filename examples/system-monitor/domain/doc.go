// Package domain holds the value types the system monitor passes between its
// collector and its UI. Every snapshot is immutable: a collector builds a
// value, publishes it, and never writes to it again. Nothing in this package
// reads the machine.
//
// Units are fixed here. Cumulative CPU times are nanoseconds, CPU percentages
// run 0 to 100 across all logical cores, memory and disk values are bytes,
// rates are bytes per second, temperatures are degrees Celsius, and counts are
// plain numbers. A Value with Valid false means the source could not read that
// item, which is not the same as a measured zero.
package domain
