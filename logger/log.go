package logger

import (
	"log"
	"os"
)

//Logger pcap_reader formatted log.
var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "pcap-reader: ", log.Ldate|log.Lmicroseconds)
	// Logger = log.New(os.Stdout, "pcap-reader:", log.Ldate|log.Lmicroseconds|log.Llongfile)
}
