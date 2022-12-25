package fix

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io"
	"net"
)

func Connect(addr string, cert []byte) (conn io.ReadWriteCloser, err error) {
	if cert != nil {
		pool := x509.NewCertPool()
		if cert != nil {
			ok := pool.AppendCertsFromPEM(cert)
			if !ok {
				err = errors.New("cannot parse PEM certificate for tcp dial")
				return
			}
		}
		conn, err = tls.Dial("tcp", addr, &tls.Config{
			RootCAs:            pool,
			InsecureSkipVerify: true,
		})
		return
	}

	conn, err = net.Dial("tcp", addr)
	return
}
