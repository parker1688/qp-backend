package common

import (
	"bootpkg/common/tool"
	"testing"
)

func TestRedisEncryptionPassword(t *testing.T) {
	Password := ""
	t.Log(tool.EncryptAESPrefixRandKeySalt(Password, "Password"))
}

func TestRedisDPassword(t *testing.T) {
	Password := ""
	t.Log(tool.DecryptAESPrefixRandKeySalt(Password, "Password"))
}

func TestMysqlEncryptionDSN(t *testing.T) {
	DSN := "root:y.NAq.NCBZa#*7eZ}K2Y@(mysql1.svc.cluster.local:3306)/zbgame?charset=utf8&parseTime=true&loc=Local"
	t.Log(tool.EncryptAESPrefixRandKeySalt(DSN, "DSN"))
}

func TestMysqlDncryptionDSN(t *testing.T) {
	DSN := "NTETFHNMQVEQAJOCPE48aUl0UNPZGnjCWfVdwwBLUjuBeCBf4jkapGUXAmEeNz4C0ZwZJwRcJkeRDBIkFo6er/MPILn1OELEvBB+w94bYmcyeLEBb48hDsQeLOyUjUlgJELSjyjirsKnBcMe0FSF4vzLI283dqE9EqoBZw=="
	t.Log(tool.DecryptAESPrefixRandKeySalt(DSN, "DSN"))
}
