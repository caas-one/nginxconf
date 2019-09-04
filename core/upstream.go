package core

import (
	"errors"
	"strings"
)

// determine whether it is a upstream directive
func isUpstreamDirective(directive string) bool {
	if isEqualString(directive, UpstreamDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_timeout directive
func isUpstreamKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, UpstreamKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// upstream name
func processUpstreamName(args []string) (string, error) {
	if len(args) == 1 {
		return args[0], nil
	}
	return "", errors.New("Bad number of args in parsed struct within upstream module")
}

// determine whether it is a load blance directive
func isUpstreamLBMethodDirecitive(block InnerBlock) bool {
	if isEqualString(block.Directive, IPHash) || isEqualString(block.Directive, LeastConn) {
		return true
	}
	return false
}

// process upstream load blance
//  load blance include  ip_hash, least_conn，rr is default
func processUpstreamLBMethod(block InnerBlock) (string, error) {
	if !isEqualString(block.Directive, "") {
		return block.Directive, nil
	}
	return "", errors.New("Bad directive of LBMethod within upstream module")
}

// determine whether it is a server directive
func isUpstreamServerDirective(block InnerBlock) bool {
	if isEqualString(block.Directive, UpstreamServerDirective) {
		return true
	}
	return false
}

// determine whether it is a Server directive
// there are three types of address: 10.10.12.45:80, app.example.com:80, unix:/tmp/backend3
func isUpstreamServerAddress(arg string) bool {
	if strings.Contains(arg, ":") {
		ss := strings.Split(arg, ":")
		if len(ss) == 2 {
			if isValidHost(ss[0]) && isValidPort(ss[1]) {
				return true
			}
		}
	}

	if isValidHost(arg) {
		return true
	}

	if isUnixProtocol(arg) {
		return true
	}
	return false
}

// determine whether it is a upstream server parameters
func processUpstreamServerParameters(parameters *UpstreamServerParameters, arg string) error {
	ss := strings.Split(arg, "=")
	// eg.: "weight=99", "max_fails=2"
	if len(ss) == 2 {
		// weight
		if isEqualString(ss[0], UpstreamServerWeightDirective) {
			parameters.Weight = ss[1]
			return nil
		}

		// max_fails
		if isEqualString(ss[0], UpstreamServerMaxFailsDirective) {
			parameters.MaxFails = ss[1]
			return nil
		}

		// fail_timeout
		if isEqualString(ss[0], UpstreamServerFailTimeoutDirective) {
			parameters.FailTimeout = ss[1]
			return nil
		}

		// max_conns
		if isEqualString(ss[0], UpstreamServerMaxConnsDirective) {
			parameters.MaxConns = ss[1]
			return nil
		}
	}

	// down, backup
	param := strings.Trim(arg, " \n\t")
	if isEqualString(param, UpstreamServerDown) {
		parameters.Down = UpstreamServerDown
		return nil
	}

	if isEqualString(param, UpstreamServerBackup) {
		parameters.Backup = UpstreamServerBackup
		return nil
	}

	return errors.New("No args match the server parameters within upstream module")
}

// process upstream module server directive
func processUpstreamServerDirective(block InnerBlock) (*UpstreamServer, error) {
	if !isUpstreamServerDirective(block) {
		return nil, errors.New("Bad directive of server within upstream module")
	}

	server := &UpstreamServer{}
	parameters := &UpstreamServerParameters{}

	for _, arg := range block.Args {
		if isUpstreamServerAddress(arg) {
			server.Address = arg
		}

		processUpstreamServerParameters(parameters, arg)
	}

	server.Params = *parameters
	return server, nil
}

// determine whether it is a keepalive directive
func isUpstreamKeepaliveDirective(block InnerBlock) bool {
	if isEqualString(block.Directive, UpstreamKeepaliveDirective) {
		return true
	}
	return false
}

// process keepalive directive
func processKeepaliveValue(args []string) (string, error) {
	if len(args) == 1 {
		return args[0], nil
	}
	return "", errors.New("Bad number of args in keepalive part within upstream module")
}

func processUpstreamKeepaliveDirective(block InnerBlock) (*UpstreamKeepalive, error) {
	if !isUpstreamKeepaliveDirective(block) {
		return nil, errors.New("Bad directive of keepalive within upstream module")
	}

	keepalive := &UpstreamKeepalive{}
	time, err := processKeepaliveValue(block.Args)
	if err != nil {
		return nil, err
	}

	keepalive.Time = time
	return keepalive, nil
}

// ProcessUpstream process upstream  module
func ProcessUpstream(block *Block) (*Upstream, error) {
	if !isUpstreamDirective(block.Directive) {
		return nil, errors.New("Not upstream directive")
	}

	// process upsteam name
	upstream := NewUpstream()
	name, err := processUpstreamName(block.Args)
	if err != nil {
		return nil, err
	}

	upstream.Name = name

	for _, innerBlock := range block.InnerBlocks {

		if isUpstreamLBMethodDirecitive(innerBlock) {
			method, err := processUpstreamLBMethod(innerBlock)
			if err != nil {
				return upstream, err
			}
			upstream.Method = method
		}

		if isUpstreamKeepaliveTimeoutDirective(innerBlock.Directive) {
			upstream.KeepaliveTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		// process keepalive directive
		if isUpstreamKeepaliveDirective(innerBlock) {
			keepalive, err := processUpstreamKeepaliveDirective(innerBlock)
			if err != nil {
				return upstream, err
			}
			upstream.UpstreamKeepalive = *keepalive
		}

		if isUpstreamServerDirective(innerBlock) {
			server, err := processUpstreamServerDirective(innerBlock)
			if err != nil {
				return upstream, err
			}
			upstream.Servers = append(upstream.Servers, *server)
		}
	}

	return upstream, nil
}
