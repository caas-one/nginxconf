package nginx

import (
	"errors"
	"strings"
)

// 判断是否是proxy_ssl_password_file指令
func isServerProxySslPasswordFileDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslPasswordFileDirective) {
		return true
	}
	return false
}

// 判断是否是read_ahead指令
func isServerReadAheadDirective(directive string) bool {
	if isEqualString(directive, ServerReadAheadDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffering指令
func isServerProxyBufferingDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBufferingDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_connect_timeout指令
func isServerProxyConnectTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyConnectTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ignore_client_abort指令
func isServerProxyIgnoreClientAbortDirective(directive string) bool {
	if isEqualString(directive, ServerProxyIgnoreClientAbortDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_disable指令
func isServerKeepaliveDisableDirective(directive string) bool {
	if isEqualString(directive, ServerKeepaliveDisableDirective) {
		return true
	}
	return false
}

// 判断是否是send_timeout指令
func isServerSendTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerSendTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream_timeout指令
func isServerProxyNextUpstreamTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNextUpstreamTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_temp_file_write_size指令
func isServerProxyTempFileWriteSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyTempFileWriteSizeDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_timeout指令
func isServerKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_close指令
func isServerLingeringCloseDirective(directive string) bool {
	if isEqualString(directive, ServerLingeringCloseDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_use_stale指令
func isServerProxyCacheUseStaleDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheUseStaleDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_read_timeout指令
func isServerProxyReadTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyReadTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_bypass指令
func isServerProxyCacheBypassDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheBypassDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_certificate指令
func isServerProxySslCertificateDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCertificateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_certificate_key指令
func isServerProxySslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCertificateKeyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_protocols指令
func isServerProxySslProtocolsDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslProtocolsDirective) {
		return true
	}
	return false
}

// 判断是否是rewrite指令
func isServerRewriteDirective(directive string) bool {
	if isEqualString(directive, ServerRewriteDirective) {
		return true
	}
	return false
}

// 判断是否是keepalive_requests指令
func isServerKeepaliveRequestsDirective(directive string) bool {
	if isEqualString(directive, ServerKeepaliveRequestsDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_buffer_size指令
func isServerClientBodyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffer_size指令
func isServerProxyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_trusted_certificate指令
func isServerProxySslTrustedCertificateDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslTrustedCertificateDirective) {
		return true
	}
	return false
}

// 判断是否是server_name指令
func isServerNameDirective(directive string) bool {
	if isEqualString(directive, ServerNameDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cookie_path指令
func isServerProxyCookiePathDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCookiePathDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_http_version指令
func isServerProxyHttpVersionDirective(directive string) bool {
	if isEqualString(directive, ServerProxyHTTPVersionDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_ciphers指令
func isServerProxySslCiphersDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCiphersDirective) {
		return true
	}
	return false
}

// 判断是否是msie_padding指令
func isServerMsiePaddingDirective(directive string) bool {
	if isEqualString(directive, ServerMsiePaddingDirective) {
		return true
	}
	return false
}

// 判断是否是request_pool_size指令
func isServerRequestPoolSizeDirective(directive string) bool {
	if isEqualString(directive, ServerRequestPoolSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock_age指令
func isServerProxyCacheLockAgeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheLockAgeDirective) {
		return true
	}
	return false
}

// 判断是否是try_files指令
func isServerTryFilesDirective(directive string) bool {
	if isEqualString(directive, ServerTryFilesDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_max_temp_file_size指令
func isServerProxyMaxTempFileSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyMaxTempFileSizeDirective) {
		return true
	}
	return false
}

// 判断是否是absolute_redirect指令
func isServerAbsoluteRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerAbsoluteRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是chunked_transfer_encoding指令
func isServerChunkedTransferEncodingDirective(directive string) bool {
	if isEqualString(directive, ServerChunkedTransferEncodingDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_buffers指令
func isServerProxyBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cookie_domain指令
func isServerProxyCookieDomainDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCookieDomainDirective) {
		return true
	}
	return false
}

// 判断是否是merge_slashes指令
func isServerMergeSlashesDirective(directive string) bool {
	if isEqualString(directive, ServerMergeSlashesDirective) {
		return true
	}
	return false
}

// 判断是否是msie_refresh指令
func isServerMsieRefreshDirective(directive string) bool {
	if isEqualString(directive, ServerMsieRefreshDirective) {
		return true
	}
	return false
}

// 判断是否是subrequest_output_buffer_size指令
func isServerSubrequestOutputBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerSubrequestOutputBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是types_hash_max_size指令
func isServerTypesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, ServerTypesHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_request_body指令
func isServerProxyPassRequestBodyDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassRequestBodyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_send_timeout指令
func isServerProxySendTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxySendTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是ignore_invalid_headers指令
func isServerIgnoreInvalidHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerIgnoreInvalidHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是ssl_prefer_server_ciphers指令
func isServerSslPreferServerCiphersDirective(directive string) bool {
	if isEqualString(directive, ServerSslPreferServerCiphersDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache指令
func isServerOpenFileCacheDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheDirective) {
		return true
	}
	return false
}

// 判断是否是underscores_in_headers指令
func isServerUnderscoresInHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerUnderscoresInHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是types_hash_bucket_size指令
func isServerTypesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, ServerTypesHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是output_buffers指令
func isServerOutputBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerOutputBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_request_headers指令
func isServerProxyPassRequestHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassRequestHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_verify指令
func isServerProxySslVerifyDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslVerifyDirective) {
		return true
	}
	return false
}

// 判断是否是ssl_certificate_key指令
func isServerSslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, ServerSslCertificateKeyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock指令
func isServerProxyCacheLockDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheLockDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_valid指令
func isServerProxyCacheValidDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheValidDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_verify_depth指令
func isServerProxySslVerifyDepthDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslVerifyDepthDirective) {
		return true
	}
	return false
}

// 判断是否是deny指令
func isServerDenyDirective(directive string) bool {
	if isEqualString(directive, ServerDenyDirective) {
		return true
	}
	return false
}

// 判断是否是listen指令
func isServerListenDirective(directive string) bool {
	if isEqualString(directive, ServerListenDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_timeout指令
func isServerLingeringTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerLingeringTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_bind指令
func isServerProxyBindDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBindDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_key指令
func isServerProxyCacheKeyDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheKeyDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_socket_keepalive指令
func isServerProxySocketKeepaliveDirective(directive string) bool {
	if isEqualString(directive, ServerProxySocketKeepaliveDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_store_access指令
func isServerProxyStoreAccessDirective(directive string) bool {
	if isEqualString(directive, ServerProxyStoreAccessDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_session_reuse指令
func isServerProxySslSessionReuseDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslSessionReuseDirective) {
		return true
	}
	return false
}

// 判断是否是if指令
func isServerIfBlocksDirective(directive string) bool {
	if isEqualString(directive, ServerIfBlocksDirective) {
		return true
	}
	return false
}

// 判断是否是client_header_timeout指令
func isServerClientHeaderTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerClientHeaderTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_methods指令
func isServerProxyCacheMethodsDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheMethodsDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_force_ranges指令
func isServerProxyForceRangesDirective(directive string) bool {
	if isEqualString(directive, ServerProxyForceRangesDirective) {
		return true
	}
	return false
}

// 判断是否是return指令
func isServerReturnDirective(directive string) bool {
	if isEqualString(directive, ServerReturnDirective) {
		return true
	}
	return false
}

// 判断是否是port_in_redirect指令
func isServerPortInRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerPortInRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是directio_alignment指令
func isServerDirectioAlignmentDirective(directive string) bool {
	if isEqualString(directive, ServerDirectioAlignmentDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass指令
func isServerProxyPassDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassDirective) {
		return true
	}
	return false
}

// 判断是否是limit_rate指令
func isServerLimitRateDirective(directive string) bool {
	if isEqualString(directive, ServerLimitRateDirective) {
		return true
	}
	return false
}

// 判断是否是postpone_output指令
func isServerPostponeOutputDirective(directive string) bool {
	if isEqualString(directive, ServerPostponeOutputDirective) {
		return true
	}
	return false
}

// 判断是否是aio指令
func isServerAioDirective(directive string) bool {
	if isEqualString(directive, ServerAioDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_headers_hash_max_size指令
func isServerProxyHeadersHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyHeadersHashMaxSizeDirective) {
		return true
	}
	return false
}

// 判断是否是aio_write指令
func isServerAioWriteDirective(directive string) bool {
	if isEqualString(directive, ServerAioWriteDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_in_single_buffer指令
func isServerClientBodyInSingleBufferDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyInSingleBufferDirective) {
		return true
	}
	return false
}

// 判断是否是error_page指令
func isServerErrorPageDirective(directive string) bool {
	if isEqualString(directive, ServerErrorPageDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_set_body指令
func isServerProxySetBodyDirective(directive string) bool {
	if isEqualString(directive, ServerProxySetBodyDirective) {
		return true
	}
	return false
}

// 判断是否是ssl指令
func isServerSslDirective(directive string) bool {
	if isEqualString(directive, ServerSslDirective) {
		return true
	}
	return false
}

// 判断是否是limit_conn指令
func isServerLimitConnDirective(directive string) bool {
	if isEqualString(directive, ServerLimitConnDirective) {
		return true
	}
	return false
}

// 判断是否是limit_req指令
func isServerLimitReqDirective(directive string) bool {
	if isEqualString(directive, ServerLimitReqDirective) {
		return true
	}
	return false
}

// 判断是否是send_lowat指令
func isServerSendLowatDirective(directive string) bool {
	if isEqualString(directive, ServerSendLowatDirective) {
		return true
	}
	return false
}

// 判断是否是sendfile指令
func isServerSendFileDirective(directive string) bool {
	if isEqualString(directive, ServerSendFileDirective) {
		return true
	}
	return false
}

// 判断是否是tcp_nopush指令
func isServerTcpNopushDirective(directive string) bool {
	if isEqualString(directive, ServerTCPNopushDirective) {
		return true
	}
	return false
}

// 判断是否是directio指令
func isServerDirectIODirective(directive string) bool {
	if isEqualString(directive, ServerDirectIODirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ignore_headers指令
func isServerProxyIgnoreHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerProxyIgnoreHeadersDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_pass_header指令
func isServerProxyPassHeaderDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassHeaderDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache指令
func isServerProxyCacheDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_min_uses指令
func isServerProxyCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheMinUsesDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_path指令
func isServerProxyCachePathDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCachePathDirective) {
		return true
	}
	return false
}

// 判断是否是server指令
func isServerDirective(directive string) bool {
	if isEqualString(directive, ServerDirective) {
		return true
	}
	return false
}

// 判断是否是root指令
func isServerRootDirective(directive string) bool {
	if isEqualString(directive, ServerRootDirective) {
		return true
	}
	return false
}

// 判断是否是large_client_header_buffers指令
func isServerLargeClientHeaderBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerLargeClientHeaderBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是server_name_in_redirect指令
func isServerServerNameInRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerServerNameInRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_temp_path指令
func isServerClientBodyTempPathDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyTempPathDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_headers_hash_bucket_size指令
func isServerProxyHeadersHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyHeadersHashBucketSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_name指令
func isServerProxySslNameDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslNameDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_store指令
func isServerProxyStoreDirective(directive string) bool {
	if isEqualString(directive, ServerProxyStoreDirective) {
		return true
	}
	return false
}

// 判断是否是access_log指令
func isServerAccessLogDirective(directive string) bool {
	if isEqualString(directive, ServerAccessLogDirective) {
		return true
	}
	return false
}

// 判断是否是client_max_body_size指令
func isServerClientMaxBodySizeDirective(directive string) bool {
	if isEqualString(directive, ServerClientMaxBodySizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_temp_path指令
func isServerProxyTempPathDirective(directive string) bool {
	if isEqualString(directive, ServerProxyTempPathDirective) {
		return true
	}
	return false
}

// 判断是否是ssl_session_timeout指令
func isServerSslSessionTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerSslSessionTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是lingering_time指令
func isServerLingeringTimeDirective(directive string) bool {
	if isEqualString(directive, ServerLingeringTimeDirective) {
		return true
	}
	return false
}

// 判断是否是disable_symlinks指令
func isServerDisableSymlinksDirective(directive string) bool {
	if isEqualString(directive, ServerDisableSymlinksDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream_tries指令
func isServerProxyNextUpstreamTriesDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNextUpstreamTriesDirective) {
		return true
	}
	return false
}

// 判断是否是ssl_certificate指令
func isServerSslCertificateDirective(directive string) bool {
	if isEqualString(directive, ServerSslCertificateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_background_update指令
func isServerProxyCacheBackgroundUpdateDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheBackgroundUpdateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_crl指令
func isServerProxySslCrlDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCrlDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_redirect指令
func isServerProxyRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerProxyRedirectDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_request_buffering指令
func isServerProxyRequestBufferingDirective(directive string) bool {
	if isEqualString(directive, ServerProxyRequestBufferingDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_ssl_server_name指令
func isServerProxySslServerNameDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslServerNameDirective) {
		return true
	}
	return false
}

// 判断是否是location指令
func isServerLocationsDirective(directive string) bool {
	if isEqualString(directive, ServerLocationsDirective) {
		return true
	}
	return false
}

// 判断是否是log_not_found指令
func isServerLogNotFoundDirective(directive string) bool {
	if isEqualString(directive, ServerLogNotFoundDirective) {
		return true
	}
	return false
}

// 判断是否是cp_nodelay指令
func isServerTcpNodelayDirective(directive string) bool {
	if isEqualString(directive, ServerTCPNodelayDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_purge指令
func isServerProxyCachePurgeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCachePurgeDirective) {
		return true
	}
	return false
}

// 判断是否是ssl_protocols指令
func isServerSslProtocolsDirective(directive string) bool {
	if isEqualString(directive, ServerSslProtocolsDirective) {
		return true
	}
	return false
}

// 判断是否是resolver指令
func isServerResolverDirective(directive string) bool {
	if isEqualString(directive, ServerResolverDirective) {
		return true
	}
	return false
}

// 判断是否是sendfile_max_chunk指令
func isServerSendFileMaxChunkDirective(directive string) bool {
	if isEqualString(directive, ServerSendFileMaxChunkDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_lock_timeout指令
func isServerProxyCacheLockTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheLockTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是connection_pool_size指令
func isServerConnectionPoolSizeDirective(directive string) bool {
	if isEqualString(directive, ServerConnectionPoolSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_no_cache指令
func isServerProxyNoCacheDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNoCacheDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_timeout指令
func isServerClientBodyTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyTimeoutDirective) {
		return true
	}
	return false
}

// 判断是否是satisfy指令
func isServerSatisfyDirective(directive string) bool {
	if isEqualString(directive, ServerSatisfyDirective) {
		return true
	}
	return false
}

// 判断是否是client_header_buffer_size指令
func isServerClientHeaderBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerClientHeaderBufferSizeDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_intercept_errors指令
func isServerProxyInterceptErrorsDirective(directive string) bool {
	if isEqualString(directive, ServerProxyInterceptErrorsDirective) {
		return true
	}
	return false
}

// 判断是否是ssl_ciphers指令
func isServerSslCiphersDirective(directive string) bool {
	if isEqualString(directive, ServerSslCiphersDirective) {
		return true
	}
	return false
}

// 判断是否是limit_rate_after指令
func isServerLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, ServerLimitRateAfterDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_min_uses指令
func isServerOpenFileCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheMinUsesDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_valid指令
func isServerOpenFileCacheValidDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheValidDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_set_header指令
func isServerProxySetHeaderDirective(directive string) bool {
	if isEqualString(directive, ServerProxySetHeaderDirective) {
		return true
	}
	return false
}

// 判断是否是allow指令
func isServerAllowDirective(directive string) bool {
	if isEqualString(directive, ServerAllowDirective) {
		return true
	}
	return false
}

// 判断是否是log_subrequest指令
func isServerLogSubrequestDirective(directive string) bool {
	if isEqualString(directive, ServerLogSubrequestDirective) {
		return true
	}
	return false
}

// 判断是否是max_ranges指令
func isServerMaxRangesDirective(directive string) bool {
	if isEqualString(directive, ServerMaxRangesDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_max_range_offset指令
func isServerProxyCacheMaxRangeOffsetDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheMaxRangeOffsetDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_send_lowat指令
func isServerProxySendLowatDirective(directive string) bool {
	if isEqualString(directive, ServerProxySendLowatDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_cache_convert_head指令
func isServerProxyCacheConvertHeadDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheConvertHeadDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_limit_rate指令
func isServerProxyLimitRateDirective(directive string) bool {
	if isEqualString(directive, ServerProxyLimitRateDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_method指令
func isServerProxyMethodDirective(directive string) bool {
	if isEqualString(directive, ServerProxyMethodDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_next_upstream指令
func isServerProxyNextUpstreamDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNextUpstreamDirective) {
		return true
	}
	return false
}

// 判断是否是open_file_cache_errors指令
func isServerOpenFileCacheErrorsDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheErrorsDirective) {
		return true
	}
	return false
}

// 判断是否是recursive_error_pages指令
func isServerRecursiveErrorPagesDirective(directive string) bool {
	if isEqualString(directive, ServerRecursiveErrorPagesDirective) {
		return true
	}
	return false
}

// 判断是否是client_body_in_file_only指令
func isServerClientBodyInFileOnlyDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyInFileOnlyDirective) {
		return true
	}
	return false
}

// 判断是否是set指令
func isServerSetDirective(directive string) bool {
	if isEqualString(directive, ServerSetDirecrive) {
		return true
	}
	return false
}

// 判断是否是proxy_busy_buffers_size指令
func isServerProxyBusyBuffersSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBusyBuffersSizeDirective) {
		return true
	}
	return false
}

// 判断是否是add_header指令
func isServerAddHeaderDirective(directive string) bool {
	if isEqualString(directive, AddHeaderDirective) {
		return true
	}
	return false
}

// 判断是否gzip指令
func isServerGzipDirective(directive string) bool {
	if isEqualString(directive, ServerGzipDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_min_length指令
func isServerGzipMinLengthDirective(directive string) bool {
	if isEqualString(directive, ServerGzipMinLengthDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_buffers指令
func isServerGzipBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerGzipBuffersDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_comp_level指令
func isServerGzipCompLevelDirective(directive string) bool {
	if isEqualString(directive, ServerGzipCompLevelDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_types指令
func isServerGzipTypesDirective(directive string) bool {
	if isEqualString(directive, ServerGzipTypesDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_vary指令
func isServerGzipVaryDirective(directive string) bool {
	if isEqualString(directive, ServerGzipVaryDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_disable指令
func isServerGzipDisableDirective(directive string) bool {
	if isEqualString(directive, ServerGzipDisableDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_http_version指令
func isServerGzipHttpVersionDirective(directive string) bool {
	if isEqualString(directive, ServerGzipHTTPVersionDirective) {
		return true
	}
	return false
}

// 判断是否是gzip_proxied指令
func isServerGzipProxiedDirective(directive string) bool {
	if isEqualString(directive, ServerGzipProxiedDirective) {
		return true
	}
	return false
}

// 判断是否是proxy_hide_header指令
func isServerProxyHideHeaderDirective(directive string) bool {
	if isEqualString(directive, ServerGzipProxiedDirective) {
		return true
	}
	return false
}

// 处理Server中的指令
// 只处理如下用到的Server指令
/*
error_page
if
listen
location
return
rewrite
server_name
ssl
ssl_certificate
ssl_certificate_key
ssl_ciphers
ssl_prefer_server_ciphers
ssl_protocols
ssl_session_timeout
*/

func ProcessServer(block *Block) (*Server, error) {
	if !isServerDirective(block.Directive) {
		return nil, errors.New("Not server directive")
	}

	server := NewServer()
	for _, innerBlock := range block.InnerBlocks {

		// 处理listen指令
		if isServerListenDirective(innerBlock.Directive) {
			server.Listen = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理return指令
		if isServerReturnDirective(innerBlock.Directive) {
			server.Return = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理server_name指令
		if isServerNameDirective(innerBlock.Directive) {
			server.Name = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理ssl指令
		if isServerSslDirective(innerBlock.Directive) {
			server.Ssl = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理ssl_ciphers指令
		if isServerSslCiphersDirective(innerBlock.Directive) {
			ss := processArgsArray(innerBlock.Args)

			// 判断是否需要去除换行符
			server.SslCiphers = trimNewline(ss)

			// ToDo: add '${{sslCipher}}'
			// server.SslCiphers = "'" + server.SslCiphers + "'"

			continue
		}

		// 处理ssl_protocols指令
		if isServerSslProtocolsDirective(innerBlock.Directive) {
			server.SslProtocols = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理ssl_session_timeout指令
		if isServerSslSessionTimeoutDirective(innerBlock.Directive) {
			server.SslSessionTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理error_page指令
		if isServerErrorPageDirective(innerBlock.Directive) {
			server.ErrorPage = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理ssl_certificate_key指令
		if isServerSslCertificateKeyDirective(innerBlock.Directive) {
			server.SslCertificateKey = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理ssl_prefer_server_ciphers指令
		if isServerSslPreferServerCiphersDirective(innerBlock.Directive) {
			server.SslPreferServerCiphers = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理ssl_certificate指令
		if isServerSslCertificateDirective(innerBlock.Directive) {
			server.SslCertificate = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySslPasswordFileDirective(innerBlock.Directive) {
			server.ProxySslPasswordFile = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerReadAheadDirective(innerBlock.Directive) {
			server.ReadAhead = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyBufferingDirective(innerBlock.Directive) {
			server.ProxyBuffering = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerLimitConnDirective(innerBlock.Directive) {
			server.LimitConn = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerLimitReqDirective(innerBlock.Directive) {
			server.LimitReq = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyConnectTimeoutDirective(innerBlock.Directive) {
			server.ProxyConnectTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyIgnoreClientAbortDirective(innerBlock.Directive) {
			server.ProxyIgnoreClientAbort = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerKeepaliveDisableDirective(innerBlock.Directive) {
			server.KeepaliveDisable = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerSendTimeoutDirective(innerBlock.Directive) {
			server.SendTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyNextUpstreamTimeoutDirective(innerBlock.Directive) {
			server.ProxyNextUpstreamTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyTempFileWriteSizeDirective(innerBlock.Directive) {
			server.ProxyTempFileWriteSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerKeepaliveTimeoutDirective(innerBlock.Directive) {
			server.KeepaliveTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerLingeringCloseDirective(innerBlock.Directive) {
			server.LingeringClose = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyCacheUseStaleDirective(innerBlock.Directive) {
			server.ProxyCacheUseStale = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyReadTimeoutDirective(innerBlock.Directive) {
			server.ProxyReadTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyCacheBypassDirective(innerBlock.Directive) {
			server.ProxyCacheBypass = append(server.ProxyCacheBypass, processArgsArray(innerBlock.Args))
			continue
		}

		if isServerProxySslCertificateDirective(innerBlock.Directive) {
			server.ProxySslCertificate = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySslCertificateKeyDirective(innerBlock.Directive) {
			server.ProxySslCertificateKey = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySslProtocolsDirective(innerBlock.Directive) {
			server.ProxySslProtocols = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerKeepaliveRequestsDirective(innerBlock.Directive) {
			server.KeepaliveRequests = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerClientBodyBufferSizeDirective(innerBlock.Directive) {
			server.ClientBodyBufferSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyBufferSizeDirective(innerBlock.Directive) {
			server.ProxyBufferSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySslTrustedCertificateDirective(innerBlock.Directive) {
			server.ProxySslTrustedCertificate = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyCookiePathDirective(innerBlock.Directive) {
			server.ProxyCookiePath = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyHttpVersionDirective(innerBlock.Directive) {
			server.ProxyHttpVersion = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySslCiphersDirective(innerBlock.Directive) {
			server.ProxySslCiphers = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerMsiePaddingDirective(innerBlock.Directive) {
			server.MsiePadding = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerRequestPoolSizeDirective(innerBlock.Directive) {
			server.RequestPoolSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyCacheLockAgeDirective(innerBlock.Directive) {
			server.ProxyCacheLockAge = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerTryFilesDirective(innerBlock.Directive) {
			server.TryFiles = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyMaxTempFileSizeDirective(innerBlock.Directive) {
			server.ProxyMaxTempFileSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerAbsoluteRedirectDirective(innerBlock.Directive) {
			server.AbsoluteRedirect = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerChunkedTransferEncodingDirective(innerBlock.Directive) {
			server.ChunkedTransferEncoding = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyBuffersDirective(innerBlock.Directive) {
			server.ProxyBuffers = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyCookieDomainDirective(innerBlock.Directive) {
			server.ProxyCookieDomain = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerMergeSlashesDirective(innerBlock.Directive) {
			server.MergeSlashes = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerMsieRefreshDirective(innerBlock.Directive) {
			server.MsieRefresh = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerSubrequestOutputBufferSizeDirective(innerBlock.Directive) {
			server.SubrequestOutputBufferSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerTypesHashMaxSizeDirective(innerBlock.Directive) {
			server.TypesHashMaxSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyPassRequestBodyDirective(innerBlock.Directive) {
			server.ProxyPassRequestBody = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySendTimeoutDirective(innerBlock.Directive) {
			server.ProxySendTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerIgnoreInvalidHeadersDirective(innerBlock.Directive) {
			server.IgnoreInvalidHeaders = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerOpenFileCacheDirective(innerBlock.Directive) {
			server.OpenFileCache = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerUnderscoresInHeadersDirective(innerBlock.Directive) {
			server.UnderscoresInHeaders = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerTypesHashBucketSizeDirective(innerBlock.Directive) {
			server.TypesHashBucketSize = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerOutputBuffersDirective(innerBlock.Directive) {
			server.OutputBuffers = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyPassRequestHeadersDirective(innerBlock.Directive) {
			server.ProxyPassRequestHeaders = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySslVerifyDirective(innerBlock.Directive) {
			server.ProxySslVerify = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyCacheLockDirective(innerBlock.Directive) {
			server.ProxyCacheLock = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheValidDirective(innerBlock.Directive) {
			server.ProxyCacheValid = append(server.ProxyCacheValid, processArgsArray(innerBlock.Args))
			continue
		}
		if isServerProxySslVerifyDepthDirective(innerBlock.Directive) {
			server.ProxySslVerifyDepth = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerLingeringTimeoutDirective(innerBlock.Directive) {
			server.LingeringTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyBindDirective(innerBlock.Directive) {
			server.ProxyBind = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheKeyDirective(innerBlock.Directive) {
			server.ProxyCacheKey = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxySocketKeepaliveDirective(innerBlock.Directive) {
			server.ProxySocketKeepalive = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyStoreAccessDirective(innerBlock.Directive) {
			server.ProxyStoreAccess = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxySslSessionReuseDirective(innerBlock.Directive) {
			server.ProxySslSessionReuse = processArgsArray(innerBlock.Args)
			continue

		}
		if isServerClientHeaderTimeoutDirective(innerBlock.Directive) {
			server.ClientHeaderTimeout = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheMethodsDirective(innerBlock.Directive) {
			server.ProxyCacheMethods = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyForceRangesDirective(innerBlock.Directive) {
			server.ProxyForceRanges = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerPortInRedirectDirective(innerBlock.Directive) {
			server.PortInRedirect = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerDirectioAlignmentDirective(innerBlock.Directive) {
			server.DirectioAlignment = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyPassDirective(innerBlock.Directive) {
			server.ProxyPass = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerLimitRateDirective(innerBlock.Directive) {
			server.LimitRate = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerPostponeOutputDirective(innerBlock.Directive) {
			server.PostponeOutput = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerAioDirective(innerBlock.Directive) {
			server.Aio = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyHeadersHashMaxSizeDirective(innerBlock.Directive) {
			server.ProxyHeadersHashMaxSize = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerAioWriteDirective(innerBlock.Directive) {
			server.AioWrite = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerClientBodyInSingleBufferDirective(innerBlock.Directive) {
			server.ClientBodyInSingleBuffer = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxySetBodyDirective(innerBlock.Directive) {
			server.ProxySetBody = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerSendLowatDirective(innerBlock.Directive) {
			server.SendLowat = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerSendFileDirective(innerBlock.Directive) {
			server.SendFile = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerTcpNopushDirective(innerBlock.Directive) {
			server.TcpNopush = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerDirectIODirective(innerBlock.Directive) {
			server.DirectIO = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyIgnoreHeadersDirective(innerBlock.Directive) {
			server.ProxyIgnoreHeaders = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyPassHeaderDirective(innerBlock.Directive) {
			server.ProxyPassHeader = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheDirective(innerBlock.Directive) {
			server.ProxyCache = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheMinUsesDirective(innerBlock.Directive) {
			server.ProxyCacheMinUses = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCachePathDirective(innerBlock.Directive) {
			server.ProxyCachePath = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerRootDirective(innerBlock.Directive) {
			server.Root = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerLargeClientHeaderBuffersDirective(innerBlock.Directive) {
			server.LargeClientHeaderBuffers = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerServerNameInRedirectDirective(innerBlock.Directive) {
			server.ServerNameInRedirect = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerClientBodyTempPathDirective(innerBlock.Directive) {
			server.ClientBodyTempPath = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyHeadersHashBucketSizeDirective(innerBlock.Directive) {
			server.ProxyHeadersHashBucketSize = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxySslNameDirective(innerBlock.Directive) {
			server.ProxySslName = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyStoreDirective(innerBlock.Directive) {
			server.ProxyStore = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerAccessLogDirective(innerBlock.Directive) {
			server.AccessLog = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerClientMaxBodySizeDirective(innerBlock.Directive) {
			server.ClientMaxBodySize = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyTempPathDirective(innerBlock.Directive) {
			server.ProxyTempPath = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerLingeringTimeoutDirective(innerBlock.Directive) {
			server.LingeringTimeout = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerDisableSymlinksDirective(innerBlock.Directive) {
			server.DisableSymlinks = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyNextUpstreamTriesDirective(innerBlock.Directive) {
			server.ProxyNextUpstreamTries = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheBackgroundUpdateDirective(innerBlock.Directive) {
			server.ProxyCacheBackgroundUpdate = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxySslCrlDirective(innerBlock.Directive) {
			server.ProxySslCrl = processArgsArray(innerBlock.Args)
			continue
		}

		if isServerProxyRedirectDirective(innerBlock.Directive) {
			server.ProxyRedirect = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyRequestBufferingDirective(innerBlock.Directive) {
			server.ProxyRequestBuffering = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxySslServerNameDirective(innerBlock.Directive) {
			server.ProxySslServerName = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerLogNotFoundDirective(innerBlock.Directive) {
			server.LogNotFound = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerTcpNodelayDirective(innerBlock.Directive) {
			server.TcpNodelay = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCachePurgeDirective(innerBlock.Directive) {
			server.ProxyCachePurge = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerResolverDirective(innerBlock.Directive) {
			server.Resolver = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerSendFileMaxChunkDirective(innerBlock.Directive) {
			server.SendFileMaxChunk = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheLockTimeoutDirective(innerBlock.Directive) {
			server.ProxyCacheLockTimeout = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerConnectionPoolSizeDirective(innerBlock.Directive) {
			server.ConnectionPoolSize = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyNoCacheDirective(innerBlock.Directive) {
			server.ProxyNoCache = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerClientBodyTimeoutDirective(innerBlock.Directive) {
			server.ClientBodyTimeout = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerSatisfyDirective(innerBlock.Directive) {
			server.Satisfy = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerClientHeaderBufferSizeDirective(innerBlock.Directive) {
			server.ClientHeaderBufferSize = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyInterceptErrorsDirective(innerBlock.Directive) {
			server.ProxyInterceptErrors = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerLimitRateAfterDirective(innerBlock.Directive) {
			server.LimitRateAfter = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerOpenFileCacheMinUsesDirective(innerBlock.Directive) {
			server.OpenFileCacheMinUses = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerOpenFileCacheValidDirective(innerBlock.Directive) {
			server.OpenFileCacheValid = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxySetHeaderDirective(innerBlock.Directive) {
			server.ProxySetHeader = append(server.ProxySetHeader, processArgsArray(innerBlock.Args))
			continue
		}
		if isServerLogSubrequestDirective(innerBlock.Directive) {
			server.LogSubrequest = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerMaxRangesDirective(innerBlock.Directive) {
			server.MaxRanges = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheMaxRangeOffsetDirective(innerBlock.Directive) {
			server.ProxyCacheMaxRangeOffset = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxySendLowatDirective(innerBlock.Directive) {
			server.ProxySendLowat = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyCacheConvertHeadDirective(innerBlock.Directive) {
			server.ProxyCacheConvertHead = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyLimitRateDirective(innerBlock.Directive) {
			server.ProxyLimitRate = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyMethodDirective(innerBlock.Directive) {
			server.ProxyMethod = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerOpenFileCacheErrorsDirective(innerBlock.Directive) {
			server.OpenFileCacheErrors = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerRecursiveErrorPagesDirective(innerBlock.Directive) {
			server.RecursiveErrorPages = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerClientBodyInFileOnlyDirective(innerBlock.Directive) {
			server.ClientBodyInFileOnly = processArgsArray(innerBlock.Args)
			continue
		}
		if isServerProxyBusyBuffersSizeDirective(innerBlock.Directive) {
			server.ProxyBusyBuffersSize = processArgsArray(innerBlock.Args)
			continue
		}

		// 处理gzip指令
		if isServerGzipDirective(innerBlock.Directive) {
			server.Gzip = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_min_length指令
		if isServerGzipMinLengthDirective(innerBlock.Directive) {
			server.GzipMinLength = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_buffers指令
		if isServerGzipBuffersDirective(innerBlock.Directive) {
			server.GzipBuffers = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_comp_level指令
		if isServerGzipCompLevelDirective(innerBlock.Directive) {
			server.GzipCompLevel = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_types指令
		if isServerGzipTypesDirective(innerBlock.Directive) {
			server.GzipTypes = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_vary指令
		if isServerGzipVaryDirective(innerBlock.Directive) {
			server.GzipVary = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_disable指令
		if isServerGzipDisableDirective(innerBlock.Directive) {
			server.GzipDisable = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_http_version指令
		if isServerGzipHttpVersionDirective(innerBlock.Directive) {
			server.GzipHttpVersion = processArgsArray(innerBlock.Args)
		}

		// 处理gzip_proxied指令
		if isServerGzipProxiedDirective(innerBlock.Directive) {
			server.GzipProxied = processArgsArray(innerBlock.Args)
		}

		// 处理proxy_hide_header指令
		if isServerProxyHideHeaderDirective(innerBlock.Directive) {
			server.ProxyHideHeader = processArgsArray(innerBlock.Args)
		}

		// 处理allow指令
		if isServerAllowDirective(innerBlock.Directive) {
			server.Allow = append(server.Allow, processArgsArray(innerBlock.Args))
			continue
		}

		// 处理deny指令
		if isServerDenyDirective(innerBlock.Directive) {
			server.Deny = append(server.Deny, processArgsArray(innerBlock.Args))
			continue
		}

		// 处理set指令
		if isServerSetDirective(innerBlock.Directive) {
			if len(innerBlock.Args) > 1 {
				for index := range innerBlock.Args {
					if strings.EqualFold(innerBlock.Args[index], "") {
						innerBlock.Args[index] = "''"
					}
				}
				server.Set = append(server.Set, processArgsArray(innerBlock.Args))
			}
			continue
		}

		// 处理if指令
		if isServerIfBlocksDirective(innerBlock.Directive) {
			ifBlocks, err := ProcessServerIfBlocks(innerBlock)
			if err != nil {
				return server, err
			}
			server.IfBlocks = append(server.IfBlocks, *ifBlocks)
		}

		// 处理location指令
		if isServerLocationsDirective(innerBlock.Directive) {
			location, err := ProcessLocation(&innerBlock)
			if err != nil {
				return server, err
			}
			server.Locations = append(server.Locations, *location)
		}

		// 处理rewrite指令
		if isServerRewriteDirective(innerBlock.Directive) {
			args := processQuote(innerBlock.Args)
			server.Rewrite = append(server.Rewrite, processArgsArray(args))
		}

		// add_header
		if isServerAddHeaderDirective(innerBlock.Directive) {
			// 判断add_header指令里是否需要加引号
			for i := range innerBlock.Args {
				innerBlock.Args[i] = nonescapeQuotation(innerBlock.Args[i])
			}
			args := processQuote(innerBlock.Args)
			server.AddHeader = append(server.AddHeader, processArgsArray(args))
		}
	}
	return server, nil
}
