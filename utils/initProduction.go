package utils

import (
	"encoding/json"
	"os"

	"gitlab.com/arthurnavah/Production-Arthur/PreInscripcionRG-API/databases"
)

func InitProduction() {
	os.Mkdir("databases", 0777)
	os.Mkdir("security", 0777)
	os.Mkdir("security/keys", 0777)

	file, _ := os.Create("databases/configDB.json")
	file2, _ := os.Create("security/keys/private.rsa")
	file3, _ := os.Create("security/keys/public.rsa")

	j := databases.ConfigurationDB{
		Server:   "localhost",
		Port:     "5432",
		User:     "tuUsuario",
		Password: "tuPassword",
		Database: "baseDeDatos",
		SSLMode:  "disable",
	}

	h, _ := json.MarshalIndent(j, "", "\t")

	file.Write(h)

	file2.WriteString("-----BEGIN RSA PRIVATE KEY-----\nMIICXwIBAAKBgQDGZt2nkAy8oyGMe6FkPU4BJkL8NnH+uzWhrjQFRa9wBwuzVAbm\n9858NvzJRTsKSTA8MXeag+tZPQbK5xHniZaf8qy+Gu2VbM5EYnvc/JEBxxaUPJMp\nPbaeIAWwhSxORT1xo+c/1U1BZEdZDmLen11uG2RAfSRPkWsZY1aXXlD9FwIDAQAB\nAoGBAK+eQFfhR5T2pWyvxqVvKowT7TlJjFBaMFgEVmHQVEHKys2a9F0gPzNujQYv\n9Nne/QZbFy671Ohx/4A9V3jnLl6pEnOv1TwpZkvNtE1RX7l1pPFyNAQZ4URtDAdO\n7cW4y+TjiPGemSc8DkizX6GdlfpNmxXu8NJF9GvVpLudhYt5AkEA7qsJhXmkcCcN\nePjvieuu4B2a2sa2rnd2sMz+Bw52xakSaWHNenBhP0092QTVyxiHfZ5ql6WQ/EN8\nq96xF5u25QJBANTPQe8W1jp2o+Pd5oi+5cyRbHnGznARSfFAaxh6HSQgjirVfhzy\nE34gDwxBN8p4QmK0Ix+CRz56D3QIDnEaSEsCQQDNSMxNxE1OMikCbowKw2+NUamK\nl/U1p4etlwTAqQ48AarWfcsxj0v2GMgjzGbf499Wi1X/zechNCMd6dPFDKiJAkEA\niLWAQMyl4LbYgAjMESq5S1pcmjpp4bppiu78HaRM1imHyuKNeqDUfuCvagr3gT0S\nQIBCU/qINA2LodmPFX9x0wJBAIYc91qc0M+taFphD85ZerUtbFIV28slz1tkb2/v\niRVwKrSqyFLF1p678u7g2rdNsTbHVQ2+05guxGeiZ1jFJtY=\n-----END RSA PRIVATE KEY-----")

	file3.WriteString("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDGZt2nkAy8oyGMe6FkPU4BJkL8\nNnH+uzWhrjQFRa9wBwuzVAbm9858NvzJRTsKSTA8MXeag+tZPQbK5xHniZaf8qy+\nGu2VbM5EYnvc/JEBxxaUPJMpPbaeIAWwhSxORT1xo+c/1U1BZEdZDmLen11uG2RA\nfSRPkWsZY1aXXlD9FwIDAQAB\n-----END PUBLIC KEY-----")
}
