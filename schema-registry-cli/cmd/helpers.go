package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/viper"

	schemaregistry "github.com/wildbeavers/schema-registry"
)

func stdinToString() string {
	bs, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	return string(bs)
}

func printSchema(sch schemaregistry.Schema) {
	log.Printf("version: %d\n", sch.Version)
	log.Printf("id: %d\n", sch.ID)

	pretty, err := prettyjson.Format([]byte(sch.Schema))
	if err != nil {
		fmt.Println(sch.Schema) //isn't a json object, which is legal
		return
	}
	os.Stdout.Write(pretty)
	os.Stdout.WriteString("\n")
}

func getByID(id int) error {
	cl := assertClient()
	sch, err := cl.GetSchemaByID(id)
	if err != nil {
		return err
	}
	fmt.Println(sch)
	return nil
}

func getLatestBySubject(subj string) error {
	cl := assertClient()
	sch, err := cl.GetLatestSchema(subj)
	if err != nil {
		return err
	}
	printSchema(sch)
	return nil
}

func getBySubjectVersion(subj string, ver int) error {
	cl := assertClient()
	sch, err := cl.GetSchemaBySubject(subj, ver)
	if err != nil {
		return err
	}
	printSchema(sch)
	return nil
}

func printConfig(cfg schemaregistry.ConfigGet, subj string) {
	if subj == "" {
		subj = "global"
	}
	if cfg.CompatibilityType == "" {
		cfg.CompatibilityType = "not defined, using global"
	}
	fmt.Printf("%s compatibility type: %s\n", subj, cfg.CompatibilityType)
}

func getConfig(subj string) error {
	cl := assertClient()
	cfg, err := cl.GetConfig(subj)
	if err != nil {
		return err
	}
	printConfig(cfg, subj)
	return nil
}

func deleteSubject(subj string) error {
	cl := assertClient()
	_, err := cl.DeleteSubject(subj)
	if err != nil {
		return err
	}
	return nil
}

func assertClient() *schemaregistry.Client {
	var c *schemaregistry.Client
	var err error
	if viper.GetString("basic_auth_user") != "" && viper.GetString("basic_auth_pass") != "" {
		// Use Basic Authentication
		c, err = schemaregistry.NewClientWithBasicAuth(viper.GetString("url"),
			viper.GetString("basic_auth_user"), viper.GetString("basic_auth_pass"))
	} else {
		c, err = schemaregistry.NewClient(viper.GetString("url"))
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return c
}
