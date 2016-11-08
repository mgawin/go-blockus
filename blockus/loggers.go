package blockus

import (
        "log"
        "os"
    )

var (
    ErrorLog *log.Logger
    DebugLog *log.Logger
)

func init() {

    ErrorLog = log.New(os.Stdout,
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
    
    
    
     DebugLog = log.New(os.Stderr,
        "MEINDEBUG: ",
        log.Ldate|log.Ltime|log.Lshortfile)
    }
    

