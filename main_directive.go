package nginx

import "fmt"

// 判断是否是env指令
func isMainEnvDirective(directive string) bool {
	if isEqualString(directive, MainEnvDirective) {
		return true
	}
	return false
}

// 判断是否是include指令
func isMainIncludeDirective(directive string) bool {
	if isEqualString(directive, MainIncludeDirective) {
		return true
	}
	return false
}

// 判断是否是lock_file指令
func isMainLockFileDirective(directive string) bool {
	if isEqualString(directive, MainLockFileDirective) {
		return true
	}
	return false
}

// 判断是否是pcre_jit指令
func isMainPcreJITDirective(directive string) bool {
	if isEqualString(directive, MainPcreJITDirective) {
		return true
	}
	return false
}

// 判断是否是thread_pool指令
func isMainThreadPoolDirective(directive string) bool {
	if isEqualString(directive, MainThreadPoolDirective) {
		return true
	}
	return false
}

// 判断是否是user指令
func isMainUserDirective(directive string) bool {
	if isEqualString(directive, MainUserDirective) {
		return true
	}
	return false
}

// 判断是否是worker_rlimit_nofile指令
func isMainWorkerRlimitNofileDirective(directive string) bool {
	if isEqualString(directive, MainWorkerRlimitNofileDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_bind指令
func isMainProxyBindDirective(directive string) bool {
	if isEqualString(directive, MainProxyBindDirective) {
		return true
	}
	return false
}

// 判断是否是daemon指令
func isMainDaemonDirective(directive string) bool {
	if isEqualString(directive, MainDaemonDirective) {
		return true
	}
	return false
}

// 判断是否是events指令
func isMainEventsDirective(directive string) bool {
	if isEqualString(directive, MainEventsDirective) {
		return true
	}
	return false
}

// 判断是否是load_module指令
func isMainLoadModuleDirective(directive string) bool {
	if isEqualString(directive, MainLoadModuleDirective) {
		return true
	}
	return false
}

// 判断是否是worker_shutdown_timeout指令
func isMainWorkerShutdownTimeoutDirective(directive string) bool {
	if isEqualString(directive, MainWorkerShutdownTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是main指令
func isMainDirective(directive string) bool {
	if isEqualString(directive, MainDirective) {
		return true
	}
	return false
}

// 判断是否是timer_resolution指令
func isMainTimerResolutionDirective(directive string) bool {
	if isEqualString(directive, MainTimerResolutionDirective) {
		return true
	}
	return false
}

// 判断是否是worker_cpu_affinity指令
func isMainWorkerCPUAffinityDirective(directive string) bool {
	if isEqualString(directive, MainWorkerCPUAffinityDirective) {
		return true
	}
	return false
}

// 判断是否是worker_priority指令
func isMainWorkerPriorityDirective(directive string) bool {
	if isEqualString(directive, MainWorkerPriorityDirective) {
		return true
	}
	return false
}

// 判断是否是worker_processes指令
func isMainWorkerProcessesDirective(directive string) bool {
	if isEqualString(directive, MainWorkerProcessesDirective) {
		return true
	}
	return false
}

// 判断是否是worker_rlimit_core指令
func isMainWorkerRlimitCoreDirective(directive string) bool {
	if isEqualString(directive, MainWorkerRlimitCoreDirective) {
		return true
	}
	return false
}

// 判断是否是working_directory指令
func isMainWorkingDirectoryDirective(directive string) bool {
	if isEqualString(directive, MainWorkingDirectoryDirective) {
		return true
	}
	return false
}

// 判断是否是error_page指令
func isMainErrorPageDirective(directive string) bool {
	if isEqualString(directive, MainErrorPageDirective) {
		return true
	}
	return false
}

// 判断是否是master_process指令
func isMainMasterProcessDirective(directive string) bool {
	if isEqualString(directive, MainMasterProcessDirective) {
		return true
	}
	return false
}

// 判断是否是pid指令
func isMainPidDirective(directive string) bool {
	if isEqualString(directive, MainPidDirective) {
		return true
	}
	return false
}

// 判断是否是ssl_engine指令
func isMainSslEngineDirective(directive string) bool {
	if isEqualString(directive, MainSslEngineDirective) {
		return true
	}
	return false
}

// 判断是否是http指令
func isMainHTTPDirective(directive string) bool {
	if isEqualString(directive, MainHTTPDirective) {
		return true
	}
	return false
}

// 处理Main指令
func ProcessMain(config *Config) (*Main, error) {
	main := NewMain()
	for _, parsed := range config.Parsed {
		fmt.Printf("Directive: %s\n", parsed.Directive)
		// 处理worker_cpu_affinity指令
		if isMainWorkerCPUAffinityDirective(parsed.Directive) {
			main.WorkerCpuAffinity = processArgsArray(parsed.Args)
			continue
		}

		// 处理worker_rlimit_core指令
		if isMainWorkerRlimitCoreDirective(parsed.Directive) {
			main.WorkerRlimitCore = processArgsArray(parsed.Args)
			continue
		}

		// 处理daemon指令
		if isMainDaemonDirective(parsed.Directive) {
			main.Daemon = processArgsArray(parsed.Args)
			continue
		}

		// 处理lock_file指令
		if isMainLockFileDirective(parsed.Directive) {
			main.LockFile = processArgsArray(parsed.Args)
			continue
		}

		// 处理master_process指令
		if isMainMasterProcessDirective(parsed.Directive) {
			main.MasterProcess = processArgsArray(parsed.Args)
			continue
		}

		// 处理pcre_jit指令
		if isMainPcreJITDirective(parsed.Directive) {
			main.PcreJIT = processArgsArray(parsed.Args)
			continue
		}

		// 处理ssl_engine指令
		if isMainSslEngineDirective(parsed.Directive) {
			main.SslEngine = processArgsArray(parsed.Args)
			continue
		}

		// 处理working_directory指令
		if isMainWorkingDirectoryDirective(parsed.Directive) {
			main.WorkingDirectory = processArgsArray(parsed.Args)
			continue
		}

		// 处理pid指令
		if isMainPidDirective(parsed.Directive) {
			main.Pid = processArgsArray(parsed.Args)
			continue
		}

		// 处理thread_pool指令
		if isMainThreadPoolDirective(parsed.Directive) {
			main.ThreadPool = processArgsArray(parsed.Args)
			continue
		}

		// 处理worker_shutdown_timeout指令
		if isMainWorkerShutdownTimeoutDirective(parsed.Directive) {
			main.WorkerShutdownTimeout = processArgsArray(parsed.Args)
			continue
		}

		// 处理proxy_bind指令
		if isMainProxyBindDirective(parsed.Directive) {
			main.ProxyBind = processArgsArray(parsed.Args)
			continue
		}

		// 处理timer_resolution指令
		if isMainTimerResolutionDirective(parsed.Directive) {
			main.TimerResolution = processArgsArray(parsed.Args)
			continue
		}

		// 处理user指令
		if isMainUserDirective(parsed.Directive) {
			main.User = processArgsArray(parsed.Args)
			continue
		}

		// 处理worker_priority指令
		if isMainWorkerPriorityDirective(parsed.Directive) {
			main.WorkerPriority = processArgsArray(parsed.Args)
			continue
		}

		// 处理worker_rlimit_nofile指令
		if isMainWorkerRlimitNofileDirective(parsed.Directive) {
			main.WorkerRlimitNofile = processArgsArray(parsed.Args)
			continue
		}

		// 处理error_page指令
		if isMainErrorPageDirective(parsed.Directive) {
			main.ErrorPage = processArgsArray(parsed.Args)
			continue
		}

		// 处理events指令
		if isMainEventsDirective(parsed.Directive) {
			event, err := ProcessEvent(&parsed)
			if err != nil {
				return main, err
			}
			main.Events = *event
			continue
		}

		// 处理load_module指令
		if isMainLoadModuleDirective(parsed.Directive) {
			main.LoadModule = processArgsArray(parsed.Args)
			continue
		}

		// 处理worker_processes指令
		if isMainWorkerProcessesDirective(parsed.Directive) {
			main.WorkerProcesses = processArgsArray(parsed.Args)
			continue
		}

		// 处理env指令
		if isMainEnvDirective(parsed.Directive) {
			main.Env = append(main.Env, processArgsArray(parsed.Args))
		}

		// 处理include指令
		if isMainIncludeDirective(parsed.Directive) {
			main.Include = append(main.Include, processArgsArray(parsed.Args))
		}

		// 处理http指令
		if isMainHTTPDirective(parsed.Directive) {
			fmt.Println("===================In http==============")
			http, err := ProcessHttp(&parsed)
			if err != nil {
				return main, err
			}
			main.Http = append(main.Http, *http)
		}
	}

	return main, nil
}
