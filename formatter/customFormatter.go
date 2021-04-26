package log

import (
    "bytes"
    "github.com/sirupsen/logrus"
)

type CustomFormatter struct {
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
    time := entry.Time.Format("2006-01-02 15:04:05")
    buffer := bytes.Buffer{}
    buffer.WriteString(time)
    buffer.WriteString(" [")
    buffer.WriteString(entry.Level.String())
    buffer.WriteString("]-")
    buffer.WriteString(entry.Data["line"].(string))
    buffer.WriteString(":")
    buffer.WriteString(entry.Message)
    buffer.WriteByte('\n')
    return buffer.Bytes(), nil
}
