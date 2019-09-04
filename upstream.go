package nginx

import (
	"errors"
	"strings"
)

// 判断是否是upstream的指令
func isUpstreamDirective(directive string) bool {
	if isEqualString(directive, UpstreamDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_timeout指令
func isUpstreamKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, UpstreamKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// 获取upstream的名字
// upstream模块中，upstream关键字后面的字符串是upsteram的名字
// 在parsed结构体下面第一级的args中，理论上args数组只含有一个元素
func processUpstreamName(args []string) (string, error) {
	if len(args) == 1 {
		return args[0], nil
	}
	return "", errors.New("Bad number of args in parsed struct within upstream module")
}

// 判断是否是负载均衡指令
func isUpstreamLBMethodDirecitive(block InnerBlock) bool {
	if isEqualString(block.Directive, IPHash) || isEqualString(block.Directive, LeastConn) {
		return true
	}
	return false
}

// 处理upstream的负载均衡策略
// 负载均衡策略包括ip_hash和least_conn两种，rr轮询的方式不需要显式指定
func processUpstreamLBMethod(block InnerBlock) (string, error) {
	if !isEqualString(block.Directive, "") {
		return block.Directive, nil
	}
	return "", errors.New("Bad directive of LBMethod within upstream module")
}

// 判断是否是server指令
func isUpstreamServerDirective(block InnerBlock) bool {
	if isEqualString(block.Directive, UpstreamServerDirective) {
		return true
	}
	return false
}

// 判断是否是Server指令中的地址
// 地址的组成大概分为三种：10.10.12.45:80, app.example.com:80, unix:/tmp/backend3
func isUpstreamServerAddress(arg string) bool {
	// 是否包括":"字符
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

// 判断是否是upstream中server字段的参数
// 并给相应的参数赋值
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

// 处理upstream中的server指令
// server指令主要判断args里面的各个部分，并给UpstreamServerParameters结构体赋值
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

		// 处理其他参数: weight, max_conns等
		processUpstreamServerParameters(parameters, arg)
	}

	server.Params = *parameters
	return server, nil
}

// 判断是否是keepalive指令
// 在keepalive指令部分，args数组包含一个值
func isUpstreamKeepaliveDirective(block InnerBlock) bool {
	if isEqualString(block.Directive, UpstreamKeepaliveDirective) {
		return true
	}
	return false
}

// 处理keepalive指令的参数，提取里面的值
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

func ProcessUpstream(block *Block) (*Upstream, error) {
	if !isUpstreamDirective(block.Directive) {
		return nil, errors.New("Not upstream directive")
	}

	// 处理upsteam的name
	upstream := NewUpstream()
	name, err := processUpstreamName(block.Args)
	if err != nil {
		return nil, err
	}

	upstream.Name = name

	// 处理Parsed.Blocks，里面包括upstream中的server, lbmethod等
	for _, innerBlock := range block.InnerBlocks {

		// 处理负载均衡策略
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

		// 处理keepalive指令
		if isUpstreamKeepaliveDirective(innerBlock) {
			keepalive, err := processUpstreamKeepaliveDirective(innerBlock)
			if err != nil {
				return upstream, err
			}
			upstream.UpstreamKeepalive = *keepalive
		}

		// 处理Server部分，里面包含了UpstreamName的解析
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
