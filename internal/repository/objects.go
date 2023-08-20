package repository

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"os"
	"strconv"
)

// # NOTE: hash-object : converts and existing file into a git object
// # NOTE: cat-file: prints an existing git object to the standard output.

type GitObject struct {
	data []byte
}

type GitBlob struct {
	fmtStr string
}

func NewGitBlob() *GitBlob {
	return &GitBlob{
		fmtStr: "blob",
	}
}

func (b *GitBlob) Serialize() []byte {
	return nil
}

func (b *GitBlob) Deserialize(data []byte) {

}

func InitGitObject(data []byte) {
	object := &GitObject{}

	if data != nil {
		object.Deserialize(data)
	} else {
		object.Init()
	}
}

func (object *GitObject) Serialize(repo *GitRepository) ([]byte, error) {

	return nil, fmt.Errorf("Unimplemented!")
}

func (object *GitObject) Deserialize(data []byte) error {

	return fmt.Errorf("Unimplemented")
}

func (object *GitObject) Init() {

}

func ObjectRead(repo *GitRepository, sha string) (*GitObject, error) {
	path, err := RepoFile(repo, false, "objects", sha[0:2], sha[2:])

	if path == "" {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decompress, err := zlib.NewReader(file)

	if err != nil {
		return nil, err
	}

	defer decompress.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(decompress)

	if err != nil {
		return nil, err
	}

	raw := buf.Bytes()
	x := bytes.IndexByte(raw, ' ')
	fmtStr := string(raw[:x])

	y := bytes.IndexByte(raw[x:], 0)
	sizeStr := string(raw[x+1 : x+y])
	size, err := strconv.Atoi(sizeStr)

	if err != nil {
		return nil, err
	}

	if size != len(raw)-y-1 {
		return nil, fmt.Errorf("Malformed object %s: bad length", sha)
	}

	var obj GitObject

	switch fmtStr {
	case "commit":
		obj = &GitCommit{}
	case "tree":
		obj = &GitTree{}
	case "tag":
		obj = &GitTag{}
	case "blob":
		obj = &GitBlob{}
	default:
		return nil, fmt.Errorf("Unkown type %s for object %s", fmtStr, sha)
	}

	err = obj.Deserialize(raw[y+1:])

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func ObjectWrite(obj *GitObject, repo *GitRepository) (string, error) {
	data, err := obj.Serialize(repo)

	if err != nil {
		return "", err
	}

	header := []byte(fmt.Sprintf("%s %d\x00", obj.fmtStr, len(data)))
	result := append(header, data...)

	sha := fmt.Sprintf("%x", sha1.Sum(result))
	if repo != nil {
		path, err := RepoFile(repo, true, "objects", sha[0:2], sha[2:])

		if err != nil {
			return "", err
		}

		if _, err := os.Stat(path); os.IsNotExist(err) {

			file, err := os.Create(path)

			if err != nil {
				return "", err
			}

			defer file.Close()

			zlibWriter := zlib.NewWriter(file)
			defer zlibWriter.Close()

			_, err = zlibWriter.Write(result)
			if err != nil {
				return "", err
			}
		}
	}

	return sha, nil
}

func objectFind(repo *GitRepository, name string, fmt string, follow bool) string {
	return name
}
