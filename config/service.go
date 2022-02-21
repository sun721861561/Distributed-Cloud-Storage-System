package config

const (
	// UploadServiceHost : Address of the upload service listener
	UploadServiceHost = "0.0.0.0:8080"
	// UploadLBHost: Upload service LB address
	UploadLBHost = "http://upload.fileserver.com"
	// DownloadLBHost: Download Service LB Address
	DownloadLBHost = "http://download.fileserver.com"
	// TracerAgentHost: tracing agent address
	TracerAgentHost = "127.0.0.1:6831"
)
