package core

import "fmt"

// determine whether it is a env directive
func isGlobalEnvDirective(directive string) bool {
	if isEqualString(directive, GlobalEnvDirective) {
		return true
	}
	return false
}

// determine whether it is a include directive
func isGlobalIncludeDirective(directive string) bool {
	if isEqualString(directive, GlobalIncludeDirective) {
		return true
	}
	return false
}

// determine whether it is a lock_file directive
func isGlobalLockFileDirective(directive string) bool {
	if isEqualString(directive, GlobalLockFileDirective) {
		return true
	}
	return false
}

// determine whether it is a pcre_jit directive
func isGlobalPcreJITDirective(directive string) bool {
	if isEqualString(directive, GlobalPcreJITDirective) {
		return true
	}
	return false
}

// determine whether it is a thread_pool directive
func isGlobalThreadPoolDirective(directive string) bool {
	if isEqualString(directive, GlobalThreadPoolDirective) {
		return true
	}
	return false
}

// determine whether it is a user directive
func isGlobalUserDirective(directive string) bool {
	if isEqualString(directive, GlobalUserDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_rlimit_nofile directive
func isGlobalWorkerRlimitNofileDirective(directive string) bool {
	if isEqualString(directive, GlobalWorkerRlimitNofileDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_bind directive
func isGlobalProxyBindDirective(directive string) bool {
	if isEqualString(directive, GlobalProxyBindDirective) {
		return true
	}
	return false
}

// determine whether it is a daemon directive
func isGlobalDaemonDirective(directive string) bool {
	if isEqualString(directive, GlobalDaemonDirective) {
		return true
	}
	return false
}

// determine whether it is a events directive
func isGlobalEventsDirective(directive string) bool {
	if isEqualString(directive, GlobalEventsDirective) {
		return true
	}
	return false
}

// determine whether it is a load_module directive
func isGlobalLoadModuleDirective(directive string) bool {
	if isEqualString(directive, GlobalLoadModuleDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_shutdown_timeout directive
func isGlobalWorkerShutdownTimeoutDirective(directive string) bool {
	if isEqualString(directive, GlobalWorkerShutdownTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a Global directive
func isGlobalDirective(directive string) bool {
	if isEqualString(directive, GlobalDirective) {
		return true
	}
	return false
}

// determine whether it is a timer_resolution directive
func isGlobalTimerResolutionDirective(directive string) bool {
	if isEqualString(directive, GlobalTimerResolutionDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_cpu_affinity directive
func isGlobalWorkerCPUAffinityDirective(directive string) bool {
	if isEqualString(directive, GlobalWorkerCPUAffinityDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_priority directive
func isGlobalWorkerPriorityDirective(directive string) bool {
	if isEqualString(directive, GlobalWorkerPriorityDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_processes directive
func isGlobalWorkerProcessesDirective(directive string) bool {
	if isEqualString(directive, GlobalWorkerProcessesDirective) {
		return true
	}
	return false
}

// determine whether it is a worker_rlimit_core directive
func isGlobalWorkerRlimitCoreDirective(directive string) bool {
	if isEqualString(directive, GlobalWorkerRlimitCoreDirective) {
		return true
	}
	return false
}

// determine whether it is a working_directory directive
func isGlobalWorkingDirectoryDirective(directive string) bool {
	if isEqualString(directive, GlobalWorkingDirectoryDirective) {
		return true
	}
	return false
}

// determine whether it is a error_page directive
func isGlobalErrorPageDirective(directive string) bool {
	if isEqualString(directive, GlobalErrorPageDirective) {
		return true
	}
	return false
}

// determine whether it is a master_process directive
func isGlobalMasterProcessDirective(directive string) bool {
	if isEqualString(directive, GlobalMasterProcessDirective) {
		return true
	}
	return false
}

// determine whether it is a pid directive
func isGlobalPidDirective(directive string) bool {
	if isEqualString(directive, GlobalPidDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl_engine directive
func isGlobalSslEngineDirective(directive string) bool {
	if isEqualString(directive, GlobalSslEngineDirective) {
		return true
	}
	return false
}

// determine whether it is a http directive
func isGlobalHTTPDirective(directive string) bool {
	if isEqualString(directive, GlobalHTTPDirective) {
		return true
	}
	return false
}

// ProcessGlobal process Global directive
func ProcessGlobal(config *Config) (*Global, error) {
	global := NewGlobal()
	for _, parsed := range config.Parsed {
		fmt.Printf("Directive: %s\n", parsed.Directive)
		// process worker_cpu_affinity directive
		if isGlobalWorkerCPUAffinityDirective(parsed.Directive) {
			global.WorkerCpuAffinity = processArgsArray(parsed.Args)
			continue
		}

		// process worker_rlimit_core directive
		if isGlobalWorkerRlimitCoreDirective(parsed.Directive) {
			global.WorkerRlimitCore = processArgsArray(parsed.Args)
			continue
		}

		// process daemon directive
		if isGlobalDaemonDirective(parsed.Directive) {
			global.Daemon = processArgsArray(parsed.Args)
			continue
		}

		// process lock_file directive
		if isGlobalLockFileDirective(parsed.Directive) {
			global.LockFile = processArgsArray(parsed.Args)
			continue
		}

		// process master_process directive
		if isGlobalMasterProcessDirective(parsed.Directive) {
			global.MasterProcess = processArgsArray(parsed.Args)
			continue
		}

		// process pcre_jit directive
		if isGlobalPcreJITDirective(parsed.Directive) {
			global.PcreJIT = processArgsArray(parsed.Args)
			continue
		}

		// process ssl_engine directive
		if isGlobalSslEngineDirective(parsed.Directive) {
			global.SslEngine = processArgsArray(parsed.Args)
			continue
		}

		// process working_directory directive
		if isGlobalWorkingDirectoryDirective(parsed.Directive) {
			global.WorkingDirectory = processArgsArray(parsed.Args)
			continue
		}

		// process pid directive
		if isGlobalPidDirective(parsed.Directive) {
			global.Pid = processArgsArray(parsed.Args)
			continue
		}

		// process thread_pool directive
		if isGlobalThreadPoolDirective(parsed.Directive) {
			global.ThreadPool = processArgsArray(parsed.Args)
			continue
		}

		// process worker_shutdown_timeout directive
		if isGlobalWorkerShutdownTimeoutDirective(parsed.Directive) {
			global.WorkerShutdownTimeout = processArgsArray(parsed.Args)
			continue
		}

		// process proxy_bind directive
		if isGlobalProxyBindDirective(parsed.Directive) {
			global.ProxyBind = processArgsArray(parsed.Args)
			continue
		}

		// process timer_resolution directive
		if isGlobalTimerResolutionDirective(parsed.Directive) {
			global.TimerResolution = processArgsArray(parsed.Args)
			continue
		}

		// process user directive
		if isGlobalUserDirective(parsed.Directive) {
			global.User = processArgsArray(parsed.Args)
			continue
		}

		// process worker_priority directive
		if isGlobalWorkerPriorityDirective(parsed.Directive) {
			global.WorkerPriority = processArgsArray(parsed.Args)
			continue
		}

		// process worker_rlimit_nofile directive
		if isGlobalWorkerRlimitNofileDirective(parsed.Directive) {
			global.WorkerRlimitNofile = processArgsArray(parsed.Args)
			continue
		}

		// process error_page directive
		if isGlobalErrorPageDirective(parsed.Directive) {
			global.ErrorPage = processArgsArray(parsed.Args)
			continue
		}

		// process events directive
		if isGlobalEventsDirective(parsed.Directive) {
			event, err := ProcessEvent(&parsed)
			if err != nil {
				return global, err
			}
			global.Events = *event
			continue
		}

		// process load_module directive
		if isGlobalLoadModuleDirective(parsed.Directive) {
			global.LoadModule = processArgsArray(parsed.Args)
			continue
		}

		// process worker_processes directive
		if isGlobalWorkerProcessesDirective(parsed.Directive) {
			global.WorkerProcesses = processArgsArray(parsed.Args)
			continue
		}

		// process env directive
		if isGlobalEnvDirective(parsed.Directive) {
			global.Env = append(global.Env, processArgsArray(parsed.Args))
		}

		// process include directive
		if isGlobalIncludeDirective(parsed.Directive) {
			global.Include = append(global.Include, processArgsArray(parsed.Args))
		}

		// process http directive
		if isGlobalHTTPDirective(parsed.Directive) {
			fmt.Println("===================In http==============")
			http, err := ProcessHTTP(&parsed)
			if err != nil {
				return global, err
			}
			global.Http = append(global.Http, *http)
		}
	}

	return global, nil
}
