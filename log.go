package gobatis

import (
	"fmt"
)

//Info level
func (b *Batis) Info(message string, args ...interface{}) {
	b.olog.Println(fmt.Sprintf(message, args...))
}

//Warn level
func (b *Batis) Warn(message string, args ...interface{}) {
	b.olog.Println(fmt.Sprintf(message, args...))
}

//Error level
func (b *Batis) Error(message string, args ...interface{}) {
	b.elog.Println(fmt.Sprintf(message, args...))
}
