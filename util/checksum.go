package util

import (
  "crypto/sha256"
  "io"
  "os"
)

func Checksum(path string) ([]byte, error) {
  var result []byte
  file, err := os.Open(path)

  if err != nil {
    return result, err
  }

  defer file.Close()

  hash := sha256.New()

  if _, err := io.Copy(hash, file); err != nil {
    return result, err
  }

  return hash.Sum(result), nil
}
