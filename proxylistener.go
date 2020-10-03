package proxylistener // import "github.com/anisse/proxylistener"

import (
	"bufio"
	"net"
	"time"

	proxyproto "github.com/pires/go-proxyproto"
)

type proxyListener struct {
	l net.Listener
}

type proxyConn struct {
	c net.Conn
	h *proxyproto.Header
	r *bufio.Reader
}

func Listen(network, address string) (net.Listener, error) {
	l, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return &proxyListener{l}, nil
}

func ListenTCP(network string, laddr *net.TCPAddr) (net.Listener, error) {
	l, err := net.ListenTCP(network, laddr)
	if err != nil {
		return nil, err
	}
	return &proxyListener{l}, nil
}

func (p *proxyListener) Accept() (net.Conn, error) {
	conn, err := p.l.Accept()
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(conn)

	h, err := proxyproto.Read(r)
	if err != nil {
		return nil, err
	}
	return &proxyConn{conn, h, r}, nil
}

// Functions below are "just" wrapped

func (p *proxyListener) Close() error   { return p.l.Close() }
func (p *proxyListener) Addr() net.Addr { return p.l.Addr() }

func (c *proxyConn) Read(b []byte) (n int, err error) {
	return c.r.Read(b)
}
func (c *proxyConn) LocalAddr() net.Addr {
	return &net.TCPAddr{
		IP:   c.h.DestinationAddress,
		Port: int(c.h.DestinationPort),
	}
}
func (c *proxyConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{
		IP:   c.h.SourceAddress,
		Port: int(c.h.SourcePort),
	}
}

// Functions below are "just" wrapped

func (c *proxyConn) Write(b []byte) (n int, err error)  { return c.c.Write(b) }
func (c *proxyConn) Close() error                       { return c.c.Close() }
func (c *proxyConn) SetDeadline(t time.Time) error      { return c.c.SetDeadline(t) }
func (c *proxyConn) SetReadDeadline(t time.Time) error  { return c.c.SetReadDeadline(t) }
func (c *proxyConn) SetWriteDeadline(t time.Time) error { return c.c.SetWriteDeadline(t) }
