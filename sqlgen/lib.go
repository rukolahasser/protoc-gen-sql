// Use protoc-gen-sql as a library
package sqlgen

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// ExecSQLGenerator must be run after installing protoc-gen-sql plugin
func ExecSQLGenerator(outDir, binDir, fileName, packageName, owner, contractName, version, protoPath, protoFile string) {
	cmd := exec.Command(
		"protoc",
		fmt.Sprintf("--sql_out=%s", outDir),
		fmt.Sprintf("--plugin=%s/protoc-gen-sql", binDir),
		fmt.Sprintf("--sql_opt=paths=source_relative,file=%s,contract=%s,version=%s,package=%s,owner=%s", fileName, contractName, version, packageName, owner),
		fmt.Sprintf("--proto_path=%s", protoPath),
		protoFile,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed generating sql file: %v", stderr.String())
	}
}
