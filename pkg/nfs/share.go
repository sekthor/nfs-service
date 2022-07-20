package nfs

import (
	"errors"
	"strings"
)

type Share struct {
	Directory string `json:"name"`
	Hosts     []Host `json:"hosts"`
}

type Host struct {
	Host    string   `json:"host"`
	Options []string `json:"options"`
}

func Temp() string {
	serialized, err := serializeShares(Shares)
	if err != nil {

	}
	return serialized
}

func serializeShares(s []Share) (string, error) {
	var shares = "# Do not change! File mangaged by nfs-service\n"
	for _, share := range s {
		serializedShare, err := serializeShare(&share)
		if err != nil {

		}
		shares = shares + "\n" + serializedShare
	}

	return shares, nil
}

func serializeShare(s *Share) (string, error) {
	var share = s.Directory
	for _, host := range s.Hosts {
		share = share + " " + host.Host + "(" + strings.Join(host.Options, ",") + ")"
	}
	return share, nil
}

/*
func DeserializeConfig(f *os.File) {
	scanner := bufio.NewScanner(f) // f is the *os.File
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		if
	}
	if err := scanner.Err(); err != nil {
	   // handle error
	}
}
*/

func DeserializeConfigLine(s *string) (Share, error) {
	var share Share
	slices := strings.Split(*s, " ")
	if len(slices) < 2 {
		// error: wrong format. no Host or something?
		return share, errors.New("Whoops")
	}
	share.Directory = slices[0]
	share.Hosts = []Host{}
	for _, host := range slices[1:] {
		newHost := Host{}
		hostElements := strings.Split(host, "(")
		options := strings.Trim(hostElements[1], ")")
		newHost.Host = hostElements[0]
		newHost.Options = strings.Split(options, ",")
		share.Hosts = append(share.Hosts, newHost)
	}
	return share, nil
}
