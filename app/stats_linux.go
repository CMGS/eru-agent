package app

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/HunanTV/eru-agent/common"
	"github.com/HunanTV/eru-agent/logs"
)

func GetNetStats(pid string, result map[string]uint64) (err error) {
	cmd := exec.Command("nsenter", "-t", pid, "-n", "cat", "/proc/net/dev")

	outr, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	//FIXME ignore stderr

	cmd.Start()
	s := bufio.NewScanner(outr)
	var d uint64
	for s.Scan() {
		var name string
		var n [8]uint64
		text := s.Text()
		if strings.Index(text, ":") < 1 {
			continue
		}
		ts := strings.Split(text, ":")
		fmt.Sscanf(ts[0], "%s", &name)
		if !strings.HasPrefix(name, common.VLAN_PREFIX) {
			continue
		}
		fmt.Sscanf(ts[1],
			"%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
			&n[0], &n[1], &n[2], &n[3], &d, &d, &d, &d,
			&n[4], &n[5], &n[6], &n[7], &d, &d, &d, &d,
		)
		result[name+".inbytes"] = n[0]
		result[name+".inpackets"] = n[1]
		result[name+".inerrs"] = n[2]
		result[name+".indrop"] = n[3]
		result[name+".outbytes"] = n[4]
		result[name+".outpackets"] = n[5]
		result[name+".outerrs"] = n[6]
		result[name+".outdrop"] = n[7]
	}
	err = cmd.Wait()
	logs.Debug("Container net status", result)
	return
}
