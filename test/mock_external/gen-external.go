package mock_external

//go:generate mockgen -destination=kubernetes.go -package $GOPACKAGE k8s.io/client-go/kubernetes Interface

