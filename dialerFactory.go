package gremcos

import (
	"context"
	"net"
	"net/http"
	"time"

	gorilla "github.com/gorilla/websocket"
	"github.com/rs/zerolog"
	"github.com/supplyon/gremcos/interfaces"
)

// WebsocketDialer is a function type for dialing/ connecting to a websocket server and creating a WebsocketConnection
type websocketDialer func(urlStr string, requestHeader http.Header) (interfaces.WebsocketConnection, *http.Response, error)

// websocketDialerFactory is a function type that is able to create WebsocketDialer's
type websocketDialerFactory func(writeBufferSize, readBufferSize int, handshakeTimout time.Duration) websocketDialer

// gorillaWebsocketDialerFactory is a function that is able to create WebsocketDialer's using the websocket implementation
// of github.com/gorilla/websocket
var gorillaWebsocketDialerFactory = func(writeBufferSize, readBufferSize int, handshakeTimout time.Duration) websocketDialer {
	// create the gorilla websocket dialer
	dialer := gorilla.Dialer{
		WriteBufferSize:  writeBufferSize,
		ReadBufferSize:   readBufferSize,
		HandshakeTimeout: handshakeTimout,
	}

	// return the WebsocketDialer, wrapping the gorilla websocket dial call
	return func(urlStr string, requestHeader http.Header) (interfaces.WebsocketConnection, *http.Response, error) {
		return dialer.Dial(urlStr, requestHeader)
	}

}

var DebugGorillaWebsocketDialerFactory = func(logger zerolog.Logger) func(writeBufferSize, readBufferSize int, handshakeTimout time.Duration) websocketDialer {
	return func(writeBufferSize, readBufferSize int, handshakeTimout time.Duration) websocketDialer {
		netDialer := &net.Dialer{}
		netDial := func(network, addr string) (net.Conn, error) {
			dialStart := time.Now()
			conn, err := netDialer.DialContext(context.Background(), network, addr)
			if err != nil {
				dialDur := time.Since(dialStart)
				logger.Warn().Err(err).Msgf("Dialing %#v %#v failed: %v after %s", network, addr, err, dialDur)
			}
			return conn, err
		}

		// create the gorilla websocket dialer
		dialer := gorilla.Dialer{
			WriteBufferSize:  writeBufferSize,
			ReadBufferSize:   readBufferSize,
			HandshakeTimeout: handshakeTimout,
			NetDial:          netDial,
		}

		// return the WebsocketDialer, wrapping the gorilla websocket dial call
		return func(urlStr string, requestHeader http.Header) (interfaces.WebsocketConnection, *http.Response, error) {
			return dialer.Dial(urlStr, requestHeader)
		}
	}
}
