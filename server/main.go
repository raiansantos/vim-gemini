package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github/raiansantos/vim-gemini/gemini"
	"log"
	"net"
	"os"
)

type Client struct {
	conn     net.Conn
	gmClient *gemini.Client
}

func (c *Client) createVimAnswer(ctx context.Context, data string) []byte {
	resp := make(map[string]any)
	resp["answer"] = data

	dataBts, err := json.Marshal(resp)
	if err != nil {
		return nil
	}

	return dataBts
}

func (c *Client) executeClient(ctx context.Context, command map[string]any) []byte {
	cmd, ok := command["command"].(string)
	if !ok {
		fmt.Println("invalid command")
		return []byte{}
	}

	filetype, ok := command["filetype"].(string)
	if !ok {
		fmt.Println("invalid filetype")
		return []byte{}
	}

	data, ok := command["data"].(string)
	if !ok {
		fmt.Println("invalid data")
		return []byte{}
	}

	if cmd == "explain" {
		answer, err := c.gmClient.ExplainCode(ctx, filetype, data)
		if err != nil {
			fmt.Println(err)
			return []byte{}
		}
		return c.createVimAnswer(ctx, answer)
	}

	if cmd == "debug" {
		answer, err := c.gmClient.DebugCode(ctx, filetype, data)
		if err != nil {
			fmt.Println(err)
			return []byte{}
		}
		return c.createVimAnswer(ctx, answer)
	}

	return []byte{}
}

func (c *Client) handleRequest(ctx context.Context) {
	defer c.conn.Close()

	buf := make([]byte, 1024)

	n, err := c.conn.Read(buf)
	if err != nil {
		log.Println("error reading data", err)
		return
	}

	var data map[string]any
	err = json.Unmarshal(buf[:n], &data)
	if err != nil {
		log.Println("invalid JSON", string(buf[:n]), err)
		return
	}

	_, err = c.conn.Write(c.executeClient(ctx, data))
	if err != nil {
		fmt.Println("error returning", err)
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "32000"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server started. Listening at", fmt.Sprintf("%s:%s", "0.0.0.0", port))

	gmClient, err := gemini.New(os.Getenv("GEMINI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn:     conn,
			gmClient: gmClient,
		}
		go client.handleRequest(context.Background())
	}
}
