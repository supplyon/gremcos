package gremcos

import (
	"context"
	"net"
	"net/http"
	"time"

	gorilla "github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/supplyon/gremcos/interfaces"
)

// WebsocketDialer is a function type for dialing/ connecting to a websocket server and creating a WebsocketConnection
type websocketDialer func(urlStr string, requestHeader http.Header) (interfaces.WebsocketConnection, *http.Response, error)

// websocketDialerFactory is a function type that is able to create WebsocketDialer's
type websocketDialerFactory func(writeBufferSize, readBufferSize int, handshakeTimout time.Duration) websocketDialer

// gorillaWebsocketDialerFactory is a function that is able to create WebsocketDialer's using the websocket implementation
// of github.com/gorilla/websocket
var gorillaWebsocketDialerFactory = func(logger zerolog.Logger) func(writeBufferSize, readBufferSize int, handshakeTimout time.Duration) websocketDialer {
	return func(writeBufferSize, readBufferSize int, handshakeTimout time.Duration) websocketDialer {
		// create the gorilla websocket dialer
		dialer := gorilla.Dialer{
			WriteBufferSize:  writeBufferSize,
			ReadBufferSize:   readBufferSize,
			HandshakeTimeout: handshakeTimout,
			NetDial:          verboseDialerfunc(logger),
		}

		// return the WebsocketDialer, wrapping the gorilla websocket dial call
		return func(urlStr string, requestHeader http.Header) (interfaces.WebsocketConnection, *http.Response, error) {
			return dialer.Dial(urlStr, requestHeader)
		}
	}
}

func dialWithTimeout(ctx context.Context, network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return d.DialContext(ctx, network, address)
}

func verboseDialerfunc(logger zerolog.Logger) func(network, addr string) (net.Conn, error) {
	dnsTimeout := time.Second
	return func(network, addr string) (net.Conn, error) {
		resolver := &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				dialStart := time.Now()

				conn, err := dialWithTimeout(ctx, network, address, dnsTimeout)
				if err != nil {
					logger.Error().Err(err).Msgf("DNS Lookup %#v %#v (err=%v) took %s", network, address, err, time.Since(dialStart))

					// try to re-resolve using cloudflare dns
					reResolveAddress := "8.8.8.8:53"
					dialStartReresolve := time.Now()
					_, errReResolve := dialWithTimeout(ctx, network, reResolveAddress, dnsTimeout)
					logger.Error().Err(err).Msgf("Retry DNS Lookup %#v %#v (err=%v) took %s", network, reResolveAddress, errReResolve, time.Since(dialStartReresolve))
					return nil, errors.Wrap(err, "connecting to dns during reresolve")
				}
				logger.Info().Msgf("DNS Lookup %#v %#v took %s", network, address, time.Since(dialStart))

				return conn, errors.Wrap(err, "connecting to dns")

			},
		}
		netDialer := &net.Dialer{
			Resolver: resolver,
		}

		dialStart := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		conn, err := netDialer.DialContext(ctx, network, addr)
		if err != nil {
			dialDur := time.Since(dialStart)
			logger.Warn().Err(err).Msgf("Dialing %#v %#v failed: %v after %s", network, addr, err, dialDur)
		}

		return conn, errors.Wrap(err, "dialing to cosmos")
	}
}
