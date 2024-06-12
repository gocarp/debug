// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gocarp/encoding/hash"
	"github.com/gocarp/errors"
)

// BinVersion returns the version of current running binary.
// It uses hash.BKDRHash+BASE36 algorithm to calculate the unique version of the binary.
func BinVersion() string {
	if binaryVersion == "" {
		binaryContent, _ := os.ReadFile(selfPath)
		binaryVersion = strconv.FormatInt(
			int64(hash.BKDR(binaryContent)),
			36,
		)
	}
	return binaryVersion
}

// BinVersionMd5 returns the version of current running binary.
// It uses MD5 algorithm to calculate the unique version of the binary.
func BinVersionMd5() string {
	if binaryVersionMd5 == "" {
		binaryVersionMd5, _ = md5File(selfPath)
	}
	return binaryVersionMd5
}

// md5File encrypts file content of `path` using MD5 algorithms.
func md5File(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = errors.Wrapf(err, `os.Open failed for name "%s"`, path)
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = errors.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
