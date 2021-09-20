package leveldb

import "github.com/zoroqi/transcribe-code/goleveldb/leveldb/opt"

type DB struct {

}

func OpenFile(path string, o *opt.Options)(db *DB,err error) {
	//stor, err := storage.OpenFile(path,nil)
	//if err != nil {
	//	return
	//}
	//
	//db,err = Open(stor,o)
	//if err != nil {
	//	stor.Close()
	//} else {
	//	db.closer=stor
	//}
	//return
}

