package whatcango

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"context"
	"database/sql"
	"io"
	"net"
	"net/http"
)

func Wrong() {
	// archive/tar
	tar.ErrHeader = nil
	tar.ErrWriteTooLong = nil
	tar.ErrFieldTooLong = nil
	tar.ErrWriteAfterClose = nil

	// archive/zip
	zip.ErrFormat = nil
	zip.ErrAlgorithm = nil
	zip.ErrChecksum = nil

	// bufio
	bufio.ErrInvalidUnreadByte = nil
	bufio.ErrInvalidUnreadRune = nil
	bufio.ErrBufferFull = nil
	bufio.ErrNegativeCount = nil
	bufio.ErrTooLong = nil
	bufio.ErrNegativeAdvance = nil
	bufio.ErrAdvanceTooFar = nil
	bufio.ErrBadReadCount = nil

	// bytes
	bytes.ErrTooLarge = nil

	// compress/gzip
	gzip.ErrChecksum = nil
	gzip.ErrHeader = nil

	// compress/zlib
	zlib.ErrChecksum = nil
	zlib.ErrDictionary = nil
	zlib.ErrHeader = nil

	// context
	context.Canceled = nil
	context.DeadlineExceeded = nil

	// TODO: crypto/*

	// database/sql
	sql.ErrConnDone = nil
	sql.ErrNoRows = nil
	sql.ErrTxDone = nil

	// TODO: database/sql/driver
	// TODO: debug/*
	// TODO: encoding/*

	io.EOF = nil
	io.ErrUnexpectedEOF = nil
	io.ErrShortWrite = nil

	// net
	net.ErrClosed = nil
	net.ErrWriteToConnected = nil

	http.ErrBodyNotAllowed = nil

	// TODO: everything else
}
