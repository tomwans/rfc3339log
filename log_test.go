package rfc3339log

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"
)

func TestLoggingToBuffer(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(buf, "", log.LstdFlags)
	l.Println("hey there")
	if !strings.HasSuffix(buf.String(), "hey there\n") {
		t.Fatalf("logging the actual text to the buffer don't work: %s", buf.String())
	}
}

func TestEnsureTimeIsLoggedAsRFC3339(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(buf, "", log.LstdFlags)
	l.Println("ok")
	x := buf.String()
	xs := strings.Split(x, " ")
	_, err := time.Parse(time.RFC3339, strings.TrimSpace(xs[0]))
	if err != nil {
		t.Fatalf("expected first non-space characters to be the RFC3339 timestamp: %s", err)
	}
}

func TestEnsurePrefixCanBeSet(t *testing.T) {
	goodpref := "testing!"
	buf := &bytes.Buffer{}
	l := New(buf, goodpref, log.LstdFlags)
	l.Println("ok")
	x := buf.String()
	if !strings.HasPrefix(x, goodpref) {
		t.Fatalf("expected prefix %s in '%s'", goodpref, x)
	}
}

func TestPrint(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(buf, "", log.LstdFlags)
	l.Print("ok")
	x := buf.String()
	if !strings.HasSuffix(x, "ok") {
		t.Fatalf("expected suffix %s in '%s'", "ok", x)
	}
}

func TestPrintf(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(buf, "testing! ", log.LstdFlags)
	l.Printf("%s %d", "ok", 1)
	x := buf.String()
	if !strings.HasSuffix(x, "ok 1") {
		t.Fatalf("expected suffix %s in '%s'", "ok 1", x)
	}
}

func TestPrintln(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New(buf, "", log.LstdFlags)
	l.Println("um", "ok")
	x := buf.String()
	if !strings.HasSuffix(x, "um ok\n") {
		t.Fatalf("expected suffix %s in '%s'", "um ok", x)
	}
}

func TestZeroValuePrintln(t *testing.T) {
	var l *Logger
	l.Println("um", "ok")
	// expect no panics
}

func BenchmarkOutputWithPrefix(b *testing.B) {
	buf := &bytes.Buffer{}
	l := New(buf, "testing! ", log.LstdFlags)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Output(2, "test")
	}
}

func BenchmarkOutputWithPrefixParallel(b *testing.B) {
	buf := &bytes.Buffer{}
	l := New(buf, "testing! ", log.LstdFlags)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Output(2, "test")
		}
	})
}
