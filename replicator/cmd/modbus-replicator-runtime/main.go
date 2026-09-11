package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/tamzrod/MCS.OSJS/replicator"
)

const maxMessage = 1024 * 1024

func main() {
	root, err := replicator.ConfigRootFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	manager := replicator.NewRuntimeManager(replicator.Store{Root: root})
	if err := manager.Boot(); err != nil {
		log.Fatalf("restore Replicator runtime: %v", err)
	}
	defer manager.Stop()

	socketPath := replicator.RuntimeSocketPath(root)
	if err := os.MkdirAll(filepath.Dir(socketPath), 0o755); err != nil {
		log.Fatal(err)
	}
	_ = os.Remove(socketPath)
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = listener.Close()
		_ = os.Remove(socketPath)
	}()
	if err := os.Chmod(socketPath, 0o666); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	go func() {
		<-ctx.Done()
		_ = listener.Close()
	}()
	log.Printf("Replicator runtime listening on %s", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Printf("accept Replicator runtime request: %v", err)
			continue
		}
		go handle(conn, manager)
	}
}

func handle(conn net.Conn, manager *replicator.RuntimeManager) {
	defer conn.Close()
	var header [4]byte
	if _, err := io.ReadFull(conn, header[:]); err != nil {
		return
	}
	size := binary.BigEndian.Uint32(header[:])
	if size == 0 || size > maxMessage {
		_ = writeResponse(conn, replicator.RuntimeResponse{
			Version: replicator.RuntimeProtocolVersion,
			Error:   &replicator.RuntimeError{Code: "INVALID_REQUEST", Message: "invalid request size"},
		})
		return
	}
	body := make([]byte, int(size))
	if _, err := io.ReadFull(conn, body); err != nil {
		return
	}
	var request replicator.RuntimeRequest
	if err := json.Unmarshal(body, &request); err != nil {
		_ = writeResponse(conn, replicator.RuntimeResponse{
			Version: replicator.RuntimeProtocolVersion,
			Error:   &replicator.RuntimeError{Code: "INVALID_REQUEST", Message: err.Error()},
		})
		return
	}
	_ = writeResponse(conn, replicator.HandleRuntimeRequest(manager, request))
}

func writeResponse(conn net.Conn, response replicator.RuntimeResponse) error {
	body, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if len(body) > maxMessage {
		return fmt.Errorf("response exceeds maximum message size")
	}
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], uint32(len(body)))
	if _, err := conn.Write(header[:]); err != nil {
		return err
	}
	_, err = conn.Write(body)
	return err
}
