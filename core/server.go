package core

import (
	"errors"
	"strings"
)

// determine whether it is a proxy_ssl_password_file directive
func isServerProxySslPasswordFileDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslPasswordFileDirective) {
		return true
	}
	return false
}

// determine whether it is a read_ahead directive
func isServerReadAheadDirective(directive string) bool {
	if isEqualString(directive, ServerReadAheadDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffering directive
func isServerProxyBufferingDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBufferingDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_connect_timeout directive
func isServerProxyConnectTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyConnectTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ignore_client_abort directive
func isServerProxyIgnoreClientAbortDirective(directive string) bool {
	if isEqualString(directive, ServerProxyIgnoreClientAbortDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_disable directive
func isServerKeepaliveDisableDirective(directive string) bool {
	if isEqualString(directive, ServerKeepaliveDisableDirective) {
		return true
	}
	return false
}

// determine whether it is a send_timeout directive
func isServerSendTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerSendTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream_timeout directive
func isServerProxyNextUpstreamTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNextUpstreamTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_temp_file_write_size directive
func isServerProxyTempFileWriteSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyTempFileWriteSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_timeout directive
func isServerKeepaliveTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerKeepaliveTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_close directive
func isServerLingeringCloseDirective(directive string) bool {
	if isEqualString(directive, ServerLingeringCloseDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_use_stale directive
func isServerProxyCacheUseStaleDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheUseStaleDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_read_timeout directive
func isServerProxyReadTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyReadTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_bypass directive
func isServerProxyCacheBypassDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheBypassDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_certificate directive
func isServerProxySslCertificateDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCertificateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_certificate_key directive
func isServerProxySslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCertificateKeyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_protocols directive
func isServerProxySslProtocolsDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslProtocolsDirective) {
		return true
	}
	return false
}

// determine whether it is a rewrite directive
func isServerRewriteDirective(directive string) bool {
	if isEqualString(directive, ServerRewriteDirective) {
		return true
	}
	return false
}

// determine whether it is a keepalive_requests directive
func isServerKeepaliveRequestsDirective(directive string) bool {
	if isEqualString(directive, ServerKeepaliveRequestsDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_buffer_size directive
func isServerClientBodyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffer_size directive
func isServerProxyBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_trusted_certificate directive
func isServerProxySslTrustedCertificateDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslTrustedCertificateDirective) {
		return true
	}
	return false
}

// determine whether it is a server_name directive
func isServerNameDirective(directive string) bool {
	if isEqualString(directive, ServerNameDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cookie_path directive
func isServerProxyCookiePathDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCookiePathDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_http_version directive
func isServerProxyHTTPVersionDirective(directive string) bool {
	if isEqualString(directive, ServerProxyHTTPVersionDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_ciphers directive
func isServerProxySslCiphersDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCiphersDirective) {
		return true
	}
	return false
}

// determine whether it is a msie_padding directive
func isServerMsiePaddingDirective(directive string) bool {
	if isEqualString(directive, ServerMsiePaddingDirective) {
		return true
	}
	return false
}

// determine whether it is a request_pool_size directive
func isServerRequestPoolSizeDirective(directive string) bool {
	if isEqualString(directive, ServerRequestPoolSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock_age directive
func isServerProxyCacheLockAgeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheLockAgeDirective) {
		return true
	}
	return false
}

// determine whether it is a try_files directive
func isServerTryFilesDirective(directive string) bool {
	if isEqualString(directive, ServerTryFilesDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_max_temp_file_size directive
func isServerProxyMaxTempFileSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyMaxTempFileSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a absolute_redirect directive
func isServerAbsoluteRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerAbsoluteRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a chunked_transfer_encoding directive
func isServerChunkedTransferEncodingDirective(directive string) bool {
	if isEqualString(directive, ServerChunkedTransferEncodingDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_buffers directive
func isServerProxyBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cookie_domain directive
func isServerProxyCookieDomainDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCookieDomainDirective) {
		return true
	}
	return false
}

// determine whether it is a merge_slashes directive
func isServerMergeSlashesDirective(directive string) bool {
	if isEqualString(directive, ServerMergeSlashesDirective) {
		return true
	}
	return false
}

// determine whether it is a msie_refresh directive
func isServerMsieRefreshDirective(directive string) bool {
	if isEqualString(directive, ServerMsieRefreshDirective) {
		return true
	}
	return false
}

// determine whether it is a subrequest_output_buffer_size directive
func isServerSubrequestOutputBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerSubrequestOutputBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a types_hash_max_size directive
func isServerTypesHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, ServerTypesHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_request_body directive
func isServerProxyPassRequestBodyDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassRequestBodyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_send_timeout directive
func isServerProxySendTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxySendTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a ignore_invalid_headers directive
func isServerIgnoreInvalidHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerIgnoreInvalidHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl_prefer_server_ciphers directive
func isServerSslPreferServerCiphersDirective(directive string) bool {
	if isEqualString(directive, ServerSslPreferServerCiphersDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache directive
func isServerOpenFileCacheDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a underscores_in_headers directive
func isServerUnderscoresInHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerUnderscoresInHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a types_hash_bucket_size directive
func isServerTypesHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, ServerTypesHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a output_buffers directive
func isServerOutputBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerOutputBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_request_headers directive
func isServerProxyPassRequestHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassRequestHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_verify directive
func isServerProxySslVerifyDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslVerifyDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl_certificate_key directive
func isServerSslCertificateKeyDirective(directive string) bool {
	if isEqualString(directive, ServerSslCertificateKeyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock directive
func isServerProxyCacheLockDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheLockDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_valid directive
func isServerProxyCacheValidDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheValidDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_verify_depth directive
func isServerProxySslVerifyDepthDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslVerifyDepthDirective) {
		return true
	}
	return false
}

// determine whether it is a deny directive
func isServerDenyDirective(directive string) bool {
	if isEqualString(directive, ServerDenyDirective) {
		return true
	}
	return false
}

// determine whether it is a listen directive
func isServerListenDirective(directive string) bool {
	if isEqualString(directive, ServerListenDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_timeout directive
func isServerLingeringTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerLingeringTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_bind directive
func isServerProxyBindDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBindDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_key directive
func isServerProxyCacheKeyDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheKeyDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_socket_keepalive directive
func isServerProxySocketKeepaliveDirective(directive string) bool {
	if isEqualString(directive, ServerProxySocketKeepaliveDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_store_access directive
func isServerProxyStoreAccessDirective(directive string) bool {
	if isEqualString(directive, ServerProxyStoreAccessDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_session_reuse directive
func isServerProxySslSessionReuseDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslSessionReuseDirective) {
		return true
	}
	return false
}

// determine whether it is a if directive
func isServerIfBlocksDirective(directive string) bool {
	if isEqualString(directive, ServerIfBlocksDirective) {
		return true
	}
	return false
}

// determine whether it is a client_header_timeout directive
func isServerClientHeaderTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerClientHeaderTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_methods directive
func isServerProxyCacheMethodsDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheMethodsDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_force_ranges directive
func isServerProxyForceRangesDirective(directive string) bool {
	if isEqualString(directive, ServerProxyForceRangesDirective) {
		return true
	}
	return false
}

// determine whether it is a return directive
func isServerReturnDirective(directive string) bool {
	if isEqualString(directive, ServerReturnDirective) {
		return true
	}
	return false
}

// determine whether it is a port_in_redirect directive
func isServerPortInRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerPortInRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a directio_alignment directive
func isServerDirectioAlignmentDirective(directive string) bool {
	if isEqualString(directive, ServerDirectioAlignmentDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass directive
func isServerProxyPassDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_rate directive
func isServerLimitRateDirective(directive string) bool {
	if isEqualString(directive, ServerLimitRateDirective) {
		return true
	}
	return false
}

// determine whether it is a postpone_output directive
func isServerPostponeOutputDirective(directive string) bool {
	if isEqualString(directive, ServerPostponeOutputDirective) {
		return true
	}
	return false
}

// determine whether it is a aio directive
func isServerAioDirective(directive string) bool {
	if isEqualString(directive, ServerAioDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_headers_hash_max_size directive
func isServerProxyHeadersHashMaxSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyHeadersHashMaxSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a aio_write directive
func isServerAioWriteDirective(directive string) bool {
	if isEqualString(directive, ServerAioWriteDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_in_single_buffer directive
func isServerClientBodyInSingleBufferDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyInSingleBufferDirective) {
		return true
	}
	return false
}

// determine whether it is a error_page directive
func isServerErrorPageDirective(directive string) bool {
	if isEqualString(directive, ServerErrorPageDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_set_body directive
func isServerProxySetBodyDirective(directive string) bool {
	if isEqualString(directive, ServerProxySetBodyDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl directive
func isServerSslDirective(directive string) bool {
	if isEqualString(directive, ServerSslDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_conn directive
func isServerLimitConnDirective(directive string) bool {
	if isEqualString(directive, ServerLimitConnDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_req directive
func isServerLimitReqDirective(directive string) bool {
	if isEqualString(directive, ServerLimitReqDirective) {
		return true
	}
	return false
}

// determine whether it is a send_lowat directive
func isServerSendLowatDirective(directive string) bool {
	if isEqualString(directive, ServerSendLowatDirective) {
		return true
	}
	return false
}

// determine whether it is a sendfile directive
func isServerSendFileDirective(directive string) bool {
	if isEqualString(directive, ServerSendFileDirective) {
		return true
	}
	return false
}

// determine whether it is a tcp_nopush directive
func isServerTCPNopushDirective(directive string) bool {
	if isEqualString(directive, ServerTCPNopushDirective) {
		return true
	}
	return false
}

// determine whether it is a directio directive
func isServerDirectIODirective(directive string) bool {
	if isEqualString(directive, ServerDirectIODirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ignore_headers directive
func isServerProxyIgnoreHeadersDirective(directive string) bool {
	if isEqualString(directive, ServerProxyIgnoreHeadersDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_pass_header directive
func isServerProxyPassHeaderDirective(directive string) bool {
	if isEqualString(directive, ServerProxyPassHeaderDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache directive
func isServerProxyCacheDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_min_uses directive
func isServerProxyCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheMinUsesDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_path directive
func isServerProxyCachePathDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCachePathDirective) {
		return true
	}
	return false
}

// determine whether it is a server directive
func isServerDirective(directive string) bool {
	if isEqualString(directive, ServerDirective) {
		return true
	}
	return false
}

// determine whether it is a root directive
func isServerRootDirective(directive string) bool {
	if isEqualString(directive, ServerRootDirective) {
		return true
	}
	return false
}

// determine whether it is a large_client_header_buffers directive
func isServerLargeClientHeaderBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerLargeClientHeaderBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a server_name_in_redirect directive
func isServerServerNameInRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerServerNameInRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_temp_path directive
func isServerClientBodyTempPathDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyTempPathDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_headers_hash_bucket_size directive
func isServerProxyHeadersHashBucketSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyHeadersHashBucketSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_name directive
func isServerProxySslNameDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslNameDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_store directive
func isServerProxyStoreDirective(directive string) bool {
	if isEqualString(directive, ServerProxyStoreDirective) {
		return true
	}
	return false
}

// determine whether it is a access_log directive
func isServerAccessLogDirective(directive string) bool {
	if isEqualString(directive, ServerAccessLogDirective) {
		return true
	}
	return false
}

// determine whether it is a client_max_body_size directive
func isServerClientMaxBodySizeDirective(directive string) bool {
	if isEqualString(directive, ServerClientMaxBodySizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_temp_path directive
func isServerProxyTempPathDirective(directive string) bool {
	if isEqualString(directive, ServerProxyTempPathDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl_session_timeout directive
func isServerSslSessionTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerSslSessionTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a lingering_time directive
func isServerLingeringTimeDirective(directive string) bool {
	if isEqualString(directive, ServerLingeringTimeDirective) {
		return true
	}
	return false
}

// determine whether it is a disable_symlinks directive
func isServerDisableSymlinksDirective(directive string) bool {
	if isEqualString(directive, ServerDisableSymlinksDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream_tries directive
func isServerProxyNextUpstreamTriesDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNextUpstreamTriesDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl_certificate directive
func isServerSslCertificateDirective(directive string) bool {
	if isEqualString(directive, ServerSslCertificateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_background_update directive
func isServerProxyCacheBackgroundUpdateDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheBackgroundUpdateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_crl directive
func isServerProxySslCrlDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslCrlDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_redirect directive
func isServerProxyRedirectDirective(directive string) bool {
	if isEqualString(directive, ServerProxyRedirectDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_request_buffering directive
func isServerProxyRequestBufferingDirective(directive string) bool {
	if isEqualString(directive, ServerProxyRequestBufferingDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_ssl_server_name directive
func isServerProxySslServerNameDirective(directive string) bool {
	if isEqualString(directive, ServerProxySslServerNameDirective) {
		return true
	}
	return false
}

// determine whether it is a location directive
func isServerLocationsDirective(directive string) bool {
	if isEqualString(directive, ServerLocationsDirective) {
		return true
	}
	return false
}

// determine whether it is a log_not_found directive
func isServerLogNotFoundDirective(directive string) bool {
	if isEqualString(directive, ServerLogNotFoundDirective) {
		return true
	}
	return false
}

// determine whether it is a cp_nodelay directive
func isServerTCPNodelayDirective(directive string) bool {
	if isEqualString(directive, ServerTCPNodelayDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_purge directive
func isServerProxyCachePurgeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCachePurgeDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl_protocols directive
func isServerSslProtocolsDirective(directive string) bool {
	if isEqualString(directive, ServerSslProtocolsDirective) {
		return true
	}
	return false
}

// determine whether it is a resolver directive
func isServerResolverDirective(directive string) bool {
	if isEqualString(directive, ServerResolverDirective) {
		return true
	}
	return false
}

// determine whether it is a sendfile_max_chunk directive
func isServerSendFileMaxChunkDirective(directive string) bool {
	if isEqualString(directive, ServerSendFileMaxChunkDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_lock_timeout directive
func isServerProxyCacheLockTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheLockTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a connection_pool_size directive
func isServerConnectionPoolSizeDirective(directive string) bool {
	if isEqualString(directive, ServerConnectionPoolSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_no_cache directive
func isServerProxyNoCacheDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNoCacheDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_timeout directive
func isServerClientBodyTimeoutDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyTimeoutDirective) {
		return true
	}
	return false
}

// determine whether it is a satisfy directive
func isServerSatisfyDirective(directive string) bool {
	if isEqualString(directive, ServerSatisfyDirective) {
		return true
	}
	return false
}

// determine whether it is a client_header_buffer_size directive
func isServerClientHeaderBufferSizeDirective(directive string) bool {
	if isEqualString(directive, ServerClientHeaderBufferSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_intercept_errors directive
func isServerProxyInterceptErrorsDirective(directive string) bool {
	if isEqualString(directive, ServerProxyInterceptErrorsDirective) {
		return true
	}
	return false
}

// determine whether it is a ssl_ciphers directive
func isServerSslCiphersDirective(directive string) bool {
	if isEqualString(directive, ServerSslCiphersDirective) {
		return true
	}
	return false
}

// determine whether it is a limit_rate_after directive
func isServerLimitRateAfterDirective(directive string) bool {
	if isEqualString(directive, ServerLimitRateAfterDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_min_uses directive
func isServerOpenFileCacheMinUsesDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheMinUsesDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_valid directive
func isServerOpenFileCacheValidDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheValidDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_set_header directive
func isServerProxySetHeaderDirective(directive string) bool {
	if isEqualString(directive, ServerProxySetHeaderDirective) {
		return true
	}
	return false
}

// determine whether it is a allow directive
func isServerAllowDirective(directive string) bool {
	if isEqualString(directive, ServerAllowDirective) {
		return true
	}
	return false
}

// determine whether it is a log_subrequest directive
func isServerLogSubrequestDirective(directive string) bool {
	if isEqualString(directive, ServerLogSubrequestDirective) {
		return true
	}
	return false
}

// determine whether it is a max_ranges directive
func isServerMaxRangesDirective(directive string) bool {
	if isEqualString(directive, ServerMaxRangesDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_max_range_offset directive
func isServerProxyCacheMaxRangeOffsetDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheMaxRangeOffsetDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_send_lowat directive
func isServerProxySendLowatDirective(directive string) bool {
	if isEqualString(directive, ServerProxySendLowatDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_cache_convert_head directive
func isServerProxyCacheConvertHeadDirective(directive string) bool {
	if isEqualString(directive, ServerProxyCacheConvertHeadDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_limit_rate directive
func isServerProxyLimitRateDirective(directive string) bool {
	if isEqualString(directive, ServerProxyLimitRateDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_method directive
func isServerProxyMethodDirective(directive string) bool {
	if isEqualString(directive, ServerProxyMethodDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_next_upstream directive
func isServerProxyNextUpstreamDirective(directive string) bool {
	if isEqualString(directive, ServerProxyNextUpstreamDirective) {
		return true
	}
	return false
}

// determine whether it is a open_file_cache_errors directive
func isServerOpenFileCacheErrorsDirective(directive string) bool {
	if isEqualString(directive, ServerOpenFileCacheErrorsDirective) {
		return true
	}
	return false
}

// determine whether it is a recursive_error_pages directive
func isServerRecursiveErrorPagesDirective(directive string) bool {
	if isEqualString(directive, ServerRecursiveErrorPagesDirective) {
		return true
	}
	return false
}

// determine whether it is a client_body_in_file_only directive
func isServerClientBodyInFileOnlyDirective(directive string) bool {
	if isEqualString(directive, ServerClientBodyInFileOnlyDirective) {
		return true
	}
	return false
}

// determine whether it is a set directive
func isServerSetDirective(directive string) bool {
	if isEqualString(directive, ServerSetDirecrive) {
		return true
	}
	return false
}

// determine whether it is a proxy_busy_buffers_size directive
func isServerProxyBusyBuffersSizeDirective(directive string) bool {
	if isEqualString(directive, ServerProxyBusyBuffersSizeDirective) {
		return true
	}
	return false
}

// determine whether it is a add_header directive
func isServerAddHeaderDirective(directive string) bool {
	if isEqualString(directive, AddHeaderDirective) {
		return true
	}
	return false
}

// gzip directive
func isServerGzipDirective(directive string) bool {
	if isEqualString(directive, ServerGzipDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_min_length directive
func isServerGzipMinLengthDirective(directive string) bool {
	if isEqualString(directive, ServerGzipMinLengthDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_buffers directive
func isServerGzipBuffersDirective(directive string) bool {
	if isEqualString(directive, ServerGzipBuffersDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_comp_level directive
func isServerGzipCompLevelDirective(directive string) bool {
	if isEqualString(directive, ServerGzipCompLevelDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_types directive
func isServerGzipTypesDirective(directive string) bool {
	if isEqualString(directive, ServerGzipTypesDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_vary directive
func isServerGzipVaryDirective(directive string) bool {
	if isEqualString(directive, ServerGzipVaryDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_disable directive
func isServerGzipDisableDirective(directive string) bool {
	if isEqualString(directive, ServerGzipDisableDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_http_version directive
func isServerGzipHTTPVersionDirective(directive string) bool {
	if isEqualString(directive, ServerGzipHTTPVersionDirective) {
		return true
	}
	return false
}

// determine whether it is a gzip_proxied directive
func isServerGzipProxiedDirective(directive string) bool {
	if isEqualString(directive, ServerGzipProxiedDirective) {
		return true
	}
	return false
}

// determine whether it is a proxy_hide_header directive
func isServerProxyHideHeaderDirective(directive string) bool {
	if isEqualString(directive, ServerGzipProxiedDirective) {
		return true
	}
	return false
}

// ProcessServer process server module
func ProcessServer(block *Block) (*Server, error) {
	if !isServerDirective(block.Directive) {
		return nil, errors.New("Not server directive")
	}

	server := NewServer()
	for _, innerBlock := range block.InnerBlocks {

		// process listen directive
		if isServerListenDirective(innerBlock.Directive) {
			server.Listen = processArgsArray(innerBlock.Args)
			continue
		}

		// process return directive
		if isServerReturnDirective(innerBlock.Directive) {
			server.Return = processArgsArray(innerBlock.Args)
			continue
		}

		// process server_name directive
		if isServerNameDirective(innerBlock.Directive) {
			server.Name = processArgsArray(innerBlock.Args)
			continue
		}

		// process ssl directive
		if isServerSslDirective(innerBlock.Directive) {
			server.Ssl = processArgsArray(innerBlock.Args)
			continue
		}

		// process ssl_ciphers directive
		if isServerSslCiphersDirective(innerBlock.Directive) {
			ss := processArgsArray(innerBlock.Args)

			server.SslCiphers = trimNewline(ss)
			continue
		}

		// process ssl_protocols directive
		if isServerSslProtocolsDirective(innerBlock.Directive) {
			server.SslProtocols = processArgsArray(innerBlock.Args)
			continue
		}

		// process ssl_session_timeout directive
		if isServerSslSessionTimeoutDirective(innerBlock.Directive) {
			server.SslSessionTimeout = processArgsArray(innerBlock.Args)
			continue
		}

		// process error_page directive
		if isServerErrorPageDirective(innerBlock.Directive) {
			server.ErrorPage = processArgsArray(innerBlock.Args)
			continue
		}

		// process ssl_certificate_key directive
		if isServerSslCertificateKeyDirective(innerBlock.Directive) {
			server.SslCertificateKey = processArgsArray(innerBlock.Args)
			continue
		}

		// process ssl_prefer_server_ciphers directive
		if isServerSslPreferServerCiphersDirective(innerBlock.Directive) {
			server.SslPreferServerCiphers = processArgsArray(innerBlock.Args)
			continue
		}

		// process ssl_certificate directive
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

		if isServerProxyHTTPVersionDirective(innerBlock.Directive) {
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
		if isServerTCPNopushDirective(innerBlock.Directive) {
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
		if isServerTCPNodelayDirective(innerBlock.Directive) {
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

		// process gzip directive
		if isServerGzipDirective(innerBlock.Directive) {
			server.Gzip = processArgsArray(innerBlock.Args)
		}

		// process gzip_min_length directive
		if isServerGzipMinLengthDirective(innerBlock.Directive) {
			server.GzipMinLength = processArgsArray(innerBlock.Args)
		}

		// process gzip_buffers directive
		if isServerGzipBuffersDirective(innerBlock.Directive) {
			server.GzipBuffers = processArgsArray(innerBlock.Args)
		}

		// process gzip_comp_level directive
		if isServerGzipCompLevelDirective(innerBlock.Directive) {
			server.GzipCompLevel = processArgsArray(innerBlock.Args)
		}

		// process gzip_types directive
		if isServerGzipTypesDirective(innerBlock.Directive) {
			server.GzipTypes = processArgsArray(innerBlock.Args)
		}

		// process gzip_vary directive
		if isServerGzipVaryDirective(innerBlock.Directive) {
			server.GzipVary = processArgsArray(innerBlock.Args)
		}

		// process gzip_disable directive
		if isServerGzipDisableDirective(innerBlock.Directive) {
			server.GzipDisable = processArgsArray(innerBlock.Args)
		}

		// process gzip_http_version directive
		if isServerGzipHTTPVersionDirective(innerBlock.Directive) {
			server.GzipHttpVersion = processArgsArray(innerBlock.Args)
		}

		// process gzip_proxied directive
		if isServerGzipProxiedDirective(innerBlock.Directive) {
			server.GzipProxied = processArgsArray(innerBlock.Args)
		}

		// process proxy_hide_header directive
		if isServerProxyHideHeaderDirective(innerBlock.Directive) {
			server.ProxyHideHeader = processArgsArray(innerBlock.Args)
		}

		// process allow directive
		if isServerAllowDirective(innerBlock.Directive) {
			server.Allow = append(server.Allow, processArgsArray(innerBlock.Args))
			continue
		}

		// process deny directive
		if isServerDenyDirective(innerBlock.Directive) {
			server.Deny = append(server.Deny, processArgsArray(innerBlock.Args))
			continue
		}

		// process set directive
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

		// process if directive
		if isServerIfBlocksDirective(innerBlock.Directive) {
			ifBlocks, err := ProcessServerIfBlocks(innerBlock)
			if err != nil {
				return server, err
			}
			server.IfBlocks = append(server.IfBlocks, *ifBlocks)
		}

		// process location directive
		if isServerLocationsDirective(innerBlock.Directive) {
			location, err := ProcessLocation(&innerBlock)
			if err != nil {
				return server, err
			}
			server.Locations = append(server.Locations, *location)
		}

		// process rewrite directive
		if isServerRewriteDirective(innerBlock.Directive) {
			args := processQuote(innerBlock.Args)
			server.Rewrite = append(server.Rewrite, processArgsArray(args))
		}

		// add_header
		if isServerAddHeaderDirective(innerBlock.Directive) {
			for i := range innerBlock.Args {
				innerBlock.Args[i] = nonescapeQuotation(innerBlock.Args[i])
			}
			args := processQuote(innerBlock.Args)
			server.AddHeader = append(server.AddHeader, processArgsArray(args))
		}
	}
	return server, nil
}
