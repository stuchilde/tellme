package cmd

import (
	"fmt"
	"github.com/ip2location/ip2location-go/v9"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

var (
	IP          = ""
	version     = "1.0.0"
	cmdName     = "ip2location"
	shortDesc   = "Convert ip address to location"
	GetIpURL    = "https://ifconfig.me"
	binFilePath = "./files/IP2LOCATION-LITE-DB3.BIN"
)

const (
	Country = "Country"
	Region  = "Region"
	City    = "City"
)

var ip2LocationCmd = &cobra.Command{
	Use:     cmdName,
	Version: version,
	Short:   shortDesc,
	Run: func(cmd *cobra.Command, args []string) {
		ip2Location(IP)
	},
}

func init() {
	ip2LocationCmd.Flags().BoolP(version, "v", false, "version of ip2location")
	ip2LocationCmd.Flags().StringVarP(&IP, "ip", "", GetPublicIP(), "ip address")
}

func ip2Location(ip string) {
	db, err := ip2location.OpenDB(binFilePath)
	if err != nil {
		fmt.Print(err)
		return
	}
	results, err := db.Get_all(ip)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%-15s: %s\n", Country, results.Country_long)
	fmt.Printf("%-15s: %s\n", Region, results.Region)
	fmt.Printf("%-15s: %s\n", City, results.City)
	db.Close()
}

// Execute command
func Execute() {
	if err := ip2LocationCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//GetPublicIP Get the public ip address
func GetPublicIP() string {
	resp, err := http.Get(GetIpURL)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bodyByte)
}
