package utility

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

/*************************************************
 *	specification;
 *	name 			= HashStr
 *	Function 	= generate Hash used specific hash func
 *	note			= umd5, sha1, sha256, sha512
 *						= default is sha256
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00/V1.10
 *  input 		= trg string  string value
 * 						= alg string: Hash function
 *  output    = hashed string: string
 *  end of specification;
**************************************************/
func HashStr(trg, alg string) string {
	hashed := ""
	b := []byte(trg)

	switch alg {
	case "md5":
		md5 := md5.Sum(b)
		hashed = hex.EncodeToString(md5[:])
	case "sha1":
		sha1 := sha1.Sum(b)
		hashed = hex.EncodeToString(sha1[:])
	case "sha512":
		sha512 := sha512.Sum512(b)
		hashed = hex.EncodeToString(sha512[:])
	default:
		sha256 := sha256.Sum256(b)
		hashed = hex.EncodeToString(sha256[:])
	}

	return hashed
}
