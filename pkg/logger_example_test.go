package seekable_test

import (
	"fmt"

	seekable "github.com/rjasonadams/zstd-seekable-format-go/pkg"
)

// ZapLogger demonstrates how to adapt zap.Logger to the seekable.Logger interface.
// Uncomment this example if you want to use zap:
//
// import "go.uber.org/zap"
//
// type ZapLogger struct {
//     *zap.Logger
// }
//
// func (z ZapLogger) Debug(msg string, fields ...seekable.Field) {
//     zapFields := make([]zap.Field, len(fields))
//     for i, f := range fields {
//         switch v := f.Value.(type) {
//         case string:
//             zapFields[i] = zap.String(f.Key, v)
//         case int:
//             zapFields[i] = zap.Int(f.Key, v)
//         case uint64:
//             zapFields[i] = zap.Uint64(f.Key, v)
//         case uint32:
//             zapFields[i] = zap.Uint32(f.Key, v)
//         default:
//             zapFields[i] = zap.Any(f.Key, v)
//         }
//     }
//     z.Logger.Debug(msg, zapFields...)
// }
//
// Usage:
//     zapLogger, _ := zap.NewDevelopment()
//     logger := ZapLogger{zapLogger}
//     reader, err := seekable.NewReader(rs, decoder, seekable.WithRLogger(logger))

// SimpleLogger demonstrates a basic logger implementation that prints to stdout.
type SimpleLogger struct{}

func (l SimpleLogger) Debug(msg string, fields ...seekable.Field) {
	fmt.Print(msg)
	for _, f := range fields {
		fmt.Printf(" %s=%v", f.Key, f.Value)
	}
	fmt.Println()
}

// Example showing how to use a custom logger with the reader.
func ExampleWithRLogger() {
	// Create a simple logger
	logger := SimpleLogger{}

	// Use it when creating a reader
	// reader, err := seekable.NewReader(rs, decoder, seekable.WithRLogger(logger))
	_ = logger
}

// Example showing how to use a custom logger with the writer.
func ExampleWithWLogger() {
	// Create a simple logger
	logger := SimpleLogger{}

	// Use it when creating a writer
	// writer, err := seekable.NewWriter(w, encoder, seekable.WithWLogger(logger))
	_ = logger
}
